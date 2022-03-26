package com.github.sklrsn.rabbitmq;

import com.github.sklrsn.rabbitmq.consumer.RMQConsumer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
    private static final Logger logger = LoggerFactory.getLogger(Main.class.getName());

    public static void main(String[] args) {
        logger.info("starting consumer ..");
        RMQConsumer consumer = new RMQConsumer();
        consumer.start();

        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            System.out.println("consumer is shutting down");
            System.exit(0);
        }));
    }
}
