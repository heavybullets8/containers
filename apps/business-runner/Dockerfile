# syntax=docker/dockerfile:1

FROM public.ecr.aws/aws-cli/aws-cli:2.37.11@sha256:3dacc5db57c923c4223e949795f538ecf1f2212b2b7d5a028b47b97f91564c0d AS aws-cli

FROM ghcr.io/home-operations/actions-runner:2.338.0@sha256:c867ecdf82d4e8537ea8a41acc6addae09f069daec137e5614b87e4678a953e8

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
