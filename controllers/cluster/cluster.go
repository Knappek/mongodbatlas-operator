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

package cluster

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	mongodbatlasv1alpha1 "github.com/knappek/mongodbatlas-operator/api/v1alpha1"
	"github.com/knappek/mongodbatlas-operator/util"
)

// MongoDBAtlasClusterReconciler reconciles a MongoDBAtlasCluster object
type MongoDBAtlasClusterReconciler struct {
	client.Client
	Log                  logr.Logger
	Scheme               *runtime.Scheme
	AtlasClient          *ma.Client
	ReconciliationConfig *util.ReconciliationConfig
}

// +kubebuilder:rbac:groups=mongodbatlas.knappek.com,resources=mongodbatlasclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=mongodbatlas.knappek.com,resources=mongodbatlasclusters/status,verbs=get;update;patch

func (r *MongoDBAtlasClusterReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()

	// Fetch the MongoDBAtlasCluster atlasCluster
	atlasCluster := &mongodbatlasv1alpha1.MongoDBAtlasCluster{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, atlasCluster)
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

	projectName := atlasCluster.Spec.ProjectName
	atlasProject := &mongodbatlasv1alpha1.MongoDBAtlasProject{}
	atlasProjectNamespacedName := types.NamespacedName{
		Name:      projectName,
		Namespace: atlasCluster.Namespace,
	}
	err = r.Client.Get(context.TODO(), atlasProjectNamespacedName, atlasProject)
	if err != nil {
		return ctrl.Result{}, err
	}
	groupID := atlasProject.Status.ID
	// Define default logger
	reqLogger := r.Log.WithValues("mongodbatlascluster", req.NamespacedName, "mongodbatlascluster.name", req.Name, "mongodbatlascluster.groupID", groupID)

	// Check if the MongoDBAtlasCluster CR was marked to be deleted
	isMongoDBAtlasClusterToBeDeleted := atlasCluster.GetDeletionTimestamp() != nil
	if isMongoDBAtlasClusterToBeDeleted {
		// check if Delete request has already been sent to the MongoDB Atlas API
		if atlasCluster.Status.StateName != "DELETING" && atlasCluster.Status.StateName != "DELETED" {
			err := deleteMongoDBAtlasCluster(reqLogger, r.AtlasClient, atlasCluster)
			if err != nil {
				return ctrl.Result{}, err
			}
			atlasCluster.Status.StateName = "DELETING"
			err = r.Client.Status().Update(context.TODO(), atlasCluster)
			if err != nil {
				return ctrl.Result{}, err
			}
			reqLogger.Info("Wait until Cluster has been deleted.")
			// Requeue after 20 seconds and check again for the status until CR can be deleted
			return ctrl.Result{RequeueAfter: time.Second * 20}, nil
		}

		// wait until cluster has been deleted successfully
		_, resp, err := r.AtlasClient.Clusters.Get(groupID, atlasCluster.Name)
		if err != nil {
			if resp.StatusCode == 404 {
				reqLogger.Info("Cluster deleted.")
				// Update finalizer to allow delete CR
				atlasCluster.SetFinalizers(nil)
				// Update CR
				err = r.Client.Update(context.TODO(), atlasCluster)
				if err != nil {
					return ctrl.Result{}, err
				}
				// MongoDB Atlas Cluster successfully deleted
				return ctrl.Result{}, nil
			}
			return ctrl.Result{}, err
		}
		// if err == nil, cluster still exists - Requeue after 20 seconds
		return ctrl.Result{RequeueAfter: time.Second * 20}, nil
	}

	// Create a new cluster
	isMongoDBAtlasClusterToBeCreated := reflect.DeepEqual(atlasCluster.Status, mongodbatlasv1alpha1.MongoDBAtlasClusterStatus{})
	if isMongoDBAtlasClusterToBeCreated {
		// check if Create request has already been sent to the MongoDB Atlas API
		if atlasCluster.Status.StateName != "CREATING" {
			err = createMongoDBAtlasCluster(reqLogger, r.AtlasClient, atlasCluster, atlasProject)
			if err != nil {
				return ctrl.Result{}, err
			}
			atlasCluster.Status.StateName = "CREATING"
			err = r.Client.Status().Update(context.TODO(), atlasCluster)
			if err != nil {
				return ctrl.Result{}, err
			}
			// Add finalizer for this CR
			if err := r.addFinalizer(reqLogger, atlasCluster); err != nil {
				return ctrl.Result{}, err
			}
			// Requeue to periodically reconcile the CR MongoDBAtlasCluster in order to recreate a manually deleted Atlas cluster
			return ctrl.Result{RequeueAfter: r.ReconciliationConfig.Time}, nil
		}
	}

	// update existing cluster
	isMongoDBAtlasClusterToBeUpdated := mongodbatlasv1alpha1.IsMongoDBAtlasClusterToBeUpdated(atlasCluster.Spec.MongoDBAtlasClusterRequestBody, atlasCluster.Status.MongoDBAtlasClusterRequestBody)
	if isMongoDBAtlasClusterToBeUpdated {
		// check if Update request has already been sent to the MongoDB Atlas API
		if atlasCluster.Status.StateName != "UPDATING" {
			err = updateMongoDBAtlasCluster(reqLogger, r.AtlasClient, atlasCluster, atlasProject)
			if err != nil {
				return ctrl.Result{}, err
			}
			atlasCluster.Status.StateName = "UPDATING"
			err = r.Client.Status().Update(context.TODO(), atlasCluster)
			if err != nil {
				return ctrl.Result{}, err
			}
			// Requeue to periodically reconcile the CR MongoDBAtlasCluster in order to recreate a manually deleted Atlas cluster
			return ctrl.Result{RequeueAfter: r.ReconciliationConfig.Time}, nil
		}
	}

	// if no Create/Update/Delete command apply, then fetch the status
	c, _, err := r.AtlasClient.Clusters.Get(groupID, atlasCluster.Name)
	if err != nil {
		return ctrl.Result{}, err
	}
	err = updateCRStatus(reqLogger, atlasCluster, c)
	if err != nil {
		return ctrl.Result{}, err
	}
	// Update CR Status
	err = r.Client.Status().Update(context.TODO(), atlasCluster)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Requeue to periodically reconcile the CR MongoDBAtlasCluster in order to recreate a manually deleted Atlas cluster
	return ctrl.Result{RequeueAfter: r.ReconciliationConfig.Time}, nil

}

