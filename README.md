# matchoshop
E-commerce for men's products

## Requirements

- [Golang](https://golang.org/) as main programming language.
- [Go Module](https://go.dev/blog/using-go-modules) for package management.
- [Goose](https://github.com/steinbacher/goose/) as migration tool.
- [Docker-compose](https://docs.docker.com/compose/) for running PostgreSQL Database locally.

## Setup
Prepare necessary environemt by rename .env.example to .env

```bash
HOST=
PORT=9000
DATABASE_URL=postgres://postgres:mypass@localhost:8010/matchoshop
DB_SSL_MODE=disable
TIMEZONE=Asia/Jakarta
```

Build Database Environment Container

```bash
docker-compose up
```

## Run the service

Get Go packages

```bash
go get .
```

Update Go package vendor

```bash
go mod vendor
```

Build the app

```bash
go build -o bin/matchoshop -v .
```


Run the proggram

```bash
./bin/matchoshop
```

## migration
### Create new migration
goose create AddSomeColumns

### Up migration
goose up

### Down migration
goose down

### Check migration status
goose status