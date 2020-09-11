package e2e

// import (
// 	goctx "context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	knappekv1alpha1 "github.com/Knappek/mongodbatlas-operator/pkg/apis/knappek/v1alpha1"
// 	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"

// 	framework "github.com/operator-framework/operator-sdk/pkg/test"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/apimachinery/pkg/types"
// 	"k8s.io/apimachinery/pkg/util/wait"
// )

// var (
// 	eventTypeName = "OUTSIDE_METRIC_THRESHOLD"
// 	enabled       = true
// 	notifications = []ma.Notification{ma.Notification{
// 		TypeName:     "GROUP",
// 		IntervalMin:  5,
// 		DelayMin:     0,
// 		SMSEnabled:   false,
// 		EmailEnabled: true,
// 	}}
// 	metricThreshold = ma.MetricThreshold{
// 		MetricName: "QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED",
// 		Mode:       "AVERAGE",
// 		Operator:   "GREATER_THAN",
// 		Threshold:  500.0,
// 		Units:      "RAW",
// 	}
// )

// func MongoDBAtlasAlertConfiguration(t *testing.T, ctx *framework.TestCtx, f *framework.Framework, namespace string) {
// 	resourceName := "e2etest-testalertconfiguration"
// 	exampleMongoDBAtlasAlertConfiguration := &knappekv1alpha1.MongoDBAtlasAlertConfiguration{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      resourceName,
// 			Namespace: namespace,
// 		},
// 		Spec: knappekv1alpha1.MongoDBAtlasAlertConfigurationSpec{
// 			ProjectName: atlasProjectName,
// 			MongoDBAtlasAlertConfigurationRequestBody: knappekv1alpha1.MongoDBAtlasAlertConfigurationRequestBody{
// 				EventTypeName:   eventTypeName,
// 				Enabled:         enabled,
// 				Notifications:   notifications,
// 				MetricThreshold: metricThreshold,
// 			},
// 		},
// 	}
// 	err := f.Client.Create(goctx.TODO(), exampleMongoDBAtlasAlertConfiguration, &framework.CleanupOptions{TestContext: ctx, Timeout: time.Second * 5, RetryInterval: time.Second * 1})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("wait for creating AlertConfiguration: %v\n", exampleMongoDBAtlasAlertConfiguration.ObjectMeta.Name)
// 	err = waitForMongoDBAtlasAlertConfiguration(t, f, exampleMongoDBAtlasAlertConfiguration, 500.0)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("AlertConfiguration %v successfully created\n", exampleMongoDBAtlasAlertConfiguration.ObjectMeta.Name)

// 	// update resource
// 	updatedMetricThreshold := ma.MetricThreshold{
// 		MetricName: "QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED",
// 		Mode:       "AVERAGE",
// 		Operator:   "GREATER_THAN",
// 		Threshold:  1000.0,
// 		Units:      "RAW",
// 	}

// 	exampleMongoDBAtlasAlertConfiguration.Spec.MetricThreshold = updatedMetricThreshold
// 	err = f.Client.Update(goctx.TODO(), exampleMongoDBAtlasAlertConfiguration)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("wait for updating AlertConfiguration: %v\n", exampleMongoDBAtlasAlertConfiguration.ObjectMeta.Name)
// 	err = waitForMongoDBAtlasAlertConfiguration(t, f, exampleMongoDBAtlasAlertConfiguration, 1000.0)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("AlertConfiguration %v successfully updated\n", exampleMongoDBAtlasAlertConfiguration.ObjectMeta.Name)
// }

// func waitForMongoDBAtlasAlertConfiguration(t *testing.T, f *framework.Framework, p *knappekv1alpha1.MongoDBAtlasAlertConfiguration, desiredState float64) error {
// 	retryInterval := time.Second * 5
// 	timeout := time.Minute * 5
// 	err := wait.Poll(retryInterval, timeout, func() (done bool, err error) {
// 		err = f.Client.Get(goctx.TODO(), types.NamespacedName{Name: p.Name, Namespace: p.Namespace}, p)
// 		return isInDesiredState(t, err, p.Name, p.Kind, p.Status.MetricThreshold.Threshold, desiredState)
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
