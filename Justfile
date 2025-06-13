default:
    just --list

lint:
    buf lint
    golangci-lint run ./internal/... ./cmd/...
    pnpm -r lint

fmt:
    golangci-lint fmt ./internal/... ./cmd/...
    pnpm -r format

gen:
    buf generate
    just gen-ent
    go mod tidy
    just fmt

gen-ent:
    go run -mod=mod entgo.io/ent/cmd/ent generate ./internal/ent/schema

hooks-set:
    git config --local core.hooksPath .githooks

test-backend:
    docker run --name veripass-test-db -e POSTGRES_USER=veripass -e POSTGRES_PASSWORD=veripass -e POSTGRES_DB=veripass -p 5432:5432 -d --rm postgres:latest
    until docker exec veripass-test-db pg_isready -U veripass; do sleep 1; done
    go test -v ./internal/...
    docker stop veripass-test-db