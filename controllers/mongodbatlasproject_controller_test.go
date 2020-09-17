package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"

	testutil "github.com/Knappek/mongodbatlas-operator/pkg/controller/test"
	mongodbatlasv1alpha1 "github.com/knappek/mongodbatlas-operator/api/v1alpha1"
	"github.com/knappek/mongodbatlas-operator/util"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var (
	projectName    = "unittest-project"
	groupID        = "5a0a1e7e0f2912c554080ae6"
	namespace      = "mongodbatlas"
	organizationID = "testOrgID"
	created        = "2016-07-14T14:19:33Z"
	clusterCount   = 0
)

func TestNonExistingMongoDBAtlasProjectCR(t *testing.T) {
	// Set the logger to development mode for verbose logs.
	logf.SetLogger(logf.ZapLogger(true))

	// A MongoDBAtlasProject resource with metadata and spec.
	mongodbatlasproject := &mongodbatlasv1alpha1.MongoDBAtlasProject{
		ObjectMeta: metav1.ObjectMeta{
			Name:      projectName,
			Namespace: namespace,
		},
		Spec: mongodbatlasv1alpha1.MongoDBAtlasProjectSpec{
			OrgID: organizationID,
		},
	}
	// Objects to track in the fake client.
	objs := []runtime.Object{
		mongodbatlasproject,
	}

	// Register operator types with the runtime scheme.
	s := scheme.Scheme
	s.AddKnownTypes(mongodbatlasv1alpha1.GroupVersion, mongodbatlasproject)

	// Create a fake k8s client to mock API calls.
	k8sClient := fake.NewFakeClient(objs...)
	// Create a fake atlas client to mock API calls.
	// atlasClient, server := test.NewAtlasFakeClient(t)
	httpClient, mux, server := testutil.Server()
	defer server.Close()
	// getByName: assert that there is no existing project
	mux.HandleFunc("/api/atlas/v1.0/groups/byName/"+projectName, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "GET", r)
		fmt.Fprintf(w, "")
	})
	// Post
	mux.HandleFunc("/api/atlas/v1.0/groups/", func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"clusterCount": `+strconv.Itoa(clusterCount)+`, "created":"`+created+`", "id":"`+groupID+`", "links":[], "name":"`+projectName+`", "orgId":"`+organizationID+`"}`)
	})
	atlasClient := ma.NewClient(httpClient)

	// Create a MongoDBAtlasProjectReconciler object with the scheme and fake client.
	r := &MongoDBAtlasProjectReconciler{
		Client:               k8sClient,
		Log:                  ctrl.Log,
		Scheme:               s,
		AtlasClient:          atlasClient,
		ReconciliationConfig: util.GetReconcilitationConfig(),
	}

	// Mock request with non-existing project
	req := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      "non-existent-project",
			Namespace: namespace,
		},
	}

	res, err := r.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}
	assert.Equal(t, false, res.Requeue)
	assert.Equal(t, ctrl.Result{}, res)
}

func TestCreateMongoDBAtlasProject(t *testing.T) {
	// Set the logger to development mode for verbose logs.
	logf.SetLogger(logf.ZapLogger(true))

	// A MongoDBAtlasProject resource with metadata and spec.
	mongodbatlasproject := &mongodbatlasv1alpha1.MongoDBAtlasProject{
		ObjectMeta: metav1.ObjectMeta{
			Name:      projectName,
			Namespace: namespace,
		},
		Spec: mongodbatlasv1alpha1.MongoDBAtlasProjectSpec{
			OrgID: organizationID,
		},
	}
	// Objects to track in the fake client.
	objs := []runtime.Object{
		mongodbatlasproject,
	}

	// Register operator types with the runtime scheme.
	s := scheme.Scheme
	s.AddKnownTypes(mongodbatlasv1alpha1.GroupVersion, mongodbatlasproject)

	// Create a fake k8s client to mock API calls.
	k8sClient := fake.NewFakeClient(objs...)
	// Create a fake atlas client to mock API calls.
	// atlasClient, server := test.NewAtlasFakeClient(t)
	httpClient, mux, server := testutil.Server()
	defer server.Close()
	// getByName: assert that there is no existing project
	mux.HandleFunc("/api/atlas/v1.0/groups/byName/"+projectName, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "GET", r)
		fmt.Fprintf(w, "")
	})
	// Post
	mux.HandleFunc("/api/atlas/v1.0/groups/", func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"clusterCount": `+strconv.Itoa(clusterCount)+`, "created":"`+created+`", "id":"`+groupID+`", "links":[], "name":"`+projectName+`", "orgId":"`+organizationID+`"}`)
	})
	atlasClient := ma.NewClient(httpClient)

	// Create a MongoDBAtlasProjectReconciler object with the scheme and fake client.
	r := &MongoDBAtlasProjectReconciler{
		Client:               k8sClient,
		Log:                  ctrl.Log,
		Scheme:               s,
		AtlasClient:          atlasClient,
		ReconciliationConfig: util.GetReconcilitationConfig(),
	}

	// Mock request to simulate Reconcile() being called on an event for a
	// watched resource .
	req := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      projectName,
			Namespace: namespace,
		},
	}
	res, err := r.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}
	assert.Equal(t, time.Second*120, res.RequeueAfter)

	// Check if the CR has been created and has the correct status.
	cr := &mongodbatlasv1alpha1.MongoDBAtlasProject{}
	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
	if err != nil {
		t.Fatalf("get MongoDBAtlasProject: (%v)", err)
	}
	assert.Equal(t, "finalizer.knappek.com", cr.ObjectMeta.GetFinalizers()[0], "The finalizer in the CR is not as expected")
	assert.Equal(t, organizationID, cr.Spec.OrgID, "The orgID in the Spec block is not as expected")
	assert.Equal(t, groupID, cr.Status.ID, "The id in the Status block is not as expected")
	assert.Equal(t, projectName, cr.Status.Name, "The name in the Status block is not as expected")
	assert.Equal(t, organizationID, cr.Status.OrgID, "The orgId in the Status block is not as expected")
	assert.Equal(t, created, cr.Status.Created, "The create in the Status block is not as expected")
	assert.Equal(t, clusterCount, cr.Status.ClusterCount, "The clusterCount in the Status block is not as expected")
}

