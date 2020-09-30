/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	mongodbatlasv1alpha1 "github.com/knappek/mongodbatlas-operator/api/v1alpha1"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var cfg *rest.Config
var k8sClient client.Client
var testEnv *envtest.Environment

var (
	e2eAtlasProjectName = "e2etest-project"
	// e2eOrganizationID   = flag.String("organizationID", "", "MongoDB Atlas Organization ID")
	e2eOrganizationID = "5c4a2a55553855344780cf5f"

	k8sNamespace = &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}
	exampleMongoDBAtlasProject = &mongodbatlasv1alpha1.MongoDBAtlasProject{
		ObjectMeta: metav1.ObjectMeta{
			Name:      e2eAtlasProjectName,
			Namespace: namespace,
		},
		Spec: mongodbatlasv1alpha1.MongoDBAtlasProjectSpec{
			OrgID: e2eOrganizationID,
		},
	}
)

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"Controller Suite",
		[]Reporter{printer.NewlineReporter{}})
}

var _ = BeforeSuite(func(done Done) {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))

	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{filepath.Join("..", "config", "crd", "bases")},
	}

	var err error
	cfg, err = testEnv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(cfg).ToNot(BeNil())

	err = mongodbatlasv1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	// +kubebuilder:scaffold:scheme

	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).ToNot(HaveOccurred())
	Expect(k8sClient).ToNot(BeNil())

	// start operator

	By("creating a MongoDB Atlas project")
	err = k8sClient.Create(context.TODO(), k8sNamespace)
	Expect(err).ToNot(HaveOccurred())
	err = k8sClient.Create(context.TODO(), exampleMongoDBAtlasProject)
	Expect(err).ToNot(HaveOccurred())

	By("verifying the MongoDB Atlas project")
	objkey := types.NamespacedName{
		Name:      e2eAtlasProjectName,
		Namespace: namespace,
	}
	result := &mongodbatlasv1alpha1.MongoDBAtlasProject{}
	err = k8sClient.Get(context.TODO(), objkey, result)
	Expect(err).NotTo(HaveOccurred())
	
	// wait for status to come up and assert the value
	// p.Status.Name, "e2etest-project


	close(done)
}, 60)

var _ = AfterSuite(func() {
	By("deleting MongoDB Atlas project")
	err := k8sClient.Delete(context.TODO(), exampleMongoDBAtlasProject)
	Expect(err).ToNot(HaveOccurred())
	err = k8sClient.Delete(context.TODO(), k8sNamespace)
	Expect(err).ToNot(HaveOccurred())

	By("tearing down the test environment")
	err = testEnv.Stop()
	Expect(err).ToNot(HaveOccurred())
})
