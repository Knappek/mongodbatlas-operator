package mongodbatlasproject

import (
	"context"
	"fmt"

	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
	utils "github.com/Knappek/mongodbatlas-operator/pkg/utils"

	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_mongodbatlasproject")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new MongoDBAtlasProject Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMongoDBAtlasProject{client: mgr.GetClient(), scheme: mgr.GetScheme()}
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

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner MongoDBAtlasProject
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &knappekv1alpha1.MongoDBAtlasProject{},
	})
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
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a MongoDBAtlasProject object and makes changes based on the state read
// and what is in the MongoDBAtlasProject.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMongoDBAtlasProject) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling MongoDBAtlasProject")

	// Fetch the MongoDBAtlasProject instance
	instance := &knappekv1alpha1.MongoDBAtlasProject{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
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

	// Define a new Pod object
	err = newMongoDBAtlasProject(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	// TODO: Continue here
	// Set MongoDBAtlasProject instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Pod already exists
	found := &corev1.Pod{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Pod created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Pod already exists - don't requeue
	reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}

func newMongoDBAtlasProject(cr *knappekv1alpha1.MongoDBAtlasProject) error {
	// create MongoDB Atlas client
	config := utils.Config{
		AtlasUsername: cr.Spec.Username,
		AtlasAPIKey:   cr.Spec.APIKey,
	}
	client := config.NewClient()

	params := ma.Project{
		OrgID: cr.Spec.OrgID,
		Name:  cr.Name,
	}
	p, _, err := client.Projects.Create(&params)
	if err != nil {
		return fmt.Errorf("Error creating MongoDB Project IP Project %v: %s", cr.Name, err)
	}

	cr.Status.ID = p.ID

	p, resp, err := client.Projects.Get(cr.Status.ID)
	if err != nil {
		if resp.StatusCode == 404 {
			cr.Status.ID = ""
			return fmt.Errorf("(%v) MongoDB Project %s not found: %s", resp.StatusCode, cr.Status.ID, err)
		}
		return fmt.Errorf("Error reading MongoDB Project %s: %s", cr.Status.ID, err)
	}

	cr.Status.OrgID = p.OrgID
	cr.Status.Name = p.Name
	cr.Status.Status = p.Created
	cr.Status.ClusterCount = p.ClusterCount

	return nil
}
