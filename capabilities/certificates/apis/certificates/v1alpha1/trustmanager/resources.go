/*
Copyright 2025.

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

package trustmanager

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	certificatesv1alpha1 "github.com/klearwave/operators-capability/capabilities/certificates/apis/certificates/v1alpha1"
)

// sampleTrustManager is a sample containing all fields
const sampleTrustManager = `apiVersion: certificates.platform.klearwave.io/v1alpha1
kind: TrustManager
metadata:
  name: trustmanager-sample
spec:
  namespace: "klearwave-certificates-system"
  controller:
    replicas: 1
    image: "quay.io/jetstack/trust-manager:v0.17.1"
    resources:
      requests:
        cpu: "25m"
        memory: "32Mi"
      limits:
        memory: "64Mi"
`

// sampleTrustManagerRequired is a sample containing only required fields
const sampleTrustManagerRequired = `apiVersion: certificates.platform.klearwave.io/v1alpha1
kind: TrustManager
metadata:
  name: trustmanager-sample
spec:
`

// Sample returns the sample manifest for this custom resource.
func Sample(requiredOnly bool) string {
	if requiredOnly {
		return sampleTrustManagerRequired
	}

	return sampleTrustManager
}

// Generate returns the child resources that are associated with this workload given
// appropriate structured inputs.
func Generate(
	workloadObj certificatesv1alpha1.TrustManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {
	resourceObjects := []client.Object{}

	for _, f := range CreateFuncs {
		resources, err := f(&workloadObj, reconciler, req)

		if err != nil {
			return nil, err
		}

		resourceObjects = append(resourceObjects, resources...)
	}

	return resourceObjects, nil
}

// GenerateForCLI returns the child resources that are associated with this workload given
// appropriate YAML manifest files.
func GenerateForCLI(workloadFile []byte) ([]client.Object, error) {
	var workloadObj certificatesv1alpha1.TrustManager
	if err := yaml.Unmarshal(workloadFile, &workloadObj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml into workload, %w", err)
	}

	if err := workload.Validate(&workloadObj); err != nil {
		return nil, fmt.Errorf("error validating workload yaml, %w", err)
	}

	return Generate(workloadObj, nil, nil)
}

// CreateFuncs is an array of functions that are called to create the child resources for the controller
// in memory during the reconciliation loop prior to persisting the changes or updates to the Kubernetes
// database.
var CreateFuncs = []func(
	*certificatesv1alpha1.TrustManager,
	workload.Reconciler,
	*workload.Request,
) ([]client.Object, error){
	CreateCRDBundlesTrustCertManagerIo,
	CreateServiceAccountNamespaceTrustManager,
	CreateClusterRoleTrustManager,
	CreateClusterRoleBindingTrustManager,
	CreateRoleNamespaceTrustManager,
	CreateRoleNamespaceTrustManagerLeaderelection,
	CreateRoleBindingNamespaceTrustManager,
	CreateRoleBindingNamespaceTrustManagerLeaderelection,
	CreateCertKlearwaveCertificatesSystemTrustManager,
	CreateDeploymentNamespaceTrustManager,
	CreateServiceNamespaceTrustManagerMetrics,
	CreateServiceNamespaceTrustManager,
	CreateValidatingWebhookTrustManager,
}

// InitFuncs is an array of functions that are called prior to starting the controller manager.  This is
// necessary in instances which the controller needs to "own" objects which depend on resources to
// pre-exist in the cluster. A common use case for this is the need to own a custom resource.
// If the controller needs to own a custom resource type, the CRD that defines it must
// first exist. In this case, the InitFunc will create the CRD so that the controller
// can own custom resources of that type.  Without the InitFunc the controller will
// crash loop because when it tries to own a non-existent resource type during manager
// setup, it will fail.
var InitFuncs = []func(
	*certificatesv1alpha1.TrustManager,
	workload.Reconciler,
	*workload.Request,
) ([]client.Object, error){
	CreateCRDBundlesTrustCertManagerIo,
}

func ConvertWorkload(component workload.Workload) (*certificatesv1alpha1.TrustManager, error) {
	p, ok := component.(*certificatesv1alpha1.TrustManager)
	if !ok {
		return nil, certificatesv1alpha1.ErrUnableToConvertTrustManager
	}

	return p, nil
}
