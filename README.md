**This repository is a detached fork of the upstream repo:
[onedr0p/containers](https://github.com/onedr0p/containers), which has now been
archived in favor of
[home-operations/containers](https://github.com/home-operations/containers)**

This repository was made when I originally wanted a container to test nzbget,
now I add some additional applications that upstream has removed, or deprecated.

For the *arr applications, upstream only offers `development` flavors, where I
was wanting `stable`.

## Deprecations

Anything added to the upstream repo mentioned above, will be deprecated and
removed from this repository, I do not really have a timeline for when that is
going to happen. This is not a very "official" repositry. Which is why you
should use the upstream repository whenver possible.

## Tag Immutability

The containers built here do not use immutable tags, at least not in the more
common way you have seen from [linuxserver.io](https://fleet.linuxserver.io/) or
[Bitnami](https://bitnami.com/stacks/containers).

We do take a similar approach but instead of appending a `-ls69` or `-r420`
prefix to the tag, we instead insist on pinning to the sha256 digest of the
image. While this is not as pretty, it is just as functional in making the
images immutable.

| Container                                                 | Immutable |
| --------------------------------------------------------- | --------- |
| `ghcr.io/Heavybullets8/sonarr:rolling`                    | ❌        |
| `ghcr.io/Heavybullets8/sonarr:4.0.14.2939`                | ❌        |
| `ghcr.io/Heavybullets8/sonarr:rolling@sha256:3527...`     | ✅        |
| `ghcr.io/Heavybullets8/sonarr:4.0.14.2939@sha256:3527...` | ✅        |

_If pinning an image to the sha256 digest, tools like
[Renovate](https://github.com/renovatebot/renovate) support updating the
container on a digest or application version change._

## Rootless

To run these containers as non-root, make sure you update your configuration to
the user and group you want.

### Docker Compose

```yaml
networks:
  sonarr:
    name: sonarr
    external: true
services:
  sonarr:
    image: ghcr.io/heavybullets8/sonarr:4.0.14.2939
    container_name: sonarr
    user: 568:568
    # ...
```

### Kubernetes

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sonarr
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

## Passing arguments to an application

Some applications do not support defining configuration via environment
variables and instead only allow certain config to be set in the command line
arguments for the app. To circumvent this, for applications that have an
`entrypoint.sh` read below.

1. First read the Kubernetes docs on
   [defining command and arguments for a Container](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/).
2. Look up the documentation for the application and find a argument you would
   like to set.
3. Set the extra arguments in the `args` section like below.

   ```yaml
   args:
     - --port
     - "8080"
   ```

## Configuration volume

For applications that need to have persistent configuration data the config
volume is hardcoded to `/config` inside the container. This is not able to be
changed in most cases.

## Credits

All credit goes to [onedr0p](https://github.com/onedr0p),
[bjw-s](https://github.com/bjw-s), [buroa](https://github.com/buroa), and all of
the other contributors to the upstream project.
