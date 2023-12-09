# OAuth Server Implementation with Role Based Access Control

## Technical Overview

#### Technologies Used

- Golang
- SQLite
- Open Policy Agent (OPA)

### Breif Introduction to OAuth

OAuth is an open standard for access delegation, commonly used as a way for Internet users to grant websites or applications access to their information on other websites but without giving them the passwords.

#### OAuth Roles

- **Resource Owner:** An entity capable of granting access to a protected resource. When the resource owner is a person, it is referred to as an end-user.
- **Resource Server:** The server hosting the protected resources, capable of accepting and responding to protected resource requests using access tokens.
- **Client:** An application making protected resource requests on behalf of the resource owner and with its authorization. The term "client" does not imply any particular implementation characteristics (e.g., whether the application executes on a server, a desktop, or other devices).

Designed specifically to work with Hypertext Transfer Protocol (HTTP), OAuth essentially allows access tokens to be issued to third-party clients by an authorization server, with the approval of the resource owner, or end-user. The client then uses the access token to access the protected resources hosted by the resource server.

### OAuth Scopes or Roles

- OAuth scopes let you specify exactly how your app needs to access a resource owned by a user. Scopes limit access for OAuth tokens. They do not grant any permission themselves.
- In this project we specify roles in the scope parameter of the OAuth token request. Each token has a role associated with it. The role is used to grant access to the resources.
- Resource server checks the role of the token and grants access to the resource if the role is authorized to access the resource.

### Open Policy Agent

- Open Policy Agent (OPA) is an open source, general-purpose policy engine that enables unified, context-aware policy enforcement across the entire stack.
- OPA provides a high-level declarative language that let's you specify policy as code and simple APIs to offload policy decision-making from your software.
- This project uses OPA to implement Role Based Access Control (RBAC) for the Resource Server.
