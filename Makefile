createdb:
    docker exec -it senpaislist-backend_postgres_1 createdb --username=postgres --owner=postgres anime

dropdb:
    docker exec -it senpaislist-backend_postgres_1 dropdb anime

.PHONY: createdb dropdb