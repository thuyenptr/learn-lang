package com.study.monads.optional;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

class Student {
    String name;
    Student() {
        this.name = "student name";
    }
}
public class LearnOptional {

    private static Student getStudent() {
        return Math.random() % 2 == 0 ? null : new Student();
    }

    public static void preJava8CheckNull() {
        Student student = getStudent();
        if (student != null) {
            System.out.println("Student info: " + student.name);
        } else {
            System.out.println("NULL");
        }
    }

    public static void Java8CheckNull() {
        Student student = getStudent();
        Optional<Student> optionalStudent = Optional.ofNullable(student);

        optionalStudent.ifPresent(std -> System.out.println(std.name));
    }

    private static List<String> getList() {
        return null;
    }

    private static Optional<List<String>> getOptionalList() {
        return Optional.empty(); // return null
    }

    public static void preJava8CheckNullAndReturn() {
        List<String> list = getList() != null ? getList() : new ArrayList<>();
        System.out.println("List value: " + list);
    }

    public static void Java8CheckNullAndReturn() {
        List<String> list = getOptionalList().orElse(new ArrayList<>());
        List<String> list1 = getOptionalList().orElseGet(ArrayList::new);

        System.out.println("List value: " + list);
        System.out.println("List1 value: " + list1);
    }

    public static void main(String[] args) {
    }
}
