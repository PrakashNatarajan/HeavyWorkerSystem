##To Start or Stop RabbitMQ##
sudo service rabbitmq-server start
sudo service rabbitmq-server stop
sudo service rabbitmq-server restart
sudo service rabbitmq-server status

https://www.rabbitmq.com/rabbitmqctl.8.html
rabbitmqctl stop_app
rabbitmqctl reset
rabbitmqctl start_app
rabbitmqctl delete_vhost test
rabbitmqctl add_vhost test
rabbitmqctl set_permissions -p /myvhost tonyg "^.*" ".*" ".*"
rabbitmqctl add_user tonyg changeit
rabbitmqctl set_user_tags tonyg administrator
rabbitmqctl set_permissions -p /myvhost tonyg ".*" ".*" ".*"
rabbitmqctl status