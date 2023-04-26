terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

# provision to ap-southeast-2 region
provider "aws" {
    region = "ap-southeast-2"
}

resource "aws_ecr_repository" "ai_ci_ecr_repo" {
  name = "ai-ci-ecr-repo"
}

resource "aws_ecs_cluster" "ai_ci_cluster" {
  name = "ai-ci-cluster"
}

resource "aws_ecs_task_definition" "ai_ci_task" {
  family = "ai-ci-task"
  container_definitions = jsonencode([
    {
      name = "${var.server_bin_name}"
      image = "${aws_ecr_repository.ai_ci_ecr_repo.repository_url}"
      essential = true
      portMappings = [
        {
          containerPort = 8080
          hostPort = 8080
        }
      ]
      memory = 512
      cpu = 256
    }
  ])
  requires_compatibilities = [ "FARGATE" ]
  network_mode = "awsvpc"
  memory = 512
  cpu = 256
  execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
}

resource "aws_iam_role" "ecs_task_execution_role" {
  name = "ecs_task_execution_role"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
}

data "aws_iam_policy_document" "assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_ecs_service" "ai_ci_service" {
  name = "ai-ci-service"
  cluster = aws_ecs_cluster.ai_ci_cluster.id
  task_definition = aws_ecs_task_definition.ai_ci_task.arn
  launch_type = "FARGATE"
  desired_count = 3

  network_configuration {
    subnets = [ aws_default_subnet.subnet_a.id, aws_default_subnet.subnet_b.id, aws_default_subnet.subnet_c.id ]
    assign_public_ip = true
  }
}

resource "aws_default_subnet" "subnet_a" {
  availability_zone = "ap-southeast-2a"
}

resource "aws_default_subnet" "subnet_b" {
  availability_zone = "ap-southeast-2b"
}

resource "aws_default_subnet" "subnet_c" {
  availability_zone = "ap-southeast-2c"
}