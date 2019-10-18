package mongodbatlasalertconfiguration

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"

	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
	"github.com/Knappek/mongodbatlas-operator/pkg/config"
	testutil "github.com/Knappek/mongodbatlas-operator/pkg/controller/test"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var (
	namespace      = "mongodbatlas"
	organizationID = "testOrgID"
	projectName    = "unittest-project"
	groupID        = "5a0a1e7e0f2912c554080ae6"
	resourceName   = "testalert"
	id = "57b76ddc96e8215c017ceafb"
	eventTypeName = "OUTSIDE_METRIC_THRESHOLD"
	enabled = true
	notifications = []ma.Notification{ma.Notification{
		TypeName: "GROUP",
		IntervalMin: 5,
		DelayMin: 0,
		SMSEnabled: false,
		EmailEnabled: true,
	}}
	metricThreshold = ma.MetricThreshold{
		MetricName: "QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED",
		Mode: "AVERAGE",
		Operator: "GREATER_THAN",
		Threshold: 500.0,
		Units: "RAW",
	}
)

func TestCreatemongodbatlasalertconfiguration(t *testing.T) {
	// Set the logger to development mode for verbose logs.
	logf.SetLogger(logf.ZapLogger(true))

	// A MongoDBAtlasProject resource with metadata and spec.
	mongodbatlasproject := testutil.CreateAtlasProject(projectName, groupID, namespace, organizationID)

	// A mongodbatlasalertconfiguration resource with metadata and spec.
	mongodbatlasalertconfiguration := &knappekv1alpha1.MongoDBAtlasAlertConfiguration{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resourceName,
			Namespace: namespace,
		},
		Spec: knappekv1alpha1.MongoDBAtlasAlertConfigurationSpec{
			ProjectName: projectName,
			MongoDBAtlasAlertConfigurationRequestBody: knappekv1alpha1.MongoDBAtlasAlertConfigurationRequestBody{
				EventTypeName: eventTypeName,
				Enabled: enabled,
				Notifications: notifications,
				MetricThreshold: metricThreshold,
			},
		},
	}

	// Objects to track in the fake client.
	objs := []runtime.Object{
		mongodbatlasalertconfiguration,
		mongodbatlasproject,
	}

	// Register operator types with the runtime scheme.
	s := scheme.Scheme
	s.AddKnownTypes(knappekv1alpha1.SchemeGroupVersion, mongodbatlasalertconfiguration, mongodbatlasproject)

	// Create a fake k8s client to mock API calls.
	k8sClient := fake.NewFakeClient(objs...)
	// Create a fake atlas client to mock API calls.
	// atlasClient, server := test.NewAtlasFakeClient(t)
	httpClient, mux, server := testutil.Server()
	defer server.Close()
	atlasClient := ma.NewClient(httpClient)

	// Post request for MongoDBAtlasAlertConfiguration
	mux.HandleFunc("/api/atlas/v1.0/groups/"+groupID+"/alertConfigs", func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"id" : "`+id+`",
			"enabled" : `+strconv.FormatBool(enabled)+`,
			"eventTypeName" : "`+eventTypeName+`",
			"groupId" : "`+groupID+`",
			"matchers" : [ ],
			"notifications" : [ {
			  "delayMin" : `+strconv.Itoa(notifications[0].DelayMin)+`,
			  "emailEnabled" : `+strconv.FormatBool(notifications[0].EmailEnabled)+`,
			  "intervalMin" : `+strconv.Itoa(notifications[0].IntervalMin)+`,
			  "smsEnabled" : `+strconv.FormatBool(notifications[0].SMSEnabled)+`,
			  "typeName" : "`+notifications[0].TypeName+`"
			} ],
			"metricThreshold" : {
			  "metricName" : "QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED",
			  "mode" : "`+metricThreshold.Mode+`",
			  "operator" : "`+metricThreshold.Operator+`",
			  "threshold" : `+fmt.Sprintf("%f", metricThreshold.Threshold)+`,
			  "units" : "`+metricThreshold.Units+`"
			}
		  }`)
	})

	// Create a ReconcileMongoDBAtlasAlertConfiguration object with the scheme and fake client.
	r := &ReconcileMongoDBAtlasAlertConfiguration{
		client:               k8sClient,
		scheme:               s,
		atlasClient:          atlasClient,
		reconciliationConfig: config.GetReconcilitationConfig(),
	}

	// Mock request to simulate Reconcile() being called on an event for a
	// watched resource .
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      resourceName,
			Namespace: namespace,
		},
	}
	res, err := r.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}
	assert.Equal(t, time.Second*120, res.RequeueAfter)

	// Check if the CR has been created and has the correct status.
	cr := &knappekv1alpha1.MongoDBAtlasAlertConfiguration{}
	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
	if err != nil {
		t.Fatalf("get mongodbatlasalertconfiguration: (%v)", err)
	}
	assert.Equal(t, groupID, cr.Status.GroupID)
	assert.Equal(t, eventTypeName, cr.Status.EventTypeName)
	assert.Equal(t, enabled, cr.Status.Enabled)
	assert.Equal(t, notifications, cr.Status.Notifications)
	assert.Equal(t, metricThreshold, cr.Status.MetricThreshold)
	assert.Empty(t, cr.Status.Matchers)
}

func TestDeletemongodbatlasalertconfiguration(t *testing.T) {
	// Set the logger to development mode for verbose logs.
	logf.SetLogger(logf.ZapLogger(true))

	// A MongoDBAtlasProject resource with metadata and spec.
	mongodbatlasproject := testutil.CreateAtlasProject(projectName, groupID, namespace, organizationID)

	// A mongodbatlasalertconfiguration resource with metadata and spec.
	mongodbatlasalertconfiguration := &knappekv1alpha1.MongoDBAtlasAlertConfiguration{
		ObjectMeta: metav1.ObjectMeta{
			Name:              resourceName,
			Namespace:         namespace,
			DeletionTimestamp: &metav1.Time{Time: time.Now()},
			Finalizers:        []string{"finalizer.knappek.com"},
		},
		Spec: knappekv1alpha1.MongoDBAtlasAlertConfigurationSpec{
			ProjectName: projectName,
			MongoDBAtlasAlertConfigurationRequestBody: knappekv1alpha1.MongoDBAtlasAlertConfigurationRequestBody{
				EventTypeName: eventTypeName,
				Enabled: enabled,
				Notifications: notifications,
				MetricThreshold: metricThreshold,
			},
		},
		Status: knappekv1alpha1.MongoDBAtlasAlertConfigurationStatus{
			ID: id,
			GroupID: groupID,
			MongoDBAtlasAlertConfigurationRequestBody: knappekv1alpha1.MongoDBAtlasAlertConfigurationRequestBody{
				EventTypeName: eventTypeName,
				Enabled: enabled,
				Notifications: notifications,
				MetricThreshold: metricThreshold,
			},
		},
	}

	// Objects to track in the fake client.
	objs := []runtime.Object{
		mongodbatlasalertconfiguration,
		mongodbatlasproject,
	}

	// Register operator types with the runtime scheme.
	s := scheme.Scheme
	s.AddKnownTypes(knappekv1alpha1.SchemeGroupVersion, mongodbatlasalertconfiguration, mongodbatlasproject)

	// Create a fake k8s client to mock API calls.
	k8sClient := fake.NewFakeClient(objs...)
	// Create a fake atlas client to mock API calls.
	// atlasClient, server := test.NewAtlasFakeClient(t)
	httpClient, mux, server := testutil.Server()
	defer server.Close()
	atlasClient := ma.NewClient(httpClient)

	// Delete
	mux.HandleFunc("/api/atlas/v1.0/groups/"+groupID+"/alertConfigs/"+id, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "DELETE", r)
		fmt.Fprintf(w, `{}`)
	})

	// Create a ReconcileMongoDBAtlasAlertConfiguration object with the scheme and fake client.
	r := &ReconcileMongoDBAtlasAlertConfiguration{
		client:               k8sClient,
		scheme:               s,
		atlasClient:          atlasClient,
		reconciliationConfig: config.GetReconcilitationConfig(),
	}

	// Mock request to simulate Reconcile() being called on an event for a
	// watched resource .
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      resourceName,
			Namespace: namespace,
		},
	}
	res, err := r.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}
	assert.Equal(t, time.Second*120, res.RequeueAfter)

	// Check if the CR has been deleted
	cr := &knappekv1alpha1.MongoDBAtlasAlertConfiguration{}
	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
	assert.Nil(t, err)
	assert.Nil(t, cr.ObjectMeta.GetFinalizers())
}

// func TestUpdatemongodbatlasalertconfiguration(t *testing.T) {
// 	// Set the logger to development mode for verbose logs.
// 	logf.SetLogger(logf.ZapLogger(true))

// 	// A MongoDBAtlasProject resource with metadata and spec.
// 	mongodbatlasproject := testutil.CreateAtlasProject(projectName, groupID, namespace, organizationID)

// 	// updates
// 	updatedMetricThreshold := ma.MetricThreshold{
// 		MetricName: "QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED",
// 		Mode: "AVERAGE",
// 		Operator: "LOWER_THAN",
// 		Threshold: 100.0,
// 		Units: "RAW",
// 	}

// 	// A mongodbatlasalertconfiguration resource with metadata and spec.
// 	mongodbatlasalertconfiguration := &knappekv1alpha1.MongoDBAtlasAlertConfiguration{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      resourceName,
// 			Namespace: namespace,
// 		},
// 		Spec: knappekv1alpha1.MongoDBAtlasAlertConfigurationSpec{
// 			ProjectName: projectName,
// 			MongoDBAtlasAlertConfigurationRequestBody: knappekv1alpha1.MongoDBAtlasAlertConfigurationRequestBody{
// 				EventTypeName: eventTypeName,
// 				Enabled: enabled,
// 				Notifications: notifications,
// 				MetricThreshold: updatedMetricThreshold,
// 			},
// 		},
// 		Status: knappekv1alpha1.MongoDBAtlasAlertConfigurationStatus{
// 			ID: id,
// 			GroupID: groupID,
// 			MongoDBAtlasAlertConfigurationRequestBody: knappekv1alpha1.MongoDBAtlasAlertConfigurationRequestBody{
// 				EventTypeName: eventTypeName,
// 				Enabled: enabled,
// 				Notifications: notifications,
// 				MetricThreshold: metricThreshold,
// 			},
// 		},
// 	}

// 	// Objects to track in the fake client.
// 	objs := []runtime.Object{
// 		mongodbatlasalertconfiguration,
// 		mongodbatlasproject,
// 	}

// 	// Register operator types with the runtime scheme.
// 	s := scheme.Scheme
// 	s.AddKnownTypes(knappekv1alpha1.SchemeGroupVersion, mongodbatlasalertconfiguration, mongodbatlasproject)

// 	// Create a fake k8s client to mock API calls.
// 	k8sClient := fake.NewFakeClient(objs...)
// 	// Create a fake atlas client to mock API calls.
// 	// atlasClient, server := test.NewAtlasFakeClient(t)
// 	httpClient, mux, server := testutil.Server()
// 	defer server.Close()
// 	atlasClient := ma.NewClient(httpClient)
// 	// Construct Update API call
// 	mux.HandleFunc("/api/atlas/v1.0/groups/<TODO>", func(w http.ResponseWriter, r *http.Request) {
// 		testutil.AssertMethod(t, "PUT", r)
// 		w.Header().Set("Content-Type", "application/json")
// 		fmt.Fprintf(w, `{
// 			"id" : "`+id+`",
// 			"enabled" : `+strconv.FormatBool(enabled)+`,
// 			"eventTypeName" : "`+eventTypeName+`",
// 			"groupId" : "`+groupID+`",
// 			"matchers" : [ ],
// 			"notifications" : [ {
// 				"delayMin" : `+strconv.Itoa(notifications[0].DelayMin)+`,
// 				"emailEnabled" : `+strconv.FormatBool(notifications[0].EmailEnabled)+`,
// 				"intervalMin" : `+strconv.Itoa(notifications[0].IntervalMin)+`,
// 				"smsEnabled" : `+strconv.FormatBool(notifications[0].SMSEnabled)+`,
// 				"typeName" : "`+notifications[0].TypeName+`"
// 			} ],
// 			"metricThreshold" : {
// 				"metricName" : "QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED",
// 				"mode" : "`+updatedMetricThreshold.Mode+`",
// 				"operator" : "`+updatedMetricThreshold.Operator+`",
// 				"threshold" : `+fmt.Sprintf("%f", updatedMetricThreshold.Threshold)+`,
// 				"units" : "`+updatedMetricThreshold.Units+`"
// 			}
// 		}`)
// 	})
// 	// Create a ReconcileMongoDBAtlasAlertConfiguration object with the scheme and fake client.
// 	r := &ReconcileMongoDBAtlasAlertConfiguration{
// 		client:               k8sClient,
// 		scheme:               s,
// 		atlasClient:          atlasClient,
// 		reconciliationConfig: config.GetReconcilitationConfig(),
// 	}

// 	// Mock request to simulate Reconcile() being called on an event for a
// 	// watched resource .
// 	req := reconcile.Request{
// 		NamespacedName: types.NamespacedName{
// 			Name:      resourceName,
// 			Namespace: namespace,
// 		},
// 	}
// 	res, err := r.Reconcile(req)
// 	if err != nil {
// 		t.Fatalf("reconcile: (%v)", err)
// 	}
// 	assert.Equal(t, time.Second*120, res.RequeueAfter)

// 	// Check if the CR has been created and has the correct status.
// 	cr := &knappekv1alpha1.MongoDBAtlasAlertConfiguration{}
// 	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
// 	if err != nil {
// 		t.Fatalf("get mongodbatlasalertconfiguration: (%v)", err)
// 	}
// 	assert.Equal(t, updatedMetricThreshold, cr.Status.MetricThreshold)
// }

// func TestNoUpdatemongodbatlasalertconfiguration(t *testing.T) {
// 	// Set the logger to development mode for verbose logs.
// 	logf.SetLogger(logf.ZapLogger(true))

// 	// A MongoDBAtlasProject resource with metadata and spec.
// 	mongodbatlasproject := testutil.CreateAtlasProject(projectName, groupID, namespace, organizationID)

// 	// A mongodbatlasalertconfiguration resource with metadata and spec. This Spec contains only the bare minimum, other values
// 	// will be filled with default values
// 	mongodbatlasalertconfiguration := &knappekv1alpha1.MongoDBAtlasAlertConfiguration{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      resourceName,
// 			Namespace: namespace,
// 		},
// 		Spec: knappekv1alpha1.MongoDBAtlasAlertConfigurationSpec{
// 			ProjectName: projectName,
// 			MongoDBAtlasAlertConfigurationRequestBody: knappekv1alpha1.MongoDBAtlasAlertConfigurationRequestBody{
// 				//
// 				// TODO: minimum requirements for the spec
// 				//
// 			},
// 		},
// 		Status: knappekv1alpha1.MongoDBAtlasAlertConfigurationStatus{
// 			//
// 			// TODO: some other read only values
// 			//
// 			MongoDBAtlasAlertConfigurationRequestBody: knappekv1alpha1.MongoDBAtlasAlertConfigurationRequestBody{
// 				//
// 				// TODO
// 				//
// 			},
// 		},
// 	}

// 	// Objects to track in the fake client.
// 	objs := []runtime.Object{
// 		mongodbatlasalertconfiguration,
// 		mongodbatlasproject,
// 	}

// 	// Register operator types with the runtime scheme.
// 	s := scheme.Scheme
// 	s.AddKnownTypes(knappekv1alpha1.SchemeGroupVersion, mongodbatlasalertconfiguration, mongodbatlasproject)

// 	// Create a fake k8s client to mock API calls.
// 	k8sClient := fake.NewFakeClient(objs...)
// 	// Create a fake atlas client to mock API calls.
// 	// atlasClient, server := test.NewAtlasFakeClient(t)
// 	httpClient, mux, server := testutil.Server()
// 	defer server.Close()
// 	atlasClient := ma.NewClient(httpClient)
// 	// Construct Update API call
// 	mux.HandleFunc("/api/atlas/v1.0/groups/<TODO>", func(w http.ResponseWriter, r *http.Request) {
// 		testutil.AssertMethod(t, "GET", r)
// 		w.Header().Set("Content-Type", "application/json")
// 		fmt.Fprintf(w, `{
// 			//
// 			// TODO
// 			//
// 		}`)
// 	})
// 	// Create a ReconcileMongoDBAtlasAlertConfiguration object with the scheme and fake client.
// 	r := &ReconcileMongoDBAtlasAlertConfiguration{
// 		client:               k8sClient,
// 		scheme:               s,
// 		atlasClient:          atlasClient,
// 		reconciliationConfig: config.GetReconcilitationConfig(),
// 	}

// 	// Mock request to simulate Reconcile() being called on an event for a
// 	// watched resource .
// 	req := reconcile.Request{
// 		NamespacedName: types.NamespacedName{
// 			Name:      resourceName,
// 			Namespace: namespace,
// 		},
// 	}
// 	res, err := r.Reconcile(req)
// 	if err != nil {
// 		t.Fatalf("reconcile: (%v)", err)
// 	}
// 	assert.Equal(t, time.Second*120, res.RequeueAfter)

// 	// Check if the CR has been created and has the correct status.
// 	cr := &knappekv1alpha1.MongoDBAtlasAlertConfiguration{}
// 	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
// 	if err != nil {
// 		t.Fatalf("get mongodbatlasalertconfiguration: (%v)", err)
// 	}
// 	//
// 	// TODO: assert that resource has not been updated
// 	//
// }
