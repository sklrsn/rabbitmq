package com.github.sklrsn.rabbitmq.streams;

public interface MessageWriter {
    void open();

    void write();

    void close();
}
