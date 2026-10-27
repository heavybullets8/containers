# syntax=docker/dockerfile:1

FROM public.ecr.aws/aws-cli/aws-cli:2.39.0@sha256:6014c77975a4ae237395924cc821333f20bf87c91964a45dd4ed5f3b13b424f5 AS aws-cli

FROM ghcr.io/home-operations/actions-runner:2.338.1@sha256:9011bec82009308072bb8677229181791f96bada721df8a91108290ad4a3fdc5

USER root

ADD --checksum=sha256:0144068502a1eddd2a0280ede10ef607d1ec592ce819940991203941564e8e76 \
    https://www.postgresql.org/media/keys/ACCC4CF8.asc /usr/share/keyrings/postgresql.asc

RUN printf '%s\n' \
    'deb [signed-by=/usr/share/keyrings/postgresql.asc] https://apt.postgresql.org/pub/repos/apt noble-pgdg main' \
    > /etc/apt/sources.list.d/pgdg.list \
    && chmod 644 /usr/share/keyrings/postgresql.asc

RUN apt-get update \
    && apt-get install --yes --no-install-recommends \
    age lftp poppler-utils postgresql-client-17 \
    && rm -rf /var/lib/apt/lists/*

COPY --from=aws-cli /usr/local/aws-cli/ /usr/local/aws-cli/
RUN ln -s /usr/local/aws-cli/v2/current/bin/aws /usr/local/bin/aws

USER runner
