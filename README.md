<!---
NOTE: AUTO-GENERATED FILE
to edit this file, instead edit its template at: ./scripts/templates/README.md.j2
-->
<div align="center">

**This repository is a detached fork of the upstream repo: [onedr0p/containers](https://github.com/onedr0p/containers)**

## Containers

_An opinionated collection of container images_

</div>

<div align="center">

</div>

Welcome to my container images. If you're looking for a container, start by [browsing the GitHub Packages page for this repo's packages](https://github.com/Heavybullets8?tab=packages&repo_name=containers).

## Mission Statement

The goal of this project is to support [semantically versioned](https://semver.org/), [rootless](https://rootlesscontaine.rs/), and [multiple architecture](https://www.docker.com/blog/multi-arch-build-and-images-the-simple-way/) containers for various applications.

It also adheres to a [KISS principle](https://en.wikipedia.org/wiki/KISS_principle), logging to stdout, [one process per container](https://testdriven.io/tips/59de3279-4a2d-4556-9cd0-b444249ed31e/), no [s6-overlay](https://github.com/just-containers/s6-overlay), and all images are built on top of [Alpine](https://hub.docker.com/_/alpine) or [Ubuntu](https://hub.docker.com/_/ubuntu).

## Tag Immutability

The containers built here do not use immutable tags, at least not in the more common way you have seen from [linuxserver.io](https://fleet.linuxserver.io/) or [Bitnami](https://bitnami.com/stacks/containers).

We do take a similar approach but instead of appending a `-ls69` or `-r420` prefix to the tag, we instead insist on pinning to the sha256 digest of the image. While this is not as pretty, it is just as functional in making the images immutable.

| Container                                                | Immutable |
|----------------------------------------------------------|-----------|
| `ghcr.io/Heavybullets8/nzbget:rolling`                   | ❌        |
| `ghcr.io/Heavybullets8/nzbget:24.4`                      | ❌        |
| `ghcr.io/Heavybullets8/nzbget:rolling@sha256:3527...`    | ✅        |
| `ghcr.io/Heavybullets8/nzbget:24.4@sha256:3527...`       | ✅        |

_If pinning an image to the sha256 digest, tools like [Renovate](https://github.com/renovatebot/renovate) support updating the container on a digest or application version change._

## Rootless

To run these containers as non-root, make sure you update your configuration to the user and group you want.

### Docker Compose

```yaml
networks:
  nzbget:
    name: nzbget
    external: true
services:
  nzbget:
    image: ghcr.io/Heavybullets8/nzbget:24.4
    container_name: nzbget
    user: 568:568
    # ...
```

### Kubernetes

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nzbget
# ...
spec:
  # ...
  template:
    # ...
    spec:
      # ...
      securityContext:
        runAsUser: 568
        runAsGroup: 568
        fsGroup: 568
        fsGroupChangePolicy: OnRootMismatch
# ...
```

## Passing arguments to a application

Some applications do not support defining configuration via environment variables and instead only allow certain config to be set in the command line arguments for the app. To circumvent this, for applications that have an `entrypoint.sh` read below.

1. First read the Kubernetes docs on [defining command and arguments for a Container](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/).
2. Look up the documentation for the application and find a argument you would like to set.
3. Set the extra arguments in the `args` section like below.

    ```yaml
    args:
      - --port
      - "8080"
    ```

## Configuration volume

For applications that need to have persistent configuration data the config volume is hardcoded to `/config` inside the container. This is not able to be changed in most cases.

## Available Images

Each Image will be built with a `rolling` tag, along with tags specific to it's version. Available Images Below

Container | Channel | Image
--- | --- | ---
[nzbget](https://github.com/Heavybullets8/containers/pkgs/container/nzbget) | stable | ghcr.io/Heavybullets8/nzbget


## Credits

All credit goes to [onedr0p](https://github.com/onedr0p), [bjw-s](https://github.com/bjw-s), and all of the other contributors to the upstream project.