cmake_minimum_required(VERSION 3.15)
project(ball)

set(CMAKE_CXX_STANDARD 11)

add_executable(ball main.cpp src/server.cpp include/server.h)