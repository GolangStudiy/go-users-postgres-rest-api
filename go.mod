module github.com/GolangStudiy/go-users-postgres-rest-api/server

go 1.18

replace github.com/GolangStudiy/go-users-postgres-rest-api/databaseclient => ./src/configurations/databaseclient

replace github.com/GolangStudiy/go-users-postgres-rest-api/postgrescontainer => ./tests/utils/postgrescontainer
