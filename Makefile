
ACCOUNT_ID = 3


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
migrate_sql:
	migrate create -ext sql -dir db/migration -seq add_users
sqlc:
	sqlc generate
mock:
	mockgen -package mockdb -destination db/mock/store.go simplebank/db/sqlc Store

get_accounts:
	docker exec -it postgres psql -U postgres -d simple_bank -c "SELECT * FROM accounts"
get_transfers:
	docker exec -it postgres psql -U postgres -d simple_bank -c "SELECT * FROM transfers"
get_entries:
	docker exec -it postgres psql -U postgres -d simple_bank -c "SELECT * FROM entries"
get_users:
	docker exec -it postgres psql -U postgres -d simple_bank -c "SELECT * FROM users"
delete_accounts:
	docker exec -it postgres psql -U postgres -d simple_bank -c "TRUNCATE accounts RESTART IDENTITY CASCADE"
delete_entries:
	docker exec -it postgres psql -U postgres -d simple_bank -c "TRUNCATE entries RESTART IDENTITY CASCADE"
delete_transfers:
	docker exec -it postgres psql -U postgres -d simple_bank -c "TRUNCATE transfers RESTART IDENTITY CASCADE"
update_accounts:
	docker exec -it postgres psql -U postgres -d simple_bank -c "UPDATE accounts SET currency='EUR' WHERE ID = $(ACCOUNT_ID)"


test:
	go test -v -cover ./...
server:
	go run main.go

.PHONY: createdb dropdb postgres get_postgres migrateup migratedown migrate_sql sqlc mock
.PHONY: get_accounts get_transfers get_entries delete_accounts delete_entries delete_transfers update_accounts
.PHONY: test server