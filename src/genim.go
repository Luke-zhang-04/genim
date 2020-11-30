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
	outfile := ""

	app := &cli.App{
		Name:      "genim",
		Usage:     "Generate an image from a string",
		Copyright: "2020 Luke Zhang; BSD-3-Clause License",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "out",
				Aliases:     []string{"o"},
				Usage:       "Output generated PNG to `FILE`",
				Destination: &outfile,
			},
		},
		Action: func(c *cli.Context) error {
			instring := strconv.FormatInt(time.Now().UnixNano(), 10)

			if c.NArg() > 0 {
				instring = c.Args().Get(0)
			}

			hasher := sha512.New()
			_, err := hasher.Write([]byte(instring))

			if err != nil {
				log.Fatal(err)
			}

			hashedString := hex.EncodeToString(hasher.Sum(nil))

			err = Generate(outfile, hashedString)

			return err
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
