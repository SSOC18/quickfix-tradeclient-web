# quickfix-tradeclient-web
Golang Web Application that can receives an order input and sends the message to RabbitMQ server.

Installation
------------

Make sure latest golang and rabbitmq are installed. Navigate to the quickfix-tradeclient-web repository directory, and run command

`go run webui.go`


Example
--------

Run the web application, and got to page:localhost:5004

Make sure the RabbitMQ-server is running, fill the form and sumbit. The order will be sent to exchange="" and queue="orders"

To see the message in queue from terminal, use rabbitmqadmin as follows:

>>>rabbitmqadmin get queue="orders" requeue=false


Other rabbitmqadmin useful commands:

To list available exchanges: rabbitmqadmin -V test list exchanges

To list details of available queues: rabbitmqadmin -f long -d 3 list queues

See other commands:

rabbitmqadmin --bash-completion
rabbitmqadmin --help
rabbitmqadmin help subcommands