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

package certmanager

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	certificatesv1alpha1 "github.com/klearwave/operators-capability/capabilities/certificates/apis/certificates/v1alpha1"
	"github.com/klearwave/operators-capability/capabilities/certificates/apis/certificates/v1alpha1/certmanager/mutate"
)

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceCertManagerCainjector creates the Service resource with name cert-manager-cainjector.
func CreateServiceNamespaceCertManagerCainjector(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "cert-manager-cainjector",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":                 "cainjector",
					"app.kubernetes.io/version":              "v1.17.2",
					"platform.klearwave.io/capability":       "certificates",
					"platform.klearwave.io/component":        "cert-manager",
					"platform.klearwave.io/operator-version": "unstable",
					"platform.klearwave.io/platform-version": "unstable",
					"app.kubernetes.io/part-of":              "platform",
					"app.kubernetes.io/managed-by":           "certificates-operator",
				},
			},
			"spec": map[string]interface{}{
				"type": "ClusterIP",
				"ports": []interface{}{
					map[string]interface{}{
						"protocol": "TCP",
						"port":     9402,
						"name":     "http-metrics",
					},
				},
				"selector": map[string]interface{}{
					"app.kubernetes.io/name":                 "cainjector",
					"platform.klearwave.io/capability":       "certificates",
					"platform.klearwave.io/component":        "cert-manager",
					"platform.klearwave.io/operator-version": "unstable",
					"platform.klearwave.io/platform-version": "unstable",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceCertManagerCainjector(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceCertManager creates the Service resource with name cert-manager.
func CreateServiceNamespaceCertManager(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "cert-manager",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":                 "cert-manager",
					"app.kubernetes.io/version":              "v1.17.2",
					"platform.klearwave.io/capability":       "certificates",
					"platform.klearwave.io/component":        "cert-manager",
					"platform.klearwave.io/operator-version": "unstable",
					"platform.klearwave.io/platform-version": "unstable",
					"app.kubernetes.io/part-of":              "platform",
					"app.kubernetes.io/managed-by":           "certificates-operator",
				},
			},
			"spec": map[string]interface{}{
				"type": "ClusterIP",
				"ports": []interface{}{
					map[string]interface{}{
						"protocol":   "TCP",
						"port":       9402,
						"name":       "tcp-prometheus-servicemonitor",
						"targetPort": 9402,
					},
				},
				"selector": map[string]interface{}{
					"app.kubernetes.io/name":                 "cert-manager",
					"platform.klearwave.io/capability":       "certificates",
					"platform.klearwave.io/component":        "cert-manager",
					"platform.klearwave.io/operator-version": "unstable",
					"platform.klearwave.io/platform-version": "unstable",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceCertManager(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceCertManagerWebhook creates the Service resource with name cert-manager-webhook.
func CreateServiceNamespaceCertManagerWebhook(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "cert-manager-webhook",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":                 "webhook",
					"app.kubernetes.io/version":              "v1.17.2",
					"platform.klearwave.io/capability":       "certificates",
					"platform.klearwave.io/component":        "cert-manager",
					"platform.klearwave.io/operator-version": "unstable",
					"platform.klearwave.io/platform-version": "unstable",
					"app.kubernetes.io/part-of":              "platform",
					"app.kubernetes.io/managed-by":           "certificates-operator",
				},
			},
			"spec": map[string]interface{}{
				"type": "ClusterIP",
				"ports": []interface{}{
					map[string]interface{}{
						"name":       "https",
						"port":       443,
						"protocol":   "TCP",
						"targetPort": "https",
					},
					map[string]interface{}{
						"name":       "metrics",
						"port":       9402,
						"protocol":   "TCP",
						"targetPort": "http-metrics",
					},
				},
				"selector": map[string]interface{}{
					"app.kubernetes.io/name":                 "webhook",
					"platform.klearwave.io/capability":       "certificates",
					"platform.klearwave.io/component":        "cert-manager",
					"platform.klearwave.io/operator-version": "unstable",
					"platform.klearwave.io/platform-version": "unstable",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceCertManagerWebhook(resourceObj, parent, reconciler, req)
}
