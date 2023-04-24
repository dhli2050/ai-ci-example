.PHONY: lint build validate deploy destroy unit-test interface-test

lint:
	golangci-lint run

build:
	GOOS=linux GOARCH=amd64 go build -o albumsvr

validate:
	cd terraform && terraform validate

deploy:
	cd terraform && terraform apply -auto-approve

destroy:
	cd terraform && terraform destroy -auto-approve

unit-test:
	GIN_MODE=release go test -v ./...

interface-test:
	go test -v -run Interface
