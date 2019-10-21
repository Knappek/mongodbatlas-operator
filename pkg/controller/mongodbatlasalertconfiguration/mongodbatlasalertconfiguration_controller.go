package mongodbatlasalertconfiguration

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

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

var log = logf.Log.WithName("controller_mongodbatlasalertconfiguration")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new MongoDBAtlasAlertConfiguration Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMongoDBAtlasAlertConfiguration{
		client:               mgr.GetClient(),
		scheme:               mgr.GetScheme(),
		atlasClient:          config.GetAtlasClient(),
		reconciliationConfig: config.GetReconcilitationConfig(),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("mongodbatlasalertconfiguration-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource MongoDBAtlasAlertConfiguration
	err = c.Watch(&source.Kind{Type: &knappekv1alpha1.MongoDBAtlasAlertConfiguration{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileMongoDBAtlasAlertConfiguration{}

// ReconcileMongoDBAtlasAlertConfiguration reconciles a MongoDBAtlasAlertConfiguration object
type ReconcileMongoDBAtlasAlertConfiguration struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client               client.Client
	scheme               *runtime.Scheme
	atlasClient          *ma.Client
	reconciliationConfig *config.ReconciliationConfig
}

// Reconcile reads that state of the MongoDBAtlasAlertConfiguration object and makes changes based on the state read
// and what is in the MongoDBAtlasAlertConfiguration.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMongoDBAtlasAlertConfiguration) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// Fetch the MongoDBAtlasAlertConfiguration atlasAlertConfiguration
	atlasAlertConfiguration := &knappekv1alpha1.MongoDBAtlasAlertConfiguration{}
	err := r.client.Get(context.TODO(), request.NamespacedName, atlasAlertConfiguration)
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

	projectName := atlasAlertConfiguration.Spec.ProjectName
	atlasProject := &knappekv1alpha1.MongoDBAtlasProject{}
	atlasProjectNamespacedName := types.NamespacedName{
		Name:      projectName,
		Namespace: atlasAlertConfiguration.Namespace,
	}

	err = r.client.Get(context.TODO(), atlasProjectNamespacedName, atlasProject)
	if err != nil {
		return reconcile.Result{}, err
	}

	groupID := atlasProject.Status.ID
	// Define default logger
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "MongoDBAtlasAlertConfiguration.Name", request.Name, "MongoDBAtlasAlertConfiguration.GroupID", groupID)

	// Check if the MongoDBAtlasAlertConfiguration CR was marked to be deleted
	isMongoDBAtlasAlertConfigurationToBeDeleted := atlasAlertConfiguration.GetDeletionTimestamp() != nil
	if isMongoDBAtlasAlertConfigurationToBeDeleted {
		err := deleteMongoDBAtlasAlertConfiguration(reqLogger, r.atlasClient, atlasAlertConfiguration)
		if err != nil {
			return reconcile.Result{}, err
		}
		err = r.client.Update(context.TODO(), atlasAlertConfiguration)
		if err != nil {
			return reconcile.Result{}, err
		}
		// Requeue to periodically reconcile the CR MongoDBAtlasAlertConfiguration in order to recreate a manually deleted Atlas AlertConfiguration
		return reconcile.Result{RequeueAfter: r.reconciliationConfig.Time}, nil
	}

	// Create a new MongoDBAtlasAlertConfiguration
	isMongoDBAtlasAlertConfigurationToBeCreated := reflect.DeepEqual(atlasAlertConfiguration.Status, knappekv1alpha1.MongoDBAtlasAlertConfigurationStatus{})
	if isMongoDBAtlasAlertConfigurationToBeCreated {
		err = createMongoDBAtlasAlertConfiguration(reqLogger, r.atlasClient, atlasAlertConfiguration, atlasProject)
		if err != nil {
			return reconcile.Result{}, err
		}
		err = r.client.Status().Update(context.TODO(), atlasAlertConfiguration)
		if err != nil {
			return reconcile.Result{}, err
		}
		// Add finalizer for this CR
		if err := r.addFinalizer(reqLogger, atlasAlertConfiguration); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{RequeueAfter: r.reconciliationConfig.Time}, nil
	}

	// update existing MongoDBAtlasAlertConfiguration
	isMongoDBAtlasAlertConfigurationToBeUpdated := knappekv1alpha1.IsMongoDBAtlasAlertConfigurationToBeUpdated(atlasAlertConfiguration.Spec.MongoDBAtlasAlertConfigurationRequestBody, atlasAlertConfiguration.Status.MongoDBAtlasAlertConfigurationRequestBody)
	if isMongoDBAtlasAlertConfigurationToBeUpdated {
		err = updateMongoDBAtlasAlertConfiguration(reqLogger, r.atlasClient, atlasAlertConfiguration, atlasProject)
		if err != nil {
			return reconcile.Result{}, err
		}
		err = r.client.Status().Update(context.TODO(), atlasAlertConfiguration)
		if err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{RequeueAfter: r.reconciliationConfig.Time}, nil
	}

	// if no Create/Update/Delete command apply, then fetch the status
	err = getMongoDBAtlasAlertConfiguration(reqLogger, r.atlasClient, atlasAlertConfiguration)
	if err != nil {
		return reconcile.Result{}, err
	}
	err = r.client.Status().Update(context.TODO(), atlasAlertConfiguration)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Requeue to periodically reconcile the CR MongoDBAtlasAlertConfiguration in order to recreate a manually deleted Atlas AlertConfiguration
	return reconcile.Result{RequeueAfter: r.reconciliationConfig.Time}, nil
}

func createMongoDBAtlasAlertConfiguration(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasAlertConfiguration, ap *knappekv1alpha1.MongoDBAtlasProject) error {
	groupID := ap.Status.ID
	name := cr.Name
	params := getAlertConfigurationParams(cr)
	c, resp, err := atlasClient.AlertConfigurations.Create(groupID, &params)
	if err != nil {
		return fmt.Errorf("Error creating AlertConfiguration %v: %s", name, err)
	}
	if resp.StatusCode == http.StatusOK {
		reqLogger.Info("AlertConfiguration created.")
		return updateCRStatus(reqLogger, cr, c)
	}
	return fmt.Errorf("(%v) Error creating AlertConfiguration %s: %s", resp.StatusCode, name, err)
}

func updateMongoDBAtlasAlertConfiguration(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasAlertConfiguration, ap *knappekv1alpha1.MongoDBAtlasProject) error {
	groupID := ap.Status.ID
	id := cr.Status.ID
	name := cr.Name
	params := getAlertConfigurationParams(cr)
	c, resp, err := atlasClient.AlertConfigurations.Update(groupID, id, &params)
	if err != nil {
		return fmt.Errorf("(%v) Error updating AlertConfiguration %v: %s", resp.StatusCode, name, err)
	}
	if resp.StatusCode == http.StatusOK {
		reqLogger.Info("AlertConfiguration updated.")
		return updateCRStatus(reqLogger, cr, c)
	}
	return fmt.Errorf("(%v) Error updating AlertConfiguration %s: %s", resp.StatusCode, name, err)
}

func deleteMongoDBAtlasAlertConfiguration(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasAlertConfiguration) error {
	groupID := cr.Status.GroupID
	id := cr.Status.ID
	name := cr.Name
	// cluster exists and can be deleted
	resp, err := atlasClient.AlertConfigurations.Delete(groupID, id)
	if err != nil {
		if resp.StatusCode == http.StatusNotFound {
			reqLogger.Info("AlertConfiguration does not exist in Atlas. Deleting CR.")
			// Update finalizer to allow delete CR
			cr.SetFinalizers(nil)
			// CR can be deleted - Requeue
			return nil
		}
		return fmt.Errorf("(%v) Error deleting AlertConfiguration %s: %s", resp.StatusCode, name, err)
	}
	if resp.StatusCode == http.StatusOK {
		// Update finalizer to allow delete CR
		cr.SetFinalizers(nil)
		reqLogger.Info("AlertConfiguration deleted.")
		return nil
	}
	return fmt.Errorf("(%v) Error deleting AlertConfiguration %s: %s", resp.StatusCode, name, err)
}

func getMongoDBAtlasAlertConfiguration(reqLogger logr.Logger, atlasClient *ma.Client, cr *knappekv1alpha1.MongoDBAtlasAlertConfiguration) error {
	groupID := cr.Status.GroupID
	id := cr.Status.ID
	name := cr.Name
	c, resp, err := atlasClient.AlertConfigurations.Get(groupID, id)
	if err != nil {
		return fmt.Errorf("(%v) Error fetching AlertConfiguration information %s: %s", resp.StatusCode, name, err)
	}
	err = updateCRStatus(reqLogger, cr, c)
	if err != nil {
		return fmt.Errorf("Error updating AlertConfiguration CR Status: %s", err)
	}
	return nil
}

func getAlertConfigurationParams(cr *knappekv1alpha1.MongoDBAtlasAlertConfiguration) ma.AlertConfiguration {
	return ma.AlertConfiguration{
		EventTypeName:   cr.Spec.EventTypeName,
		Enabled:         cr.Spec.Enabled,
		Notifications:   cr.Spec.Notifications,
		MetricThreshold: cr.Spec.MetricThreshold,
		Matchers:        cr.Spec.Matchers,
	}
}

func updateCRStatus(reqLogger logr.Logger, cr *knappekv1alpha1.MongoDBAtlasAlertConfiguration, c *ma.AlertConfiguration) error {
	// update status field in CR
	cr.Status.ID = c.ID
	cr.Status.GroupID = c.GroupID
	cr.Status.EventTypeName = c.EventTypeName
	cr.Status.Enabled = c.Enabled
	cr.Status.Notifications = c.Notifications
	cr.Status.MetricThreshold = c.MetricThreshold
	cr.Status.Matchers = c.Matchers
	return nil
}

func (r *ReconcileMongoDBAtlasAlertConfiguration) addFinalizer(reqLogger logr.Logger, cr *knappekv1alpha1.MongoDBAtlasAlertConfiguration) error {
	if len(cr.GetFinalizers()) < 1 && cr.GetDeletionTimestamp() == nil {
		cr.SetFinalizers([]string{"finalizer.knappek.com"})

		// Update CR
		err := r.client.Update(context.TODO(), cr)
		if err != nil {
			reqLogger.Error(err, "Failed to update AlertConfiguration with finalizer")
			return err
		}
	}
	return nil
}
