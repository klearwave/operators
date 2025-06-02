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

package awspodidentitywebhook

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	identityv1alpha1 "github.com/klearwave/operators-capability/capabilities/identity/apis/identity/v1alpha1"
	"github.com/klearwave/operators-capability/capabilities/identity/apis/identity/v1alpha1/awspodidentitywebhook/mutate"
)

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateMutatingWebhookNamespaceAwsPodIdentityWebhook creates the MutatingWebhookConfiguration resource with name aws-pod-identity-webhook.
func CreateMutatingWebhookNamespaceAwsPodIdentityWebhook(
	parent *identityv1alpha1.AWSPodIdentityWebhook,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "MutatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name":      "aws-pod-identity-webhook",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"annotations": map[string]interface{}{
					// controlled by field: namespace
					"cert-manager.io/inject-ca-from": "" + parent.Spec.Namespace + "/aws-pod-identity-webhook-cert",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"name":          "pod-identity-webhook.amazonaws.com",
					"failurePolicy": "Ignore",
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "aws-pod-identity-webhook",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/mutate",
						},
					},
					"objectSelector": map[string]interface{}{
						"matchExpressions": []interface{}{
							map[string]interface{}{
								"key":      "eks.amazonaws.com/skip-pod-identity-webhook",
								"operator": "DoesNotExist",
								"values":   []interface{}{},
							},
						},
					},
					"rules": []interface{}{
						map[string]interface{}{
							"operations": []interface{}{
								"CREATE",
							},
							"apiGroups": []interface{}{
								"",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"resources": []interface{}{
								"pods",
							},
						},
					},
					"sideEffects": "None",
					"admissionReviewVersions": []interface{}{
						"v1beta1",
					},
				},
			},
		},
	}

	return mutate.MutateMutatingWebhookNamespaceAwsPodIdentityWebhook(resourceObj, parent, reconciler, req)
}
