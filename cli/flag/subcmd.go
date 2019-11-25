package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func flagUsage() {
	usageText := `subcmd is an example cli tool.

Usage:
subcmd command [arguments]
The commands are:
convhex    convert Number to Hex
convbinary convert Number to binary
Use "subcmd [command] --help" for more information about command`

	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}

func main() {
	flag.Usage = flagUsage
	convHexCmd := flag.NewFlagSet("convhex", flag.ExitOnError)
	convBinaryCmd := flag.NewFlagSet("convbinary", flag.ExitOnError)

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	switch os.Args[1] {
	case "convhex":
		i := convHexCmd.Int64("i", 0, "Convert nuber to hex")
		convHexCmd.Parse(os.Args[2:])
		fmt.Println(strconv.FormatInt(*i, 16))
	case "convbinary":
		i := convBinaryCmd.Int64("i", 0, "Convert nuber to binary")
		convBinaryCmd.Parse(os.Args[2:])
		fmt.Println(strconv.FormatInt(*i, 2))
	default:
		flag.Usage()
	}
}
