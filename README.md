# OpenRegistry - An open, decentralized, and reliable Container Registry

## Checks and Badges

| Type | Status |
|------|--------|
| OCI - Push | [![OCI Push](https://github.com/containerish/openregistry/actions/workflows/oci-dist-spec-push.yml/badge.svg?branch=main)](https://github.com/containerish/openregistry/actions/workflows/oci-dist-spec-push.yml)|
| OCI - Pull | [![OCI Pull](https://github.com/containerish/openregistry/actions/workflows/oci-dist-spec-pull.yml/badge.svg?branch=main)](https://github.com/containerish/openregistry/actions/workflows/oci-dist-spec-pull.yml)|
| OCI - Content Management | [![OCI Content Management](https://github.com/containerish/openregistry/actions/workflows/oci-dist-spec-content-management.yml/badge.svg?branch=main)](https://github.com/containerish/openregistry/actions/workflows/oci-dist-spec-content-management.yml)|
| OCI - Content Discovery | [![OCI Content Discovery](https://github.com/containerish/openregistry/actions/workflows/oci-dist-spec-content-discovery.yml/badge.svg?branch=main)](https://github.com/containerish/openregistry/actions/workflows/oci-dist-spec-content-discovery.yml)|
| Linter | [![OCI Push](https://github.com/containerish/openregistry/actions/workflows/golangci-lint.yml/badge.svg?branch=main)](https://github.com/containerish/openregistry/actions/workflows/golangci-lint.yml)|
| CodeQL | [![CodeQL](https://github.com/containerish/OpenRegistry/actions/workflows/codeql-analysis.yml/badge.svg?branch=main)](https://github.com/containerish/OpenRegistry/actions/workflows/codeql-analysis.yml)|
| Freshping | <a href="http://freshworks.com/website-monitoring?utm_source=status_badge&utm_medium=status_badge" target="_blank"><img src="https://statuspage.freshping.io/badge/91e4eb06-289b-4b4c-8beb-a0e5804959f4?0.56759354585684"/> </a>|
| Certifications | <a href="https://conformance.opencontainers.org/#openregistry" alt="OpenRegistry on OpenContainers" target="_blank"><img src="https://raw.githubusercontent.com/opencontainers/artwork/main/certified/oci_certified_color.svg" width=80 /><a/> |

> Disclaimer: Please refrain from using the **main** branch to run OpenRegistry in Production. The branch is highly experimental and not stable for Production use. Please only use the [released versions](https://github.com/containerish/OpenRegistry/releases) 

## Introduction
OpenRegistry is an open source, decentralized container registry which is fully compliant with [OCI Container Distribution Specification](https://github.com/opencontainers/distribution-spec/blob/main/spec.md).
The specification provides similar capabilities as that of the Docker Registry HTTP API V2 protocol.

## Why OpenRegistry?
For the longest time, we have relied on DockerHub to host and distribute our container images (both private and public). OpenRegistry tries to provide a decentralized alternative to that by running a community driven container registry, for People by People.

OpenRegitry uses [Akash Network](https://akash.network) as it's compute layer and IPFS, Filebase, or Storj for storage. Since AkashNetwork provides a spot like compute market, fault tolerance, Scalability and Resiliency are our priorities from day one.
	
## Getting Started
Working with OpenRegistry is no different than working with any other container registry. Following are the steps to get started:

### Sign-up: 
Head over to [Parachute by OpenRegistry](https://parachute.openregistry.dev) and sign yourself up. The sign process is essential as pushing to container repositories is a restricted operation and requires proper authorization.
> Currently we're only accepting registrations for a closed Beta program, Kindly register for Beta [here](https://parachute.openregistry.dev)

### Push an Image:
When using Docker CLI, the images are pushed to DockerHub by default. For Pushing images to OpenRegistry instead, follow the below steps:
* change the name of your image, e.g if you have an image named janedoe/alpine:latest, change it like so:
```bash
docker tag janedoe/alpine:latest openregistry.dev/janedoe/alipne:latest
docker push openregistry.dev/janedoe/alpine:latest
```

### Pull an Image:
Assuming you've pushed an image using the above method:
```bash
docker pull openregistry.dev/janedoe/alpine:latest
```

### How to Run this project locally:

If you'd like to run OpenRegistry locally or contribute a change/feature/bug fix or code changes, please follow this
guide on [how to set it up for Development](./docs/contributing/development-environment-setup.md)
