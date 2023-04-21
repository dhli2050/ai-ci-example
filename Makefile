.PHONY: lint build deploy destroy

lint:
	golangci-lint run

build:
	GOOS=linux GOARCH=amd64 go build -o albumsvr

deploy:
	cd terraform && terraform apply -auto-approve

destroy:
	cd terraform && terraform destroy -auto-approve
