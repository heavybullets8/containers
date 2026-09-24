# syntax=docker/dockerfile:1

FROM public.ecr.aws/aws-cli/aws-cli:2.37.2@sha256:e38214027df83cb6631adcf980a092a98d1d29788789bff2a0f424e87e3da8ed AS aws-cli

FROM ghcr.io/home-operations/actions-runner:2.337.0@sha256:8b8e8326b5a08e658a0a41cfd2867afab69ca4194a152f986aa6fcd5430d09e7

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
