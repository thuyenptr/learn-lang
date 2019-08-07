package com.study.recursion;

public class Factorial {
    public static long calFactory(int num) {
        // 0! = 1
        if (0 == num) {
            System.out.println("Put to stack value: " + num);
            System.out.println("Base case: " + num);
            return 1;
        }
        System.out.println("Put to stack value: " + num);

        long result = calFactory(num - 1);

        System.out.println("Get from stack value: " + num);

        return num * result;
    }

    public static long calFactory2(int num) {
        if (0 == num) {
            return 1;
        }
        return num * calFactory2(num - 1);
    }
}
