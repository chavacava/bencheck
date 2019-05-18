package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	deltaRE          = `([\+|\-]\d+(?:\.\d+)?)(?:%)`
	defaultTolerance = 0
)

func main() {
	tolerance := flag.Float64("t", defaultTolerance, "delta tolerance")
	flag.Parse()

	in := os.Stdin
	out := os.Stdout
	re := regexp.MustCompile(deltaRE)
	ec := 0 // exit code

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		delta := re.FindString(line)
		if delta == "" {
			continue
		}

		value, err := strconv.ParseFloat(strings.TrimSuffix(delta, "%"), 16)
		if err != nil {
			continue
		}

		if value > *tolerance {
			fmt.Fprintf(out, "%s delta is greater than %v\n", line, *tolerance)
			ec = 1
		}
	}

	os.Exit(ec)
}
