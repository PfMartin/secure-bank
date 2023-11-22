# ----------------------------- #
# Postgres container management #
# ----------------------------- #
connect-to-database:
	docker exec -it secure-bank-db psql -U root secure_bank

create-container:
	docker run --name secure-bank-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

destroy-container:
	docker stop secure-bank-db && docker rm secure-bank-db	

start-container:
	docker start secure-bank-db

tail-container-logs:
	docker logs -f secure-bank-db

# ------------------- #
# Database management #
# ------------------- #
createdb:
	docker exec -it secure-bank-db createdb --username=root --owner=root secure_bank

dropdb:
	docker exec -it secure-bank-db dropdb secure_bank

migrate-down:
	 migrate --path db/migration --database "postgresql://root:secret@localhost:5432/secure_bank?sslmode=disable" --verbose down

migrate-up:
	 migrate --path db/migration --database "postgresql://root:secret@localhost:5432/secure_bank?sslmode=disable" --verbose up

# ---- #
# SQLC #
# ---- #
sqlc:
	sqlc generate


.PHONY: connect-to-database createdb create-container destroy-container dropdb migrate-down migrate-up start-container sqlc tail-container-logs   