package main

import (
	"bufio"
	"encoding/base32"
	"flag"
	"fmt"
	"os"
)

// constants for the program input source
const (
	// input from pipe
	pipe = "pipe"

	// input as program argument
	arg = "arg"
)

func encode(source []byte) string {
	encoded := base32.StdEncoding.EncodeToString(source)
	return encoded
}

func decode(source []byte) string {
	decoded, err := base32.StdEncoding.DecodeString(string(source))
	if err != nil {
		// skip decoding
		fmt.Println(fmt.Errorf("could not decode input [%s], error: [%s]", source, err))
		os.Exit(-1)
	}
	return string(decoded)
}

func checkSource() (string, error) {
	// check whether the input comes from pipe or as program arg
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		return arg, nil
	} else if info.Size() > 0 {
		return pipe, nil
	} else {
		return "", fmt.Errorf("couldn't determine the source")
	}
}

func readInput(src string) string {
	var input string
	switch src {
	case arg:
		if flag.NArg() < 1 {
			flag.Usage()
			os.Exit(-1)
		}
		input = flag.Arg(0)
	case pipe:
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()

		if err != nil {
			fmt.Println(fmt.Errorf("could not read input, error: [%s]", err))
			os.Exit(-1)
		}

		input = string(line)
	default:
		fmt.Println(fmt.Errorf("unsupported source"))
		os.Exit(-1)
	}
	return input

}

func main() {
	// handling program flags
	dec := flag.Bool("d", false, "false or missing - encode / true - encode")
	flag.Parse()

	// determine the input source: pipe or program argument
	source, err := checkSource()
	if err != nil {
		fmt.Println("Error:", err)
		flag.Usage()
		os.Exit(-1)
	}

	// read the input
	input := readInput(source)

	if input == "" {
		fmt.Println("missing input")
		flag.Usage()
		os.Exit(-1)
	}

	if *dec == false {
		fmt.Println(encode([]byte(input)))
	} else {
		fmt.Println(decode([]byte(input)))
	}

}
