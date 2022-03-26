package com.github.sklrsn.rabbitmq.streams;

public interface MessageReader {
    void open();

    void read();

    void close();
}
