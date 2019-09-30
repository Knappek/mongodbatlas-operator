package mongodbatlasdatabaseuser

import (
	"context"
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
		// check if Delete request has already been sent to the MongoDB Atlas API
		//
		// TODO
		//

		// wait until MongoDBAtlasDatabaseUser has been deleted successfully
		//
		// TODO
		//
		// if err == nil, MongoDBAtlasDatabaseUser still exists - Requeue after 20 seconds
		return reconcile.Result{RequeueAfter: time.Second * 20}, nil
	}

	// Create a new MongoDBAtlasDatabaseUser
	isMongoDBAtlasDatabaseUserToBeCreated := reflect.DeepEqual(atlasDatabaseUser.Status, knappekv1alpha1.MongoDBAtlasDatabaseUserStatus{})
	if isMongoDBAtlasDatabaseUserToBeCreated {
		// check if Create request has already been sent to the MongoDB Atlas API
		if atlasDatabaseUser.Status.StateName != "CREATING" {
			err = createMongoDBAtlasDatabaseUser(reqLogger, r.atlasClient, atlasDatabaseUser, atlasProject)
			if err != nil {
				return reconcile.Result{}, err
			}
			atlasDatabaseUser.Status.StateName = "CREATING"
			err = r.client.Status().Update(context.TODO(), atlasDatabaseUser)
			if err != nil {
				return reconcile.Result{}, err
			}
			// Add finalizer for this CR
			if err := r.addFinalizer(reqLogger, atlasDatabaseUser); err != nil {
				return reconcile.Result{}, err
			}
			// Requeue after 30 seconds and check again for the status until CR can be deleted
			return reconcile.Result{RequeueAfter: time.Second * 30}, nil
		}
	}

	// update existing MongoDBAtlasDatabaseUser
	isMongoDBAtlasDatabaseUserToBeUpdated := knappekv1alpha1.IsMongoDBAtlasDatabaseUserToBeUpdated(atlasDatabaseUser.Spec.MongoDBAtlasDatabaseUserRequestBody, atlasDatabaseUser.Status.MongoDBAtlasDatabaseUserRequestBody)
	if isMongoDBAtlasDatabaseUserToBeUpdated {
		// check if Update request has already been sent to the MongoDB Atlas API
		if atlasDatabaseUser.Status.StateName != "UPDATING" {
			err = updateMongoDBAtlasDatabaseUser(reqLogger, r.atlasClient, atlasDatabaseUser, atlasProject)
			if err != nil {
				return reconcile.Result{}, err
			}
			atlasDatabaseUser.Status.StateName = "UPDATING"
			err = r.client.Status().Update(context.TODO(), atlasDatabaseUser)
			if err != nil {
				return reconcile.Result{}, err
			}
			// Requeue after 30 seconds and check again for the status until CR can be deleted
			return reconcile.Result{RequeueAfter: time.Second * 30}, nil
		}
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
	//
	// TODO
	//
	if err != nil {
		return fmt.Errorf("Error creating DatabaseUser %v: %s", cr.Name, err)
	}
	reqLogger.Info("Sent request to create DatabaseUser.", "MongoDBAtlasDatabaseUser.GroupID", groupID)
	return updateCRStatus(reqLogger, cr, c)
}

func updateMongoDBAtlasDatabaseUser(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasDatabaseUser, ap *knappekv1alpha1.MongoDBAtlasProject) error {
	groupID := ap.Status.ID
	//
	// TODO
	//
	if err != nil {
		return fmt.Errorf("Error updating DatabaseUser %v: %s", cr.Name, err)
	}
	reqLogger.Info("Sent request to update DatabaseUser.", "MongoDBAtlasDatabaseUser.GroupID", groupID)
	return updateCRStatus(reqLogger, cr, c)
}

func deleteMongoDBAtlasDatabaseUser(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasDatabaseUser) error {
	groupID := cr.Status.GroupID
	//
	// TODO
	//
	if err != nil {
		return fmt.Errorf("Error deleting DatabaseUser %v: %s", cr.Name, err)
	}
	reqLogger.Info("Sent request to delete DatabaseUser.", "MongoDBAtlasDatabaseUser.GroupID", groupID)
	return nil
}

func getMongoDBAtlasDatabaseUser(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasDatabaseUser) (*DatabaseUser, *http.Response, error) {
	//
	// TODO
	//
	return kindShort, resp, err
}

func updateCRStatus(reqLogger logr.Logger, cr *knappekv1alpha1.MongoDBAtlasDatabaseUser, c *ma.DatabaseUser) error {
	// update status field in CR
	cr.Status.ID = kindShort.ID
	cr.Status.GroupID = kindShort.GroupID
	cr.Status.Name = kindShort.Name
	//
	// TODO
	//
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
