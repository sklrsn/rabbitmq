package com.github.sklrsn.rabbitmq.consumer;

import java.io.IOException;
import java.util.concurrent.TimeoutException;

import com.github.sklrsn.rabbitmq.config.Config;
import com.github.sklrsn.rabbitmq.connection.RMQConnection;
import com.github.sklrsn.rabbitmq.logger.ConsoleLogger;
import com.github.sklrsn.rabbitmq.streams.MessageReader;
import com.github.sklrsn.rabbitmq.streams.MessageWriter;
import com.github.sklrsn.rabbitmq.streams.RabbitMQMessageReader;
import com.github.sklrsn.rabbitmq.streams.RabbitMQMessageWriter;
import com.rabbitmq.client.Connection;
import lombok.NonNull;

public class RMQConsumer {
    private static ConsoleLogger logger = ConsoleLogger.getInstance();
    private Config config;
    private MessageReader rmqReader;
    private MessageWriter rmqWriter;

    public RMQConsumer(@NonNull Config config) {
        this.config = config;
    }

    public void start() {
        try {
            Connection rmqConnection = RMQConnection.getConnection(config.getRabbitMQHost(),
                    config.getRabbitMQUser(), config.getRabbitMQPass());
            logger.info("rabbitmq consumer connected ");

            rmqReader = new RabbitMQMessageReader(config, rmqConnection);
            rmqWriter = new RabbitMQMessageWriter(config, rmqConnection);

            rmqReader.open();
            rmqWriter.open();

        } catch (IOException e) {
            e.printStackTrace();
        } catch (TimeoutException e) {
            e.printStackTrace();
        }
    }
}
