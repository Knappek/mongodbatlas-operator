package mongodbatlasproject

import (
	"context"
	"fmt"
	"time"

	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
	utils "github.com/Knappek/mongodbatlas-operator/pkg/utils"

	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
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

	// Creates a new MongoDB Atlas Project
	err = newMongoDBAtlasProject(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	err = r.client.Status().Update(context.TODO(), instance)
	if err != nil {
		return reconcile.Result{}, err
	}
	// MongoDB Atlas Project successfully created
	// Requeue to periodically reconcile the CR MongoDBAtlasProject in order to recreate a manually deleted Atlas project
	return reconcile.Result{RequeueAfter: time.Second * 30}, nil
}

func newMongoDBAtlasProject(cr *knappekv1alpha1.MongoDBAtlasProject) error {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	orgID, err := clientset.CoreV1().Secrets(cr.Namespace).Get(cr.Spec.OrgID.SecretName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("Error fetching OrgID secret %v: %s", cr.Spec.OrgID.SecretName, err)
	}
	apiKey, err := clientset.CoreV1().Secrets(cr.Namespace).Get(cr.Spec.APIKey.SecretName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("Error fetching APIKey secret %v: %s", cr.Spec.APIKey.SecretName, err)
	}

	// create MongoDB Atlas client
	atlasConfig := utils.Config{
		AtlasUsername: cr.Spec.Username,
		AtlasAPIKey:   string(apiKey.Data[cr.Spec.APIKey.Key]),
	}
	atlasClient := atlasConfig.NewClient()

	params := ma.Project{
		OrgID: string(orgID.Data[cr.Spec.OrgID.Key]),
		Name:  cr.Name,
	}
	// check if project already exists
	_, _, err = atlasClient.Projects.GetByName(cr.Name)
	if err != nil {
		p, _, err := atlasClient.Projects.Create(&params)
		if err != nil {
			return fmt.Errorf("Error creating MongoDB Atlas Project %v: %s", cr.Name, err)
		}

		cr.Status.ID = p.ID

		p, resp, err := atlasClient.Projects.Get(cr.Status.ID)
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
	// project already exists, try to update it
	return nil

}
