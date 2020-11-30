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
// outfile - output file of image
// hashedString - hashed string with hex numbers to generate image from
// isSymmetric - if image should be symmetric along the y axis
// width, height - dimensions of the image
// blocksize - size of each square
func Generate(
	outfile,
	hashedString string,
	isSymmetric bool,
	width,
	height,
	blocksize int,
) error {
	sliceStart := hexToBase10(string(hashedString[0])) // Start of slice to get colour
	color := hashedString[sliceStart : sliceStart+6]   // The colour of the blocks
	threshold := sliceStart                            // The amount before a block should be drawn
	const background string = "fff"                    // Background color
	curX, curY := 0, 0                                 // Current x and y values
	index := 0                                         // Index of hashed string

	if threshold >= 13 { // If threshold is too high, lower it
		threshold = 12
	} else if threshold <= 3 { // If it's too low, raise is
		threshold = 4
	}

	img := gg.NewContext(width, height)                      // New canvas
	img.DrawRectangle(0, 0, float64(width), float64(height)) // background
	img.SetHexColor(background)
	img.Fill()

	for curY < height { // For each row
		leftSide := []bool{}                    // Left side of drawing for symmetry purposes
		reverseIndex := width/(blocksize*2) - 1 // Reversed index for going backwards

		for curX < width { // For each column
			if isSymmetric && curX >= width/2 {
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
