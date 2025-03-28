# matchoshop
E-commerce for men's product

## Requirements

- [Golang](https://golang.org/) as main programming language.
- [Go Module](https://go.dev/blog/using-go-modules) for package management.
- [Goose](https://github.com/steinbacher/goose/) as migration tool.
- [Postgresql](https://www.postgresql.org/) as database driver.
- [Docker-compose](https://docs.docker.com/compose/) for running database container locally.
- [Mockery](https://github.com/vektra/mockery/) for generate mockup object

## Setup
### Prepare necessary environment by rename .env.example to .env

```bash
HOST=
PORT=9000
DATABASE_URL=postgres://postgres:mypass@localhost:8010/matchoshop
DB_SSL_MODE=disable
TIMEZONE=Asia/Jakarta
```

### Export database environment for migration config
```bash
export DATABASE_URL=postgres://postgres:mypass@localhost:8010/matchoshop
export DB_SSL_MODE=disable
```

### Run database container

```bash
docker-compose up
```

## Run the App

### Get packages

```bash
go get .
```

### Delete unused packages if necessary

```bash
go mod tidy
```

### Update package vendor

```bash
go mod vendor
```

### Build the app

```bash
go build -o bin/matchoshop -v .
```

### Run the App

```bash
./bin/matchoshop
```

## Migration

### Create new migration
```bash
./migration.sh create AddSomeColumns
```

### Up migration
```bash
./migration.sh up
```

### Down migration
```bash
./migration.sh down
```

### Check migration status
```bash
./migration.sh status
```

## Mockup

### Generate new mockup object

```bash
mockery --all --dir=internal  --output=internal/mocks
```

## Unit Test

### Run unit test
go test -v ./path/to/test_file

```bash
go test -v ./internal/core/service
```

### Run unit test specific function
go test -v [function name]

```bash
go test -v -run TestProductCategory_Create_Success
```

## Deployment

### Deploy to heroku server

```bash
git push heroku master
```
