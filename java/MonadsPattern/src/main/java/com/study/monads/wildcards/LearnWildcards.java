package com.study.monads.wildcards;

import java.util.ArrayList;
import java.util.Collection;

class SinkImpl<T> implements Sink<T> {

    @Override
    public void flush(T t) {
        System.out.println("Flushing....");
    }
}

public class LearnWildcards {

    // <T> is type parameter
    public static <T> T writeAll(Collection<? extends T> collection, Sink<? super T> sink) {
        T t = null;

        for (T col : collection) {
            t = col;
            sink.flush(t);
        }
        return t;
    }

    public static void main(String[] args) {
        Sink<Object> sink = new SinkImpl<>();
        Collection<String> col = new ArrayList<>();

        String str = writeAll(col, sink);
    }
}
