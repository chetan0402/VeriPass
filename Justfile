default:
    just --list

lint:
    buf lint
    golangci-lint run ./...
    pnpm -r lint

fmt:
    golangci-lint fmt ./...
    pnpm -r format

gen:
    buf generate
