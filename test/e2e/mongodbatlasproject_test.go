package e2e

import (
	goctx "context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
)

func MongoDBAtlasProject(t *testing.T, ctx *framework.TestCtx, f *framework.Framework, namespace string) {
	// create MongoDBAtlasProject custom resource
	exampleMongoDBAtlasProject := &knappekv1alpha1.MongoDBAtlasProject{
		ObjectMeta: metav1.ObjectMeta{
			Name:      atlasProjectName,
			Namespace: namespace,
		},
		Spec: knappekv1alpha1.MongoDBAtlasProjectSpec{
			OrgID: *organizationID,
		},
	}
	err := f.Client.Create(goctx.TODO(), exampleMongoDBAtlasProject, &framework.CleanupOptions{TestContext: ctx, Timeout: time.Second * 5, RetryInterval: time.Second * 1})
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
	retryInterval := time.Second * 5
	timeout := time.Second * 60
	err := wait.Poll(retryInterval, timeout, func() (done bool, err error) {
		err = f.Client.Get(goctx.TODO(), types.NamespacedName{Name: p.Name, Namespace: p.Namespace}, p)
		return waitForManifestStatus(t, err, p.Name, p.Kind, p.Status, knappekv1alpha1.MongoDBAtlasProjectStatus{})
	})
	if err != nil {
		return err
	}
	return nil
}
