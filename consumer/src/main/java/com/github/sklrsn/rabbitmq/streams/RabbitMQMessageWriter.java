package com.github.sklrsn.rabbitmq.streams;

import com.rabbitmq.client.Connection;

public class RabbitMQMessageWriter implements MessageWriter {

    private Connection rmqConnection;

    public RabbitMQMessageWriter(Connection connection) {
        this.rmqConnection = connection;
    }

    @Override
    public void open() {

    }

    @Override
    public void write() {

    }

    @Override
    public void close() {

    }
}
