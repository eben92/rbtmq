publish:
	@go run main.go

consume: 
	@go run consumer.go


rabbit:
	@docker run -d --hostname rbtmq-go --name rbtmq-go -p 15672:15672 -p 5672:5672 rabbitmq:3-management
