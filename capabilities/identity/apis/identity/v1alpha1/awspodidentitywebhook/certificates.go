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

// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=get;list;watch;create;update;patch;delete

// CreateCertNamespaceAwsPodIdentityWebhook creates the Certificate resource with name aws-pod-identity-webhook.
func CreateCertNamespaceAwsPodIdentityWebhook(
	parent *identityv1alpha1.AWSPodIdentityWebhook,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "cert-manager.io/v1",
			"kind":       "Certificate",
			"metadata": map[string]interface{}{
				"name":      "aws-pod-identity-webhook",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"secretName": "aws-pod-identity-webhook-cert",
				"commonName": "aws-pod-identity-webhook.klearwave-identity-system.svc",
				"dnsNames": []interface{}{
					"aws-pod-identity-webhook",
					"aws-pod-identity-webhook.klearwave-identity-system",
					"aws-pod-identity-webhook.klearwave-identity-system.svc",
					"aws-pod-identity-webhook.klearwave-identity-system.svc.local",
				},
				"isCA":        true,
				"duration":    "2160h",
				"renewBefore": "360h",
				"issuerRef": map[string]interface{}{
					"name": "internal",
					"kind": "ClusterIssuer",
				},
			},
		},
	}

	return mutate.MutateCertNamespaceAwsPodIdentityWebhook(resourceObj, parent, reconciler, req)
}
