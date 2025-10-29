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

start-test-database: stop-test-database
    docker run --rm -d --name veripass-test-db -e POSTGRES_USER=veripass -e POSTGRES_PASSWORD=veripass -e POSTGRES_DB=veripass -p 5432:5432 postgres:latest -c logging_collector=on -c log_statement=all -c log_filename=postgresql.log
    until docker exec veripass-test-db pg_isready -U veripass; do sleep 1; done

stop-test-database:
    docker stop veripass-test-db || true

test-backend: start-test-database && stop-test-database
    go test -v ./internal/...