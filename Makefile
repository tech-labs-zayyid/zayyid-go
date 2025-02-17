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
create-migration $$(enter):
	@read -p "Migration name:" migration_name; \
	dbmate -d "./migrations" new $$migration_name
migration-up:
	dbmate -d "./migrations" -e "DATABASE_URL" up
npx-migration-up:
	npx dbmate -d "./migrations" -e "DATABASE_URL" up