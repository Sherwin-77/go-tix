# Golang Echo Template

This is a simple golang echo template that implement golang template project structure

## Getting Started
1. Use this template to create a new project
2. Modify rename-module.sh (or rename-module.ps1) variable to match your project name
3. Run `rename-module.sh` or `rename-module.ps1`
4. Copy .env.example and fill in your environment variables

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


## Progress

| Name                          | Status |
|-------------------------------|--------|
| Core                          | 🚧     |
| Unit Test                     | 🚧     |
| Docker                        | ✔️     |
| Migration                     | ✔️     |
| API Documentation             | 🚧      |
| Auto handle role & permission | ❌      |
| Github Workflow               | ❌      |

If you want to contribute, feel free to open pull request
