# Guance Community hack GuideLines

## What is this?

The hack directory contains scripts used to build, test, package, and release this project.

Many Kubernetes projects use this pattern. We use it in Guance Cloud to ensure that all contributors use the same build and release process.

## How to use it?

`make` will call the mage step to complete the workflow. Follow commands is supported.

```bash
Package main includes the Makefile of Iacker.

Targets:
  dev:d2         build svg from d2 files
  dev:fmt        format the code
  dev:install    install the binary into local environment
  dev:lint       lint the code
```

## References

* [Kubernetes hack GuideLines](https://github.com/kubernetes/kubernetes/tree/v1.26.1/hack)
* [Dagger hack Overview](https://github.com/dagger/dagger/tree/main/hack)
