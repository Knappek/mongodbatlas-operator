package mongodbatlasproject

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
	utils "github.com/Knappek/mongodbatlas-operator/pkg/utils"
	"github.com/go-logr/logr"

	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
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

	// Creates a new MongoDB Atlas Project
	err = createMongoDBAtlasProject(reqLogger, atlasProject)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Check if the APP CR was marked to be deleted
	isAtlasProjectToBeDeleted := atlasProject.GetDeletionTimestamp() != nil
	if isAtlasProjectToBeDeleted {
		// TODO(user): Add the cleanup steps that the operator needs to do before the CR can be deleted
		err := deleteMongoDBAtlasProject(reqLogger, atlasProject)
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
		return reconcile.Result{}, nil
	}

	// Add finalizer for this CR
	if err := r.addFinalizer(reqLogger, atlasProject); err != nil {
		return reconcile.Result{}, err
	}

	err = r.client.Status().Update(context.TODO(), atlasProject)
	if err != nil {
		return reconcile.Result{}, err
	}
	// MongoDB Atlas Project successfully created
	// Requeue to periodically reconcile the CR MongoDBAtlasProject in order to recreate a manually deleted Atlas project
	return reconcile.Result{RequeueAfter: time.Second * 30}, nil
}

func createMongoDBAtlasProject(reqLogger logr.Logger, cr *knappekv1alpha1.MongoDBAtlasProject) error {
	clientset, err := getKubernetesClient()
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
		cr.Status.OrgID = p.OrgID
		cr.Status.Name = p.Name
		cr.Status.Status = p.Created
		cr.Status.ClusterCount = p.ClusterCount
		reqLogger.Info("MongoDB Atlas Project %s created. Project ID: %s", p.Name, p.ID)

		return nil
	}
	// project already exists
	return nil
}

func deleteMongoDBAtlasProject(reqLogger logr.Logger, cr *knappekv1alpha1.MongoDBAtlasProject) error {
	clientset, err := getKubernetesClient()
	if err != nil {
		panic(err.Error())
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
	var p *ma.Project
	var resp *http.Response

	atlasProjectID := cr.Status.ID
	if atlasProjectID == "" {
		reqLogger.Info("MongoDBAtlasProject CustomResource %s has empty .status.id. Searching Project by name and try to delete Project afterwards", cr.Name)
		p, resp, err = atlasClient.Projects.GetByName(cr.Name)
		if err != nil {
			if resp.StatusCode == 404 {
				reqLogger.Info("MongoDBAtlasProject CustomResource %s does not exist in Atlas. Deleting CR.")
				return nil
			}
			return fmt.Errorf("Error getting MongoDB Project by Name %s: %s", cr.Name, err)
		}
		atlasProjectID = p.ID
	}
	// check if project exists
	_, _, err = atlasClient.Projects.Get(atlasProjectID)
	if err != nil {
		// project does not exist, skip doing something
		reqLogger.Info("MongoDB Atlas Project %s (ID: %s) does not exist in Atlas. Just remove the Custom Resource.", cr.Name, atlasProjectID)
		return nil
	}
	// project exists and can be deleted
	resp, err = atlasClient.Projects.Delete(atlasProjectID)
	if err != nil {
		return fmt.Errorf("(%v) Error deleting MongoDB Project %s: %s", resp.StatusCode, atlasProjectID, err)
	}
	return nil
}

//addFinalizer will add this attribute to the Memcached CR
func (r *ReconcileMongoDBAtlasProject) addFinalizer(reqLogger logr.Logger, cr *knappekv1alpha1.MongoDBAtlasProject) error {
	if len(cr.GetFinalizers()) < 1 && cr.GetDeletionTimestamp() == nil {
		reqLogger.Info("Adding Finalizer for the MongoDB Atlas Project")
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

func getKubernetesClient() (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		// creates the in-cluster config
		config, err = rest.InClusterConfig()
	} else {
		// creates out-of-cluster config
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	return kubernetes.NewForConfig(config)
}
