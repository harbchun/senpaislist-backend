dcupbuild:
	docker-compose up --build -d

dcdownvolume:
	docker-compose down -v

createdb:
	docker exec -it senpaislist-backend_postgres_1 createdb -U postgres --owner=postgres senpaislist

dropdb:
	docker exec -it senpaislist-backend_postgres_1 dropdb -U postgres senpaislist

sqlc:
	sqlc generate

migrateup:
	migrate -path db/migration -database "postgres://postgres:championsclub123@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://postgres:championsclub123@localhost:5432/postgres?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

builddbpy:
	python ./scripts/builddb.py

.PHONY:
	dcupbuild dcdownvolume createdb dropdb sqlc migrateup migratedown test builddbpy
