# Summary

This document describes the certicates capability.


## Components

- [cert-manager](https://cert-manager.io/docs/)
- [trust-manager](https://cert-manager.io/docs/trust/trust-manager/)


## Architecture

TBD


## Use Cases

### Self-Signed Certificate Authority

For simple use cases and in scenarios where development and testing is occurring, a 
self-signed certificate is a good choice.  This is the default option if no other 
options are passed.

**Related CLI Options**:

* `--certificates-internal-ca-cert`
* `--certificates-internal-ca-key`
* `--certificates-ingress-ca-cert`
* `--certificates-ingress-ca-key`


### Custom Certificate Authority

For custom use cases where a organization-specific certificate authority is required.  You 
may pass the certificate authority certificate and key.

**Related CLI Options**:

* `--certificates-internal-ca-cert`
* `--certificates-internal-ca-key`
* `--certificates-ingress-ca-cert`
* `--certificates-ingress-ca-key`


### LetsEncrypt Certificate Authority

For custom use cases where a public certificate is required.

**Related CLI Options**:

* `--certificates-ingress-ca-cert`
* `--certificates-ingress-ca-key`
