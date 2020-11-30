package main

/*
Copyright (c) 2020 Luke Zhang luke-zhang-04.github.io/
BSD-3-Clause License
*/

import (
	"crypto/sha512"
	"encoding/hex"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	var outfile string     // Output file
	var isSymmetrical bool // If image should be symmetrical
	var width, height int  // Dimensions of the image
	var blocksize int      // Size of each square

	app := &cli.App{ // CLI Config
		Name:      "genim",
		Usage:     "Generate an image from a string",
		Copyright: "2020 Luke Zhang; BSD-3-Clause License",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "out",
				Aliases:     []string{"o"},
				Usage:       "Output generated PNG to `FILE`",
				Value:       "out.png",
				Destination: &outfile,
			},
			&cli.BoolFlag{
				Name:        "symmetrical",
				Aliases:     []string{"s", "sym"},
				Value:       false,
				Usage:       "If image should be symmetrical on the y-axis",
				Destination: &isSymmetrical,
			},
			&cli.IntFlag{
				Name:        "width",
				Aliases:     []string{"W"},
				Usage:       "Specify the `WIDTH` of the image",
				Value:       1024,
				Destination: &width,
			},
			&cli.IntFlag{
				Name:        "height",
				Aliases:     []string{"H"},
				Usage:       "Specify the `HEIGHT` of the image",
				Value:       1024,
				Destination: &height,
			},
			&cli.IntFlag{
				Name:        "blocksize",
				Aliases:     []string{"b", "block"},
				Usage:       "Specify the `SIZE` of each square",
				Value:       128,
				Destination: &blocksize,
			},
		},
		Action: func(c *cli.Context) error {
			// The input string as unix timestamp
			instring := strconv.FormatInt(time.Now().UnixNano(), 10)

			// If user provied an input string, use that instead
			if c.NArg() > 0 {
				instring = c.Args().Get(0)
			}

			// Hash the string with sha512
			hasher := sha512.New()
			_, err := hasher.Write([]byte(instring))

			if err != nil {
				log.Fatal(err)
			}

			// Encode hashed stirng to hex
			hashedString := hex.EncodeToString(hasher.Sum(nil))

			// Run generate function
			err = Generate(
				outfile,
				hashedString,
				isSymmetrical,
				width,
				height,
				blocksize,
			)

			return err
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
