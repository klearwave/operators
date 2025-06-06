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

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceCertManagerCainjector creates the Deployment resource with name cert-manager-cainjector.
func CreateDeploymentNamespaceCertManagerCainjector(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
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
				"replicas": 1,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":                 "cainjector",
						"platform.klearwave.io/capability":       "certificates",
						"platform.klearwave.io/component":        "cert-manager",
						"platform.klearwave.io/operator-version": "unstable",
						"platform.klearwave.io/platform-version": "unstable",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
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
						"annotations": map[string]interface{}{
							"prometheus.io/path":   "/metrics",
							"prometheus.io/scrape": "true",
							"prometheus.io/port":   "9402",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName": "cert-manager-cainjector",
						"enableServiceLinks": false,
						"securityContext": map[string]interface{}{
							"fsGroup":      1001,
							"runAsUser":    1001,
							"runAsGroup":   1001,
							"runAsNonRoot": true,
						},
						"containers": []interface{}{
							map[string]interface{}{
								"name": "cert-manager-cainjector",
								// controlled by field: injector.image
								//  Image to use for cert-manager CA injector deployment.
								"image":           parent.Spec.Injector.Image,
								"imagePullPolicy": "IfNotPresent",
								"args": []interface{}{
									"--v=2",
									"--leader-election-namespace=$(POD_NAMESPACE)",
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 9402,
										"name":          "http-metrics",
										"protocol":      "TCP",
									},
								},
								"env": []interface{}{
									map[string]interface{}{
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
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
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										// controlled by field: injector.resources.requests.cpu
										//  CPU requests to use for cert-manager CA injector deployment.
										"cpu": parent.Spec.Injector.Resources.Requests.Cpu,
										// controlled by field: injector.resources.requests.memory
										//  Memory requests to use for cert-manager CA injector deployment.
										"memory": parent.Spec.Injector.Resources.Requests.Memory,
									},
									"limits": map[string]interface{}{
										// controlled by field: injector.resources.limits.memory
										//  Memory limits to use for cert-manager CA injector deployment.
										"memory": parent.Spec.Injector.Resources.Limits.Memory,
									},
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os":                "linux",
							"platform.klearwave.io/node-type": "platform",
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceCertManagerCainjector(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceCertManager creates the Deployment resource with name cert-manager.
func CreateDeploymentNamespaceCertManager(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
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
				"replicas": 1,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":                 "cert-manager",
						"platform.klearwave.io/capability":       "certificates",
						"platform.klearwave.io/component":        "cert-manager",
						"platform.klearwave.io/operator-version": "unstable",
						"platform.klearwave.io/platform-version": "unstable",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
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
						"annotations": map[string]interface{}{
							"prometheus.io/path":   "/metrics",
							"prometheus.io/scrape": "true",
							"prometheus.io/port":   "9402",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName": "cert-manager",
						"enableServiceLinks": false,
						"securityContext": map[string]interface{}{
							"fsGroup":      1001,
							"runAsUser":    1001,
							"runAsGroup":   1001,
							"runAsNonRoot": true,
						},
						"containers": []interface{}{
							map[string]interface{}{
								"name": "cert-manager-controller",
								// controlled by field: controller.image
								//  Image to use for cert-manager controller deployment.
								"image":           parent.Spec.Controller.Image,
								"imagePullPolicy": "IfNotPresent",
								"args": []interface{}{
									"--v=2",
									"--cluster-resource-namespace=$(POD_NAMESPACE)",
									"--leader-election-namespace=$(POD_NAMESPACE)",
									"--acme-http01-solver-image=quay.io/jetstack/cert-manager-acmesolver:v1.17.2",
									"--max-concurrent-challenges=60",
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 9402,
										"name":          "http-metrics",
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"containerPort": 9403,
										"name":          "http-healthz",
										"protocol":      "TCP",
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
								"env": []interface{}{
									map[string]interface{}{
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
								},
								"livenessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"port":   "http-healthz",
										"path":   "/livez",
										"scheme": "HTTP",
									},
									"initialDelaySeconds": 10,
									"periodSeconds":       10,
									"timeoutSeconds":      15,
									"successThreshold":    1,
									"failureThreshold":    8,
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										// controlled by field: controller.resources.requests.cpu
										//  CPU requests to use for cert-manager controller deployment.
										"cpu": parent.Spec.Controller.Resources.Requests.Cpu,
										// controlled by field: controller.resources.requests.memory
										//  Memory requests to use for cert-manager controller deployment.
										"memory": parent.Spec.Controller.Resources.Requests.Memory,
									},
									"limits": map[string]interface{}{
										// controlled by field: controller.resources.limits.memory
										//  Memory limits to use for cert-manager controller deployment.
										"memory": parent.Spec.Controller.Resources.Limits.Memory,
									},
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os":                "linux",
							"platform.klearwave.io/node-type": "platform",
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
															"cert-manager",
														},
													},
												},
											},
										},
									},
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
															"cert-manager-cainjector",
														},
													},
												},
											},
										},
									},
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
															"cert-manager-webhook",
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

	return mutate.MutateDeploymentNamespaceCertManager(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceCertManagerWebhook creates the Deployment resource with name cert-manager-webhook.
func CreateDeploymentNamespaceCertManagerWebhook(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
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
				"replicas": 1,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":                 "webhook",
						"platform.klearwave.io/capability":       "certificates",
						"platform.klearwave.io/component":        "cert-manager",
						"platform.klearwave.io/operator-version": "unstable",
						"platform.klearwave.io/platform-version": "unstable",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
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
							"prometheus.io/path":   "/metrics",
							"prometheus.io/scrape": "true",
							"prometheus.io/port":   "9402",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName": "cert-manager-webhook",
						"enableServiceLinks": false,
						"securityContext": map[string]interface{}{
							"fsGroup":      1001,
							"runAsUser":    1001,
							"runAsGroup":   1001,
							"runAsNonRoot": true,
						},
						"containers": []interface{}{
							map[string]interface{}{
								"name": "cert-manager-webhook",
								// controlled by field: webhook.image
								//  Image to use for cert-manager webhook deployment.
								"image":           parent.Spec.Webhook.Image,
								"imagePullPolicy": "IfNotPresent",
								"args": []interface{}{
									"--v=2",
									"--secure-port=10250",
									"--dynamic-serving-ca-secret-namespace=$(POD_NAMESPACE)",
									"--dynamic-serving-ca-secret-name=cert-manager-webhook-ca",
									"--dynamic-serving-dns-names=cert-manager-webhook",
									"--dynamic-serving-dns-names=cert-manager-webhook.$(POD_NAMESPACE)",
									"--dynamic-serving-dns-names=cert-manager-webhook.$(POD_NAMESPACE).svc",
								},
								"ports": []interface{}{
									map[string]interface{}{
										"name":          "https",
										"protocol":      "TCP",
										"containerPort": 10250,
									},
									map[string]interface{}{
										"name":          "healthcheck",
										"protocol":      "TCP",
										"containerPort": 6080,
									},
									map[string]interface{}{
										"containerPort": 9402,
										"name":          "http-metrics",
										"protocol":      "TCP",
									},
								},
								"livenessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path":   "/livez",
										"port":   6080,
										"scheme": "HTTP",
									},
									"initialDelaySeconds": 60,
									"periodSeconds":       10,
									"timeoutSeconds":      1,
									"successThreshold":    1,
									"failureThreshold":    3,
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path":   "/healthz",
										"port":   6080,
										"scheme": "HTTP",
									},
									"initialDelaySeconds": 5,
									"periodSeconds":       5,
									"timeoutSeconds":      1,
									"successThreshold":    1,
									"failureThreshold":    3,
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
								"env": []interface{}{
									map[string]interface{}{
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										// controlled by field: webhook.resources.requests.cpu
										//  CPU requests to use for cert-manager webhook deployment.
										"cpu": parent.Spec.Webhook.Resources.Requests.Cpu,
										// controlled by field: webhook.resources.requests.memory
										//  Memory requests to use for cert-manager webhook deployment.
										"memory": parent.Spec.Webhook.Resources.Requests.Memory,
									},
									"limits": map[string]interface{}{
										// controlled by field: webhook.resources.limits.memory
										//  Memory limits to use for cert-manager webhook deployment.
										"memory": parent.Spec.Webhook.Resources.Limits.Memory,
									},
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os":                "linux",
							"platform.klearwave.io/node-type": "platform",
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceCertManagerWebhook(resourceObj, parent, reconciler, req)
}
