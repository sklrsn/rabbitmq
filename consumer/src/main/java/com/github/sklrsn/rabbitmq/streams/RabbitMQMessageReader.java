package com.github.sklrsn.rabbitmq.streams;

import com.rabbitmq.client.Channel;
import com.rabbitmq.client.Connection;

import java.io.IOException;
import java.util.Optional;

public class RabbitMQMessageReader implements MessageReader {
    private Connection rmqConnection;
    private Optional<Channel> channel;

    public RabbitMQMessageReader(Connection connection) {
        this.rmqConnection = connection;
    }

    @Override
    public void open() {
        try {
            this.channel = rmqConnection.openChannel();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    @Override
    public void read() {

    }

    @Override
    public void close() {

    }
}
