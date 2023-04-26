variable "ec2_name" {
  type = string
  default = "SimpleAlbumServer"
}

variable "server_bin_name" {
  type = string
  default = "albumsvr"
}

variable "region" {
  type = string
  default = "ap-southeast-2"
}

variable "image_version" {
  type = string
  default = "latest"
}