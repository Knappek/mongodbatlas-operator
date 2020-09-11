package e2e

// import (
// 	"flag"
// 	"fmt"
// 	"testing"
// 	"time"

// 	apis "github.com/Knappek/mongodbatlas-operator/pkg/apis"
// 	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
// 	f "github.com/operator-framework/operator-sdk/pkg/test"
// 	"github.com/operator-framework/operator-sdk/pkg/test/e2eutil"
// )

// func TestMain(m *testing.M) {
// 	f.MainEntry(m)
// }

// var (
// 	atlasProjectName = "e2etest-project"
// 	organizationID   = flag.String("organizationID", "", "MongoDB Atlas Organization ID")
// )

// func TestMongoDBAtlas(t *testing.T) {
// 	err := registerTypes(t)
// 	if err != nil {
// 		t.Fatalf("failed to add custom resource scheme to framework: %v", err)
// 	}
// 	ctx := f.NewTestCtx(t)
// 	defer ctx.Cleanup()

// 	err = ctx.InitializeClusterResources(&f.CleanupOptions{TestContext: ctx, Timeout: time.Second * 5, RetryInterval: time.Second * 1})
// 	if err != nil {
// 		t.Fatalf("failed to initialize cluster resources: %v", err)
// 	}
// 	// get namespace
// 	namespace, err := ctx.GetNamespace()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	// get global framework variables
// 	f := f.Global
// 	// wait for mongodbatlas-operator to be ready
// 	err = e2eutil.WaitForOperatorDeployment(t, f.KubeClient, namespace, "mongodbatlas-operator", 1, time.Second*5, time.Second*30)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	MongoDBAtlasProject(t, ctx, f, namespace)
// 	MongoDBAtlasCluster(t, ctx, f, namespace)
// 	MongoDBAtlasDatabaseUser(t, ctx, f, namespace)
// 	MongoDBAtlasAlertConfiguration(t, ctx, f, namespace)
// 	fmt.Println("Cleanup resources...")
// }

// func registerTypes(t *testing.T) error {
// 	// MongoDBAtlasProject
// 	mongoDBAtlasProjectList := &knappekv1alpha1.MongoDBAtlasProjectList{}
// 	err := f.AddToFrameworkScheme(apis.AddToScheme, mongoDBAtlasProjectList)
// 	if err != nil {
// 		return err
// 	}

// 	// MongoDBAtlasCluster
// 	mongoDBAtlasClusterList := &knappekv1alpha1.MongoDBAtlasClusterList{}
// 	err = f.AddToFrameworkScheme(apis.AddToScheme, mongoDBAtlasClusterList)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
