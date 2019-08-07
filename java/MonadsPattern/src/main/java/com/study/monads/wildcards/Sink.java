package com.study.monads.wildcards;

public interface Sink<T> {
    void flush(T t);
}
