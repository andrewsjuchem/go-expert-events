# Go Expert - Event Dispatch + RabbitMQ

RabbitMQ - Starting up container
```
docker-compose up -d --build
```

RabbitMQ - Consumer
```
go run cmd/consumer/main.go 
```

RabbitMQ - Producer
```
go run cmd/producer/main.go 
```

Event Dispacter - Testing
```
go run cmd/event_dispatcher/main.go 
```