create-container:
	docker run --name secure-bank-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

start-container:
	docker start secure-bank-db

destroy-container:
	docker stop secure-bank-db && docker rm secure-bank-db	

tail-container-logs:
	docker logs -f secure-bank-db

connect-to-database:
	docker exec -it secure-bank-db psql -U root secure_bank

createdb:
	docker exec -it secure-bank-db createdb --username=root --owner=root secure_bank

dropdb:
	docker exec -it secure-bank-db dropdb secure_bank

migrate-up:
	 migrate --path migration --database "postgresql://root:secret@localhost:5432/secure_bank?sslmode=disable" --verbose up

migrate-down:
	 migrate --path migration --database "postgresql://root:secret@localhost:5432/secure_bank?sslmode=disable" --verbose down

.PHONY: create-container start-container destroy-container tail-container-logs createdb dropdb migrate-up migrate-down