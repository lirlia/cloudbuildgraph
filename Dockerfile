FROM golang AS build-env
COPY . /go-src/
WORKDIR /go-src/
RUN go build -o /cloudbuildgraph .

FROM ubuntu:latest

LABEL org.opencontainers.image.source="https://github.com/lirlia/cloudbuildgraph"

RUN \
  # Install Graphviz
  # Note that the lack of a "lock" mechanism for apt dependencies
  # currently prevents a fully reproducible build
  apt-get update \
  && apt-get install -y --no-install-recommends \
  graphviz \
  && rm -rf /var/lib/apt/lists/* \
  \
  # Add a non-root user
  && useradd app

# Run as non-root user
USER app

COPY --from=build-env /cloudbuildgraph /

ENTRYPOINT ["/cloudbuildgraph"]
