# Streaming RabbitMQ
This repository is an example of how streaming works in RabbitMQ in Go.

## Requirements
You will need an rabbitMQ with Advertised server settings and enabled Streaming plugin

```bash
docker run -d --name rabbitmq -p 5552:5552 -p 5672:5672 -p 15672:15672 -e RABBITMQ_SERVER_ADDITIONAL_ERL_ARGS='-rabbitmq_stream advertised_host localhost -rabbit loopback_users "none"' rabbitmq:3.9-management

docker exec rabbitmq rabbitmq-plugins enable rabbitmq_stream # Enable Stream plugin
docker exec rabbitmq rabbitmq-plugins enable rabbitmq_stream_management # Enable UI to manage streams
```

## Appendix
[Stream feature matrix](https://www.rabbitmq.com/stream-core-plugin-comparison.html)
[RabbitMQCTL](https://www.rabbitmq.com/rabbitmqctl.8.html)
[RabbitMQ Consumption](https://www.rabbitmq.com/streams.html#consuming)
[RabbitMQ Data safety](https://www.rabbitmq.com/streams.html#data-safety)
[SuperStreams](https://github.com/rabbitmq/rabbitmq-stream-go-client/issues/149)