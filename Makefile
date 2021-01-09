createdb:
    docker exec -it senpaislist-backend_postgres_1 createdb -U postgres --owner=postgres anime

dropdb:
    docker exec -it senpaislist-backend_postgres_1 dropdb -U postgres anime

sqlc:
    sqlc generate

migrateup:
    migrate -path db/migration -database "postgres://postgres:championsclub123@postgres:5432/postgres?sslmode=disable" -verbose up

migratedown:
    migrate -path db/migration -database "postgres://postgres:championsclub123@postgres:5432/postgres?sslmode=disable" -verbose down

.PHONY: 
    createdb dropdb sqlc migrateup migratedown
