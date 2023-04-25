output "container_id" {
  description = "ID of the Docker container"
  value = docker_container.albumsvr_golang.id
}

output "image_id" {
  description = "ID of the Docker image"
  value = docker_image.albumsvr_golang.id
}