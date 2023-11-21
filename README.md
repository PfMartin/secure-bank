# secure-bank

Example project for developing a banking api with golang, grpc and postgres

## Database Setup

### Postgres Container

- Start the container

```zsh
docker run --name secure-bank-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Jahnel01 -d postgres:16-alpine
```

- Connect to the container
- Password is not required when connecting locally

```zsh
docker exec -it secure-bank-db psql -U root
```

- Check if it works

```zsh
select now();
```

- Exit the container by exiting `psql`

```bash
\q
```

- Access logs of the container

```zsh
docker logs secure-bank-db
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

#### Installation

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
DROP TABLE IF EXISTS entries;
DROP TABLE IF EXISTS transfers;
DROP TABLE IF EXISTS accounts;
```

- Apply the migration

```zsh
migrate --path db/migration --database "postgresql://root:secret@localhost:5432/secure_bank?sslmode=disable" --verbose up
```
