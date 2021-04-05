POSTGRES_URL='postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable'
migrate -database $POSTGRES_URL -path ./schema up