func createMongoDBAtlasCluster(reqLogger logr.Logger, atlasClient *ma.Client, cr *mongodbatlasv1alpha1.MongoDBAtlasCluster, ap *mongodbatlasv1alpha1.MongoDBAtlasProject) error {
	groupID := ap.Status.ID
	params := getClusterParams(cr)

	c, _, err := atlasClient.Clusters.Create(groupID, &params)
	if err != nil {
		return fmt.Errorf("Error creating Cluster %v: %s", cr.Name, err)
	}
	reqLogger.Info("Sent request to create Cluster.")
	return updateCRStatus(reqLogger, cr, c)
}

func updateMongoDBAtlasCluster(reqLogger logr.Logger, atlasClient *ma.Client, cr *mongodbatlasv1alpha1.MongoDBAtlasCluster, ap *mongodbatlasv1alpha1.MongoDBAtlasProject) error {
	groupID := ap.Status.ID
	params := getClusterParams(cr)
	c, _, err := atlasClient.Clusters.Update(groupID, cr.Name, &params)
	if err != nil {
		return fmt.Errorf("Error updating Cluster %v: %s", cr.Name, err)
	}
	reqLogger.Info("Sent request to update Cluster.")
	return updateCRStatus(reqLogger, cr, c)
}

func deleteMongoDBAtlasCluster(reqLogger logr.Logger, atlasClient *ma.Client, cr *mongodbatlasv1alpha1.MongoDBAtlasCluster) error {
	groupID := cr.Status.GroupID
	clusterName := cr.Status.Name
	// cluster exists and can be deleted
	resp, err := atlasClient.Clusters.Delete(groupID, clusterName)
	if err != nil {
		if resp.StatusCode == http.StatusNotFound {
			reqLogger.Info("Cluster does not exist in Atlas. Deleting CR.")
			// CR can be deleted - Requeue
			return nil
		}
		return fmt.Errorf("(%v) Error deleting Cluster %s: %s", resp.StatusCode, clusterName, err)
	}
	reqLogger.Info("Sent request to delete Cluster.")
	return nil
}

