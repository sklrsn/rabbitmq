package com.github.sklrsn.rabbitmq;

import com.github.sklrsn.rabbitmq.consumer.RMQConsumer;
import com.github.sklrsn.rabbitmq.logger.ConsoleLogger;

public class Main {
    private static ConsoleLogger logger = ConsoleLogger.getInstance();

    public static void main(String[] args) {
        logger.info("starting consumer ..");

        RMQConsumer consumer = new RMQConsumer();
        consumer.start();

        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            logger.info("consumer is shutting down");
            Runtime.getRuntime().halt(1);
        }));
    }
}
