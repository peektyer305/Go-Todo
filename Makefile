ARG = foo

go-install:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/cespare/reflex@latest
	go install github.com/go-delve/delve/cmd/dlv@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

wire-gen:
	wire ./di/container.go

migrate-status:
	goose --dir ./migrations status

migrate-up:
	goose --dir ./migrations up

migrate-down:
	goose --dir ./migrations down

migrate-reset:
	goose --dir ./migrations reset

migrate-refresh: migrate-reset migrate-up

migrate-create:
	goose --dir ./migrations create ${ARG} sql

migrate-create-go:
	goose --dir ./migrations create ${ARG} go

