# Useful command which help to you testing and debugging process

- Create migration:
- `migrate create -ext sql -dir .\db\migration\ -seq init_schema`
- Go to inside postgres container:
- `docker exec -it postgres /bin/bash`
- Create database inside postgres container:
- `createdb --username=postgres --owner=postgres simple_bank`
- Go to inside psql:
- `docker exec -it postgres psql -U postgres -d postgres`
- Delete database:
- `docker exec -it postgres dropdb -U postgres simple_bank`
- Create database:
- `docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank`
- Send migration to database:
- `migrate -path db/migration -database "postgresql://postgres:123321@127.0.0.1:6432/simple_bank?sslmode=disable" -verbose up`
