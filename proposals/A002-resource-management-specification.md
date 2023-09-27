A002: Resource Management Specification
----
* Author(s): @yufeiminds
* Approver: @yufeiminds 
* Status: Implemented
* Last updated: 2023-03-06
* Discussion at: #1

## Abstract

This proposal describes a new solution to manage the API resources by the unified specification.

The focus points of this proposal are to design a specification to describe the resource schema and the resource management API. Achieve the following goals:

1. **Describe The Resource Schema**. Describe the resource name, title, description, attributes, and relationships of resources. Includes Internationalization (i18n) support.
2. **Describe The Type System of Resource**. Describe the attribute type system of the resource. Include primitive types (string, number, boolean, etc.), complex types (object, array, etc.), and reference types (resource reference, etc.).
3. **Extend the Resource Schema**. Design the extension mechanism of the resource schema. The extended metadata can be used at the specified artifact, such as Terraform Provider, API Explorer, etc.

So we define a new specification called "Resource Management Specification" (RMS). It is a specification that describes the resource name, title, description, attributes, and relationships of resources. It is a unified specification for all resources in the cloud APIs.

This proposal is a specification proposal, not an implementation proposal. It will describe the concrete definition of the resource schema. The user can follow the specification to declare their resource.

## Background

In software development, "Single Trusted Source" is a concept that describes the source of truth for the system. All the system's components are consistent with the source of truth.

For Resource Management Specification (RMS), the "Resource" is a concept that describes the object that can be managed by the system—for example, a user, a workspace, a conservative policy, etc. RMS provides enough information to describe the traits of resources. So it can be the "Single Trusted Source" of the resource artifacts.

### Related Proposals:

* [A001: API as Code Overview](./A001-api-as-code-overview.md). Describe the "API as Code" concept and the goals of this specification.

## Proposal

Here will describe the details of the Resource Management Specification (RMS).

### Resource Schema

![Resource Schema](./A002_images/schema.svg)

*TODO: more details and descriptions*

### Metadata

RMS has a mechanism to extend the resource schema. The extended metadata can be used at the specified artifact.

There are common metadata properties:

**Resource**

* **datasource** (bool): A boolean value indicates whether the resource is a data source. The default value is `false`.

**Property**

* **dynamic** (bool): A boolean value indicates whether the property is computed after creation. The default value is `false`.
* **immutable** (bool): A boolean value indicates whether the property is immutable; it will force a new creation when changed. The default value is `false`.

### Versioning

In RMS, every resource has its package and version. The version is a string that follows the [PACKAGE_VERSION_SUFFIX](https://docs.buf.build/lint/rules#package_version_suffix). The version is used to describe the compatibility of the resource schema.

For example, If we want to create a new resource named `workspace`:

* We create a new resource with the package `workspace` and the version `v1alpha`.
* When the API is stable in our look, we can release version `v1beta`.
* After a long time of production, It is successful for many users. We can change it to `v1`.
* If we want to make a breaking change, we can release version `v2alpha`.

## Implementation

This proposal is a specification proposal, not an implementation proposal. Please see other proposals for implementation.
