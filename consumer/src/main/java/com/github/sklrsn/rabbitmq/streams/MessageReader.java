package com.github.sklrsn.rabbitmq.streams;

import java.io.IOException;

public interface MessageReader {
    void open();

    void read() throws IOException;

    void close();
}
