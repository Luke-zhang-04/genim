#!/bin/bash

@test "should generate an image" {
    ./genim
    [ "$(ls out.png)" ]
}
