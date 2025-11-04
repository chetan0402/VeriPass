group "default" {
  targets = [ "backend", "frontend" ]
}

target "backend" {
  context = "."
  dockerfile = "Dockerfile.backend"
  tags = [ "veripass/backend:latest" ]
  args = {
    "GO_VERSION" = "1.25.3-alpine"
  }
}

target "frontend" {
  context = "."
  dockerfile = "Dockerfile.frontend"
  tags = [ "veripass/frontend:latest" ]
}