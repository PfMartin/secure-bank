# secure-bank

Example project for developing a banking api with golang, grpc and postgres

## Database

### Setup

- Install `golang-migrate`

```zsh
cd db

make create-container
make createdb
make migrate-up
```

### Table Plus installation

[Installation Manual](https://tableplus.com/blog/2019/10/tableplus-linux-installation.html)

```zsh
# Add TablePlus gpg key
wget -qO - https://deb.tableplus.com/apt.tableplus.com.gpg.key | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/tableplus-archive.gpg > /dev/null

# Add TablePlus repo
sudo add-apt-repository "deb [arch=amd64] https://deb.tableplus.com/debian/22 tableplus main"

# Install
sudo apt update
sudo apt install tableplus
```

### golang-migrate

#### Install golang-migrate

[Installation Manual](https://www.geeksforgeeks.org/how-to-install-golang-migrate-on-ubuntu/)

```zsh
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash

sudo apt udpate
sudo apt install migrate
```

#### Initialization

```zsh
mkdir -p db/migration

migrate create --ext sql --dir db/migration --seq init_schema
```

- This creates up an down migrations
- If we call `migrate up` it will upgrade the database to the last number of up migrations
- If we call `migrate down` it will downgrade the database from the last to the first of down migrations

#### Apply migrations

- Create migrations file with [dbdiagram.io](https://dbdiagram.io/d)
- Copy the content to the `up`-migration
- Add the reverting commands to the `down`-migration

```sql
DROP TABLE IF EXISTS entries CASCADE;
DROP TABLE IF EXISTS transfers CASCADE;
DROP TABLE IF EXISTS accounts CASCADE;
```

## CRUD Operations

### sqlc

- Fully supports postgres
- Translates SQL to ideomatic golang std library for sql

#### Install sqlc

```zsh
sudo snap install sqlc

sqlc version
sqlc help
```

#### Configure sqlc

- Initialize sqlc --> Will create a `sqlc.yaml` in your project folder

```zsh
sqlc init
```

- Add postgresql engine and the correct file paths for queries and schema
- Remove `cloud` and `database` for using sqlc without the cloud

```yml
version: "2"
sql:
  - engine: "postgresql"
    queries: "./db/query/"
    schema: "./db/migration/"
    gen:
      go:
        package: "db"
        out: "./db/sqlc"
        sql_package: "pgx/v5"
        emit_json_tags: true
        emit_prepared_queries: false
        emit_interface: false
        emit_exact_table_names: false
```

- Write your own queries in SQL and generate code from these queries using `sqlc generate`
