FROM rabbitmq:latest

RUN rabbitmq-plugins enable --offline rabbitmq_management;

RUN rabbitmq-plugins enable --offline rabbitmq_consistent_hash_exchange

RUN rabbitmq-plugins enable --offline rabbitmq_shovel

RUN rabbitmq-plugins enable --offline rabbitmq_shovel_management

RUN  echo management_agent.disable_metrics_collector = false >> /etc/rabbitmq/conf.d/management_agent.disable_metrics_collector.conf