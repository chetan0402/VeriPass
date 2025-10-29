group "default" {
  targets = [ "backend", "frontend" ]
}

target "backend" {
  context = "."
  dockerfile = "Dockerfile.backend"
  tags = [ "veripass/backend:latest" ]
}

target "frontend" {
  context = "."
  dockerfile = "Dockerfile.frontend"
  tags = [ "veripass/frontend:latest" ]
}