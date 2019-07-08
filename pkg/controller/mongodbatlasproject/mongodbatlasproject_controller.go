package mongodbatlasproject

import (
	"context"
	"fmt"
	"time"

	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
	"github.com/Knappek/mongodbatlas-operator/pkg/config"

	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_mongodbatlasproject")

// Add creates a new MongoDBAtlasProject Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMongoDBAtlasProject{client: mgr.GetClient(), scheme: mgr.GetScheme(), atlasClient: config.GetAtlasClient()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("mongodbatlasproject-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource MongoDBAtlasProject
	err = c.Watch(&source.Kind{Type: &knappekv1alpha1.MongoDBAtlasProject{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileMongoDBAtlasProject{}

// ReconcileMongoDBAtlasProject reconciles a MongoDBAtlasProject object
type ReconcileMongoDBAtlasProject struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client      client.Client
	scheme      *runtime.Scheme
	atlasClient *ma.Client
}

// Reconcile reads that state of the cluster for a MongoDBAtlasProject object and makes changes based on the state read
// and what is in the MongoDBAtlasProject.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMongoDBAtlasProject) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "MongoDBAtlasProject.Name", request.Name)

	// Fetch the MongoDBAtlasProject instance
	atlasProject := &knappekv1alpha1.MongoDBAtlasProject{}
	err := r.client.Get(context.TODO(), request.NamespacedName, atlasProject)
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

	// Creates a new MongoDB Atlas Project with the name defined in atlasProject iff it does not yet exist
	err = createMongoDBAtlasProject(reqLogger, r.atlasClient, atlasProject)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Update CR Status
	err = r.client.Status().Update(context.TODO(), atlasProject)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Check if the MongoDBAtlasProject CR was marked to be deleted
	isMongoDBAtlasProjectToBeDeleted := atlasProject.GetDeletionTimestamp() != nil
	if isMongoDBAtlasProjectToBeDeleted {
		// TODO(user): Add the cleanup steps that the operator needs to do before the CR can be deleted
		err := deleteMongoDBAtlasProject(reqLogger, r.atlasClient, atlasProject)
		if err != nil {
			return reconcile.Result{}, err
		}
		// Update finalizer to allow delete CR
		atlasProject.SetFinalizers(nil)

		// Update CR
		err = r.client.Update(context.TODO(), atlasProject)
		if err != nil {
			return reconcile.Result{}, err
		}
		// MongoDB Atlas Project successfully deleted
		return reconcile.Result{}, nil
	}
	// Add finalizer for this CR
	if err := r.addFinalizer(reqLogger, atlasProject); err != nil {
		return reconcile.Result{}, err
	}
	// MongoDB Atlas Project successfully created
	// Requeue to periodically reconcile the CR MongoDBAtlasProject in order to recreate a manually deleted Atlas project
	return reconcile.Result{RequeueAfter: time.Second * 30}, nil
}

func createMongoDBAtlasProject(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasProject) error {
	params := ma.Project{
		OrgID: cr.Spec.OrgID,
		Name:  cr.Name,
	}
	// check if project already exists
	p, _, err := atlasClient.Projects.GetByName(cr.Name)
	if err != nil {
		p, _, err = atlasClient.Projects.Create(&params)
		if err != nil {
			return fmt.Errorf("Error creating MongoDB Atlas Project %v: %s", cr.Name, err)
		}
		reqLogger.Info("MongoDB Atlas Project created.", "MongoDBAtlasProject.ID", p.ID)
	}
	cr.Status.ID = p.ID
	cr.Status.OrgID = p.OrgID
	cr.Status.Name = p.Name
	cr.Status.Created = p.Created
	cr.Status.ClusterCount = p.ClusterCount

	return nil
}

func deleteMongoDBAtlasProject(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasProject) error {
	// check if project exists
	p, resp, err := atlasClient.Projects.GetByName(cr.Name)
	if err != nil {
		if resp.StatusCode == 404 {
			reqLogger.Info("MongoDB Atlas Project does not exist in Atlas. Deleting CR.")
			return nil
		}
		return fmt.Errorf("Error getting MongoDB Project %s: %s", cr.Name, err)
	}

	// project exists and can be deleted
	atlasProjectID := p.ID
	resp, err = atlasClient.Projects.Delete(atlasProjectID)
	if err != nil {
		return fmt.Errorf("(%v) Error deleting MongoDB Project %s: %s", resp.StatusCode, atlasProjectID, err)
	}
	reqLogger.Info("MongoDB Atlas Project deleted.", "MongoDBAtlasProject.ID", atlasProjectID)
	return nil
}

//addFinalizer will add this attribute to the Memcached CR
func (r *ReconcileMongoDBAtlasProject) addFinalizer(reqLogger logr.Logger, cr *knappekv1alpha1.MongoDBAtlasProject) error {
	if len(cr.GetFinalizers()) < 1 && cr.GetDeletionTimestamp() == nil {
		cr.SetFinalizers([]string{"finalizer.knappek.com"})

		// Update CR
		err := r.client.Update(context.TODO(), cr)
		if err != nil {
			reqLogger.Error(err, "Failed to update MongoDB Atlas Project with finalizer")
			return err
		}
	}
	return nil
}
