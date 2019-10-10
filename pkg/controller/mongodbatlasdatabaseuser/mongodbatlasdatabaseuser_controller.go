package mongodbatlasdatabaseuser

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
	"github.com/Knappek/mongodbatlas-operator/pkg/config"

	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_mongodbatlasdatabaseuser")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new MongoDBAtlasDatabaseUser Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMongoDBAtlasDatabaseUser{client: mgr.GetClient(), scheme: mgr.GetScheme(), atlasClient: config.GetAtlasClient()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("mongodbatlasdatabaseuser-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource MongoDBAtlasDatabaseUser
	err = c.Watch(&source.Kind{Type: &knappekv1alpha1.MongoDBAtlasDatabaseUser{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileMongoDBAtlasDatabaseUser{}

// ReconcileMongoDBAtlasDatabaseUser reconciles a MongoDBAtlasDatabaseUser object
type ReconcileMongoDBAtlasDatabaseUser struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client      client.Client
	scheme      *runtime.Scheme
	atlasClient *ma.Client
}

// Reconcile reads that state of the MongoDBAtlasDatabaseUser object and makes changes based on the state read
// and what is in the MongoDBAtlasDatabaseUser.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMongoDBAtlasDatabaseUser) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// Fetch the MongoDBAtlasDatabaseUser atlasDatabaseUser
	atlasDatabaseUser := &knappekv1alpha1.MongoDBAtlasDatabaseUser{}
	err := r.client.Get(context.TODO(), request.NamespacedName, atlasDatabaseUser)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	projectName := atlasDatabaseUser.Spec.ProjectName
	atlasProject := &knappekv1alpha1.MongoDBAtlasProject{}
	atlasProjectNamespacedName := types.NamespacedName{
		Name:      projectName,
		Namespace: atlasDatabaseUser.Namespace,
	}

	err = r.client.Get(context.TODO(), atlasProjectNamespacedName, atlasProject)
	if err != nil {
		return reconcile.Result{}, err
	}

	groupID := atlasProject.Status.ID
	// Define default logger
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "MongoDBAtlasDatabaseUser.Name", request.Name, "MongoDBAtlasDatabaseUser.GroupID", groupID)

	// Check if the MongoDBAtlasDatabaseUser CR was marked to be deleted
	isMongoDBAtlasDatabaseUserToBeDeleted := atlasDatabaseUser.GetDeletionTimestamp() != nil
	if isMongoDBAtlasDatabaseUserToBeDeleted {
		err := deleteMongoDBAtlasDatabaseUser(reqLogger, r.atlasClient, atlasDatabaseUser)
		if err != nil {
			return reconcile.Result{}, err
		}
		err = r.client.Status().Update(context.TODO(), atlasDatabaseUser)
		if err != nil {
			return reconcile.Result{}, err
		}
		// Requeue to periodically reconcile the CR MongoDBAtlasDatabaseUser in order to recreate a manually deleted Atlas DatabaseUser
		return reconcile.Result{RequeueAfter: time.Second * 30}, nil
	}

	// Create a new MongoDBAtlasDatabaseUser
	isMongoDBAtlasDatabaseUserToBeCreated := reflect.DeepEqual(atlasDatabaseUser.Status, knappekv1alpha1.MongoDBAtlasDatabaseUserStatus{})
	if isMongoDBAtlasDatabaseUserToBeCreated {
		err = createMongoDBAtlasDatabaseUser(reqLogger, r.atlasClient, atlasDatabaseUser, atlasProject)
		if err != nil {
			return reconcile.Result{}, err
		}
		err = r.client.Status().Update(context.TODO(), atlasDatabaseUser)
		if err != nil {
			return reconcile.Result{}, err
		}
		// Add finalizer for this CR
		if err := r.addFinalizer(reqLogger, atlasDatabaseUser); err != nil {
			return reconcile.Result{}, err
		}
		// Requeue to periodically reconcile the CR MongoDBAtlasDatabaseUser in order to recreate a manually deleted Atlas DatabaseUser
		return reconcile.Result{RequeueAfter: time.Second * 30}, nil
	}

	// update existing MongoDBAtlasDatabaseUser
	isMongoDBAtlasDatabaseUserToBeUpdated := knappekv1alpha1.IsMongoDBAtlasDatabaseUserToBeUpdated(atlasDatabaseUser.Spec.MongoDBAtlasDatabaseUserRequestBody, atlasDatabaseUser.Status)
	if isMongoDBAtlasDatabaseUserToBeUpdated {
		err = updateMongoDBAtlasDatabaseUser(reqLogger, r.atlasClient, atlasDatabaseUser, atlasProject)
		if err != nil {
			return reconcile.Result{}, err
		}
		err = r.client.Status().Update(context.TODO(), atlasDatabaseUser)
		if err != nil {
			return reconcile.Result{}, err
		}
		// Requeue to periodically reconcile the CR MongoDBAtlasDatabaseUser in order to recreate a manually deleted Atlas DatabaseUser
		return reconcile.Result{RequeueAfter: time.Second * 30}, nil
	}

	// if no Create/Update/Delete command apply, then fetch the status
	err = getMongoDBAtlasDatabaseUser(reqLogger, r.atlasClient, atlasDatabaseUser)
	if err != nil {
		return reconcile.Result{}, err
	}
	err = r.client.Status().Update(context.TODO(), atlasDatabaseUser)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Requeue to periodically reconcile the CR MongoDBAtlasDatabaseUser in order to recreate a manually deleted Atlas DatabaseUser
	return reconcile.Result{RequeueAfter: time.Second * 30}, nil
}

