package com.github.sklrsn.rabbitmq.streams;

import com.github.sklrsn.rabbitmq.config.Config;
import com.github.sklrsn.rabbitmq.logger.ConsoleLogger;
import com.rabbitmq.client.*;
import lombok.NonNull;

import java.io.IOException;
import java.util.Collection;
import java.util.Collections;
import java.util.Optional;

public class RabbitMQMessageReader implements MessageReader {
    private static ConsoleLogger logger = ConsoleLogger.getInstance();
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
        if (!this.channel.isPresent()) {
            throw new RuntimeException("channel is not open");
        }
    }

    @Override
    public void read() {
        for (String queue : config.getQueues()) {
            this.channel.ifPresent(ch -> {
                try {
                    ch.queueDeclarePassive(queue);

                    ch.basicConsume(queue, true, new Consumer() {
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

                        }

                        @Override
                        public void handleRecoverOk(String s) {

                        }

                        @Override
                        public void handleDelivery(String s, Envelope envelope, AMQP.BasicProperties basicProperties, byte[] bytes) throws IOException {
                            logger.info(envelope.getExchange());
                            logger.info(s);
                            logger.info(bytes.toString());
                        }
                    });
                } catch (IOException e) {
                    e.printStackTrace();
                }
            });
        }
    }

    @Override
    public void close() {

    }
}