func getClusterParams(cr *mongodbatlasv1alpha1.MongoDBAtlasCluster) ma.Cluster {
	return ma.Cluster{
		Name:                  cr.Name,
		MongoDBMajorVersion:   cr.Spec.MongoDBMajorVersion,
		DiskSizeGB:            util.StringToFloat64(cr.Spec.DiskSizeGB),
		BackupEnabled:         cr.Spec.BackupEnabled,
		ProviderBackupEnabled: cr.Spec.ProviderBackupEnabled,
		ReplicationSpec:       cr.Spec.ReplicationSpec,
		NumShards:             cr.Spec.NumShards,
		AutoScaling:           cr.Spec.AutoScaling,
		ProviderSettings:      cr.Spec.ProviderSettings,
	}
}

func updateCRStatus(reqLogger logr.Logger, cr *mongodbatlasv1alpha1.MongoDBAtlasCluster, c *ma.Cluster) error {
	// save old stateName for later
	oldStateName := cr.Status.StateName
	// update status field in CR
	cr.Status.ID = c.ID
	cr.Status.GroupID = c.GroupID
	cr.Status.Name = c.Name
	cr.Status.MongoDBVersion = c.MongoDBVersion
	cr.Status.MongoDBMajorVersion = c.MongoDBMajorVersion
	cr.Status.MongoURI = c.MongoURI
	cr.Status.MongoURIUpdated = c.MongoURIUpdated
	cr.Status.MongoURIWithOptions = c.MongoURIWithOptions
	cr.Status.SrvAddress = c.SrvAddress
	cr.Status.DiskSizeGB = util.Float64ToString(c.DiskSizeGB)
	cr.Status.BackupEnabled = c.BackupEnabled
	cr.Status.ProviderBackupEnabled = c.ProviderBackupEnabled
	cr.Status.StateName = c.StateName
	cr.Status.ReplicationSpec = c.ReplicationSpec
	cr.Status.NumShards = c.NumShards
	cr.Status.Paused = c.Paused
	cr.Status.AutoScaling = c.AutoScaling
	cr.Status.ProviderSettings = c.ProviderSettings
	// compare old stateName with new stateName
	newStateName := cr.Status.StateName
	if oldStateName != newStateName {
		if oldStateName == "CREATING" && newStateName == "IDLE" {
			reqLogger.Info("Cluster created.")
		}
		if oldStateName == "UPDATING" && newStateName == "IDLE" {
			reqLogger.Info("Cluster updated successfully.")
		}
	}
	return nil
}

func (r *MongoDBAtlasClusterReconciler) addFinalizer(reqLogger logr.Logger, cr *mongodbatlasv1alpha1.MongoDBAtlasCluster) error {
	if len(cr.GetFinalizers()) < 1 && cr.GetDeletionTimestamp() == nil {
		cr.SetFinalizers([]string{"finalizer.knappek.com"})

		// Update CR
		err := r.Client.Update(context.TODO(), cr)
		if err != nil {
			reqLogger.Error(err, "Failed to update Cluster with finalizer")
			return err
		}
	}
	return nil
}

func (r *MongoDBAtlasClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mongodbatlasv1alpha1.MongoDBAtlasCluster{}).
		Complete(r)
}
