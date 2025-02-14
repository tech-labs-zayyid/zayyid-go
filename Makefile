run_local:
	@echo "Run apps.."
	go run main.go
run_docker:
	@echo "Run docker.."
	docker network create nabati || @echo "network already exist! skip.."
	docker-compose build --no-cache
	docker-compose up -d
stop_docker:
	@echo "Stop docker.."
	docker-compose down
lint:
	staticcheck ./...
	gocritic check ./... 
	golangci-lint run