func TestDeleteMongoDBAtlasProject(t *testing.T) {
	// Set the logger to development mode for verbose logs.
	logf.SetLogger(logf.ZapLogger(true))

	// A MongoDBAtlasProject resource with metadata and spec.
	mongodbatlasproject := &mongodbatlasv1alpha1.MongoDBAtlasProject{
		ObjectMeta: metav1.ObjectMeta{
			Name:              projectName,
			Namespace:         namespace,
			DeletionTimestamp: &metav1.Time{Time: time.Now()},
			Finalizers:        []string{"finalizer.knappek.com"},
		},
		Spec: mongodbatlasv1alpha1.MongoDBAtlasProjectSpec{
			OrgID: organizationID,
		},
		Status: mongodbatlasv1alpha1.MongoDBAtlasProjectStatus{
			ID:           groupID,
			OrgID:        organizationID,
			Name:         projectName,
			Created:      created,
			ClusterCount: clusterCount,
		},
	}

	// Objects to track in the fake client.
	objs := []runtime.Object{
		mongodbatlasproject,
	}

	// Register operator types with the runtime scheme.
	s := scheme.Scheme
	s.AddKnownTypes(mongodbatlasv1alpha1.GroupVersion, mongodbatlasproject)

	// Create a fake k8s client to mock API calls.
	k8sClient := fake.NewFakeClient(objs...)
	// Create a fake atlas client to mock API calls.
	// atlasClient, server := test.NewAtlasFakeClient(t)
	httpClient, mux, server := testutil.Server()
	defer server.Close()
	// getByName: assert that there is no existing project
	mux.HandleFunc("/api/atlas/v1.0/groups/byName/"+projectName, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"clusterCount": 0, "created":"`+created+`", "id":"`+groupID+`", "links":[], "name":"`+projectName+`", "orgId":"`+organizationID+`"}`)
	})
	// delete
	mux.HandleFunc("/api/atlas/v1.0/groups/"+groupID, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "DELETE", r)
		fmt.Fprintf(w, `{}`)
	})
	atlasClient := ma.NewClient(httpClient)

	// Create a MongoDBAtlasProjectReconciler object with the scheme and fake client.
	r := &MongoDBAtlasProjectReconciler{
		Client:               k8sClient,
		Log:                  ctrl.Log,
		Scheme:               s,
		AtlasClient:          atlasClient,
		ReconciliationConfig: util.GetReconcilitationConfig(),
	}

	// Mock request to simulate Reconcile() being called on an event for a
	// watched resource .
	req := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      projectName,
			Namespace: namespace,
		},
	}
	res, err := r.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}
	assert.Equal(t, time.Second*120, res.RequeueAfter)

	// Check if the CR has been created and has the correct status.
	cr := &mongodbatlasv1alpha1.MongoDBAtlasProject{}
	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
	assert.Nil(t, err)
	assert.Nil(t, cr.ObjectMeta.GetFinalizers())
}
