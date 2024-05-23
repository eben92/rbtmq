run:
	@go run main.go

rabbit:
	@docker run -d --hostname rabbit-mq --name rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
