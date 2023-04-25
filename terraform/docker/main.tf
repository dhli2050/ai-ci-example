terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "3.0.2"
    }
  }
}

provider "docker" {}

resource "docker_image" "albumsvr_golang" {
  name         = "albumsvr_golang"
  build {
    context = "../.."
    dockerfile = "Dockerfile"
    build_arg = {
      tag: "album-server"
    }
    label = {
      author: "dhli"
    }
  }
  keep_locally = false
}

resource "docker_container" "albumsvr_golang" {
  image = docker_image.albumsvr_golang.image_id
  name  = var.container_name
  ports {
    internal = 8080
    external = 8080
  }
}
