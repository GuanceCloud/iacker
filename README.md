# Iacker

**NOTE**: This project is a work in progress. DON'T USE IT TO PRODUCTION.

[![Go Report Card](https://goreportcard.com/badge/github.com/GuanceCloud/iacker)](https://goreportcard.com/report/github.com/GuanceCloud/iacker)
[![GoDoc](https://godoc.org/github.com/GuanceCloud/iacker?status.svg)](https://godoc.org/github.com/GuanceCloud/iacker)
[![License](https://img.shields.io/github/license/GuanceCloud/iacker.svg)](https://www.apache.org/licenses/LICENSE-2.0)

Iacker is a development framework to help cloud-native developers hugging with the infrastructure as a code's world.

The goal is to implement the open features for service providers.

* **Open API**: Expose [Cloud Control](https://aws.amazon.com/cloudcontrolapi/) and [GraphQL](https://graphql.org/) API to the developer.
* **Open Provider**: Auto-generate provider code for [Terraform](https://www.terraform.io), [Pulumi](https://www.pulumi.com), [CDK](https://developer.hashicorp.com/terraform/cdktf), [KubeVela](https://kubevela.io), [Crossplane,](https://www.crossplane.io) etc.
* **Open Integration**: Create libraries for [Cobra](https://github.com/spf13/cobra), [Apollo,](https://www.apollographql.com) and software development kits for various languages.

## Introduction

<!-- ![Graph](./artwork/arch.svg) -->

### Features

* **Model-driven**: Define the resource schema by [CUE](https://cuelang.org/) and unified specification.
* **Developer-first**: All workflows in code, collaborate with others with [GitOps](https://www.weave.works/technologies/gitops/) workflow.
* **Generator-oriented**: Build your own generator using the unified specification. Many popular IaC tools are provided.

### Real-world example

> We used it to create the [Terraform Provider for Guance Cloud](https://github.com/GuanceCloud/terraform-provider-guance) and other related services and tools.

![Logo](./artwork/banner-ins.jpeg)

## Roadmap

All the features are described in the `proposals` folder. It is inspired by the [GRFC Template](https://github.com/grpc/proposal/blob/master/GRFC-TEMPLATE.md) used in Google. You can find the implementation status in the table below.

The proposal workflow is described in the [Guance Cloud Proposal Governance Guidelines](https://github.com/GuanceCloud/community/tree/main/proposals).

### Architecture

Architecture defines the basic concepts and principles of *Iacker* architecture.

| #    | Title                         | How-tos |
| ---- | ----------------------------- | -------------- |
| A001 | [API as Code in *Iacker*](./proposals/A001-api-as-code-overview.md)   | How *Iacker* governance our API |
| A002 | [Resource Management Specification](./proposals/A002-resource-management-specification.md) | How to define a new cloud resource |
| A003 | [Generator Framework](./proposals/A003-generator-framework.md) | How to build developer tools by *Iacker* |
| A004 | Resource Management Server | How to develop once, run as multiple protocols |

### Protocols

At *Iacker*, we use a unified resource management framework to expose various protocols for different use cases.

Protocols describe the protocols used in *Iacker*. Please follow the protocol specification (coming soon) to add a new protocol.

| #    | Title                         | How-tos |
| ---- | ----------------------------- | -------------- |
| P001 | Cloud Control API | Cloud Control API and the implementations |

### Developer Tools

Developer tools are the core of developer experience infrastructure (DXI). It provides tools to help developers build, test, and deploy their API callings.

| #    | Title                         | How-tos |
| ---- | ----------------------------- | -------------- |
| T001 | [Resource Explorer](./proposals/T001-resource-explorer.md) | How to discover and interact with cloud resource |
| T002 | Terraform Provider | How to build the engineering capability for X as Code world |

## Developer Guide

See the [hack](./hack) folder to understand how to build this project.

## Contributing

We welcome contributions to *Iacker*. Please see the [Contributing Guidelines](https://guance.io/contribution-guide/) for more information.
