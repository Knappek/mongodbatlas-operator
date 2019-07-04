package mongodbatlascluster

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
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_mongodbatlascluster")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new MongoDBAtlasCluster Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	// create MongoDB Atlas client
	atlasClient := config.GetAtlasClient()
	return &ReconcileMongoDBAtlasCluster{client: mgr.GetClient(), scheme: mgr.GetScheme(), atlasClient: atlasClient}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("mongodbatlascluster-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource MongoDBAtlasCluster
	err = c.Watch(&source.Kind{Type: &knappekv1alpha1.MongoDBAtlasCluster{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileMongoDBAtlasCluster{}

// ReconcileMongoDBAtlasCluster reconciles a MongoDBAtlasCluster object
type ReconcileMongoDBAtlasCluster struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client      client.Client
	scheme      *runtime.Scheme
	atlasClient *ma.Client
}

// Reconcile reads that state of the cluster for a MongoDBAtlasCluster object and makes changes based on the state read
// and what is in the MongoDBAtlasCluster.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMongoDBAtlasCluster) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "MongoDBAtlasCluster.Name", request.Name)

	// Fetch the MongoDBAtlasCluster atlasCluster
	atlasCluster := &knappekv1alpha1.MongoDBAtlasCluster{}
	err := r.client.Get(context.TODO(), request.NamespacedName, atlasCluster)
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

	projectName := atlasCluster.Spec.ProjectName
	atlasProject := &knappekv1alpha1.MongoDBAtlasProject{}
	atlasProjectNamespacedName := types.NamespacedName{
		Name:      projectName,
		Namespace: atlasCluster.Namespace,
	}

	err = r.client.Get(context.TODO(), atlasProjectNamespacedName, atlasProject)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Check if the MongoDBAtlasCluster CR was marked to be deleted
	isMongoDBAtlasClusterToBeDeleted := atlasCluster.GetDeletionTimestamp() != nil
	if isMongoDBAtlasClusterToBeDeleted {
		groupID := atlasProject.Status.ID
		// check if Delete request has already been sent to the MongoDB Atlas API
		if atlasCluster.Status.StateName != "DELETING" && atlasCluster.Status.StateName != "DELETED" {
			err := deleteMongoDBAtlasCluster(reqLogger, r.atlasClient, atlasCluster)
			if err != nil {
				return reconcile.Result{}, err
			}
			reqLogger.Info("Wait until Cluster has been deleted successfully.", "MongoDBAtlasCluster.GroupID", groupID)
		}

		// wait until cluster has been deleted successfully
		c, resp, err := r.atlasClient.Clusters.Get(groupID, atlasCluster.Name)
		if err != nil {
			if resp.StatusCode == 404 {
				reqLogger.Info("MongoDB Atlas Cluster has been deleted successfully.", "MongoDBAtlasCluster.GroupID", groupID)
				// Update finalizer to allow delete CR
				atlasCluster.SetFinalizers(nil)

				// Update CR
				err = r.client.Update(context.TODO(), atlasCluster)
				if err != nil {
					return reconcile.Result{}, err
				}

				// MongoDB Atlas Cluster successfully deleted
				return reconcile.Result{}, nil
			}
			return reconcile.Result{}, err
		}
		// if err == nil, cluster still exists. Update status of CR
		updateMongoDBAtlasClusterCRStatus(atlasCluster, c)
		err = r.client.Status().Update(context.TODO(), atlasCluster)
		if err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{RequeueAfter: time.Second * 20}, nil
	}

	// Creates a new MongoDB Atlas Cluster with the name defined in atlasCluster iff it does not yet exist
	err = createMongoDBAtlasCluster(reqLogger, r.atlasClient, atlasCluster, atlasProject)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Update CR Status
	err = r.client.Status().Update(context.TODO(), atlasCluster)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Add finalizer for this CR
	if err := r.addFinalizer(reqLogger, atlasCluster); err != nil {
		return reconcile.Result{}, err
	}

	// MongoDB Atlas Cluster successfully created
	// Requeue to periodically reconcile the CR MongoDBAtlasCluster in order to recreate a manually deleted Atlas cluster
	return reconcile.Result{RequeueAfter: time.Second * 30}, nil
}

func createMongoDBAtlasCluster(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasCluster, ap *knappekv1alpha1.MongoDBAtlasProject) error {
	groupID := ap.Status.ID
	params := ma.Cluster{
		GroupID:               groupID,
		Name:                  cr.Name,
		MongoDBVersion:        cr.Spec.MongoDBVersion,
		MongoDBMajorVersion:   cr.Spec.MongoDBMajorVersion,
		DiskSizeGB:            cr.Spec.DiskSizeGB,
		BackupEnabled:         cr.Spec.BackupEnabled,
		ProviderBackupEnabled: cr.Spec.ProviderBackupEnabled,
		ReplicationFactor:     cr.Spec.ReplicationFactor,
		ReplicationSpec:       cr.Spec.ReplicationSpec,
		NumShards:             cr.Spec.NumShards,
		AutoScaling:           cr.Spec.AutoScaling,
		ProviderSettings:      cr.Spec.ProviderSettings,
	}
	// check if cluster already exists
	c, _, err := atlasClient.Clusters.Get(groupID, cr.Name)
	if err != nil {
		c, _, err = atlasClient.Clusters.Create(groupID, &params)
		if err != nil {
			return fmt.Errorf("Error creating MongoDB Atlas Cluster %v: %s", cr.Name, err)
		}
		reqLogger.Info("Sent request to create MongoDB Atlas Cluster.", "MongoDBAtlasCluster.GroupID", groupID)
	}
	updateMongoDBAtlasClusterCRStatus(cr, c)

	return nil
}

func deleteMongoDBAtlasCluster(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasCluster) error {
	groupID := cr.Status.GroupID
	clusterName := cr.Status.Name
	// check if cluster exists
	_, _, err := atlasClient.Clusters.Get(groupID, clusterName)
	if err != nil {
		// cluster does not exist, skip doing something
		reqLogger.Info("MongoDB Atlas Cluster does not exist in Atlas. Deleting CR.", "MongoDBAtlasCluster.GroupID", groupID)
		return nil
	}
	// cluster exists and can be deleted
	resp, err := atlasClient.Clusters.Delete(groupID, clusterName)
	if err != nil {
		return fmt.Errorf("(%v) Error deleting MongoDB Cluster %s: %s", resp.StatusCode, clusterName, err)
	}
	reqLogger.Info("Sent request to delete MongoDB Atlas Cluster.", "MongoDBAtlasCluster.GroupID", groupID)
	return nil
}

func updateMongoDBAtlasClusterCRStatus(cr *knappekv1alpha1.MongoDBAtlasCluster, c *ma.Cluster) {
	cr.Status.ID = c.ID
	cr.Status.GroupID = c.GroupID
	cr.Status.Name = c.Name
	cr.Status.MongoDBVersion = c.MongoDBVersion
	cr.Status.MongoDBMajorVersion = c.MongoDBMajorVersion
	cr.Status.MongoURI = c.MongoURI
	cr.Status.MongoURIUpdated = c.MongoURIUpdated
	cr.Status.MongoURIWithOptions = c.MongoURIWithOptions
	cr.Status.SrvAddress = c.SrvAddress
	cr.Status.DiskSizeGB = c.DiskSizeGB
	cr.Status.BackupEnabled = c.BackupEnabled
	cr.Status.ProviderBackupEnabled = c.ProviderBackupEnabled
	cr.Status.StateName = c.StateName
	cr.Status.ReplicationFactor = c.ReplicationFactor
	cr.Status.ReplicationSpec = c.ReplicationSpec
	cr.Status.NumShards = c.NumShards
	cr.Status.Paused = c.Paused
	cr.Status.AutoScaling = c.AutoScaling
	cr.Status.ProviderSettings = c.ProviderSettings
}

//addFinalizer will add this attribute to the Memcached CR
func (r *ReconcileMongoDBAtlasCluster) addFinalizer(reqLogger logr.Logger, cr *knappekv1alpha1.MongoDBAtlasCluster) error {
	if len(cr.GetFinalizers()) < 1 && cr.GetDeletionTimestamp() == nil {
		cr.SetFinalizers([]string{"finalizer.knappek.com"})

		// Update CR
		err := r.client.Update(context.TODO(), cr)
		if err != nil {
			reqLogger.Error(err, "Failed to update MongoDB Atlas Cluster with finalizer")
			return err
		}
	}
	return nil
}
