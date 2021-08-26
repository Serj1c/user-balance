dev:
	ENV=development docker-compose -f docker-compose.yml up -d

postgres:
	docker run --name balance -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=spartak1 -d postgres:13.1-alpine

migrateup:
	migrate -path migrations/migrations -database "postgresql://root:spartak1@localhost:5432/datacademy?sslmode=disable" up

migratedown:
	migrate -path migrations/migrations -database "postgresql://root:spartak1@localhost:5432/datacademy?sslmode=disable" down

.PHONY: dev postgres migrateup migratedown