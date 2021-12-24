FROM rabbitmq:latest

RUN rabbitmq-plugins enable --offline rabbitmq_management;

RUN rabbitmq-plugins enable --offline rabbitmq_consistent_hash_exchange

RUN rabbitmq-plugins enable --offline rabbitmq_shovel

RUN rabbitmq-plugins enable --offline rabbitmq_shovel_management

RUN  rm -rf /etc/rabbitmq/conf.d/management_agent.disable_metrics_collector.conf

RUN cp /plugins/rabbitmq_management-*/priv/www/cli/rabbitmqadmin /usr/local/bin/rabbitmqadmin; \
    [ -s /usr/local/bin/rabbitmqadmin ]; \
    chmod +x /usr/local/bin/rabbitmqadmin; \
    apt-get update; \
    apt-get install -y --no-install-recommends python3; \
    rm -rf /var/lib/apt/lists/*; \
    rabbitmqadmin --version