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

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateMutatingWebhookCertManagerWebhook creates the MutatingWebhookConfiguration resource with name cert-manager-webhook.
func CreateMutatingWebhookCertManagerWebhook(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "MutatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name": "cert-manager-webhook",
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
				"annotations": map[string]interface{}{
					// controlled by field: namespace
					"cert-manager.io/inject-ca-from-secret": "" + parent.Spec.Namespace + "/cert-manager-webhook-ca",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"name": "webhook.cert-manager.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"cert-manager.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
							},
							"resources": []interface{}{
								"certificaterequests",
							},
						},
					},
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"matchPolicy":    "Equivalent",
					"timeoutSeconds": 30,
					"failurePolicy":  "Fail",
					"sideEffects":    "None",
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "cert-manager-webhook",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/mutate",
						},
					},
				},
			},
		},
	}

	return mutate.MutateMutatingWebhookCertManagerWebhook(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateValidatingWebhookCertManagerWebhook creates the ValidatingWebhookConfiguration resource with name cert-manager-webhook.
func CreateValidatingWebhookCertManagerWebhook(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "ValidatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name": "cert-manager-webhook",
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
				"annotations": map[string]interface{}{
					// controlled by field: namespace
					"cert-manager.io/inject-ca-from-secret": "" + parent.Spec.Namespace + "/cert-manager-webhook-ca",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"name": "webhook.cert-manager.io",
					"namespaceSelector": map[string]interface{}{
						"matchExpressions": []interface{}{
							map[string]interface{}{
								"key":      "cert-manager.io/disable-validation",
								"operator": "NotIn",
								"values": []interface{}{
									"true",
								},
							},
						},
					},
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"cert-manager.io",
								"acme.cert-manager.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"*/*",
							},
						},
					},
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"matchPolicy":    "Equivalent",
					"timeoutSeconds": 30,
					"failurePolicy":  "Fail",
					"sideEffects":    "None",
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "cert-manager-webhook",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate",
						},
					},
				},
			},
		},
	}

	return mutate.MutateValidatingWebhookCertManagerWebhook(resourceObj, parent, reconciler, req)
}
