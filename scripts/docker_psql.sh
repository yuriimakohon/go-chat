NAME=chat-db
sudo docker run --name=$NAME --rm -e POSTGRES_PASSWORD=postgres -dp 5436:5432 postgres
sh migrate_up.sh
sudo docker exec -it $NAME /bin/bash
