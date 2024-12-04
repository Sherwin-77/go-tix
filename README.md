# GoTix

## Running
1. Install dependency with `go mod download`
2. Run `make serve` to start server

Alternatively, you can run through docker just by running
```bash
docker compose up -d
```

## Migration
Migration handled using [golang-migrate](https://github.com/golang-migrate/migrate). You need to [install](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) it if not using docker
Migration automatically run everytime docker container start
1. To apply migration, run `make migrate` 
2. To create migration, run `make migration name=migration_name`
3. To rollback migration, run `make rollback name=migration_name`
