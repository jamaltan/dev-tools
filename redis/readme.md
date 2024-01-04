docker run -d --name redis-server -p 6379:6379 redis --requirepass "redis123" --appendonly yes

docker exec -it redis /bin/bash
docker exec -it redis-server redis-cli



docker run -d --name redis-container \
    -v /path/to/your/host/directory:/data \
    redis:latest
