package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"strings"
)

type Generation struct {
	kind           string
	sprinkle       string
	delimeter      string
	count          int
	sprinkleRarity int
	newLineAfter   int
	truncate       int
}

func main() {
	// Define command-line flags for the two modes
	words := flag.Bool("w", false, "Words")
	letters := flag.Bool("letters", false, "Letters")
	numbers := flag.Bool("numbers", false, "Numbers")
	symbols := flag.Bool("symbols", false, "Symbols")
	newLineAfter := flag.Int("nl", 0, "New line after x characters")
	truncate := flag.Int("tr", 0, "Truncate")
	delimeter := flag.String("d", "", "Specify delimeter")
	sprinkle := flag.String("sprinkles", "", "Sprinkle character")
	sprinkleRarity := flag.Int("sr", 5, "Sprinkle rarity")
	count := flag.Int("count", 10, "Number of whatever you want to generate to generate")
	helpFlag := flag.Bool("help", false, "Help")

	flag.Parse()

	// Display help message if -help is passed or no arguments are given
	if *helpFlag || flag.NFlag() == 0 {
		fmt.Println("you need help buddy no two ways about it")
		return
	}

	generation := Generation{
		sprinkle:       *sprinkle,
		delimeter:      *delimeter,
		count:          *count,
		sprinkleRarity: *sprinkleRarity,
		newLineAfter:   *newLineAfter,
		truncate:       *truncate,
	}

	// Determine which mode to run
	if *words {
		generation.kind = "words"
		generate(generation)
	} else if *letters {
		generation.kind = "letters"
		generate(generation)
	} else if *numbers {
		generation.kind = "numbers"
		generate(generation)
	} else if *letters && *numbers {
		generation.kind = "letters numbers"
		generate(generation)
	} else if *symbols && *numbers {
		generation.kind = "numbers symbols"
		generate(generation)
	} else if *letters && *symbols {
		generation.kind = "letters symbols"
		generate(generation)
	} else if *letters && *symbols && *numbers {
		generation.kind = "letters numbers symbols"
		generate(generation)
	} else {
		fmt.Println("Error: Please specify a mode.")
	}
}

func give(listType string) []string {
	numbers := strings.Split("1234567890", "")
	letters := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	symbols := strings.Split("!,\\;:@#$%^&*()-=+_/{[`~]}", "")

	switch listType {
	case "numbers":
		return numbers
	case "letters":
		return letters
	case "symbols":
		return symbols
	}
	fmt.Println("erm what the sigma")
	return numbers
}

func sprinkling(sprinkle string, rarity int) string {
	if rand.IntN(rarity) == 1 {
		return sprinkle
	}
	return ""
}

func generate(generation Generation) {
	var output string
	switch generation.kind {
	case "letters":
		letters := give("letters")
		for i := 0; i < generation.count; i++ {
			output += letters[rand.IntN(len(letters))]
			output += sprinkling(generation.sprinkle, generation.sprinkleRarity)
			output += generation.delimeter
		}
	case "numbers":
		numbers := give("numbers")
		for i := 0; i < generation.count; i++ {
			output += numbers[rand.IntN(len(numbers))]
			output += sprinkling(generation.sprinkle, generation.sprinkleRarity)
			output += generation.delimeter
		}
	case "letters numbers":
		letters := give("letters")
		numbers := give("numbers")
		for i := 0; i < generation.count; i++ {
			if rand.IntN(2) == 1 {
				output += letters[rand.IntN(len(letters))]
			} else {
				output += numbers[rand.IntN(len(numbers))]
			}
			output += sprinkling(generation.sprinkle, generation.sprinkleRarity)
			output += generation.delimeter
		}
	case "letters symbols":
		letters := give("letters")
		symbols := give("symbols")
		for i := 0; i < generation.count; i++ {
			if rand.IntN(2) == 1 {
				output += letters[rand.IntN(len(letters))]
			} else {
				output += symbols[rand.IntN(len(symbols))]
			}
			output += sprinkling(generation.sprinkle, generation.sprinkleRarity)
			output += generation.delimeter
		}
	case "numbers symbols":
		numbers := give("numbers")
		symbols := give("symbols")
		for i := 0; i < generation.count; i++ {
			if rand.IntN(2) == 1 {
				output += numbers[rand.IntN(len(numbers))]
			} else {
				output += symbols[rand.IntN(len(symbols))]
			}
			output += sprinkling(generation.sprinkle, generation.sprinkleRarity)
			output += generation.delimeter
		}
	case "letters numbers symbols":
		letters := give("letters")
		numbers := give("numbers")
		symbols := give("symbols")
		for i := 0; i < generation.count; i++ {
			lns := rand.IntN(3)
			if lns == 1 {
				output += letters[rand.IntN(len(letters))]
			} else if lns == 2 {
				output += numbers[rand.IntN(len(numbers))]
			} else {
				output += symbols[rand.IntN(len(symbols))]
			}
			output += sprinkling(generation.sprinkle, generation.sprinkleRarity)
			output += generation.delimeter
		}
	case "words":
	}

	if generation.truncate != 0 && len(output) > generation.truncate {

		output = output[:generation.truncate]
	}

	if generation.newLineAfter != 0 && len(output) > generation.newLineAfter {
		output = insertNewlines(output, generation.newLineAfter)
	}

	fmt.Println(output)
}

func insertNewlines(s string, interval int) string {
	var result string
	for i := 0; i < len(s); i += interval {
		end := i + interval
		if end > len(s) {
			end = len(s)
		}
		result += s[i:end] + "\n"
	}
	return result
}
