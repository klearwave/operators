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

package constants

// this package includes the constants which include the resource names.  it is a standalone
// package to prevent import cycle errors when attempting to reference the names from other
// packages (e.g. mutate).
const (
	CRDCertificaterequestsCertManagerIo                               = "certificaterequests.cert-manager.io"
	CRDCertificatesCertManagerIo                                      = "certificates.cert-manager.io"
	CRDChallengesAcmeCertManagerIo                                    = "challenges.acme.cert-manager.io"
	CRDClusterissuersCertManagerIo                                    = "clusterissuers.cert-manager.io"
	CRDIssuersCertManagerIo                                           = "issuers.cert-manager.io"
	CRDOrdersAcmeCertManagerIo                                        = "orders.acme.cert-manager.io"
	ServiceAccountNamespaceCertManagerCainjector                      = "cert-manager-cainjector"
	ServiceAccountNamespaceCertManager                                = "cert-manager"
	ServiceAccountNamespaceCertManagerWebhook                         = "cert-manager-webhook"
	ClusterRoleCertManagerCainjector                                  = "cert-manager-cainjector"
	ClusterRoleCertManagerControllerIssuers                           = "cert-manager-controller-issuers"
	ClusterRoleCertManagerControllerClusterissuers                    = "cert-manager-controller-clusterissuers"
	ClusterRoleCertManagerControllerCertificates                      = "cert-manager-controller-certificates"
	ClusterRoleCertManagerControllerOrders                            = "cert-manager-controller-orders"
	ClusterRoleCertManagerControllerChallenges                        = "cert-manager-controller-challenges"
	ClusterRoleCertManagerControllerIngressShim                       = "cert-manager-controller-ingress-shim"
	ClusterRoleCertManagerClusterView                                 = "cert-manager-cluster-view"
	ClusterRoleCertManagerView                                        = "cert-manager-view"
	ClusterRoleCertManagerEdit                                        = "cert-manager-edit"
	ClusterRoleCertManagerControllerApproveCertManagerIo              = "cert-manager-controller-approve:cert-manager-io"
	ClusterRoleCertManagerControllerCertificatesigningrequests        = "cert-manager-controller-certificatesigningrequests"
	ClusterRoleCertManagerWebhookSubjectaccessreviews                 = "cert-manager-webhook:subjectaccessreviews"
	ClusterRoleBindingCertManagerCainjector                           = "cert-manager-cainjector"
	ClusterRoleBindingCertManagerControllerIssuers                    = "cert-manager-controller-issuers"
	ClusterRoleBindingCertManagerControllerClusterissuers             = "cert-manager-controller-clusterissuers"
	ClusterRoleBindingCertManagerControllerCertificates               = "cert-manager-controller-certificates"
	ClusterRoleBindingCertManagerControllerOrders                     = "cert-manager-controller-orders"
	ClusterRoleBindingCertManagerControllerChallenges                 = "cert-manager-controller-challenges"
	ClusterRoleBindingCertManagerControllerIngressShim                = "cert-manager-controller-ingress-shim"
	ClusterRoleBindingCertManagerControllerApproveCertManagerIo       = "cert-manager-controller-approve:cert-manager-io"
	ClusterRoleBindingCertManagerControllerCertificatesigningrequests = "cert-manager-controller-certificatesigningrequests"
	ClusterRoleBindingCertManagerWebhookSubjectaccessreviews          = "cert-manager-webhook:subjectaccessreviews"
	RoleNamespaceCertManagerCainjectorLeaderelection                  = "cert-manager-cainjector:leaderelection"
	RoleNamespaceCertManagerLeaderelection                            = "cert-manager:leaderelection"
	RoleNamespaceCertManagerTokenrequest                              = "cert-manager-tokenrequest"
	RoleNamespaceCertManagerWebhookDynamicServing                     = "cert-manager-webhook:dynamic-serving"
	RoleBindingNamespaceCertManagerCainjectorLeaderelection           = "cert-manager-cainjector:leaderelection"
	RoleBindingNamespaceCertManagerLeaderelection                     = "cert-manager:leaderelection"
	RoleBindingNamespaceCertManagerCertManagerTokenrequest            = "cert-manager-cert-manager-tokenrequest"
	RoleBindingNamespaceCertManagerWebhookDynamicServing              = "cert-manager-webhook:dynamic-serving"
	DeploymentNamespaceCertManagerCainjector                          = "cert-manager-cainjector"
	DeploymentNamespaceCertManager                                    = "cert-manager"
	DeploymentNamespaceCertManagerWebhook                             = "cert-manager-webhook"
	ServiceNamespaceCertManagerCainjector                             = "cert-manager-cainjector"
	ServiceNamespaceCertManager                                       = "cert-manager"
	ServiceNamespaceCertManagerWebhook                                = "cert-manager-webhook"
	ClusterIssuerInternal                                             = "internal"
	MutatingWebhookCertManagerWebhook                                 = "cert-manager-webhook"
	ValidatingWebhookCertManagerWebhook                               = "cert-manager-webhook"
)
