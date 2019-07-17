package mongodbatlascluster

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
	namespace             = "mongodbatlas"
	organizationID        = "testOrgID"
	projectName           = "unittest-project"
	projectID             = "5a0a1e7e0f2912c554080ae6"
	clusterName           = "unittest-cluster"
	clusterID             = "testClusterId"
	mongoDBVersion        = "3.4"
	mongoDBMajorVersion   = "3.4"
	diskSizeGB            = 10.5
	backupEnabled         = true
	providerBackupEnabled = false
	replicationSpec       = map[string]ma.ReplicationSpec{
		"US_EAST_1": ma.ReplicationSpec{
			Priority:       7,
			ElectableNodes: 2,
			ReadOnlyNodes:  1,
			AnalyticsNodes: 1,
		},
	}
	numShards   = 1
	paused      = false
	autoscaling = ma.AutoScaling{
		DiskGBEnabled: false,
	}
	providerSettings = ma.ProviderSettings{
		ProviderName:        "AWS",
		RegionName:          "US_EAST_1",
		InstanceSizeName:    "M10",
		EncryptEBSVolume:    true,
		BackingProviderName: "",
	}
)

func TestCreateMongoDBAtlasCluster(t *testing.T) {
	// Set the logger to development mode for verbose logs.
	logf.SetLogger(logf.ZapLogger(true))

	// A MongoDBAtlasProject resource with metadata and spec.
	mongodbatlasproject := &knappekv1alpha1.MongoDBAtlasProject{
		ObjectMeta: metav1.ObjectMeta{
			Name:      projectName,
			Namespace: namespace,
		},
		Spec: knappekv1alpha1.MongoDBAtlasProjectSpec{
			OrgID: organizationID,
		},
		Status: knappekv1alpha1.MongoDBAtlasProjectStatus{
			ID:           projectID,
			Name:         projectName,
			OrgID:        organizationID,
			Created:      "2016-07-14T14:19:33Z",
			ClusterCount: 0,
		},
	}

	// A MongoDBAtlasCluster resource with metadata and spec.
	mongodbatlascluster := &knappekv1alpha1.MongoDBAtlasCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterName,
			Namespace: namespace,
		},
		Spec: knappekv1alpha1.MongoDBAtlasClusterSpec{
			ProjectName:           projectName,
			ProviderSettings:      providerSettings,
			MongoDBVersion:        mongoDBVersion,
			MongoDBMajorVersion:   mongoDBMajorVersion,
			DiskSizeGB:            diskSizeGB,
			NumShards:             numShards,
			AutoScaling:           autoscaling,
			BackupEnabled:         backupEnabled,
			ProviderBackupEnabled: providerBackupEnabled,
			ReplicationSpec:       replicationSpec,
		},
	}

	// Objects to track in the fake client.
	objs := []runtime.Object{
		mongodbatlascluster,
		mongodbatlasproject,
	}

	// Register operator types with the runtime scheme.
	s := scheme.Scheme
	s.AddKnownTypes(knappekv1alpha1.SchemeGroupVersion, mongodbatlascluster, mongodbatlasproject)

	// Create a fake k8s client to mock API calls.
	k8sClient := fake.NewFakeClient(objs...)
	// Create a fake atlas client to mock API calls.
	// atlasClient, server := test.NewAtlasFakeClient(t)
	httpClient, mux, server := testutil.Server()
	defer server.Close()
	atlasClient := ma.NewClient(httpClient)

	// Post
	mux.HandleFunc("/api/atlas/v1.0/groups/"+projectID+"/clusters", func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"name":                  clusterName,
			"mongoDBVersion":        mongoDBVersion,
			"mongoDBMajorVersion":   mongoDBMajorVersion,
			"groupId":               projectID,
			"numShards":             float64(numShards),
			"backupEnabled":         backupEnabled,
			"providerBackupEnabled": providerBackupEnabled,
			"paused":                paused,
			"diskSizeGB":            diskSizeGB,
			"autoScaling": map[string]interface{}{
				"diskGBEnabled": autoscaling.DiskGBEnabled,
			},
			"providerSettings": map[string]interface{}{
				"providerName":     providerSettings.ProviderName,
				"regionName":       providerSettings.RegionName,
				"instanceSizeName": providerSettings.InstanceSizeName,
				"encryptEBSVolume": providerSettings.EncryptEBSVolume,
			},
			"replicationSpec": map[string]interface{}{
				"US_EAST_1": map[string]interface{}{
					"priority":       float64(7),
					"electableNodes": float64(2),
					"readOnlyNodes":  float64(1),
					"analyticsNodes": float64(1),
				},
			},
		}
		testutil.AssertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{
			"id": "`+clusterID+`",
			"groupId": "`+projectID+`",
			"name":"`+clusterName+`",
			"mongoDBMajorVersion":"`+mongoDBMajorVersion+`",
			"mongoDBVersion":"`+mongoDBVersion+`",
			"numShards": `+strconv.Itoa(numShards)+`,
			"backupEnabled":`+strconv.FormatBool(backupEnabled)+`,
			"providerBackupEnabled":`+strconv.FormatBool(providerBackupEnabled)+`,
			"diskSizeGB":`+strconv.FormatFloat(diskSizeGB, 'f', 6, 64)+`,
			"paused":`+strconv.FormatBool(paused)+`,
			"stateName": "CREATING",
			"autoScaling":{
				"diskGBEnabled":`+strconv.FormatBool(autoscaling.DiskGBEnabled)+`
			},
			"providerSettings":{
				"providerName":"`+providerSettings.ProviderName+`",
				"regionName":"`+providerSettings.RegionName+`",
				"instanceSizeName":"`+providerSettings.InstanceSizeName+`",
				"encryptEBSVolume": `+strconv.FormatBool(providerSettings.EncryptEBSVolume)+`
			},
			"replicationSpec":{
				"US_EAST_1":{
					"priority":7,
					"electableNodes":2,
					"readOnlyNodes":1,
					"analyticsNodes":1
				}
			}
		}`)
	})

	// Create a ReconcileMongoDBAtlasCluster object with the scheme and fake client.
	r := &ReconcileMongoDBAtlasCluster{client: k8sClient, scheme: s, atlasClient: atlasClient}

	// Mock request to simulate Reconcile() being called on an event for a
	// watched resource .
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      clusterName,
			Namespace: namespace,
		},
	}
	res, err := r.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}
	assert.Equal(t, time.Second*30, res.RequeueAfter)

	// GET: Simulate a new reconcile where stateName changed from CREATING to IDLE
	mux.HandleFunc("/api/atlas/v1.0/groups/"+projectID+"/clusters/"+clusterName, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "GET", r)
		fmt.Fprintf(w, `{
			"id": "`+clusterID+`",
			"groupId": "`+projectID+`",
			"name":"`+clusterName+`",
			"mongoDBMajorVersion":"`+mongoDBMajorVersion+`",
			"mongoDBVersion":"`+mongoDBVersion+`",
			"numShards": `+strconv.Itoa(numShards)+`,
			"backupEnabled":`+strconv.FormatBool(backupEnabled)+`,
			"providerBackupEnabled":`+strconv.FormatBool(providerBackupEnabled)+`,
			"diskSizeGB":`+strconv.FormatFloat(diskSizeGB, 'f', 6, 64)+`,
			"paused":`+strconv.FormatBool(paused)+`,
			"stateName": "IDLE",
			"autoScaling":{
				"diskGBEnabled":`+strconv.FormatBool(autoscaling.DiskGBEnabled)+`
			},
			"providerSettings":{
				"providerName":"`+providerSettings.ProviderName+`",
				"regionName":"`+providerSettings.RegionName+`",
				"instanceSizeName":"`+providerSettings.InstanceSizeName+`",
				"encryptEBSVolume": `+strconv.FormatBool(providerSettings.EncryptEBSVolume)+`
			},
			"replicationSpec":{
				"US_EAST_1":{
					"priority":7,
					"electableNodes":2,
					"readOnlyNodes":1,
					"analyticsNodes":1
				}
			}
		}`)
	})

	// Simulate a new reconcile where stateName changed from CREATING to IDLE
	res, err = r.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}

	// Check if the CR has been created and has the correct status.
	cr := &knappekv1alpha1.MongoDBAtlasCluster{}
	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
	if err != nil {
		t.Fatalf("get MongoDBAtlasCluster: (%v)", err)
	}

	assert.Equal(t, "finalizer.knappek.com", cr.ObjectMeta.GetFinalizers()[0], "Finalizer not as expected")
	assert.Equal(t, clusterID, cr.Status.ID, "clusterID not as expected")
	assert.Equal(t, clusterName, cr.Status.Name, "clusterName not as expected")
	assert.Equal(t, projectID, cr.Status.GroupID, "projectID not as expected")
	assert.Equal(t, mongoDBVersion, cr.Status.MongoDBVersion, "mongoDBVersion not as expected")
	assert.Equal(t, mongoDBMajorVersion, cr.Status.MongoDBMajorVersion, "mongoDBMajorVersion not as expected")
	assert.Equal(t, diskSizeGB, cr.Status.DiskSizeGB, "diskSizeGB not as expected")
	assert.Equal(t, backupEnabled, cr.Status.BackupEnabled, "backupEnabled not as expected")
	assert.Equal(t, providerBackupEnabled, cr.Status.ProviderBackupEnabled, "providerBackupEnabled not as expected")
	assert.Equal(t, "IDLE", cr.Status.StateName, "stateName not as expected")
	assert.Equal(t, replicationSpec, cr.Status.ReplicationSpec, "replicationSpec not as expected")
	assert.Equal(t, numShards, cr.Status.NumShards, "numShards not as expected")
	assert.Equal(t, paused, cr.Status.Paused, "paused not as expected")
	assert.Equal(t, autoscaling, cr.Status.AutoScaling, "diskSizeGB not as expected")
	assert.Equal(t, providerSettings, cr.Status.ProviderSettings, "providerName not as expected")
}

func TestDeleteMongoDBAtlasCluster(t *testing.T) {
	// Set the logger to development mode for verbose logs.
	logf.SetLogger(logf.ZapLogger(true))

	// A MongoDBAtlasProject resource with metadata and spec.
	mongodbatlasproject := &knappekv1alpha1.MongoDBAtlasProject{
		ObjectMeta: metav1.ObjectMeta{
			Name:      projectName,
			Namespace: namespace,
		},
		Spec: knappekv1alpha1.MongoDBAtlasProjectSpec{
			OrgID: organizationID,
		},
		Status: knappekv1alpha1.MongoDBAtlasProjectStatus{
			ID:           projectID,
			Name:         projectName,
			OrgID:        organizationID,
			Created:      "2016-07-14T14:19:33Z",
			ClusterCount: 0,
		},
	}

	// A MongoDBAtlasCluster resource with metadata and spec.
	mongodbatlascluster := &knappekv1alpha1.MongoDBAtlasCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:              clusterName,
			Namespace:         namespace,
			DeletionTimestamp: &metav1.Time{Time: time.Now()},
			Finalizers:        []string{"finalizer.knappek.com"},
		},
		Spec: knappekv1alpha1.MongoDBAtlasClusterSpec{
			ProjectName:           projectName,
			ProviderSettings:      providerSettings,
			MongoDBVersion:        mongoDBVersion,
			MongoDBMajorVersion:   mongoDBMajorVersion,
			DiskSizeGB:            diskSizeGB,
			NumShards:             numShards,
			AutoScaling:           autoscaling,
			BackupEnabled:         backupEnabled,
			ProviderBackupEnabled: providerBackupEnabled,
			ReplicationSpec:       replicationSpec,
		},
		Status: knappekv1alpha1.MongoDBAtlasClusterStatus{
			GroupID:   projectID,
			Name:      clusterName,
			StateName: "IDLE",
		},
	}

	// Objects to track in the fake client.
	objs := []runtime.Object{
		mongodbatlascluster,
		mongodbatlasproject,
	}

	// Register operator types with the runtime scheme.
	s := scheme.Scheme
	s.AddKnownTypes(knappekv1alpha1.SchemeGroupVersion, mongodbatlascluster, mongodbatlasproject)

	// Create a fake k8s client to mock API calls.
	k8sClient := fake.NewFakeClient(objs...)
	// Create a fake atlas client to mock API calls.
	// atlasClient, server := test.NewAtlasFakeClient(t)
	httpClient, mux, server := testutil.Server()
	defer server.Close()
	atlasClient := ma.NewClient(httpClient)

	// Delete
	mux.HandleFunc("/api/atlas/v1.0/groups/"+projectID+"/clusters/"+clusterName, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "DELETE", r)
		fmt.Fprintf(w, `{}`)
	})

	// Create a ReconcileMongoDBAtlasCluster object with the scheme and fake client.
	r := &ReconcileMongoDBAtlasCluster{client: k8sClient, scheme: s, atlasClient: atlasClient}

	// Mock request to simulate Reconcile() being called on an event for a
	// watched resource .
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      clusterName,
			Namespace: namespace,
		},
	}
	res, err := r.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}
	assert.Equal(t, time.Second*20, res.RequeueAfter)

	// Check if the CR has been updated and has the correct status.
	cr := &knappekv1alpha1.MongoDBAtlasCluster{}
	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
	if err != nil {
		t.Fatalf("get MongoDBAtlasCluster: (%v)", err)
	}

	assert.Equal(t, "DELETING", cr.Status.StateName, "clusterID not as expected")

	httpClient2, mux2, server2 := testutil.Server()
	defer server2.Close()
	atlasClient2 := ma.NewClient(httpClient2)
	// GET: Simulate a new reconcile where cluster has been deleted successfully
	mux2.HandleFunc("/api/atlas/v1.0/groups/"+projectID+"/clusters/"+clusterName, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "GET", r)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})

	// Create a ReconcileMongoDBAtlasCluster object with the scheme and fake client.
	r2 := &ReconcileMongoDBAtlasCluster{client: k8sClient, scheme: s, atlasClient: atlasClient2}

	res2, err := r2.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}
	assert.Equal(t, reconcile.Result{}, res2)
	cr = &knappekv1alpha1.MongoDBAtlasCluster{}
	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
	if err != nil {
		t.Fatalf("get MongoDBAtlasCluster: (%v)", err)
	}
	// verify that Finalizer has been removed
	assert.Nil(t, cr.ObjectMeta.GetFinalizers())
}
