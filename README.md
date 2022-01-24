# matchoshop
E-commerce for men's products

## migration
goose -dir db/migration create initial-scheme sql
goose -dir db/migration postgres "postgres://postgres:mypass@localhost:8010/matchoshop?sslmode=disable" up