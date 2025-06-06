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

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDBundlesTrustCertManagerIo creates the CustomResourceDefinition resource with name bundles.trust.cert-manager.io.
func CreateCRDBundlesTrustCertManagerIo(
	parent *certificatesv1alpha1.TrustManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"name":        "bundles.trust.cert-manager.io",
				"annotations": map[string]interface{}{},
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
				"group": "trust.cert-manager.io",
				"names": map[string]interface{}{
					"kind":     "Bundle",
					"listKind": "BundleList",
					"plural":   "bundles",
					"singular": "bundle",
				},
				"scope": "Cluster",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Bundle ConfigMap Target Key",
								"jsonPath":    ".spec.target.configMap.key",
								"name":        "ConfigMap Target",
								"type":        "string",
							},
							map[string]interface{}{
								"description": "Bundle Secret Target Key",
								"jsonPath":    ".spec.target.secret.key",
								"name":        "Secret Target",
								"type":        "string",
							},
							map[string]interface{}{
								"description": "Bundle has been synced",
								"jsonPath":    ".status.conditions[?(@.type == \"Synced\")].status",
								"name":        "Synced",
								"type":        "string",
							},
							map[string]interface{}{
								"description": "Reason Bundle has Synced status",
								"jsonPath":    ".status.conditions[?(@.type == \"Synced\")].reason",
								"name":        "Reason",
								"type":        "string",
							},
							map[string]interface{}{
								"description": "Timestamp Bundle was created",
								"jsonPath":    ".metadata.creationTimestamp",
								"name":        "Age",
								"type":        "date",
							},
						},
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": `APIVersion defines the versioned schema of this representation of an object.
Servers should convert recognized schemas to the latest internal value, and
may reject unrecognized values.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources`,
										"type": "string",
									},
									"kind": map[string]interface{}{
										"description": `Kind is a string value representing the REST resource this object represents.
Servers may infer this from the endpoint the client submits requests to.
Cannot be updated.
In CamelCase.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds`,
										"type": "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "Desired state of the Bundle resource.",
										"properties": map[string]interface{}{
											"sources": map[string]interface{}{
												"description": "Sources is a set of references to data whose data will sync to the target.",
												"items": map[string]interface{}{
													"description": `BundleSource is the set of sources whose data will be appended and synced to
the BundleTarget in all Namespaces.`,
													"properties": map[string]interface{}{
														"configMap": map[string]interface{}{
															"description": `ConfigMap is a reference (by name) to a ConfigMap's ` + "`" + `data` + "`" + ` key(s), or to a
