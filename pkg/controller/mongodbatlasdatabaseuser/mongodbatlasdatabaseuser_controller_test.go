package mongodbatlasdatabaseuser

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"

	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
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
	namespace       = "mongodbatlas"
	organizationID  = "testOrgID"
	projectName     = "unittest-project"
	projectID       = "5a0a1e7e0f2912c554080ae6"
	resourceName    = "testuser"
	password        = "testpassword"
	databaseName    = "testdb"
	deleteAfterDate = "2100-01-01T00:00:00Z"
	roles           = []ma.Role{ma.Role{DatabaseName: databaseName, RoleName: "readWrite"}}
)

func TestCreatemongodbatlasdatabaseuser(t *testing.T) {
	// Set the logger to development mode for verbose logs.
	logf.SetLogger(logf.ZapLogger(true))

	// A MongoDBAtlasProject resource with metadata and spec.
	mongodbatlasproject := testutil.CreateAtlasProject(projectName, projectID, namespace, organizationID)

	// A mongodbatlasdatabaseuser resource with metadata and spec.
	mongodbatlasdatabaseuser := &knappekv1alpha1.MongoDBAtlasDatabaseUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resourceName,
			Namespace: namespace,
		},
		Spec: knappekv1alpha1.MongoDBAtlasDatabaseUserSpec{
			ProjectName: projectName,
			MongoDBAtlasDatabaseUserRequestBody: knappekv1alpha1.MongoDBAtlasDatabaseUserRequestBody{
				Password:        password,
				DeleteAfterDate: deleteAfterDate,
				DatabaseName:    "admin",
				Roles:           roles,
			},
		},
	}

	// Objects to track in the fake client.
	objs := []runtime.Object{
		mongodbatlasdatabaseuser,
		mongodbatlasproject,
	}

	// Register operator types with the runtime scheme.
	s := scheme.Scheme
	s.AddKnownTypes(knappekv1alpha1.SchemeGroupVersion, mongodbatlasdatabaseuser, mongodbatlasproject)

	// Create a fake k8s client to mock API calls.
	k8sClient := fake.NewFakeClient(objs...)
	// Create a fake atlas client to mock API calls.
	// atlasClient, server := test.NewAtlasFakeClient(t)
	httpClient, mux, server := testutil.Server()
	defer server.Close()
	atlasClient := ma.NewClient(httpClient)

	// Post request for MongoDBAtlasDatabaseUser
	mux.HandleFunc("/api/atlas/v1.0/groups/"+projectID+"/databaseUsers", func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"groupId":"`+projectID+`",
			"databaseName":"admin",
			"username":"`+resourceName+`",
			"roles":[{"databaseName":"`+roles[0].DatabaseName+`","roleName":"`+roles[0].RoleName+`"}]
		}`)
	})

	// Create a ReconcileMongoDBAtlasDatabaseUser object with the scheme and fake client.
	r := &ReconcileMongoDBAtlasDatabaseUser{client: k8sClient, scheme: s, atlasClient: atlasClient}

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
	assert.Equal(t, time.Second*30, res.RequeueAfter)

	// Check if the CR has been created and has the correct status.
	cr := &knappekv1alpha1.MongoDBAtlasDatabaseUser{}
	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
	if err != nil {
		t.Fatalf("get mongodbatlasdatabaseuser: (%v)", err)
	}
	assert.Equal(t, projectID, cr.Status.GroupID)
	assert.Equal(t, resourceName, cr.Status.Username)
	assert.Equal(t, "admin", cr.Status.DatabaseName)
	assert.Equal(t, roles, cr.Status.Roles)
	// assert.NotEmpty(t, cr.Status.Links)
}

// func TestDeletemongodbatlasdatabaseuser(t *testing.T) {
// 	// Set the logger to development mode for verbose logs.
// 	logf.SetLogger(logf.ZapLogger(true))

// 	// A MongoDBAtlasProject resource with metadata and spec.
// 	mongodbatlasproject := testutil.CreateAtlasProject(projectName, projectID, namespace, organizationID)

// 	// A mongodbatlasdatabaseuser resource with metadata and spec.
// 	mongodbatlasdatabaseuser := &knappekv1alpha1.MongoDBAtlasDatabaseUser{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:              resourceName,
// 			Namespace:         namespace,
// 			DeletionTimestamp: &metav1.Time{Time: time.Now()},
// 			Finalizers:        []string{"finalizer.knappek.com"},
// 		},
// 		Spec: knappekv1alpha1.MongoDBAtlasDatabaseUserSpec{
// 			ProjectName:                         projectName,
// 			MongoDBAtlasDatabaseUserRequestBody: knappekv1alpha1.MongoDBAtlasDatabaseUserRequestBody{
// 				//
// 				// TODO
// 				//
// 			},
// 		},
// 		Status: knappekv1alpha1.MongoDBAtlasDatabaseUserStatus{
// 			//
// 			// TODO: some other read only values
// 			//
// 			MongoDBAtlasDatabaseUserRequestBody: knappekv1alpha1.MongoDBAtlasDatabaseUserRequestBody{
// 				//
// 				// TODO
// 				//
// 			},
// 		},
// 	}

// 	// Objects to track in the fake client.
// 	objs := []runtime.Object{
// 		mongodbatlasdatabaseuser,
// 		mongodbatlasproject,
// 	}

// 	// Register operator types with the runtime scheme.
// 	s := scheme.Scheme
// 	s.AddKnownTypes(knappekv1alpha1.SchemeGroupVersion, mongodbatlasdatabaseuser, mongodbatlasproject)

// 	// Create a fake k8s client to mock API calls.
// 	k8sClient := fake.NewFakeClient(objs...)
// 	// Create a fake atlas client to mock API calls.
// 	// atlasClient, server := test.NewAtlasFakeClient(t)
// 	httpClient, mux, server := testutil.Server()
// 	defer server.Close()
// 	atlasClient := ma.NewClient(httpClient)

// 	// Delete
// 	mux.HandleFunc("/api/atlas/v1.0/groups/<TODO>", func(w http.ResponseWriter, r *http.Request) {
// 		testutil.AssertMethod(t, "DELETE", r)
// 		fmt.Fprintf(w, `{}`)
// 	})

// 	// Create a ReconcileMongoDBAtlasDatabaseUser object with the scheme and fake client.
// 	r := &ReconcileMongoDBAtlasDatabaseUser{client: k8sClient, scheme: s, atlasClient: atlasClient}

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
// 	assert.Equal(t, time.Second*20, res.RequeueAfter)

// 	// Check if the CR has been updated and has the correct status.
// 	cr := &knappekv1alpha1.MongoDBAtlasDatabaseUser{}
// 	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
// 	if err != nil {
// 		t.Fatalf("get mongodbatlasdatabaseuser: (%v)", err)
// 	}

// 	httpClient2, mux2, server2 := testutil.Server()
// 	defer server2.Close()
// 	atlasClient2 := ma.NewClient(httpClient2)
// 	// GET: Simulate a new reconcile where cluster has been deleted successfully
// 	mux2.HandleFunc("/api/atlas/v1.0/groups/<TODO>", func(w http.ResponseWriter, r *http.Request) {
// 		testutil.AssertMethod(t, "GET", r)
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 	})

// 	// Create a ReconcileMongoDBAtlasDatabaseUser object with the scheme and fake client.
// 	r2 := &ReconcileMongoDBAtlasDatabaseUser{client: k8sClient, scheme: s, atlasClient: atlasClient2}

// 	res2, err := r2.Reconcile(req)
// 	if err != nil {
// 		t.Fatalf("reconcile: (%v)", err)
// 	}
// 	assert.Equal(t, reconcile.Result{}, res2)
// 	cr = &knappekv1alpha1.MongoDBAtlasDatabaseUser{}
// 	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
// 	if err != nil {
// 		t.Fatalf("get mongodbatlasdatabaseuser: (%v)", err)
// 	}
// 	// verify that Finalizer has been removed
// 	assert.Nil(t, cr.ObjectMeta.GetFinalizers())
// }

// func TestUpdatemongodbatlasdatabaseuser(t *testing.T) {
// 	// Set the logger to development mode for verbose logs.
// 	logf.SetLogger(logf.ZapLogger(true))

// 	// A MongoDBAtlasProject resource with metadata and spec.
// 	mongodbatlasproject := &knappekv1alpha1.MongoDBAtlasProject{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      projectName,
// 			Namespace: namespace,
// 		},
// 		Spec: knappekv1alpha1.MongoDBAtlasProjectSpec{
// 			OrgID: organizationID,
// 		},
// 		Status: knappekv1alpha1.MongoDBAtlasProjectStatus{
// 			ID:           projectID,
// 			Name:         projectName,
// 			OrgID:        organizationID,
// 			Created:      "2016-07-14T14:19:33Z",
// 			ClusterCount: 1,
// 		},
// 	}

// 	// updates
// 	//
// 	// TODO: some updates
// 	//

// 	// A mongodbatlasdatabaseuser resource with metadata and spec.
// 	mongodbatlasdatabaseuser := &knappekv1alpha1.MongoDBAtlasDatabaseUser{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      resourceName,
// 			Namespace: namespace,
// 		},
// 		Spec: knappekv1alpha1.MongoDBAtlasDatabaseUserSpec{
// 			ProjectName:                         projectName,
// 			MongoDBAtlasDatabaseUserRequestBody: knappekv1alpha1.MongoDBAtlasDatabaseUserRequestBody{
// 				//
// 				// TODO
// 				//
// 			},
// 		},
// 		Status: knappekv1alpha1.MongoDBAtlasDatabaseUserStatus{
// 			//
// 			// TODO: some other read only values
// 			//
// 			MongoDBAtlasDatabaseUserRequestBody: knappekv1alpha1.MongoDBAtlasDatabaseUserRequestBody{
// 				//
// 				// TODO
// 				//
// 			},
// 		},
// 	}

// 	// Objects to track in the fake client.
// 	objs := []runtime.Object{
// 		mongodbatlasdatabaseuser,
// 		mongodbatlasproject,
// 	}

// 	// Register operator types with the runtime scheme.
// 	s := scheme.Scheme
// 	s.AddKnownTypes(knappekv1alpha1.SchemeGroupVersion, mongodbatlasdatabaseuser, mongodbatlasproject)

// 	// Create a fake k8s client to mock API calls.
// 	k8sClient := fake.NewFakeClient(objs...)
// 	// Create a fake atlas client to mock API calls.
// 	// atlasClient, server := test.NewAtlasFakeClient(t)
// 	httpClient, mux, server := testutil.Server()
// 	defer server.Close()
// 	atlasClient := ma.NewClient(httpClient)
// 	// Construct Update API call
// 	mux.HandleFunc("/api/atlas/v1.0/groups/<TODO>", func(w http.ResponseWriter, r *http.Request) {
// 		testutil.AssertMethod(t, "PATCH", r)
// 		w.Header().Set("Content-Type", "application/json")
// 		fmt.Fprintf(w, `{
// 			//
// 			// TODO
// 			//
// 		}`)
// 	})
// 	// Create a ReconcileMongoDBAtlasDatabaseUser object with the scheme and fake client.
// 	r := &ReconcileMongoDBAtlasDatabaseUser{client: k8sClient, scheme: s, atlasClient: atlasClient}

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
// 	assert.Equal(t, time.Second*30, res.RequeueAfter)

// 	// Check if the CR has been created and has the correct status.
// 	cr := &knappekv1alpha1.MongoDBAtlasDatabaseUser{}
// 	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
// 	if err != nil {
// 		t.Fatalf("get mongodbatlasdatabaseuser: (%v)", err)
// 	}
// 	//
// 	// TODO: some assertions
// 	//
// }

// func TestNoUpdatemongodbatlasdatabaseuser(t *testing.T) {
// 	// Set the logger to development mode for verbose logs.
// 	logf.SetLogger(logf.ZapLogger(true))

// 	// A MongoDBAtlasProject resource with metadata and spec.
// 	mongodbatlasproject := &knappekv1alpha1.MongoDBAtlasProject{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      projectName,
// 			Namespace: namespace,
// 		},
// 		Spec: knappekv1alpha1.MongoDBAtlasProjectSpec{
// 			OrgID: organizationID,
// 		},
// 		Status: knappekv1alpha1.MongoDBAtlasProjectStatus{
// 			ID:           projectID,
// 			Name:         projectName,
// 			OrgID:        organizationID,
// 			Created:      "2016-07-14T14:19:33Z",
// 			ClusterCount: 1,
// 		},
// 	}

// 	// A mongodbatlasdatabaseuser resource with metadata and spec. This Spec contains only the bare minimum, other values
// 	// will be filled with default values
// 	mongodbatlasdatabaseuser := &knappekv1alpha1.MongoDBAtlasDatabaseUser{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      resourceName,
// 			Namespace: namespace,
// 		},
// 		Spec: knappekv1alpha1.MongoDBAtlasDatabaseUserSpec{
// 			ProjectName:                         projectName,
// 			MongoDBAtlasDatabaseUserRequestBody: knappekv1alpha1.MongoDBAtlasDatabaseUserRequestBody{
// 				//
// 				// TODO: minimum requirements for the spec
// 				//
// 			},
// 		},
// 		Status: knappekv1alpha1.MongoDBAtlasDatabaseUserStatus{
// 			//
// 			// TODO: some other read only values
// 			//
// 			MongoDBAtlasDatabaseUserRequestBody: knappekv1alpha1.MongoDBAtlasDatabaseUserRequestBody{
// 				//
// 				// TODO
// 				//
// 			},
// 		},
// 	}

// 	// Objects to track in the fake client.
// 	objs := []runtime.Object{
// 		mongodbatlasdatabaseuser,
// 		mongodbatlasproject,
// 	}

// 	// Register operator types with the runtime scheme.
// 	s := scheme.Scheme
// 	s.AddKnownTypes(knappekv1alpha1.SchemeGroupVersion, mongodbatlasdatabaseuser, mongodbatlasproject)

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
// 	// Create a ReconcileMongoDBAtlasDatabaseUser object with the scheme and fake client.
// 	r := &ReconcileMongoDBAtlasDatabaseUser{client: k8sClient, scheme: s, atlasClient: atlasClient}

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
// 	assert.Equal(t, time.Second*30, res.RequeueAfter)

// 	// Check if the CR has been created and has the correct status.
// 	cr := &knappekv1alpha1.MongoDBAtlasDatabaseUser{}
// 	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
// 	if err != nil {
// 		t.Fatalf("get mongodbatlasdatabaseuser: (%v)", err)
// 	}
// 	//
// 	// TODO: assert that resource has not been updated
// 	//
// }
