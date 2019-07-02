package e2e

import (
	goctx "context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"

	framework "github.com/operator-framework/operator-sdk/pkg/test"
	corev1 "k8s.io/api/core/v1"
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
			MongoDBAtlasAuth: knappekv1alpha1.MongoDBAtlasAuth{
				PublicKey: "toppaljd",
				PrivateKey: knappekv1alpha1.PrivateKey{
					ValueFrom: &knappekv1alpha1.PrivateKeySource{
						SecretKeyRef: &corev1.SecretKeySelector{
							Key: "privateKey",
							LocalObjectReference: corev1.LocalObjectReference{
								Name: "example-monogdb-atlas-project",
							},
						},
					},
				},
			},
			ProjectName: atlasProjectName,
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
	}
	err := f.Client.Create(goctx.TODO(), exampleMongoDBAtlasCluster, &framework.CleanupOptions{TestContext: ctx, Timeout: time.Minute * 10, RetryInterval: time.Second * 30})
	if err != nil {
		t.Fatal(err)
	}

	err = waitForMongoDBAtlasCluster(t, f, exampleMongoDBAtlasCluster)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "IDLE", exampleMongoDBAtlasCluster.Status.StateName, "The Cluster StateName is not IDLE")
}

func waitForMongoDBAtlasCluster(t *testing.T, f *framework.Framework, p *knappekv1alpha1.MongoDBAtlasCluster) error {
	retryInterval := time.Second * 30
	timeout := time.Minute * 15
	err := wait.Poll(retryInterval, timeout, func() (done bool, err error) {
		err = f.Client.Get(goctx.TODO(), types.NamespacedName{Name: p.Name, Namespace: p.Namespace}, p)
		return waitForManifestStatus(t, err, p.Name, p.Kind, p.Status.StateName, "CREATING")
	})
	if err != nil {
		return err
	}
	return nil
}
