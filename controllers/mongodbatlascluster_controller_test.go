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

	mongodbatlasv1alpha1 "github.com/knappek/mongodbatlas-operator/api/v1alpha1"
	testutil "github.com/knappek/mongodbatlas-operator/controllers/test"
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
	clusterName           = "unittest-cluster"
	clusterID             = "testClusterId"
	mongoDBVersion        = "3.4"
	mongoDBMajorVersion   = "3.4"
	diskSizeGB            = "10.500000"
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
	mongodbatlasproject := testutil.CreateAtlasProject(projectName, groupID, namespace, organizationID)

	// A MongoDBAtlasCluster resource with metadata and spec.
	mongodbatlascluster := &mongodbatlasv1alpha1.MongoDBAtlasCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterName,
			Namespace: namespace,
		},
		Spec: mongodbatlasv1alpha1.MongoDBAtlasClusterSpec{
			ProjectName: projectName,
			MongoDBAtlasClusterRequestBody: mongodbatlasv1alpha1.MongoDBAtlasClusterRequestBody{
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
	s.AddKnownTypes(mongodbatlasv1alpha1.GroupVersion, mongodbatlascluster, mongodbatlasproject)

	// Create a fake k8s client to mock API calls.
	k8sClient := fake.NewFakeClient(objs...)
	// Create a fake atlas client to mock API calls.
	// atlasClient, server := test.NewAtlasFakeClient(t)
	httpClient, mux, server := testutil.Server()
	defer server.Close()
	atlasClient := ma.NewClient(httpClient)

	// Post
	mux.HandleFunc("/api/atlas/v1.0/groups/"+groupID+"/clusters", func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"autoScaling":{
				"diskGBEnabled":`+strconv.FormatBool(autoscaling.DiskGBEnabled)+`
			},
			"backupEnabled":`+strconv.FormatBool(backupEnabled)+`,
			"diskSizeGB":`+diskSizeGB+`,
			"groupId": "`+groupID+`",
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
	r := &MongoDBAtlasClusterReconciler{
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
	cr := &mongodbatlasv1alpha1.MongoDBAtlasCluster{}
	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
	if err != nil {
		t.Fatalf("get MongoDBAtlasCluster: (%v)", err)
	}
	assert.Equal(t, "CREATING", cr.Status.StateName, "stateName not as expected")

	// GET: Simulate a new reconcile where stateName changed from CREATING to IDLE
	mux.HandleFunc("/api/atlas/v1.0/groups/"+groupID+"/clusters/"+clusterName, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, "GET", r)
		fmt.Fprintf(w, `{
			"autoScaling":{
				"diskGBEnabled":`+strconv.FormatBool(autoscaling.DiskGBEnabled)+`
			},
			"backupEnabled":`+strconv.FormatBool(backupEnabled)+`,
			"diskSizeGB":`+diskSizeGB+`,
			"groupId": "`+groupID+`",
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
	cr = &mongodbatlasv1alpha1.MongoDBAtlasCluster{}
	err = k8sClient.Get(context.TODO(), req.NamespacedName, cr)
	if err != nil {
		t.Fatalf("get MongoDBAtlasCluster: (%v)", err)
	}

	assert.Equal(t, "finalizer.knappek.com", cr.ObjectMeta.GetFinalizers()[0], "Finalizer not as expected")
	assert.Equal(t, clusterID, cr.Status.ID, "clusterID not as expected")
	assert.Equal(t, clusterName, cr.Status.Name, "clusterName not as expected")
	assert.Equal(t, groupID, cr.Status.GroupID, "groupID not as expected")
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
