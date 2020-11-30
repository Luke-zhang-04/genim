package main

/*
Copyright (c) 2020 Luke Zhang luke-zhang-04.github.io/
BSD-3-Clause License
*/

import (
	"fmt"
	"strconv"

	"github.com/fogleman/gg"
)

func hexToBase10(val string) int64 {
	num, err := strconv.ParseInt(val, 16, 64)

	if err == nil {
		return num
	}

	return 0
}

// Generate generates the image
func Generate(outfile, hashedString string) error {
	sliceStart := hexToBase10(string(hashedString[0])) // Start of slice to get colour
	color := hashedString[sliceStart : sliceStart+6]   // The colour of the blocks
	threshold := sliceStart                            // The amount before a block should be drawn
	background := "fff"                                // Background color
	blocksize := 128                                   // Size of each square
	dimensions := []int{1024, 1024}                    // Dimensions of image
	curX, curY := 0, 0                                 // Current x and y values
	index := 0                                         // Index of hashed string

	if threshold >= 13 {
		threshold = 12
	} else if threshold <= 3 {
		threshold = 4
	}

	if threshold >= 8 {
		background = "000"
	}

	img := gg.NewContext(dimensions[0], dimensions[1])
	img.DrawRectangle(0, 0, float64(dimensions[0]), float64(dimensions[1]))
	img.SetHexColor(background)
	img.Fill()

	for curY < dimensions[1] {
		for curX < dimensions[0] {
			fmt.Println(hexToBase10(string(hashedString[index])))
			if threshold >= hexToBase10(string(hashedString[index])) {
				img.DrawRectangle(float64(curX), float64(curY), float64(blocksize), float64(blocksize))
			}

			curX += blocksize
			index++

			if index >= len(hashedString) {
				index = 0
			}
		}
		curY += blocksize
		curX = 0
	}

	img.SetHexColor(color)
	img.Fill()

	img.SavePNG(outfile)

	return nil
}
