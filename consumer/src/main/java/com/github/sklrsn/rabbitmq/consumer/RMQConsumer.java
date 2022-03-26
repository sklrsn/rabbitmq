package com.github.sklrsn.rabbitmq.consumer;

import java.io.IOException;
import java.util.concurrent.TimeoutException;

import com.github.sklrsn.rabbitmq.connection.RMQConnection;
import com.github.sklrsn.rabbitmq.logger.ConsoleLogger;
import com.github.sklrsn.rabbitmq.streams.MessageReader;
import com.github.sklrsn.rabbitmq.streams.MessageWriter;
import com.github.sklrsn.rabbitmq.streams.RabbitMQMessageReader;
import com.github.sklrsn.rabbitmq.streams.RabbitMQMessageWriter;
import com.rabbitmq.client.Connection;

public class RMQConsumer {
    private static ConsoleLogger logger = ConsoleLogger.getInstance();
    private MessageReader rmqReader;
    private MessageWriter rmqWriter;

    public void start() {
        try {
            Connection rmqConnection = RMQConnection.getConnection("rabbitmq",
                    "guest", "guest");
            logger.info("rabbitmq connected ");
            
            rmqReader = new RabbitMQMessageReader(rmqConnection);
            rmqWriter = new RabbitMQMessageWriter(rmqConnection);

            rmqReader.open();
            rmqWriter.open();
        } catch (IOException e) {
            e.printStackTrace();
        } catch (TimeoutException e) {
            e.printStackTrace();
        }
    }
}
