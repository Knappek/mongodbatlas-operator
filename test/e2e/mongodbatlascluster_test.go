package e2e

import (
	goctx "context"
	"fmt"
	"testing"
	"time"

	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"

	framework "github.com/operator-framework/operator-sdk/pkg/test"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
)

func MongoDBAtlasCluster(t *testing.T, ctx *framework.TestCtx, f *framework.Framework, namespace string) {
	// create MongoDBAtlasProject custom resource
	atlasClusterName := "e2etest-cluster"
	exampleMongoDBAtlasCluster := &knappekv1alpha1.MongoDBAtlasCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      atlasClusterName,
			Namespace: namespace,
		},
		Spec: knappekv1alpha1.MongoDBAtlasClusterSpec{
			ProjectName: atlasProjectName,
			MongoDBAtlasClusterRequestBody: knappekv1alpha1.MongoDBAtlasClusterRequestBody{
				ProviderSettings: ma.ProviderSettings{
					ProviderName:     "AWS",
					RegionName:       "EU_CENTRAL_1",
					InstanceSizeName: "M10",
					EncryptEBSVolume: false,
				},
				NumShards: 1,
				AutoScaling: ma.AutoScaling{
					DiskGBEnabled: false,
				},
				BackupEnabled:         false,
				ProviderBackupEnabled: false,
			},
		},
	}
	err := f.Client.Create(goctx.TODO(), exampleMongoDBAtlasCluster, &framework.CleanupOptions{TestContext: ctx, Timeout: time.Minute * 20, RetryInterval: time.Second * 30})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("wait for creating the cluster: %v\n", exampleMongoDBAtlasCluster.ObjectMeta.Name)
	err = waitForMongoDBAtlasCluster(t, f, exampleMongoDBAtlasCluster, "IDLE")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("cluster %v successfully created\n", exampleMongoDBAtlasCluster.ObjectMeta.Name)

	// update cluster
	exampleMongoDBAtlasCluster.Spec.AutoScaling.DiskGBEnabled = true
	err = f.Client.Update(goctx.TODO(), exampleMongoDBAtlasCluster)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("wait for updating the cluster: %v\n", exampleMongoDBAtlasCluster.ObjectMeta.Name)
	err = waitForMongoDBAtlasCluster(t, f, exampleMongoDBAtlasCluster, "IDLE")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("cluster %v successfully updated\n", exampleMongoDBAtlasCluster.ObjectMeta.Name)
}

func waitForMongoDBAtlasCluster(t *testing.T, f *framework.Framework, p *knappekv1alpha1.MongoDBAtlasCluster, desiredState string) error {
	retryInterval := time.Second * 30
	timeout := time.Minute * 15
	err := wait.Poll(retryInterval, timeout, func() (done bool, err error) {
		err = f.Client.Get(goctx.TODO(), types.NamespacedName{Name: p.Name, Namespace: p.Namespace}, p)
		return isInDesiredState(t, err, p.Name, p.Kind, p.Status.StateName, desiredState)
	})
	if err != nil {
		return err
	}
	return nil
}
