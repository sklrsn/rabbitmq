package com.github.sklrsn.rabbitmq.consumer;

import com.github.sklrsn.rabbitmq.connection.RMQConnection;
import com.github.sklrsn.rabbitmq.streams.MessageReader;
import com.github.sklrsn.rabbitmq.streams.MessageWriter;
import com.github.sklrsn.rabbitmq.streams.RabbitMQMessageReader;
import com.github.sklrsn.rabbitmq.streams.RabbitMQMessageWriter;
import com.rabbitmq.client.Connection;

import java.io.IOException;
import java.util.concurrent.TimeoutException;

public class RMQConsumer {
    private MessageReader rmqReader;
    private MessageWriter rmqWriter;

    public void start() {
        try {
            Connection rmqConnection = RMQConnection.getConnection("localhost", "guest", "guest");
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
