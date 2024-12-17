package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hyacinthus/mp3join"
)

func main() {
	joiner := mp3join.New()

	n := flag.Int("n", 1, "how many times to loop")
	inputFileName := flag.String("i", "", "input file to loop")
	outputFileName := flag.String("o", "output.mp3", "output file")
	flag.Parse()

	for range *n {
		func() {
			reader, err := os.Open(*inputFileName)
			defer reader.Close()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = joiner.Append(reader)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}()
	}

	dest := joiner.Reader()
	destFile, err := os.Create(*outputFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = destFile.ReadFrom(dest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("looped '%s' %d-times and saved to '%s'\n", *inputFileName, *n, *outputFileName)
}
