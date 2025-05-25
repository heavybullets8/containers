target "docker-metadata-action" {}

variable "APP" {
  default = "sonarr"
}

variable "VERSION" {
  // renovate: datasource=github-releases depName=Sonarr/Sonarr versioning=loose
  default = "4.0.14.2939"
}

variable "SOURCE" {
  default = "https://github.com/Sonarr/Sonarr"
}

group "default" {
  targets = ["image-local"]
}

target "image" {
  inherits = ["docker-metadata-action"]
  args = {
    VERSION = "${VERSION}"
  }
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
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
}
