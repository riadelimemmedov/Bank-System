postgres:
	docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123321 -p 6432:5432 -d postgres:12-alpine
get_postgres:
	docker exec -it postgres psql -U postgres -d postgres
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it postgres dropdb -U postgres simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:123321@127.0.0.1:6432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:123321@127.0.0.1:6432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate


get_accounts:
	docker exec -it postgres psql -U postgres -d simple_bank -c "SELECT * FROM accounts"
delete_accounts:
	docker exec -it postgres psql -U postgres -d simple_bank -c "TRUNCATE accounts RESTART IDENTITY CASCADE"
delete_entries:
	docker exec -it postgres psql -U postgres -d simple_bank -c "TRUNCATE entries RESTART IDENTITY CASCADE"
delete_transfers:
	docker exec -it postgres psql -U postgres -d simple_bank -c "TRUNCATE transfers RESTART IDENTITY CASCADE"



test:
	go test -v -cover -short ./...

.PHONY: createdb dropdb postgres get_postgres migrateup migratedown sqlc
.PHONY: get_accounts delete_accounts delete_entries delete_transfers
.PHONY: test