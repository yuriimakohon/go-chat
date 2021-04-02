export POSTGRES_URL='postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable'
sudo docker run --name=chat-db --rm -e POSTGRES_PASSWORD=postgres -dp 5436:5432 postgres
migrate -database $POSTGRES_URL -path ./schema up
sudo docker exec -it chat-db /bin/bash
