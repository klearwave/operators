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

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceTrustManager creates the Deployment resource with name trust-manager.
func CreateDeploymentNamespaceTrustManager(
	parent *certificatesv1alpha1.TrustManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
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
				// controlled by field: controller.replicas
				//  Number of replicas to use for the trust-manager controller deployment.
				"replicas": parent.Spec.Controller.Replicas,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"platform.klearwave.io/capability":       "certificates",
						"platform.klearwave.io/component":        "trust-manager",
						"platform.klearwave.io/operator-version": "unstable",
						"platform.klearwave.io/platform-version": "unstable",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
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
						"serviceAccountName":           "trust-manager",
						"automountServiceAccountToken": true,
						"initContainers": []interface{}{
							map[string]interface{}{
								"name":            "cert-manager-package-debian",
								"image":           "quay.io/jetstack/trust-pkg-debian-bookworm:20230311.0",
								"imagePullPolicy": "IfNotPresent",
								"args": []interface{}{
									"/copyandmaybepause",
									"/debian-package",
									"/packages",
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"mountPath": "/packages",
										"name":      "packages",
										"readOnly":  false,
									},
								},
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
									"readOnlyRootFilesystem": true,
									"runAsNonRoot":           true,
									"seccompProfile": map[string]interface{}{
										"type": "RuntimeDefault",
									},
								},
							},
						},
						"containers": []interface{}{
							map[string]interface{}{
								"name": "trust-manager",
								// controlled by field: controller.image
								//  Image to use for trust-manager controller deployment.
								"image":           parent.Spec.Controller.Image,
								"imagePullPolicy": "IfNotPresent",
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 6443,
										"name":          "webhook",
									},
									map[string]interface{}{
										"containerPort": 9402,
										"name":          "metrics",
									},
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"port": 6060,
										"path": "/readyz",
									},
									"initialDelaySeconds": 3,
									"periodSeconds":       7,
								},
								"args": []interface{}{
									"--log-format=text",
									"--log-level=1",
									"--metrics-port=9402",
									"--readiness-probe-port=6060",
									"--readiness-probe-path=/readyz",
									"--leader-elect=true",
									"--leader-election-lease-duration=15s",
									"--leader-election-renew-deadline=10s",
									"--trust-namespace=" + parent.Spec.Namespace + "", //  controlled by field: namespace
									"--webhook-host=0.0.0.0",
									"--webhook-port=6443",
									"--webhook-certificate-dir=/tls",
									"--default-package-location=/packages/cert-manager-package-debian.json",
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"mountPath": "/tls",
										"name":      "tls",
										"readOnly":  true,
									},
									map[string]interface{}{
										"mountPath": "/packages",
										"name":      "packages",
										"readOnly":  true,
									},
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										// controlled by field: controller.resources.requests.cpu
										//  CPU requests to use for trust-manager controller deployment.
										"cpu": parent.Spec.Controller.Resources.Requests.Cpu,
										// controlled by field: controller.resources.requests.memory
										//  Memory requests to use for trust-manager controller deployment.
										"memory": parent.Spec.Controller.Resources.Requests.Memory,
									},
									"limits": map[string]interface{}{
										// controlled by field: controller.resources.limits.memory
										//  Memory limits to use for trust-manager controller deployment.
										"memory": parent.Spec.Controller.Resources.Limits.Memory,
									},
								},
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"readOnlyRootFilesystem":   true,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os":                "linux",
							"platform.klearwave.io/node-type": "platform",
						},
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "packages",
								"emptyDir": map[string]interface{}{
									"sizeLimit": "50M",
								},
							},
							map[string]interface{}{
								"name": "tls",
								"secret": map[string]interface{}{
									"defaultMode": 420,
									"secretName":  "trust-manager-cert",
								},
							},
						},
						"securityContext": map[string]interface{}{
							"fsGroup":      1001,
							"runAsUser":    1001,
							"runAsGroup":   1001,
							"runAsNonRoot": true,
						},
						"affinity": map[string]interface{}{
							"podAntiAffinity": map[string]interface{}{
								"preferredDuringSchedulingIgnoredDuringExecution": []interface{}{
									map[string]interface{}{
										"weight": 100,
										"podAffinityTerm": map[string]interface{}{
											"topologyKey": "kubernetes.io/hostname",
											"labelSelector": map[string]interface{}{
												"matchExpressions": []interface{}{
													map[string]interface{}{
														"key":      "app.kubernetes.io/name",
														"operator": "In",
														"values": []interface{}{
															"trust-manager",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceTrustManager(resourceObj, parent, reconciler, req)
}
