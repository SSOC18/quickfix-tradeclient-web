QuickFix TradeClient Web User Interface
-------------------------------------------

Web user interface + quickfix trade client

The web UI takes an order from the user and passes it to RabbitMQ.

The trade client subscribes to RabbitMQ and executes orders received.

The trade client is based on the one from https://github.com/quickfixgo/examples


Installation
~~~~~~~~~~~~

```
go get github.com/streadway/amqp
make install
```

Usage
~~~~~

```
make run_tradeclient
make run_webui
```

Other
~~~~~

Setting Up the User Interface:
Can use makefile (make sure form.html is in the same bin folder)

OR
webui.go sends order through localhost:8080
RabbitMQ server running on localhost:5672
tradeclient.go back end interface


(The RabbitMQ server scripts are installed into /usr/local/sbin. This is not automatically added to your path, so you may wish to add
PATH=$PATH:/usr/local/sbin to your .bash_profile or .profile. The server can then be started with rabbitmq-server.)

The back end tradeclient receives messages from the RabbitMQ server and translates using FIX messages protocols.
