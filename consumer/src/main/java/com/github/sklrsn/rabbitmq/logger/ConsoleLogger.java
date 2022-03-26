package com.github.sklrsn.rabbitmq.logger;

public class ConsoleLogger {
    private static final ConsoleLogger INSTANCE = new ConsoleLogger();

    public static ConsoleLogger getInstance() {
        return INSTANCE;
    }

    public void info(String message) {
        System.out.println(message);
    }
}
