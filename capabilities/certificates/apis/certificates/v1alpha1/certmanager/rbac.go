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

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceCertManagerCainjector creates the ServiceAccount resource with name cert-manager-cainjector.
func CreateServiceAccountNamespaceCertManagerCainjector(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion":                   "v1",
			"kind":                         "ServiceAccount",
			"automountServiceAccountToken": true,
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
		},
	}

	return mutate.MutateServiceAccountNamespaceCertManagerCainjector(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceCertManager creates the ServiceAccount resource with name cert-manager.
func CreateServiceAccountNamespaceCertManager(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion":                   "v1",
			"kind":                         "ServiceAccount",
			"automountServiceAccountToken": true,
			"metadata": map[string]interface{}{
				"name": "cert-manager",
				// controlled by field: aws.roleARN
				//  The AWS IAM Role ARN to use for validating public DNS records for issuing public certificates.
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
		},
	}

	return mutate.MutateServiceAccountNamespaceCertManager(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceCertManagerWebhook creates the ServiceAccount resource with name cert-manager-webhook.
func CreateServiceAccountNamespaceCertManagerWebhook(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion":                   "v1",
			"kind":                         "ServiceAccount",
			"automountServiceAccountToken": true,
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
		},
	}

	return mutate.MutateServiceAccountNamespaceCertManagerWebhook(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=events,verbs=get;create;update;patch
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=apiregistration.k8s.io,resources=apiservices,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;update;patch

// CreateClusterRoleCertManagerCainjector creates the ClusterRole resource with name cert-manager-cainjector.
func CreateClusterRoleCertManagerCainjector(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "cert-manager-cainjector",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificates",
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
						"secrets",
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
						"get",
						"create",
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"admissionregistration.k8s.io",
					},
					"resources": []interface{}{
						"validatingwebhookconfigurations",
						"mutatingwebhookconfigurations",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"apiregistration.k8s.io",
					},
					"resources": []interface{}{
						"apiservices",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"apiextensions.k8s.io",
					},
					"resources": []interface{}{
						"customresourcedefinitions",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"update",
						"patch",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCertManagerCainjector(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cert-manager.io,resources=issuers,verbs=update;patch;get;list;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=issuers/status,verbs=update;patch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;delete
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch

// CreateClusterRoleCertManagerControllerIssuers creates the ClusterRole resource with name cert-manager-controller-issuers.
func CreateClusterRoleCertManagerControllerIssuers(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-issuers",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"issuers",
						"issuers/status",
					},
					"verbs": []interface{}{
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"issuers",
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
						"secrets",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"create",
						"update",
						"delete",
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

	return mutate.MutateClusterRoleCertManagerControllerIssuers(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cert-manager.io,resources=clusterissuers,verbs=update;patch;get;list;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=clusterissuers/status,verbs=update;patch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;delete
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch

// CreateClusterRoleCertManagerControllerClusterissuers creates the ClusterRole resource with name cert-manager-controller-clusterissuers.
func CreateClusterRoleCertManagerControllerClusterissuers(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-clusterissuers",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"clusterissuers",
						"clusterissuers/status",
					},
					"verbs": []interface{}{
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"clusterissuers",
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
						"secrets",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"create",
						"update",
						"delete",
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

	return mutate.MutateClusterRoleCertManagerControllerClusterissuers(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=update;patch;get;list;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates/status,verbs=update;patch
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests,verbs=update;patch;get;list;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests/status,verbs=update;patch
// +kubebuilder:rbac:groups=cert-manager.io,resources=clusterissuers,verbs=get;list;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=issuers,verbs=get;list;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates/finalizers,verbs=update
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests/finalizers,verbs=update
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=orders,verbs=create;delete;get;list;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch

// CreateClusterRoleCertManagerControllerCertificates creates the ClusterRole resource with name cert-manager-controller-certificates.
func CreateClusterRoleCertManagerControllerCertificates(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-certificates",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificates",
						"certificates/status",
						"certificaterequests",
						"certificaterequests/status",
					},
					"verbs": []interface{}{
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificates",
						"certificaterequests",
						"clusterissuers",
						"issuers",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificates/finalizers",
						"certificaterequests/finalizers",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"orders",
					},
					"verbs": []interface{}{
						"create",
						"delete",
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
						"secrets",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"create",
						"update",
						"delete",
						"patch",
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

	return mutate.MutateClusterRoleCertManagerControllerCertificates(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=orders,verbs=update;patch;get;list;watch
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=orders/status,verbs=update;patch
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=challenges,verbs=get;list;watch;create;delete
// +kubebuilder:rbac:groups=cert-manager.io,resources=clusterissuers,verbs=get;list;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=issuers,verbs=get;list;watch
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=orders/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch

// CreateClusterRoleCertManagerControllerOrders creates the ClusterRole resource with name cert-manager-controller-orders.
func CreateClusterRoleCertManagerControllerOrders(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-orders",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"orders",
						"orders/status",
					},
					"verbs": []interface{}{
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"orders",
						"challenges",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"clusterissuers",
						"issuers",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"challenges",
					},
					"verbs": []interface{}{
						"create",
						"delete",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"orders/finalizers",
					},
					"verbs": []interface{}{
						"update",
					},
				},
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

	return mutate.MutateClusterRoleCertManagerControllerOrders(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=challenges,verbs=update;patch;get;list;watch
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=challenges/status,verbs=update;patch
// +kubebuilder:rbac:groups=cert-manager.io,resources=issuers,verbs=get;list;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=clusterissuers,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;delete
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;delete;update
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=httproutes,verbs=get;list;watch;create;delete;update
// +kubebuilder:rbac:groups=route.openshift.io,resources=routes/custom-host,verbs=create
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=challenges/finalizers,verbs=update

// CreateClusterRoleCertManagerControllerChallenges creates the ClusterRole resource with name cert-manager-controller-challenges.
func CreateClusterRoleCertManagerControllerChallenges(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-challenges",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"challenges",
						"challenges/status",
					},
					"verbs": []interface{}{
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"challenges",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"issuers",
						"clusterissuers",
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
						"secrets",
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
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"pods",
						"services",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"create",
						"delete",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.k8s.io",
					},
					"resources": []interface{}{
						"ingresses",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"create",
						"delete",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gateway.networking.k8s.io",
					},
					"resources": []interface{}{
						"httproutes",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"create",
						"delete",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"route.openshift.io",
					},
					"resources": []interface{}{
						"routes/custom-host",
					},
					"verbs": []interface{}{
						"create",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"challenges/finalizers",
					},
					"verbs": []interface{}{
						"update",
					},
				},
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

	return mutate.MutateClusterRoleCertManagerControllerChallenges(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=create;update;delete;get;list;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests,verbs=create;update;delete;get;list;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=issuers,verbs=get;list;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=clusterissuers,verbs=get;list;watch
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses/finalizers,verbs=update
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=gateways,verbs=get;list;watch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=httproutes,verbs=get;list;watch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=gateways/finalizers,verbs=update
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=httproutes/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch

// CreateClusterRoleCertManagerControllerIngressShim creates the ClusterRole resource with name cert-manager-controller-ingress-shim.
func CreateClusterRoleCertManagerControllerIngressShim(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-ingress-shim",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificates",
						"certificaterequests",
					},
					"verbs": []interface{}{
						"create",
						"update",
						"delete",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificates",
						"certificaterequests",
						"issuers",
						"clusterissuers",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.k8s.io",
					},
					"resources": []interface{}{
						"ingresses",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.k8s.io",
					},
					"resources": []interface{}{
						"ingresses/finalizers",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gateway.networking.k8s.io",
					},
					"resources": []interface{}{
						"gateways",
						"httproutes",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gateway.networking.k8s.io",
					},
					"resources": []interface{}{
						"gateways/finalizers",
						"httproutes/finalizers",
					},
					"verbs": []interface{}{
						"update",
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

	return mutate.MutateClusterRoleCertManagerControllerIngressShim(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cert-manager.io,resources=clusterissuers,verbs=get;list;watch

// CreateClusterRoleCertManagerClusterView creates the ClusterRole resource with name cert-manager-cluster-view.
func CreateClusterRoleCertManagerClusterView(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "cert-manager-cluster-view",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":                                "cert-manager",
					"app.kubernetes.io/version":                             "v1.17.2",
					"rbac.authorization.k8s.io/aggregate-to-cluster-reader": "true",
					"platform.klearwave.io/capability":                      "certificates",
					"platform.klearwave.io/component":                       "cert-manager",
					"platform.klearwave.io/operator-version":                "unstable",
					"platform.klearwave.io/platform-version":                "unstable",
					"app.kubernetes.io/part-of":                             "platform",
					"app.kubernetes.io/managed-by":                          "certificates-operator",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"clusterissuers",
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

	return mutate.MutateClusterRoleCertManagerClusterView(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=get;list;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests,verbs=get;list;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=issuers,verbs=get;list;watch
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=challenges,verbs=get;list;watch
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=orders,verbs=get;list;watch

// CreateClusterRoleCertManagerView creates the ClusterRole resource with name cert-manager-view.
func CreateClusterRoleCertManagerView(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "cert-manager-view",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":                                "cert-manager",
					"app.kubernetes.io/version":                             "v1.17.2",
					"rbac.authorization.k8s.io/aggregate-to-view":           "true",
					"rbac.authorization.k8s.io/aggregate-to-edit":           "true",
					"rbac.authorization.k8s.io/aggregate-to-admin":          "true",
					"rbac.authorization.k8s.io/aggregate-to-cluster-reader": "true",
					"platform.klearwave.io/capability":                      "certificates",
					"platform.klearwave.io/component":                       "cert-manager",
					"platform.klearwave.io/operator-version":                "unstable",
					"platform.klearwave.io/platform-version":                "unstable",
					"app.kubernetes.io/part-of":                             "platform",
					"app.kubernetes.io/managed-by":                          "certificates-operator",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificates",
						"certificaterequests",
						"issuers",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"challenges",
						"orders",
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

	return mutate.MutateClusterRoleCertManagerView(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=create;delete;deletecollection;patch;update
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests,verbs=create;delete;deletecollection;patch;update
// +kubebuilder:rbac:groups=cert-manager.io,resources=issuers,verbs=create;delete;deletecollection;patch;update
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates/status,verbs=update
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=challenges,verbs=create;delete;deletecollection;patch;update
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=orders,verbs=create;delete;deletecollection;patch;update

// CreateClusterRoleCertManagerEdit creates the ClusterRole resource with name cert-manager-edit.
func CreateClusterRoleCertManagerEdit(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "cert-manager-edit",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":                       "cert-manager",
					"app.kubernetes.io/version":                    "v1.17.2",
					"rbac.authorization.k8s.io/aggregate-to-edit":  "true",
					"rbac.authorization.k8s.io/aggregate-to-admin": "true",
					"platform.klearwave.io/capability":             "certificates",
					"platform.klearwave.io/component":              "cert-manager",
					"platform.klearwave.io/operator-version":       "unstable",
					"platform.klearwave.io/platform-version":       "unstable",
					"app.kubernetes.io/part-of":                    "platform",
					"app.kubernetes.io/managed-by":                 "certificates-operator",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificates",
						"certificaterequests",
						"issuers",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"deletecollection",
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificates/status",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acme.cert-manager.io",
					},
					"resources": []interface{}{
						"challenges",
						"orders",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"deletecollection",
						"patch",
						"update",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCertManagerEdit(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cert-manager.io,resources=signers,verbs=approve

// CreateClusterRoleCertManagerControllerApproveCertManagerIo creates the ClusterRole resource with name cert-manager-controller-approve:cert-manager-io.
func CreateClusterRoleCertManagerControllerApproveCertManagerIo(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-approve:cert-manager-io",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"signers",
					},
					"verbs": []interface{}{
						"approve",
					},
					"resourceNames": []interface{}{
						"issuers.cert-manager.io/*",
						"clusterissuers.cert-manager.io/*",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCertManagerControllerApproveCertManagerIo(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests,verbs=get;list;watch;update
// +kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests/status,verbs=update;patch
// +kubebuilder:rbac:groups=certificates.k8s.io,resources=signers,verbs=sign
// +kubebuilder:rbac:groups=authorization.k8s.io,resources=subjectaccessreviews,verbs=create

// CreateClusterRoleCertManagerControllerCertificatesigningrequests creates the ClusterRole resource with name cert-manager-controller-certificatesigningrequests.
func CreateClusterRoleCertManagerControllerCertificatesigningrequests(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-certificatesigningrequests",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.k8s.io",
					},
					"resources": []interface{}{
						"certificatesigningrequests",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.k8s.io",
					},
					"resources": []interface{}{
						"certificatesigningrequests/status",
					},
					"verbs": []interface{}{
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"certificates.k8s.io",
					},
					"resources": []interface{}{
						"signers",
					},
					"resourceNames": []interface{}{
						"issuers.cert-manager.io/*",
						"clusterissuers.cert-manager.io/*",
					},
					"verbs": []interface{}{
						"sign",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"authorization.k8s.io",
					},
					"resources": []interface{}{
						"subjectaccessreviews",
					},
					"verbs": []interface{}{
						"create",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCertManagerControllerCertificatesigningrequests(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=authorization.k8s.io,resources=subjectaccessreviews,verbs=create

// CreateClusterRoleCertManagerWebhookSubjectaccessreviews creates the ClusterRole resource with name cert-manager-webhook:subjectaccessreviews.
func CreateClusterRoleCertManagerWebhookSubjectaccessreviews(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "cert-manager-webhook:subjectaccessreviews",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"authorization.k8s.io",
					},
					"resources": []interface{}{
						"subjectaccessreviews",
					},
					"verbs": []interface{}{
						"create",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCertManagerWebhookSubjectaccessreviews(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingCertManagerCainjector creates the ClusterRoleBinding resource with name cert-manager-cainjector.
func CreateClusterRoleBindingCertManagerCainjector(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "cert-manager-cainjector",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "cert-manager-cainjector",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"name":      "cert-manager-cainjector",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
					"kind":      "ServiceAccount",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingCertManagerCainjector(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingCertManagerControllerIssuers creates the ClusterRoleBinding resource with name cert-manager-controller-issuers.
func CreateClusterRoleBindingCertManagerControllerIssuers(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-issuers",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "cert-manager-controller-issuers",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"name":      "cert-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
					"kind":      "ServiceAccount",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingCertManagerControllerIssuers(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingCertManagerControllerClusterissuers creates the ClusterRoleBinding resource with name cert-manager-controller-clusterissuers.
func CreateClusterRoleBindingCertManagerControllerClusterissuers(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-clusterissuers",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "cert-manager-controller-clusterissuers",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"name":      "cert-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
					"kind":      "ServiceAccount",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingCertManagerControllerClusterissuers(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingCertManagerControllerCertificates creates the ClusterRoleBinding resource with name cert-manager-controller-certificates.
func CreateClusterRoleBindingCertManagerControllerCertificates(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-certificates",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "cert-manager-controller-certificates",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"name":      "cert-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
					"kind":      "ServiceAccount",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingCertManagerControllerCertificates(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingCertManagerControllerOrders creates the ClusterRoleBinding resource with name cert-manager-controller-orders.
func CreateClusterRoleBindingCertManagerControllerOrders(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-orders",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "cert-manager-controller-orders",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"name":      "cert-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
					"kind":      "ServiceAccount",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingCertManagerControllerOrders(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingCertManagerControllerChallenges creates the ClusterRoleBinding resource with name cert-manager-controller-challenges.
func CreateClusterRoleBindingCertManagerControllerChallenges(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-challenges",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "cert-manager-controller-challenges",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"name":      "cert-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
					"kind":      "ServiceAccount",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingCertManagerControllerChallenges(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingCertManagerControllerIngressShim creates the ClusterRoleBinding resource with name cert-manager-controller-ingress-shim.
func CreateClusterRoleBindingCertManagerControllerIngressShim(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-ingress-shim",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "cert-manager-controller-ingress-shim",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"name":      "cert-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
					"kind":      "ServiceAccount",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingCertManagerControllerIngressShim(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingCertManagerControllerApproveCertManagerIo creates the ClusterRoleBinding resource with name cert-manager-controller-approve:cert-manager-io.
func CreateClusterRoleBindingCertManagerControllerApproveCertManagerIo(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-approve:cert-manager-io",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "cert-manager-controller-approve:cert-manager-io",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"name":      "cert-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
					"kind":      "ServiceAccount",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingCertManagerControllerApproveCertManagerIo(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingCertManagerControllerCertificatesigningrequests creates the ClusterRoleBinding resource with name cert-manager-controller-certificatesigningrequests.
func CreateClusterRoleBindingCertManagerControllerCertificatesigningrequests(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "cert-manager-controller-certificatesigningrequests",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "cert-manager-controller-certificatesigningrequests",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"name":      "cert-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
					"kind":      "ServiceAccount",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingCertManagerControllerCertificatesigningrequests(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingCertManagerWebhookSubjectaccessreviews creates the ClusterRoleBinding resource with name cert-manager-webhook:subjectaccessreviews.
func CreateClusterRoleBindingCertManagerWebhookSubjectaccessreviews(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "cert-manager-webhook:subjectaccessreviews",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "cert-manager-webhook:subjectaccessreviews",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "cert-manager-webhook",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingCertManagerWebhookSubjectaccessreviews(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;update;patch;create

// CreateRoleNamespaceCertManagerCainjectorLeaderelection creates the Role resource with name cert-manager-cainjector:leaderelection.
func CreateRoleNamespaceCertManagerCainjectorLeaderelection(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "cert-manager-cainjector:leaderelection",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"coordination.k8s.io",
					},
					"resources": []interface{}{
						"leases",
					},
					"resourceNames": []interface{}{
						"cert-manager-cainjector-leader-election",
						"cert-manager-cainjector-leader-election-core",
					},
					"verbs": []interface{}{
						"get",
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"coordination.k8s.io",
					},
					"resources": []interface{}{
						"leases",
					},
					"verbs": []interface{}{
						"create",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNamespaceCertManagerCainjectorLeaderelection(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;update;patch;create

// CreateRoleNamespaceCertManagerLeaderelection creates the Role resource with name cert-manager:leaderelection.
func CreateRoleNamespaceCertManagerLeaderelection(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "cert-manager:leaderelection",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"coordination.k8s.io",
					},
					"resources": []interface{}{
						"leases",
					},
					"resourceNames": []interface{}{
						"cert-manager-controller",
					},
					"verbs": []interface{}{
						"get",
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"coordination.k8s.io",
					},
					"resources": []interface{}{
						"leases",
					},
					"verbs": []interface{}{
						"create",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNamespaceCertManagerLeaderelection(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=serviceaccounts/token,verbs=create

// CreateRoleNamespaceCertManagerTokenrequest creates the Role resource with name cert-manager-tokenrequest.
func CreateRoleNamespaceCertManagerTokenrequest(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "cert-manager-tokenrequest",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"serviceaccounts/token",
					},
					"resourceNames": []interface{}{
						"cert-manager",
					},
					"verbs": []interface{}{
						"create",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNamespaceCertManagerTokenrequest(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;update;create

// CreateRoleNamespaceCertManagerWebhookDynamicServing creates the Role resource with name cert-manager-webhook:dynamic-serving.
func CreateRoleNamespaceCertManagerWebhookDynamicServing(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "cert-manager-webhook:dynamic-serving",
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
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"secrets",
					},
					"resourceNames": []interface{}{
						"cert-manager-webhook-ca",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"secrets",
					},
					"verbs": []interface{}{
						"create",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNamespaceCertManagerWebhookDynamicServing(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNamespaceCertManagerCainjectorLeaderelection creates the RoleBinding resource with name cert-manager-cainjector:leaderelection.
func CreateRoleBindingNamespaceCertManagerCainjectorLeaderelection(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "cert-manager-cainjector:leaderelection",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "Role",
				"name":     "cert-manager-cainjector:leaderelection",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "cert-manager-cainjector",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingNamespaceCertManagerCainjectorLeaderelection(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNamespaceCertManagerLeaderelection creates the RoleBinding resource with name cert-manager:leaderelection.
func CreateRoleBindingNamespaceCertManagerLeaderelection(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "cert-manager:leaderelection",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "Role",
				"name":     "cert-manager:leaderelection",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "cert-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingNamespaceCertManagerLeaderelection(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNamespaceCertManagerCertManagerTokenrequest creates the RoleBinding resource with name cert-manager-cert-manager-tokenrequest.
func CreateRoleBindingNamespaceCertManagerCertManagerTokenrequest(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "cert-manager-cert-manager-tokenrequest",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "Role",
				"name":     "cert-manager-tokenrequest",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "cert-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingNamespaceCertManagerCertManagerTokenrequest(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNamespaceCertManagerWebhookDynamicServing creates the RoleBinding resource with name cert-manager-webhook:dynamic-serving.
func CreateRoleBindingNamespaceCertManagerWebhookDynamicServing(
	parent *certificatesv1alpha1.CertManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "cert-manager-webhook:dynamic-serving",
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
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "Role",
				"name":     "cert-manager-webhook:dynamic-serving",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "cert-manager-webhook",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingNamespaceCertManagerWebhookDynamicServing(resourceObj, parent, reconciler, req)
}