list of ConfigMap's ` + "`" + `data` + "`" + ` key(s) using label selector, in the trust Namespace.`,
															"properties": map[string]interface{}{
																"includeAllKeys": map[string]interface{}{
																	"description": `IncludeAllKeys is a flag to include all keys in the object's ` + "`" + `data` + "`" + ` field to be used. False by default.
This field must not be true when ` + "`" + `Key` + "`" + ` is set.`,
																	"type": "boolean",
																},
																"key": map[string]interface{}{
																	"description": "Key of the entry in the object's `data` field to be used.",
																	"minLength":   1,
																	"type":        "string",
																},
																"name": map[string]interface{}{
																	"description": `Name is the name of the source object in the trust Namespace.
This field must be left empty when ` + "`" + `selector` + "`" + ` is set`,
																	"minLength": 1,
																	"type":      "string",
																},
																"selector": map[string]interface{}{
																	"description": `Selector is the label selector to use to fetch a list of objects. Must not be set
when ` + "`" + `Name` + "`" + ` is set.`,
																	"properties": map[string]interface{}{
																		"matchExpressions": map[string]interface{}{
																			"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			"items": map[string]interface{}{
																				"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "key is the label key that the selector applies to.",
																						"type":        "string",
																					},
																					"operator": map[string]interface{}{
																						"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																						"type": "string",
																					},
																					"values": map[string]interface{}{
																						"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type":                   "array",
																						"x-kubernetes-list-type": "atomic",
																					},
																				},
																				"required": []interface{}{
																					"key",
																					"operator",
																				},
																				"type": "object",
																			},
																			"type":                   "array",
																			"x-kubernetes-list-type": "atomic",
																		},
																		"matchLabels": map[string]interface{}{
																			"additionalProperties": map[string]interface{}{
																				"type": "string",
																			},
																			"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																			"type": "object",
																		},
																	},
																	"type":                  "object",
																	"x-kubernetes-map-type": "atomic",
																},
															},
															"type":                  "object",
															"x-kubernetes-map-type": "atomic",
														},
														"inLine": map[string]interface{}{
															"description": "InLine is a simple string to append as the source data.",
															"type":        "string",
														},
														"secret": map[string]interface{}{
															"description": `Secret is a reference (by name) to a Secret's ` + "`" + `data` + "`" + ` key(s), or to a
list of Secret's ` + "`" + `data` + "`" + ` key(s) using label selector, in the trust Namespace.`,
															"properties": map[string]interface{}{
																"includeAllKeys": map[string]interface{}{
																	"description": `IncludeAllKeys is a flag to include all keys in the object's ` + "`" + `data` + "`" + ` field to be used. False by default.
This field must not be true when ` + "`" + `Key` + "`" + ` is set.`,
																	"type": "boolean",
																},
																"key": map[string]interface{}{
																	"description": "Key of the entry in the object's `data` field to be used.",
																	"minLength":   1,
																	"type":        "string",
																},
																"name": map[string]interface{}{
																	"description": `Name is the name of the source object in the trust Namespace.
This field must be left empty when ` + "`" + `selector` + "`" + ` is set`,
																	"minLength": 1,
																	"type":      "string",
																},
																"selector": map[string]interface{}{
																	"description": `Selector is the label selector to use to fetch a list of objects. Must not be set
when ` + "`" + `Name` + "`" + ` is set.`,
																	"properties": map[string]interface{}{
																		"matchExpressions": map[string]interface{}{
																			"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			"items": map[string]interface{}{
																				"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "key is the label key that the selector applies to.",
																						"type":        "string",
																					},
																					"operator": map[string]interface{}{
																						"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																						"type": "string",
																					},
																					"values": map[string]interface{}{
																						"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type":                   "array",
																						"x-kubernetes-list-type": "atomic",
																					},
																				},
																				"required": []interface{}{
																					"key",
																					"operator",
																				},
																				"type": "object",
																			},
																			"type":                   "array",
																			"x-kubernetes-list-type": "atomic",
																		},
																		"matchLabels": map[string]interface{}{
																			"additionalProperties": map[string]interface{}{
																				"type": "string",
																			},
																			"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																			"type": "object",
																		},
																	},
																	"type":                  "object",
																	"x-kubernetes-map-type": "atomic",
																},
															},
															"type":                  "object",
															"x-kubernetes-map-type": "atomic",
														},
														"useDefaultCAs": map[string]interface{}{
															"description": `UseDefaultCAs, when true, requests the default CA bundle to be used as a source.
Default CAs are available if trust-manager was installed via Helm
or was otherwise set up to include a package-injecting init container by using the
"--default-package-location" flag when starting the trust-manager controller.
If default CAs were not configured at start-up, any request to use the default
CAs will fail.
The version of the default CA package which is used for a Bundle is stored in the
defaultCAPackageVersion field of the Bundle's status field.`,
															"type": "boolean",
														},
													},
													"type":                  "object",
													"x-kubernetes-map-type": "atomic",
												},
												"maxItems":               100,
												"minItems":               1,
												"type":                   "array",
												"x-kubernetes-list-type": "atomic",
											},
											"target": map[string]interface{}{
												"description": "Target is the target location in all namespaces to sync source data to.",
												"properties": map[string]interface{}{
													"additionalFormats": map[string]interface{}{
														"description": "AdditionalFormats specifies any additional formats to write to the target",
														"properties": map[string]interface{}{
															"jks": map[string]interface{}{
																"description": `JKS requests a JKS-formatted binary trust bundle to be written to the target.
The bundle has "changeit" as the default password.
For more information refer to this link https://cert-manager.io/docs/faq/#keystore-passwords
Deprecated: Writing JKS is subject for removal. Please migrate to PKCS12.
PKCS#12 trust stores created by trust-manager are compatible with Java.`,
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "Key is the key of the entry in the object's `data` field to be used.",
																		"minLength":   1,
																		"type":        "string",
																	},
																	"password": map[string]interface{}{
																		"default":     "changeit",
																		"description": "Password for JKS trust store",
																		"maxLength":   128,
																		"minLength":   1,
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"key",
																},
																"type":                  "object",
																"x-kubernetes-map-type": "atomic",
															},
															"pkcs12": map[string]interface{}{
																"description": `PKCS12 requests a PKCS12-formatted binary trust bundle to be written to the target.

The bundle is by default created without a password.
For more information refer to this link https://cert-manager.io/docs/faq/#keystore-passwords`,
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "Key is the key of the entry in the object's `data` field to be used.",
																		"minLength":   1,
																		"type":        "string",
																	},
																	"password": map[string]interface{}{
																		"default":     "",
																		"description": "Password for PKCS12 trust store",
																		"maxLength":   128,
																		"type":        "string",
																	},
																	"profile": map[string]interface{}{
																		"description": `Profile specifies the certificate encryption algorithms and the HMAC algorithm
used to create the PKCS12 trust store.

If provided, allowed values are:
` + "`" + `LegacyRC2` + "`" + `: Deprecated. Not supported by default in OpenSSL 3 or Java 20.
` + "`" + `LegacyDES` + "`" + `: Less secure algorithm. Use this option for maximal compatibility.
` + "`" + `Modern2023` + "`" + `: Secure algorithm. Use this option in case you have to always use secure algorithms (e.g. because of company policy).

Default value is ` + "`" + `LegacyRC2` + "`" + ` for backward compatibility.`,
																		"enum": []interface{}{
																			"LegacyRC2",
																			"LegacyDES",
																			"Modern2023",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"key",
																},
																"type":                  "object",
																"x-kubernetes-map-type": "atomic",
															},
														},
														"type": "object",
													},
													"configMap": map[string]interface{}{
														"description": `ConfigMap is the target ConfigMap in Namespaces that all Bundle source
data will be synced to.`,
														"properties": map[string]interface{}{
															"key": map[string]interface{}{
																"description": "Key is the key of the entry in the object's `data` field to be used.",
																"minLength":   1,
																"type":        "string",
															},
															"metadata": map[string]interface{}{
																"description": "Metadata is an optional set of labels and annotations to be copied to the target.",
																"properties": map[string]interface{}{
																	"annotations": map[string]interface{}{
																		"additionalProperties": map[string]interface{}{
																			"type": "string",
																		},
																		"description": "Annotations is a key value map to be copied to the target.",
																		"type":        "object",
																	},
																	"labels": map[string]interface{}{
																		"additionalProperties": map[string]interface{}{
																			"type": "string",
																		},
																		"description": "Labels is a key value map to be copied to the target.",
																		"type":        "object",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"key",
														},
														"type": "object",
													},
													"namespaceSelector": map[string]interface{}{
														"description": `NamespaceSelector will, if set, only sync the target resource in
Namespaces which match the selector.`,
														"properties": map[string]interface{}{
															"matchExpressions": map[string]interface{}{
																"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																"items": map[string]interface{}{
																	"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																	"properties": map[string]interface{}{
																		"key": map[string]interface{}{
																			"description": "key is the label key that the selector applies to.",
																			"type":        "string",
																		},
																		"operator": map[string]interface{}{
																			"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																			"type": "string",
																		},
																		"values": map[string]interface{}{
																			"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																			"items": map[string]interface{}{
																				"type": "string",
																			},
																			"type":                   "array",
																			"x-kubernetes-list-type": "atomic",
																		},
																	},
																	"required": []interface{}{
																		"key",
																		"operator",
																	},
																	"type": "object",
																},
																"type":                   "array",
																"x-kubernetes-list-type": "atomic",
															},
															"matchLabels": map[string]interface{}{
																"additionalProperties": map[string]interface{}{
																	"type": "string",
																},
																"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																"type": "object",
															},
														},
														"type":                  "object",
														"x-kubernetes-map-type": "atomic",
													},
													"secret": map[string]interface{}{
														"description": `Secret is the target Secret that all Bundle source data will be synced to.
Using Secrets as targets is only supported if enabled at trust-manager startup.
By default, trust-manager has no permissions for writing to secrets and can only read secrets in the trust namespace.`,
														"properties": map[string]interface{}{
															"key": map[string]interface{}{
																"description": "Key is the key of the entry in the object's `data` field to be used.",
																"minLength":   1,
																"type":        "string",
															},
															"metadata": map[string]interface{}{
																"description": "Metadata is an optional set of labels and annotations to be copied to the target.",
																"properties": map[string]interface{}{
																	"annotations": map[string]interface{}{
																		"additionalProperties": map[string]interface{}{
																			"type": "string",
																		},
																		"description": "Annotations is a key value map to be copied to the target.",
																		"type":        "object",
																	},
																	"labels": map[string]interface{}{
																		"additionalProperties": map[string]interface{}{
																			"type": "string",
																		},
																		"description": "Labels is a key value map to be copied to the target.",
																		"type":        "object",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"key",
														},
														"type": "object",
													},
												},
												"type": "object",
											},
										},
										"required": []interface{}{
											"sources",
											"target",
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "Status of the Bundle. This is set and managed automatically.",
										"properties": map[string]interface{}{
											"conditions": map[string]interface{}{
												"description": `List of status conditions to indicate the status of the Bundle.
Known condition types are ` + "`" + `Bundle` + "`" + `.`,
												"items": map[string]interface{}{
													"description": "Condition contains details for one aspect of the current state of this API Resource.",
													"properties": map[string]interface{}{
														"lastTransitionTime": map[string]interface{}{
															"description": `lastTransitionTime is the last time the condition transitioned from one status to another.
This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.`,
															"format": "date-time",
															"type":   "string",
														},
														"message": map[string]interface{}{
															"description": `message is a human readable message indicating details about the transition.
This may be an empty string.`,
															"maxLength": 32768,
															"type":      "string",
														},
														"observedGeneration": map[string]interface{}{
															"description": `observedGeneration represents the .metadata.generation that the condition was set based upon.
For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date
with respect to the current state of the instance.`,
															"format":  "int64",
															"minimum": 0,
															"type":    "integer",
														},
														"reason": map[string]interface{}{
															"description": `reason contains a programmatic identifier indicating the reason for the condition's last transition.
Producers of specific condition types may define expected values and meanings for this field,
and whether the values are considered a guaranteed API.
The value should be a CamelCase string.
This field may not be empty.`,
															"maxLength": 1024,
															"minLength": 1,
															"pattern":   "^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$",
															"type":      "string",
														},
														"status": map[string]interface{}{
															"description": "status of the condition, one of True, False, Unknown.",
															"enum": []interface{}{
																"True",
																"False",
																"Unknown",
															},
															"type": "string",
														},
														"type": map[string]interface{}{
															"description": "type of condition in CamelCase or in foo.example.com/CamelCase.",
															"maxLength":   316,
															"pattern":     `^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`,
															"type":        "string",
														},
													},
													"required": []interface{}{
														"lastTransitionTime",
														"message",
														"reason",
														"status",
														"type",
													},
													"type": "object",
												},
												"type": "array",
												"x-kubernetes-list-map-keys": []interface{}{
													"type",
												},
												"x-kubernetes-list-type": "map",
											},
											"defaultCAVersion": map[string]interface{}{
												"description": `DefaultCAPackageVersion, if set and non-empty, indicates the version information
which was retrieved when the set of default CAs was requested in the bundle
source. This should only be set if useDefaultCAs was set to "true" on a source,
and will be the same for the same version of a bundle with identical certificates.`,
												"type": "string",
											},
										},
										"type": "object",
									},
								},
								"required": []interface{}{
									"spec",
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
			},
		},
	}

	return mutate.MutateCRDBundlesTrustCertManagerIo(resourceObj, parent, reconciler, req)
}
