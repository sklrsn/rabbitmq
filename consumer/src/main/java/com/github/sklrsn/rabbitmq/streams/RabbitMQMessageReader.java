package com.github.sklrsn.rabbitmq.streams;

import com.github.sklrsn.rabbitmq.config.Config;
import com.github.sklrsn.rabbitmq.logger.ConsoleLogger;
import com.rabbitmq.client.*;
import lombok.NonNull;

import java.io.IOException;

public class RabbitMQMessageReader implements MessageReader {
    private static ConsoleLogger logger = ConsoleLogger.getInstance();
    private Config config;
    private Connection rmqConnection;

    public RabbitMQMessageReader(@NonNull Config config, @NonNull Connection connection) {
        this.config = config;
        this.rmqConnection = connection;
    }

    @Override
    public void open() {
        if (!this.rmqConnection.isOpen()) {
            throw new RuntimeException("connection is closed");
        }
    }

    @Override
    public void read() throws IOException {
        for (String queue : config.getQueues()) {
            Channel channel = rmqConnection.openChannel().orElseThrow();
            channel.queueDeclarePassive(queue);
            channel.basicConsume(queue, true, "jconsumerctl", new Consumer() {
                @Override
                public void handleConsumeOk(String s) {
                }

                @Override
                public void handleCancelOk(String s) {
                }

                @Override
                public void handleCancel(String s) throws IOException {
                }

                @Override
                public void handleShutdownSignal(String s, ShutdownSignalException e) {
                    System.out.println("channel is shutting down");
                }

                @Override
                public void handleRecoverOk(String s) {
                }

                @Override
                public void handleDelivery(String s, Envelope envelope,
                        AMQP.BasicProperties basicProperties, byte[] bytes) throws IOException {
                    logger.info(bytes.toString());
                }
            });
        }

    }

    @Override
    public void close() {
    }
}
