package e2e

import (
	goctx "context"
	"flag"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	apis "github.com/Knappek/mongodbatlas-operator/pkg/apis"
	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"

	framework "github.com/operator-framework/operator-sdk/pkg/test"
	"github.com/operator-framework/operator-sdk/pkg/test/e2eutil"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
)

var (
	retryInterval        = time.Second * 5
	timeout              = time.Second * 60
	cleanupRetryInterval = time.Second * 1
	cleanupTimeout       = time.Second * 5
	organizationID = flag.String("organizationID", "", "MongoDB Atlas Organization ID")
)

func TestMongoDBAtlasProject(t *testing.T) {
	mongoDBAtlasProjectList := &knappekv1alpha1.MongoDBAtlasProjectList{}
	err := framework.AddToFrameworkScheme(apis.AddToScheme, mongoDBAtlasProjectList)
	if err != nil {
		t.Fatalf("failed to add custom resource scheme to framework: %v", err)
	}
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup()

	err = ctx.InitializeClusterResources(&framework.CleanupOptions{TestContext: ctx, Timeout: cleanupTimeout, RetryInterval: cleanupRetryInterval})
	if err != nil {
		t.Fatalf("failed to initialize cluster resources: %v", err)
	}
	// get namespace
	namespace, err := ctx.GetNamespace()
	if err != nil {
		t.Fatal(err)
	}
	// get global framework variables
	f := framework.Global
	// wait for mongodbatlas-operator to be ready
	err = e2eutil.WaitForOperatorDeployment(t, f.KubeClient, namespace, "mongodbatlas-operator", 1, time.Second*5, time.Second*30)
	if err != nil {
		t.Fatal(err)
	}

	// create MongoDBAtlasProject custom resource
	atlasProjectName := "e2etest-project"
	exampleMongoDBAtlasProject := &knappekv1alpha1.MongoDBAtlasProject{
		ObjectMeta: metav1.ObjectMeta{
			Name:      atlasProjectName,
			Namespace: namespace,
		},
		Spec: knappekv1alpha1.MongoDBAtlasProjectSpec{
			OrgID: *organizationID,
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
		},
	}

	err = f.Client.Create(goctx.TODO(), exampleMongoDBAtlasProject, &framework.CleanupOptions{TestContext: ctx, Timeout: time.Second * 5, RetryInterval: time.Second * 1})
	if err != nil {
		t.Fatal(err)
	}

	err = waitForMongoDBAtlasProject(t, f, exampleMongoDBAtlasProject)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, *organizationID, exampleMongoDBAtlasProject.Status.OrgID, "The Organization ID in the Status block is incorrect")
	assert.Equal(t, "e2etest-project", exampleMongoDBAtlasProject.Status.Name, "The Project Name in the Status block is incorrect")
}

func waitForMongoDBAtlasProject(t *testing.T, f *framework.Framework, p *knappekv1alpha1.MongoDBAtlasProject) error {
	err := wait.Poll(retryInterval, timeout, func() (done bool, err error) {
		err = f.Client.Get(goctx.TODO(), types.NamespacedName{Name: p.Name, Namespace: p.Namespace}, p)
		return waitForNonEmptyStatus(t, err, p.Name, p.Kind, p.Status, knappekv1alpha1.MongoDBAtlasProjectStatus{})
	})
	if err != nil {
		return err
	}
	return nil
}
