package main

/*
Copyright (c) 2020 Luke Zhang luke-zhang-04.github.io/
BSD-3-Clause License
*/

import (
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
func Generate(outfile, hashedString string, isSymmetric bool) error {
	sliceStart := hexToBase10(string(hashedString[0])) // Start of slice to get colour
	color := hashedString[sliceStart : sliceStart+6]   // The colour of the blocks
	threshold := sliceStart                            // The amount before a block should be drawn
	const background string = "fff"                    // Background color
	blocksize := 128                                   // Size of each square
	dimensions := []int{1024, 1024}                    // Dimensions of image
	curX, curY := 0, 0                                 // Current x and y values
	index := 0                                         // Index of hashed string

	if threshold >= 13 { // If threshold is too high, lower it
		threshold = 12
	} else if threshold <= 3 { // If it's too low, raise is
		threshold = 4
	}

	img := gg.NewContext(dimensions[0], dimensions[1])                      // New canvas
	img.DrawRectangle(0, 0, float64(dimensions[0]), float64(dimensions[1])) // background
	img.SetHexColor(background)
	img.Fill()

	for curY < dimensions[1] { // For each row
		leftSide := []bool{}                            // Left side of drawing for symmetry purposes
		reverseIndex := dimensions[0]/(blocksize*2) - 1 // Reversed index for going backwards

		for curX < dimensions[0] { // For each column
			if isSymmetric && curX >= dimensions[0]/2 {
				if leftSide[reverseIndex] {
					img.DrawRectangle(float64(curX), float64(curY), float64(blocksize), float64(blocksize))
				}

				reverseIndex--
			} else {
				shouldDrawBlock := hexToBase10(string(hashedString[index])) >= threshold
				leftSide = append(leftSide, shouldDrawBlock)

				// If current hex value is greater than threshold, draw the rectangle
				if shouldDrawBlock {
					img.DrawRectangle(float64(curX), float64(curY), float64(blocksize), float64(blocksize))
				}
			}

			// Increment current x and index
			curX += blocksize
			index++

			if index >= len(hashedString) { // Reset index if index is greater than hashed string length
				index = 0
			}
		}

		curY += blocksize // Increment current Y
		curX = 0          // Reset current x
	}

	img.SetHexColor(color)
	img.Fill()

	img.SavePNG(outfile)

	return nil
}
