package com.github.sklrsn.rabbitmq;

import com.github.sklrsn.rabbitmq.config.Config;
import com.github.sklrsn.rabbitmq.consumer.RMQConsumer;
import com.github.sklrsn.rabbitmq.logger.ConsoleLogger;

public class Main {
    private static ConsoleLogger logger = ConsoleLogger.getInstance();

    public static void main(String[] args) {
        logger.info("starting consumer ..");

        //TODO: Isolate config
        Config config = new Config();
        config.setRabbitMQHost("rabbitmq");
        config.setRabbitMQUser("guest");
        config.setRabbitMQPass("guest");
        config.setQueues(new String[]{"logs.01", "logs.02", "logs.03", "logs.04"});

        RMQConsumer consumer = new RMQConsumer(config);
        consumer.start();

        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            logger.info("consumer is shutting down");
            Runtime.getRuntime().halt(1);
        }));

    }
}
