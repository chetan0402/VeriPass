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