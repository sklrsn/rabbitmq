FROM rabbitmq:3.8.23

RUN rabbitmq-plugins enable --offline rabbitmq_management;

RUN rabbitmq-plugins enable --offline rabbitmq_consistent_hash_exchange

RUN rabbitmq-plugins enable --offline rabbitmq_shovel

RUN rabbitmq-plugins enable --offline rabbitmq_shovel_management

RUN  rm -rf /etc/rabbitmq/conf.d/management_agent.disable_metrics_collector.conf