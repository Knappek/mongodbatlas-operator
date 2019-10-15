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

func MongoDBAtlasDatabaseUser(t *testing.T, ctx *framework.TestCtx, f *framework.Framework, namespace string) {
	// create MongoDBAtlasProject custom resource
	username := "e2etest-testuser"
	exampleMongoDBAtlasDatabaseUser := &knappekv1alpha1.MongoDBAtlasDatabaseUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      username,
			Namespace: namespace,
		},
		Spec: knappekv1alpha1.MongoDBAtlasDatabaseUserSpec{
			ProjectName: atlasProjectName,
			MongoDBAtlasDatabaseUserRequestBody: knappekv1alpha1.MongoDBAtlasDatabaseUserRequestBody{
				Password:        "$upers€curep@ssword!",
				DeleteAfterDate: "2100-01-01T00:00:00Z",
				DatabaseName:    "admin",
				Roles:           []ma.Role{ma.Role{DatabaseName: "e2etestdatabase", RoleName: "readWrite"}},
			},
		},
	}
	err := f.Client.Create(goctx.TODO(), exampleMongoDBAtlasDatabaseUser, &framework.CleanupOptions{TestContext: ctx, Timeout: time.Second * 5, RetryInterval: time.Second * 1})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("wait for creating the databaseUser: %v\n", exampleMongoDBAtlasDatabaseUser.ObjectMeta.Name)
	err = waitForMongoDBAtlasDatabaseUser(t, f, exampleMongoDBAtlasDatabaseUser, "2100-01-01T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("databaseUser %v successfully created\n", exampleMongoDBAtlasDatabaseUser.ObjectMeta.Name)

	// update databaseUser
	exampleMongoDBAtlasDatabaseUser.Spec.DeleteAfterDate = "2100-02-01T00:00:00Z"
	err = f.Client.Update(goctx.TODO(), exampleMongoDBAtlasDatabaseUser)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("wait for updating the databaseUser: %v\n", exampleMongoDBAtlasDatabaseUser.ObjectMeta.Name)
	err = waitForMongoDBAtlasDatabaseUser(t, f, exampleMongoDBAtlasDatabaseUser, "2100-02-01T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("databaseUser %v successfully updated\n", exampleMongoDBAtlasDatabaseUser.ObjectMeta.Name)
}

func waitForMongoDBAtlasDatabaseUser(t *testing.T, f *framework.Framework, p *knappekv1alpha1.MongoDBAtlasDatabaseUser, desiredState string) error {
	retryInterval := time.Second * 5
	timeout := time.Second * 10
	err := wait.Poll(retryInterval, timeout, func() (done bool, err error) {
		err = f.Client.Get(goctx.TODO(), types.NamespacedName{Name: p.Name, Namespace: p.Namespace}, p)
		return isInDesiredState(t, err, p.Name, p.Kind, p.Status.DeleteAfterDate, desiredState)
	})
	if err != nil {
		return err
	}
	return nil
}
