package com.github.sklrsn.rabbitmq.streams;

import com.github.sklrsn.rabbitmq.config.Config;
import com.rabbitmq.client.Channel;
import com.rabbitmq.client.Connection;
import lombok.NonNull;

import java.io.IOException;
import java.util.Optional;

public class RabbitMQMessageReader implements MessageReader {
    private Config config;
    private Connection rmqConnection;
    private Optional<Channel> channel;

    public RabbitMQMessageReader(@NonNull Config config, @NonNull Connection connection) {
        this.config = config;
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
