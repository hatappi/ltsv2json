package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/hatappi/ltsv2json/ltsv"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	inputPath string
	tailFlag  bool
)

func init() {
	flag.StringVar(&inputPath, "i", "", "input file path")
	flag.BoolVar(&tailFlag, "f", false, "fllow stdin")
	flag.Parse()
}

func getReader() (io.Reader, error) {
	if inputPath != "" {
		file, err := os.Open(inputPath)
		if err != nil {
			return nil, err
		}
		return file, nil
	} else {
		return os.Stdin, nil
		if terminal.IsTerminal(0) {
			return nil, fmt.Errorf("please set stdin or -f")
		}
	}
	return nil, fmt.Errorf("unreachable")
}

func main() {
	r, err := getReader()
	if err != nil {
		panic(err)
	}
	reader := ltsv.NewReader(r)

	if tailFlag {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				continue
			}
			if err != nil {
				panic(err)
			}
			b, err := json.Marshal(record)
			if err != nil {
				panic(err)
			}
			os.Stdout.Write(append(b, '\n'))
		}
	} else {
		records, err := reader.ReadAll()
		if err != nil {
			panic(err)
		}
		b, err := json.Marshal(records)
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(b)
	}
}
