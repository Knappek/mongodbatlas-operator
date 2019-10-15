package mongodbatlascluster

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
	return &ReconcileMongoDBAtlasCluster{
		client:               mgr.GetClient(),
		scheme:               mgr.GetScheme(),
		atlasClient:          config.GetAtlasClient(),
		reconciliationConfig: config.GetReconcilitationConfig(),
	}
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
	client               client.Client
	scheme               *runtime.Scheme
	atlasClient          *ma.Client
	reconciliationConfig *config.ReconciliationConfig
}

// Reconcile reads that state of the cluster for a MongoDBAtlasCluster object and makes changes based on the state read
// and what is in the MongoDBAtlasCluster.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMongoDBAtlasCluster) Reconcile(request reconcile.Request) (reconcile.Result, error) {
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

	groupID := atlasProject.Status.ID
	// Define default logger
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "MongoDBAtlasCluster.Name", request.Name, "MongoDBAtlasCluster.GroupID", groupID)

	// Check if the MongoDBAtlasCluster CR was marked to be deleted
	isMongoDBAtlasClusterToBeDeleted := atlasCluster.GetDeletionTimestamp() != nil
	if isMongoDBAtlasClusterToBeDeleted {
		// check if Delete request has already been sent to the MongoDB Atlas API
		if atlasCluster.Status.StateName != "DELETING" && atlasCluster.Status.StateName != "DELETED" {
			err := deleteMongoDBAtlasCluster(reqLogger, r.atlasClient, atlasCluster)
			if err != nil {
				return reconcile.Result{}, err
			}
			atlasCluster.Status.StateName = "DELETING"
			err = r.client.Status().Update(context.TODO(), atlasCluster)
			if err != nil {
				return reconcile.Result{}, err
			}
			reqLogger.Info("Wait until Cluster has been deleted.")
			// Requeue after 20 seconds and check again for the status until CR can be deleted
			return reconcile.Result{RequeueAfter: time.Second * 20}, nil
		}

		// wait until cluster has been deleted successfully
		_, resp, err := r.atlasClient.Clusters.Get(groupID, atlasCluster.Name)
		if err != nil {
			if resp.StatusCode == 404 {
				reqLogger.Info("Cluster deleted.")
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
		// if err == nil, cluster still exists - Requeue after 20 seconds
		return reconcile.Result{RequeueAfter: time.Second * 20}, nil
	}

	// Create a new cluster
	isMongoDBAtlasClusterToBeCreated := reflect.DeepEqual(atlasCluster.Status, knappekv1alpha1.MongoDBAtlasClusterStatus{})
	if isMongoDBAtlasClusterToBeCreated {
		// check if Create request has already been sent to the MongoDB Atlas API
		if atlasCluster.Status.StateName != "CREATING" {
			err = createMongoDBAtlasCluster(reqLogger, r.atlasClient, atlasCluster, atlasProject)
			if err != nil {
				return reconcile.Result{}, err
			}
			atlasCluster.Status.StateName = "CREATING"
			err = r.client.Status().Update(context.TODO(), atlasCluster)
			if err != nil {
				return reconcile.Result{}, err
			}
			// Add finalizer for this CR
			if err := r.addFinalizer(reqLogger, atlasCluster); err != nil {
				return reconcile.Result{}, err
			}
			// Requeue to periodically reconcile the CR MongoDBAtlasCluster in order to recreate a manually deleted Atlas cluster
			return reconcile.Result{RequeueAfter: r.reconciliationConfig.Time}, nil
		}
	}

	// update existing cluster
	isMongoDBAtlasClusterToBeUpdated := knappekv1alpha1.IsMongoDBAtlasClusterToBeUpdated(atlasCluster.Spec.MongoDBAtlasClusterRequestBody, atlasCluster.Status.MongoDBAtlasClusterRequestBody)
	if isMongoDBAtlasClusterToBeUpdated {
		// check if Update request has already been sent to the MongoDB Atlas API
		if atlasCluster.Status.StateName != "UPDATING" {
			err = updateMongoDBAtlasCluster(reqLogger, r.atlasClient, atlasCluster, atlasProject)
			if err != nil {
				return reconcile.Result{}, err
			}
			atlasCluster.Status.StateName = "UPDATING"
			err = r.client.Status().Update(context.TODO(), atlasCluster)
			if err != nil {
				return reconcile.Result{}, err
			}
			// Requeue to periodically reconcile the CR MongoDBAtlasCluster in order to recreate a manually deleted Atlas cluster
			return reconcile.Result{RequeueAfter: r.reconciliationConfig.Time}, nil
		}
	}

	// if no Create/Update/Delete command apply, then fetch the status
	c, _, err := r.atlasClient.Clusters.Get(groupID, atlasCluster.Name)
	if err != nil {
		return reconcile.Result{}, err
	}
	err = updateCRStatus(reqLogger, atlasCluster, c)
	if err != nil {
		return reconcile.Result{}, err
	}
	// Update CR Status
	err = r.client.Status().Update(context.TODO(), atlasCluster)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Requeue to periodically reconcile the CR MongoDBAtlasCluster in order to recreate a manually deleted Atlas cluster
	return reconcile.Result{RequeueAfter: r.reconciliationConfig.Time}, nil
}

func createMongoDBAtlasCluster(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasCluster, ap *knappekv1alpha1.MongoDBAtlasProject) error {
	groupID := ap.Status.ID
	params := getClusterParams(cr)

	c, _, err := atlasClient.Clusters.Create(groupID, &params)
	if err != nil {
		return fmt.Errorf("Error creating Cluster %v: %s", cr.Name, err)
	}
	reqLogger.Info("Sent request to create Cluster.")
	return updateCRStatus(reqLogger, cr, c)
}

func updateMongoDBAtlasCluster(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasCluster, ap *knappekv1alpha1.MongoDBAtlasProject) error {
	groupID := ap.Status.ID
	params := getClusterParams(cr)
	c, _, err := atlasClient.Clusters.Update(groupID, cr.Name, &params)
	if err != nil {
		return fmt.Errorf("Error updating Cluster %v: %s", cr.Name, err)
	}
	reqLogger.Info("Sent request to update Cluster.")
	return updateCRStatus(reqLogger, cr, c)
}

func deleteMongoDBAtlasCluster(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasCluster) error {
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

func getClusterParams(cr *knappekv1alpha1.MongoDBAtlasCluster) ma.Cluster {
	return ma.Cluster{
		Name:                  cr.Name,
		MongoDBMajorVersion:   cr.Spec.MongoDBMajorVersion,
		DiskSizeGB:            cr.Spec.DiskSizeGB,
		BackupEnabled:         cr.Spec.BackupEnabled,
		ProviderBackupEnabled: cr.Spec.ProviderBackupEnabled,
		ReplicationSpec:       cr.Spec.ReplicationSpec,
		NumShards:             cr.Spec.NumShards,
		AutoScaling:           cr.Spec.AutoScaling,
		ProviderSettings:      cr.Spec.ProviderSettings,
	}
}

func updateCRStatus(reqLogger logr.Logger, cr *knappekv1alpha1.MongoDBAtlasCluster, c *ma.Cluster) error {
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
	cr.Status.DiskSizeGB = c.DiskSizeGB
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

func (r *ReconcileMongoDBAtlasCluster) addFinalizer(reqLogger logr.Logger, cr *knappekv1alpha1.MongoDBAtlasCluster) error {
	if len(cr.GetFinalizers()) < 1 && cr.GetDeletionTimestamp() == nil {
		cr.SetFinalizers([]string{"finalizer.knappek.com"})

		// Update CR
		err := r.client.Update(context.TODO(), cr)
		if err != nil {
			reqLogger.Error(err, "Failed to update Cluster with finalizer")
			return err
		}
	}
	return nil
}
