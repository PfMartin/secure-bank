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
DROP TABLE IF EXISTS entries CASCADE;
DROP TABLE IF EXISTS transfers CASCADE;
DROP TABLE IF EXISTS accounts CASCADE;
```
