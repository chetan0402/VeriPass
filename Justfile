default:
    just --list

lint:
    buf lint
    golangci-lint run ./internal/... ./cmd/...
    pnpm -r lint

fmt:
    buf format -w
    golangci-lint fmt ./internal/... ./cmd/...
    pnpm -r format
    go mod tidy

gen:
    just gen-proto
    just gen-ent
    go mod tidy
    just fmt

gen-proto:
    buf generate

gen-ent:
    go run -mod=mod entgo.io/ent/cmd/ent generate ./internal/ent/schema

start-test-database: stop-test-database
    docker run --rm -d --name veripass-test-db -e POSTGRES_USER=veripass -e POSTGRES_PASSWORD=veripass -e POSTGRES_DB=veripass -p 5432:5432 postgres:latest -c logging_collector=on -c log_statement=all -c log_filename=postgresql.log
    until docker exec veripass-test-db pg_isready -U veripass; do sleep 1; done

stop-test-database:
    docker stop veripass-test-db || true

start-test-dex: stop-test-dex
    docker run --rm -d --name veripass-test-dex -p 1433:1433 -v ./dex-config-testing.yaml:/etc/dex/config.yaml dexidp/dex:latest dex serve /etc/dex/config.yaml

stop-test-dex:
    docker stop veripass-test-dex || true

test-backend: start-test-database start-test-dex && stop-test-database stop-test-dex
    go test -v ./internal/...

build group="":
    docker buildx bake {{group}} --load

setup-githooks:
    git config core.hooksPath .githooks
