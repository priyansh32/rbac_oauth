# **RBAC Rego Policy Documentation**

## **Overview**

This document provides an overview and documentation for the RBAC (Role-Based Access Control) policy implemented using Rego in the resource server.

## **Rego Policy File: `roles.rego`**

### **Purpose**

The **`roles.rego`** file defines access control policies based on user roles for different resources within the resource server. The primary objective is to enforce fine-grained access control, allowing or denying access to various resources based on the user's role.

### **Roles and Hierarchy**

The RBAC system defines the following roles:

1. **SuperAdmin**
2. **Admin**
3. **Manager**
4. **Editor**
5. **Viewer**

A hierarchy is established where higher-level roles inherit permissions from lower-level roles.

### **Resources**

Access control is applied to two main resources:

1. **Users**
2. **Documents**

For the "Documents" resource, different types are considered:

- **Private**
- **Protected**
- **Public**

### **Permissions**

Permissions are specified for each role and resource type. The actions include:

- **Read**
- **Write**
- **Delete**

### **OPA (Open Policy Agent) Integration**

The Rego policy is integrated with OPA for runtime evaluation of access control rules. The **`allow`** rule, along with helper functions, checks if a user with a given role has the necessary permissions for a specific action on a resource.

## **Example Usage**

The following examples demonstrate how the RBAC policy works:

1. SuperAdmin has full access to all resources.
2. Admin can read, write, and delete both users and documents (all types).
3. Manager can read and write users and documents (all types).
4. Editor can read and write documents (all types).
5. Viewer can only read documents (all types).

## **Modification and Customization**

To modify or customize the RBAC policy, you can update the **`roles.rego`** file. Add or remove roles, adjust permissions, and extend resource types as needed. After making changes, ensure that the policy still aligns with the security requirements of your application.
