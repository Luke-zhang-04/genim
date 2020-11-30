#!/bin/bash

@test "should display CLI help" {
    output="$(./genim help)"
    [ "$(echo "$output" | grep COMMANDS)" ]
    [ "$(echo "$output" | grep OPTIONS)" ]
}

@test "should generate an image" {
    ./genim
    [ "$(ls out.png)" ]
    rm -rf out.png
}

@test "should generate an image to a specified directory" {
    ./genim -o test.png
    [ "$(ls test.png)" ]
    rm -rf test.png
}

@test "should generate an image with the --symmetrical flag" {
    ./genim --symmetrical
    [ "$(ls out.png)" ]
    rm -rf out.png
}

@test "should generate an image with a string" {
    ./genim bruh
    [ "$(ls out.png)" ]
    rm -rf out.png
}

@test "should generate an image with width and height restraints" {
    ./genim --width 512 --height 512 --block 64
    [ "$(ls out.png)" ]
    [ "$(file out.png grep | "512 x 512")" ]
    rm -rf out.png
}

@test "should be able to generate images random" {
    ./genim --rand
    [ "$(ls out.png)" ]
    rm -rf out.png
}
