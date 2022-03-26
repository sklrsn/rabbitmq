package com.github.sklrsn.rabbitmq.streams;

import com.github.sklrsn.rabbitmq.config.Config;
import com.rabbitmq.client.Connection;
import lombok.NonNull;

public class RabbitMQMessageWriter implements MessageWriter {
    private Config config;
    private Connection rmqConnection;

    public RabbitMQMessageWriter(@NonNull Config config, @NonNull Connection connection) {
        this.config = config;
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
