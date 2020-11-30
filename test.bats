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
