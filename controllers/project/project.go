/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package project

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	mongodbatlasv1alpha1 "github.com/knappek/mongodbatlas-operator/api/v1alpha1"
	"github.com/knappek/mongodbatlas-operator/util"
)

// MongoDBAtlasProjectReconciler reconciles a MongoDBAtlasProject object
type MongoDBAtlasProjectReconciler struct {
	client.Client
	Log                  logr.Logger
	Scheme               *runtime.Scheme
	AtlasClient          *ma.Client
	ReconciliationConfig *util.ReconciliationConfig
}

// +kubebuilder:rbac:groups=mongodbatlas.knappek.com,resources=mongodbatlasprojects,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=mongodbatlas.knappek.com,resources=mongodbatlasprojects/status,verbs=get;update;patch

func (r *MongoDBAtlasProjectReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	reqLogger := r.Log.WithValues("mongodbatlasproject", req.NamespacedName)

	// Fetch the MongoDBAtlasProject instance
	atlasProject := &mongodbatlasv1alpha1.MongoDBAtlasProject{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, atlasProject)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile req.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the req.
		return ctrl.Result{}, err
	}

	// Creates a new MongoDB Atlas Project with the name defined in atlasProject iff it does not yet exist
	err = createMongoDBAtlasProject(reqLogger, r.AtlasClient, atlasProject)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Update CR Status
	err = r.Client.Status().Update(context.TODO(), atlasProject)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Check if the MongoDBAtlasProject CR was marked to be deleted
	isMongoDBAtlasProjectToBeDeleted := atlasProject.GetDeletionTimestamp() != nil
	if isMongoDBAtlasProjectToBeDeleted {
		// TODO(user): Add the cleanup steps that the operator needs to do before the CR can be deleted
		err := deleteMongoDBAtlasProject(reqLogger, r.AtlasClient, atlasProject)
		if err != nil {
			return ctrl.Result{}, err
		}
		// Update CR
		err = r.Client.Update(context.TODO(), atlasProject)
		if err != nil {
			return ctrl.Result{}, err
		}
		// Requeue to periodically reconcile the CR MongoDBAtlasProject in order to recreate a manually deleted Atlas DatabaseUser
		return ctrl.Result{RequeueAfter: r.ReconciliationConfig.Time}, nil
	}
	// Add finalizer for this CR
	if err := r.addFinalizer(reqLogger, atlasProject); err != nil {
		return ctrl.Result{}, err
	}
	// MongoDB Atlas Project successfully created
	// Requeue to periodically reconcile the CR MongoDBAtlasProject in order to recreate a manually deleted Atlas project
	return ctrl.Result{RequeueAfter: r.ReconciliationConfig.Time}, nil
}

func createMongoDBAtlasProject(reqLogger logr.Logger, atlasClient *ma.Client, cr *mongodbatlasv1alpha1.MongoDBAtlasProject) error {
	params := ma.Project{
		OrgID: cr.Spec.OrgID,
		Name:  cr.Name,
	}
	// check if project already exists
	p, _, err := atlasClient.Projects.GetByName(cr.Name)
	if err != nil {
		p, _, err = atlasClient.Projects.Create(&params)
		if err != nil {
			return fmt.Errorf("Error creating Project %v: %s", cr.Name, err)
		}
		reqLogger.Info("Project created.", "MongoDBAtlasProject.ID", p.ID)
	}
	cr.Status.ID = p.ID
	cr.Status.OrgID = p.OrgID
	cr.Status.Name = p.Name
	cr.Status.Created = p.Created
	cr.Status.ClusterCount = p.ClusterCount

	return nil
}

func deleteMongoDBAtlasProject(reqLogger logr.Logger, atlasClient *ma.Client, cr *mongodbatlasv1alpha1.MongoDBAtlasProject) error {
	// check if project exists
	p, resp, err := atlasClient.Projects.GetByName(cr.Name)
	if err != nil {
		if resp.StatusCode == 404 {
			reqLogger.Info("Project does not exist in Atlas. Deleting CR.")
			// Update finalizer to allow delete CR
			cr.SetFinalizers(nil)
			return nil
		}
		return fmt.Errorf("Error getting MongoDB Project %s: %s", cr.Name, err)
	}

	// project exists and can be deleted
	atlasGroupID := p.ID
	resp, err = atlasClient.Projects.Delete(atlasGroupID)
	if err != nil {
		return fmt.Errorf("(%v) Error deleting MongoDB Project %s: %s", resp.StatusCode, atlasGroupID, err)
	}
	// Update finalizer to allow delete CR
	cr.SetFinalizers(nil)
	reqLogger.Info("Project deleted.", "MongoDBAtlasProject.ID", atlasGroupID)
	return nil
}

func (r *MongoDBAtlasProjectReconciler) addFinalizer(reqLogger logr.Logger, cr *mongodbatlasv1alpha1.MongoDBAtlasProject) error {
	if len(cr.GetFinalizers()) < 1 && cr.GetDeletionTimestamp() == nil {
		cr.SetFinalizers([]string{"finalizer.knappek.com"})

		// Update CR
		err := r.Client.Update(context.TODO(), cr)
		if err != nil {
			reqLogger.Error(err, "Failed to update Project with finalizer")
			return err
		}
	}
	return nil
}

func (r *MongoDBAtlasProjectReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mongodbatlasv1alpha1.MongoDBAtlasProject{}).
		Complete(r)
}
