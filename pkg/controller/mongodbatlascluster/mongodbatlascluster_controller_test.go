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
	mongodbatlasproject := testutil.CreateAtlasProject(projectName, projectID, namespace, organizationID)

	// A MongoDBAtlasCluster resource with metadata and spec.
	mongodbatlascluster := &knappekv1alpha1.MongoDBAtlasCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterName,
			Namespace: namespace,
		},
		Spec: knappekv1alpha1.MongoDBAtlasClusterSpec{
			ProjectName: projectName,
			MongoDBAtlasClusterRequestBody: knappekv1alpha1.MongoDBAtlasClusterRequestBody{
				ProviderSettings:      providerSettings,
				MongoDBMajorVersion:   mongoDBMajorVersion,
				DiskSizeGB:            diskSizeGB,
				NumShards:             numShards,
				AutoScaling:           autoscaling,
				BackupEnabled:         backupEnabled,
				ProviderBackupEnabled: providerBackupEnabled,
				ReplicationSpec:       replicationSpec,
			},
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
		fmt.Fprintf(w, `{
			"autoScaling":{
				"diskGBEnabled":`+strconv.FormatBool(autoscaling.DiskGBEnabled)+`
			},
			"backupEnabled":`+strconv.FormatBool(backupEnabled)+`,
			"diskSizeGB":`+strconv.FormatFloat(diskSizeGB, 'f', 6, 64)+`,
			"groupId": "`+projectID+`",
			"id": "`+clusterID+`",
			"mongoDBVersion":"`+mongoDBVersion+`",
			"mongoDBMajorVersion":"`+mongoDBMajorVersion+`",
			"name":"`+clusterName+`",
			"numShards": `+strconv.Itoa(numShards)+`,
			"paused":`+strconv.FormatBool(paused)+`,
			"providerBackupEnabled":`+strconv.FormatBool(providerBackupEnabled)+`,
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
			},
			"stateName": "CREATING"
		}`)
	})

	// Create a ReconcileMongoDBAtlasCluster object with the scheme and fake client.
	r := &ReconcileMongoDBAtlasCluster{
		client:               k8sClient,
		scheme:               s,
		atlasClient:          atlasClient,
		reconciliationConfig: config.GetReconcilitationConfig(),
	}

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
	assert.Equal(t, time.Second*120, res.RequeueAfter)

	// Check if the CR has been created and has the correct status.
	cr := &knappekv1alpha1.MongoDBAtlasCluster{}
	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
	if err != nil {
		t.Fatalf("get MongoDBAtlasCluster: (%v)", err)
	}
	assert.Equal(t, "CREATING", cr.Status.StateName, "stateName not as expected")

	// GET: Simulate a new reconcile where stateName changed from CREATING to IDLE
	mux.HandleFunc("/api/atlas/v1.0/groups/"+projectID+"/clusters/"+clusterName, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "GET", r)
		fmt.Fprintf(w, `{
			"autoScaling":{
				"diskGBEnabled":`+strconv.FormatBool(autoscaling.DiskGBEnabled)+`
			},
			"backupEnabled":`+strconv.FormatBool(backupEnabled)+`,
			"diskSizeGB":`+strconv.FormatFloat(diskSizeGB, 'f', 6, 64)+`,
			"groupId": "`+projectID+`",
			"id": "`+clusterID+`",
			"mongoDBVersion":"`+mongoDBVersion+`",
			"mongoDBMajorVersion":"`+mongoDBMajorVersion+`",
			"name":"`+clusterName+`",
			"numShards": `+strconv.Itoa(numShards)+`,
			"paused":`+strconv.FormatBool(paused)+`,
			"providerBackupEnabled":`+strconv.FormatBool(providerBackupEnabled)+`,
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
			},
			"stateName": "IDLE"
		}`)
	})

	// Simulate a new reconcile where stateName changed from CREATING to IDLE
	res, err = r.Reconcile(req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}

	// Check if the CR has been created and has the correct status.
	cr = &knappekv1alpha1.MongoDBAtlasCluster{}
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
	mongodbatlasproject := testutil.CreateAtlasProject(projectName, projectID, namespace, organizationID)

	// A MongoDBAtlasCluster resource with metadata and spec.
	mongodbatlascluster := &knappekv1alpha1.MongoDBAtlasCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:              clusterName,
			Namespace:         namespace,
			DeletionTimestamp: &metav1.Time{Time: time.Now()},
			Finalizers:        []string{"finalizer.knappek.com"},
		},
		Spec: knappekv1alpha1.MongoDBAtlasClusterSpec{
			ProjectName: projectName,
			MongoDBAtlasClusterRequestBody: knappekv1alpha1.MongoDBAtlasClusterRequestBody{
				ProviderSettings:      providerSettings,
				MongoDBMajorVersion:   mongoDBMajorVersion,
				DiskSizeGB:            diskSizeGB,
				NumShards:             numShards,
				AutoScaling:           autoscaling,
				BackupEnabled:         backupEnabled,
				ProviderBackupEnabled: providerBackupEnabled,
				ReplicationSpec:       replicationSpec,
			},
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
	r := &ReconcileMongoDBAtlasCluster{
		client:               k8sClient,
		scheme:               s,
		atlasClient:          atlasClient,
		reconciliationConfig: config.GetReconcilitationConfig(),
	}

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
	r2 := &ReconcileMongoDBAtlasCluster{
		client:               k8sClient,
		scheme:               s,
		atlasClient:          atlasClient2,
		reconciliationConfig: config.GetReconcilitationConfig(),
	}

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

func TestUpdateMongoDBAtlasCluster(t *testing.T) {
	// Set the logger to development mode for verbose logs.
	logf.SetLogger(logf.ZapLogger(true))

	// A MongoDBAtlasProject resource with metadata and spec.
	mongodbatlasproject := testutil.CreateAtlasProject(projectName, projectID, namespace, organizationID)

	// updates
	updatedDiskSizeGB := diskSizeGB + 10
	updatedInstanceSizeName := "M20"
	updatedProviderSettings := ma.ProviderSettings{
		ProviderName:        "AWS",
		RegionName:          "US_EAST_1",
		InstanceSizeName:    updatedInstanceSizeName,
		EncryptEBSVolume:    true,
		BackingProviderName: "",
	}
	// A MongoDBAtlasCluster resource with metadata and spec.
	mongodbatlascluster := &knappekv1alpha1.MongoDBAtlasCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterName,
			Namespace: namespace,
		},
		Spec: knappekv1alpha1.MongoDBAtlasClusterSpec{
			ProjectName: projectName,
			MongoDBAtlasClusterRequestBody: knappekv1alpha1.MongoDBAtlasClusterRequestBody{
				ProviderSettings:      updatedProviderSettings, // update
				MongoDBMajorVersion:   mongoDBMajorVersion,
				DiskSizeGB:            updatedDiskSizeGB, // update
				NumShards:             numShards,
				AutoScaling:           autoscaling,
				BackupEnabled:         !backupEnabled, // update
				ProviderBackupEnabled: providerBackupEnabled,
				ReplicationSpec:       replicationSpec,
			},
		},
		Status: knappekv1alpha1.MongoDBAtlasClusterStatus{
			GroupID:        projectID,
			Name:           clusterName,
			StateName:      "IDLE",
			ID:             clusterID,
			MongoDBVersion: mongoDBVersion,
			MongoDBAtlasClusterRequestBody: knappekv1alpha1.MongoDBAtlasClusterRequestBody{
				ProviderSettings:      providerSettings,
				MongoDBMajorVersion:   mongoDBMajorVersion,
				DiskSizeGB:            diskSizeGB,
				NumShards:             numShards,
				AutoScaling:           autoscaling,
				BackupEnabled:         backupEnabled,
				ProviderBackupEnabled: providerBackupEnabled,
				ReplicationSpec:       replicationSpec,
			},
			Paused: paused,
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
	// Construct Update API call
	mux.HandleFunc("/api/atlas/v1.0/groups/"+projectID+"/clusters/"+clusterName, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "PATCH", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"autoScaling":{
				"diskGBEnabled":`+strconv.FormatBool(autoscaling.DiskGBEnabled)+`
			},
			"backupEnabled":`+strconv.FormatBool(!backupEnabled)+`,
			"diskSizeGB":`+strconv.FormatFloat(updatedDiskSizeGB, 'f', 6, 64)+`,
			"groupId": "`+projectID+`",
			"id": "`+clusterID+`",
			"mongoDBVersion":"`+mongoDBVersion+`",
			"mongoDBMajorVersion":"`+mongoDBMajorVersion+`",
			"name":"`+clusterName+`",
			"numShards": `+strconv.Itoa(numShards)+`,
			"paused":`+strconv.FormatBool(paused)+`,
			"providerBackupEnabled":`+strconv.FormatBool(providerBackupEnabled)+`,
			"providerSettings":{
				"providerName":"`+providerSettings.ProviderName+`",
				"regionName":"`+providerSettings.RegionName+`",
				"instanceSizeName":"`+updatedInstanceSizeName+`",
				"encryptEBSVolume": `+strconv.FormatBool(providerSettings.EncryptEBSVolume)+`
			},
			"replicationSpec":{
				"US_EAST_1":{
					"priority":7,
					"electableNodes":2,
					"readOnlyNodes":1,
					"analyticsNodes":1
				}
			},
			"stateName": "UPDATING"
		}`)
	})
	// Create a ReconcileMongoDBAtlasCluster object with the scheme and fake client.
	r := &ReconcileMongoDBAtlasCluster{
		client:               k8sClient,
		scheme:               s,
		atlasClient:          atlasClient,
		reconciliationConfig: config.GetReconcilitationConfig(),
	}

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
	assert.Equal(t, time.Second*120, res.RequeueAfter)

	// Check if the CR has been created and has the correct status.
	cr := &knappekv1alpha1.MongoDBAtlasCluster{}
	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
	if err != nil {
		t.Fatalf("get MongoDBAtlasCluster: (%v)", err)
	}
	assert.Equal(t, "UPDATING", cr.Status.StateName, "stateName not as expected")
	assert.Equal(t, !backupEnabled, cr.Status.BackupEnabled, "backupEnabled not as expected")
	assert.Equal(t, updatedDiskSizeGB, cr.Status.DiskSizeGB, "diskSizeGB not as expected")
	assert.Equal(t, updatedInstanceSizeName, cr.Status.ProviderSettings.InstanceSizeName, "instanceSizeName not as expected")
	assert.False(t, paused, "paused not as expected")
}

func TestNoUpdateMongoDBAtlasCluster(t *testing.T) {
	// Set the logger to development mode for verbose logs.
	logf.SetLogger(logf.ZapLogger(true))

	// A MongoDBAtlasProject resource with metadata and spec.
	mongodbatlasproject := testutil.CreateAtlasProject(projectName, projectID, namespace, organizationID)

	// A MongoDBAtlasCluster resource with metadata and spec. This Spec contains only the bare minimum, other values
	// will be filled with default values
	mongodbatlascluster := &knappekv1alpha1.MongoDBAtlasCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterName,
			Namespace: namespace,
		},
		Spec: knappekv1alpha1.MongoDBAtlasClusterSpec{
			ProjectName: projectName,
			MongoDBAtlasClusterRequestBody: knappekv1alpha1.MongoDBAtlasClusterRequestBody{
				ProviderSettings: providerSettings,
				ReplicationSpec:  replicationSpec,
			},
		},
		Status: knappekv1alpha1.MongoDBAtlasClusterStatus{
			GroupID:        projectID,
			Name:           clusterName,
			StateName:      "IDLE",
			ID:             clusterID,
			MongoDBVersion: mongoDBVersion,
			MongoDBAtlasClusterRequestBody: knappekv1alpha1.MongoDBAtlasClusterRequestBody{
				ProviderSettings:      providerSettings,
				MongoDBMajorVersion:   mongoDBMajorVersion,
				DiskSizeGB:            diskSizeGB,
				NumShards:             numShards,
				AutoScaling:           autoscaling,
				BackupEnabled:         backupEnabled,
				ProviderBackupEnabled: providerBackupEnabled,
				ReplicationSpec:       replicationSpec,
			},
			Paused: paused,
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
	// Construct Update API call
	mux.HandleFunc("/api/atlas/v1.0/groups/"+projectID+"/clusters/"+clusterName, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"autoScaling":{
				"diskGBEnabled":`+strconv.FormatBool(autoscaling.DiskGBEnabled)+`
			},
			"backupEnabled":`+strconv.FormatBool(backupEnabled)+`,
			"diskSizeGB":`+strconv.FormatFloat(diskSizeGB, 'f', 6, 64)+`,
			"groupId": "`+projectID+`",
			"id": "`+clusterID+`",
			"mongoDBVersion":"`+mongoDBVersion+`",
			"mongoDBMajorVersion":"`+mongoDBMajorVersion+`",
			"name":"`+clusterName+`",
			"numShards": `+strconv.Itoa(numShards)+`,
			"paused":`+strconv.FormatBool(paused)+`,
			"providerBackupEnabled":`+strconv.FormatBool(providerBackupEnabled)+`,
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
			},
			"stateName": "IDLE"
		}`)
	})
	// Create a ReconcileMongoDBAtlasCluster object with the scheme and fake client.
	r := &ReconcileMongoDBAtlasCluster{
		client:               k8sClient,
		scheme:               s,
		atlasClient:          atlasClient,
		reconciliationConfig: config.GetReconcilitationConfig(),
	}

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
	assert.Equal(t, time.Second*120, res.RequeueAfter)

	// Check if the CR has been created and has the correct status.
	cr := &knappekv1alpha1.MongoDBAtlasCluster{}
	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
	if err != nil {
		t.Fatalf("get MongoDBAtlasCluster: (%v)", err)
	}
	assert.Equal(t, "IDLE", cr.Status.StateName, "stateName not as expected")
}
