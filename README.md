# learn-pub-sub-starter (Peril)

This is the starter code used in Boot.dev's [Learn Pub/Sub](https://learn.boot.dev/learn-pub-sub) course.

./rabbit.sh start
go run ./cmd/server

docker build -t rabbitmq-stomp .
docker run -d --rm --name rabbitmq-stomp -p 5672:5672 -p 15672:15672 -p 61613:61613 rabbitmq-stomp
nc -vz localhost 61613