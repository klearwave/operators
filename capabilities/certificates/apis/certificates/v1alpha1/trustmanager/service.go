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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	certificatesv1alpha1 "github.com/klearwave/operators-capability/capabilities/certificates/apis/certificates/v1alpha1"
	"github.com/klearwave/operators-capability/capabilities/certificates/apis/certificates/v1alpha1/trustmanager/mutate"
)

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceTrustManagerMetrics creates the Service resource with name trust-manager-metrics.
func CreateServiceNamespaceTrustManagerMetrics(
	parent *certificatesv1alpha1.TrustManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "trust-manager-metrics",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":                 "trust-manager",
					"app.kubernetes.io/version":              "v0.17.1",
					"app.kubernetes.io/managed-by":           "certificates-operator",
					"platform.klearwave.io/capability":       "certificates",
					"platform.klearwave.io/component":        "trust-manager",
					"platform.klearwave.io/operator-version": "unstable",
					"platform.klearwave.io/platform-version": "unstable",
					"app.kubernetes.io/part-of":              "platform",
				},
			},
			"spec": map[string]interface{}{
				"type": "ClusterIP",
				"ports": []interface{}{
					map[string]interface{}{
						"port":       9402,
						"targetPort": 9402,
						"protocol":   "TCP",
						"name":       "metrics",
					},
				},
				"selector": map[string]interface{}{
					"app":                                    "trust-manager",
					"platform.klearwave.io/capability":       "certificates",
					"platform.klearwave.io/component":        "trust-manager",
					"platform.klearwave.io/operator-version": "unstable",
					"platform.klearwave.io/platform-version": "unstable",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceTrustManagerMetrics(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceTrustManager creates the Service resource with name trust-manager.
func CreateServiceNamespaceTrustManager(
	parent *certificatesv1alpha1.TrustManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "trust-manager",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":                 "trust-manager",
					"app.kubernetes.io/version":              "v0.17.1",
					"app.kubernetes.io/managed-by":           "certificates-operator",
					"platform.klearwave.io/capability":       "certificates",
					"platform.klearwave.io/component":        "trust-manager",
					"platform.klearwave.io/operator-version": "unstable",
					"platform.klearwave.io/platform-version": "unstable",
					"app.kubernetes.io/part-of":              "platform",
				},
			},
			"spec": map[string]interface{}{
				"type": "ClusterIP",
				"ports": []interface{}{
					map[string]interface{}{
						"port":       443,
						"targetPort": 6443,
						"protocol":   "TCP",
						"name":       "webhook",
					},
				},
				"selector": map[string]interface{}{
					"app":                                    "trust-manager",
					"platform.klearwave.io/capability":       "certificates",
					"platform.klearwave.io/component":        "trust-manager",
					"platform.klearwave.io/operator-version": "unstable",
					"platform.klearwave.io/platform-version": "unstable",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceTrustManager(resourceObj, parent, reconciler, req)
}
