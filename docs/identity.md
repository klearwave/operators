# Summary

This document describes the identity capability.


## Components

- [aws-pod-identity-webhook](https://github.com/aws/amazon-eks-pod-identity-webhook)


## Use Cases

### Workload Authentication

The application platform comes with the ability, via the cloud provider xKS service, to 
allow pods to authenticate against the cloud to interact with cloud services via their 
service account token.  This is done via a federation that allows an external identity provider 
to sign tokens.  Each cloud provider provides their own mechanism to do this.


### User Authentication

For day 1, the Klearwave application platform will be integrated for user authentication with 
the cloud provider identity provider.

For day 2, the Klearwave application platform will allow the ability to add a custom 
identity provider (e.g. OIDC, LDAP) to integrate with other organization authentication 
systems.
