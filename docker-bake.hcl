group "default" {
  targets = [ "backend", "frontend" ]
}

group "dev" {
  targets = [ "backend", "frontend-dev" ]
}

target "backend" {
  context = "."
  dockerfile = "Dockerfile.backend"
  tags = [ "veripass/backend:latest" ]
  args = {
    "GO_VERSION" = "1.25-alpine"
  }
}

target "frontend" {
  context = "."
  dockerfile = "Dockerfile.frontend"
  tags = [ "veripass/frontend:latest" ]
  args = {
    "VITE_MOCK" = "false"
    "VITE_CLIENT_ID" = "unknown"
    "VITE_REDIRECTION_URI" = "http://unknown"
  }
}

target "frontend-dev" {
  context = "."
  dockerfile = "Dockerfile.frontend"
  tags = [ "veripass/frontend:dev" ]
  args = {
    "VITE_MOCK" = "false"
    "VITE_CLIENT_ID" = "veripass"
    "VITE_REDIRECTION_URI" = "http://localhost:5002/api/callback"
  }
}
