seed:
	go run ./cmd/seed/main.go
setup:
	go get -u ./...
	docker-compose -p ctaglearn -f ./docker/docker-compose.yaml up -d
