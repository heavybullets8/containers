# syntax=docker/dockerfile:1

FROM ghcr.io/home-operations/actions-runner:2.338.1@sha256:9011bec82009308072bb8677229181791f96bada721df8a91108290ad4a3fdc5

USER root

RUN apt-get update \
    && apt-get install --yes --no-install-recommends xz-utils \
    && rm -rf /var/lib/apt/lists/* \
    && mkdir -p /nix /etc/nix \
    && chown runner:runner /nix \
    && printf '%s\n' \
    'always-allow-substitutes = true' \
    'build-users-group =' \
    'experimental-features = nix-command flakes' \
    'max-jobs = auto' \
    'show-trace = true' \
    'trusted-users = root runner' \
    > /etc/nix/nix.conf \
    && chmod 0644 /etc/nix/nix.conf

ADD --checksum=sha256:9adda97297d9e8ab360df95c729eabff4f4f93d6db091953c3a68f29e3fb130c \
    https://releases.nixos.org/nix/nix-2.35.2/install /tmp/install-nix

RUN chown runner:runner /tmp/install-nix \
    && chmod 0644 /tmp/install-nix

USER runner

RUN sh /tmp/install-nix --no-daemon --no-channel-add \
    --nix-extra-conf-file /etc/nix/nix.conf \
    && rm /tmp/install-nix

ENV NIX_PATH=nixpkgs=channel:nixos-unstable \
    NIX_PROFILES="/nix/var/nix/profiles/default /home/runner/.nix-profile" \
    PATH="/home/runner/.nix-profile/bin:${PATH}"
