package com.github.sklrsn.rabbitmq.config;

import lombok.Getter;
import lombok.Setter;

public class Config {
    @Getter
    @Setter
    private String rabbitMQHost;
    @Getter
    @Setter
    private String rabbitMQUser;
    @Getter
    @Setter
    private String rabbitMQPass;
    @Getter
    @Setter
    private String[] queues;
}
