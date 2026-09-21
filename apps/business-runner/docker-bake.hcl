target "docker-metadata-action" {}

variable "APP" {
  default = "business-runner"
}

variable "VERSION" {
  // renovate: datasource=docker depName=ghcr.io/home-operations/actions-runner
  default = "2.337.0"
}

variable "SOURCE" {
  default = "https://github.com/heavybullets8/containers"
}

group "default" {
  targets = ["image-local"]
}

target "image" {
  inherits = ["docker-metadata-action"]
  labels = {
    "org.opencontainers.image.source" = "${SOURCE}"
  }
}

target "image-local" {
  inherits = ["image"]
  output = ["type=docker"]
  tags = ["${APP}:${VERSION}"]
}

target "image-all" {
  inherits = ["image"]
  platforms = ["linux/amd64"]
}
