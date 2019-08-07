package com.study.recursion;

public class CountDown {
    public static void countDown(int num) {
        if (num == 0) {
            System.out.println("Push to stack value: " + num);
            System.out.println("Finished.");
        } else {
            System.out.println("Push to stack value: " + num);
            countDown(num - 1);
        }

        System.out.println("Get value from stack: " + num);
    }

    public boolean isDone() {
        return true;
    }
}
