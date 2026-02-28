terraform {
  required_providers {
    koyeb = {
      source = "koyeb/koyeb"
    }
  }
}
provider "koyeb" {
  #
  # Use the KOYEB_TOKEN env variable to set your Koyeb API token.
  #
}

resource "koyeb_app" "my-app" {
  name = var.app_name
}

resource "koyeb_service" "my-service" {
  app_name = var.app_name
  definition {
    name = var.service_name
    instance_types {
      type = "free"
    }
    ports {
      port     = var.container_port
      protocol = "http"
    }
    scalings {
      min = 0
      max = 1
    }
    routes {
      path = "/"
      port = var.container_port
    }
    health_checks {
      http {
        port = var.container_port
        path = "/api/saudacoes/aleatorio"
      }
    }
    regions = ["was"]
    docker {
      image = "${var.docker_image_name}:${var.docker_image_tag}"
    }
  }

  depends_on = [
    koyeb_app.my-app
  ]
}