func createMongoDBAtlasDatabaseUser(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasDatabaseUser, ap *knappekv1alpha1.MongoDBAtlasProject) error {
	groupID := ap.Status.ID
	name := cr.Name
	params := getDatabaseUserParams(cr)
	c, resp, err := atlasClient.DatabaseUsers.Create(groupID, &params)
	if err != nil {
		return fmt.Errorf("(%v) Error creating DatabaseUser %v: %s", resp.StatusCode, name, err)
	}
	if resp.StatusCode == http.StatusOK {
		reqLogger.Info("DatabaseUser created.")
		return updateCRStatus(reqLogger, cr, c)
	}
	return fmt.Errorf("(%v) Error creating DatabaseUser %s: %s", resp.StatusCode, name, err)
}

func updateMongoDBAtlasDatabaseUser(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasDatabaseUser, ap *knappekv1alpha1.MongoDBAtlasProject) error {
	groupID := ap.Status.ID
	name := cr.Name
	params := getDatabaseUserParams(cr)
	c, resp, err := atlasClient.DatabaseUsers.Update(groupID, name, &params)
	if err != nil {
		return fmt.Errorf("Error updating DatabaseUser %v: %s", name, err)
	}
	if resp.StatusCode == http.StatusOK {
		reqLogger.Info("DatabaseUser updated.")
		return updateCRStatus(reqLogger, cr, c)
	}
	return fmt.Errorf("(%v) Error updating DatabaseUser %s: %s", resp.StatusCode, name, err)
}

func deleteMongoDBAtlasDatabaseUser(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasDatabaseUser) error {
	groupID := cr.Status.GroupID
	name := cr.Name
	// cluster exists and can be deleted
	resp, err := atlasClient.DatabaseUsers.Delete(groupID, name)
	if err != nil {
		if resp.StatusCode == http.StatusNotFound {
			reqLogger.Info("DatabaseUser does not exist in Atlas. Deleting CR.")
			// Update finalizer to allow delete CR
			cr.SetFinalizers(nil)
			// CR can be deleted - Requeue
			return nil
		}
		return fmt.Errorf("(%v) Error deleting DatabaseUser %s: %s", resp.StatusCode, name, err)
	}
	if resp.StatusCode == http.StatusOK {
		// Update finalizer to allow delete CR
		cr.SetFinalizers(nil)
		reqLogger.Info("DatabaseUser deleted.")
		return nil
	}
	return fmt.Errorf("(%v) Error deleting DatabaseUser %s: %s", resp.StatusCode, name, err)
}

func getMongoDBAtlasDatabaseUser(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasDatabaseUser) error {
	groupID := cr.Status.GroupID
	name := cr.Name
	c, resp, err := atlasClient.DatabaseUsers.Get(groupID, name)
	if err != nil {
		return fmt.Errorf("(%v) Error fetching DatabaseUser information %s: %s", resp.StatusCode, name, err)
	}
	err = updateCRStatus(reqLogger, cr, c)
	if err != nil {
		return fmt.Errorf("Error updating DatabaseUser CR Status: %s", err)
	}
	return nil
}

func getDatabaseUserParams(cr *knappekv1alpha1.MongoDBAtlasDatabaseUser) ma.DatabaseUser {
	return ma.DatabaseUser{
		Username:     cr.Name,
		Password:     cr.Spec.Password,
		DatabaseName: "admin",
		Roles: cr.Spec.Roles,
	}
}

func updateCRStatus(reqLogger logr.Logger, cr *knappekv1alpha1.MongoDBAtlasDatabaseUser, c *ma.DatabaseUser) error {
	// update status field in CR
	cr.Status.Username = c.Username
	cr.Status.GroupID = c.GroupID
	cr.Status.DatabaseName = c.DatabaseName
	cr.Status.Roles = c.Roles
	return nil
}

func (r *ReconcileMongoDBAtlasDatabaseUser) addFinalizer(reqLogger logr.Logger, cr *knappekv1alpha1.MongoDBAtlasDatabaseUser) error {
	if len(cr.GetFinalizers()) < 1 && cr.GetDeletionTimestamp() == nil {
		cr.SetFinalizers([]string{"finalizer.knappek.com"})

		// Update CR
		err := r.client.Update(context.TODO(), cr)
		if err != nil {
			reqLogger.Error(err, "Failed to update DatabaseUser with finalizer")
			return err
		}
	}
	return nil
}
