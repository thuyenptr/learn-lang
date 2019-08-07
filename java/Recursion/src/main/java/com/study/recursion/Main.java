package com.study.recursion;

import java.io.BufferedReader;
import java.io.File;
import java.io.IOException;
import java.io.InputStreamReader;

public class Main {
    public static void main(String[] args) {
//        CountDown.countDown(3);
//        long result = Factorial.calFactory(3);
//        System.out.println("Result: " + result);

        InputStreamReader inputStreamReader = new InputStreamReader(System.in);
        BufferedReader keyboard = new BufferedReader(inputStreamReader);
        try {
            String inputString = keyboard.readLine();
            System.out.println(inputString);
        } catch (IOException e) {
            e.printStackTrace();
        }

        File[] listFiles = new File(".").listFiles(File::isHidden);
    }
}
