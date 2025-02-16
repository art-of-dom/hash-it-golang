package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/art-of-dom/hash-it/internal/hashdefs"
	"io"
	"os"
	"reflect"
	"strings"
)

func reverseStr(str string) (res string) {
	for _, s := range str {
		res = string(s) + res
	}
	return
}

type HashItArgs struct {
	Help    bool
	Reverse bool
	Verify  string
	Input   string
	Name    string
	File    string
}

func getArgs() (result HashItArgs) {
	flag.BoolVar(&result.Reverse, "reverse", false, "reverse input")
	flag.BoolVar(&result.Help, "help", false, "Prints help")
	flag.StringVar(&result.Verify, "verify", "", "Verify output instead of printing")
	flag.StringVar(&result.Input, "a", "", "ASCII Input")
	flag.StringVar(&result.Name, "n", "CCITT", "Name")
	flag.StringVar(&result.File, "f", "", "File input")
	flag.Parse()

	return
}

func main() {
	var ccittCrc []byte
	args := getArgs()
	crcmap := hashdefs.MapCrcs()

	if args.Help {
		flag.PrintDefaults()
	} else {

		if args.Reverse {
			args.Input = reverseStr(args.Input)
		}

		h := hashdefs.GetHash(strings.ToUpper(args.Name))

		if h == nil {
			fmt.Println("Unknown Hash. Posible hashes are:")
			fmt.Println()
			for k := range crcmap {
				fmt.Println(k)
			}
			fmt.Println()
			os.Exit(1)
		}

		if args.File != "" {
			f, err := os.Open(args.File)
			if err != nil {
				fmt.Printf("unable to read file: %v\n", err)
				os.Exit(4)
			}
			defer f.Close()
			stat, err := f.Stat()
			buf := make([]byte, stat.Size())
			n, err := f.Read(buf)
			if err == io.EOF {
				os.Exit(3)
			}
			if err != nil {
				fmt.Println(err)
				os.Exit(3)
			}
			if n > 0 {
				h.Write([]byte(buf))
				ccittCrc = h.Sum(nil)
			}
		} else {
			h.Write([]byte(args.Input))
			ccittCrc = h.Sum(nil)
		}

		if args.Verify != "" {
			v, err := hex.DecodeString(args.Verify)
			fmt.Printf("0x%04X\n", v)

			if err != nil {
				os.Exit(3)
			}
			if !reflect.DeepEqual(v, ccittCrc) {
				os.Exit(2)
			}
		} else {
			fmt.Printf("CRC is 0x%04X\n", ccittCrc) // prints "CRC is "
		}
	}
}
