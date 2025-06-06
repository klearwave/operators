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

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceTrustManager creates the ServiceAccount resource with name trust-manager.
func CreateServiceAccountNamespaceTrustManager(
	parent *certificatesv1alpha1.TrustManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion":                   "v1",
			"kind":                         "ServiceAccount",
			"automountServiceAccountToken": true,
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
		},
	}

	return mutate.MutateServiceAccountNamespaceTrustManager(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=trust.cert-manager.io,resources=bundles,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=trust.cert-manager.io,resources=bundles/finalizers,verbs=update
// +kubebuilder:rbac:groups=trust.cert-manager.io,resources=bundles/status,verbs=patch
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;create;patch;watch;delete
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch

// CreateClusterRoleTrustManager creates the ClusterRole resource with name trust-manager.
func CreateClusterRoleTrustManager(
	parent *certificatesv1alpha1.TrustManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
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
				"name": "trust-manager",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"trust.cert-manager.io",
					},
					"resources": []interface{}{
						"bundles",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"trust.cert-manager.io",
					},
					"resources": []interface{}{
						"bundles/finalizers",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"trust.cert-manager.io",
					},
					"resources": []interface{}{
						"bundles/status",
					},
					"verbs": []interface{}{
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"configmaps",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"create",
						"patch",
						"watch",
						"delete",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"namespaces",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"events",
					},
					"verbs": []interface{}{
						"create",
						"patch",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleTrustManager(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingTrustManager creates the ClusterRoleBinding resource with name trust-manager.
func CreateClusterRoleBindingTrustManager(
	parent *certificatesv1alpha1.TrustManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
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
				"name": "trust-manager",
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "trust-manager",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "trust-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingTrustManager(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch

// CreateRoleNamespaceTrustManager creates the Role resource with name trust-manager.
func CreateRoleNamespaceTrustManager(
	parent *certificatesv1alpha1.TrustManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Role",
			"apiVersion": "rbac.authorization.k8s.io/v1",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"secrets",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNamespaceTrustManager(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;create;update;watch;list

// CreateRoleNamespaceTrustManagerLeaderelection creates the Role resource with name trust-manager:leaderelection.
func CreateRoleNamespaceTrustManagerLeaderelection(
	parent *certificatesv1alpha1.TrustManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Role",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name":      "trust-manager:leaderelection",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"coordination.k8s.io",
					},
					"resources": []interface{}{
						"leases",
					},
					"verbs": []interface{}{
						"get",
						"create",
						"update",
						"watch",
						"list",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNamespaceTrustManagerLeaderelection(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNamespaceTrustManager creates the RoleBinding resource with name trust-manager.
func CreateRoleBindingNamespaceTrustManager(
	parent *certificatesv1alpha1.TrustManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "RoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "Role",
				"name":     "trust-manager",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "trust-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingNamespaceTrustManager(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNamespaceTrustManagerLeaderelection creates the RoleBinding resource with name trust-manager:leaderelection.
func CreateRoleBindingNamespaceTrustManagerLeaderelection(
	parent *certificatesv1alpha1.TrustManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "RoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name":      "trust-manager:leaderelection",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "Role",
				"name":     "trust-manager:leaderelection",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "trust-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingNamespaceTrustManagerLeaderelection(resourceObj, parent, reconciler, req)
}
