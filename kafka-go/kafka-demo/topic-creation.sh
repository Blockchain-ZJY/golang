# 使用内部网络地址作为引导服务器
docker exec -it kafka kafka-topics --create \
    --bootstrap-server kafka:29092 \
    --topic orders.created \
    --partitions 6 \
    --replication-factor 2

docker exec -it kafka kafka-topics --create \
    --bootstrap-server localhost:29092 \
    --topic orders.processed \
    --partitions 6 \
    --replication-factor 2

docker exec -it kafka kafka-topics --create \
    --bootstrap-server localhost:29092 \
    --topic orders.failed \
    --partitions 2 \
    --replication-factor 2