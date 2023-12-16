DB_URL=postgresql://root:password@localhost:5432/simple_bank?sslmode=disable

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build:
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

migrate_up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrate_up_1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migrate_down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrate_down_1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/brizaldi/go-simple-bank/db/sqlc Store

server:
	go run main.go

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    proto/*.proto

.PHONY: up up_build down migrate_up migrate_up_1 migrate_down migrate_down_1 mock server sqlc test db_docs db_schema proto