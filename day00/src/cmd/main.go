package main

import (
	calc "anscombe/internal/calculation"
	"flag"
	"fmt"
)

func main() {
	meanFlag := flag.Bool("mean", false, "Calculate and print mean")
	medianFlag := flag.Bool("median", false, "Calculate and print median")
	modeFlag := flag.Bool("mode", false, "Calculate and print mode")
	sdFlag := flag.Bool("sd", false, "Calculate and print standard deviation")
	flag.Parse()

	numbers := calc.ParseInput()

	if !*meanFlag && !*medianFlag && !*modeFlag && !*sdFlag {
		*meanFlag = true
		*medianFlag = true
		*modeFlag = true
		*sdFlag = true
	}

	if *meanFlag {
		fmt.Println("Mean:", numbers.CalculateMean())
	}
	if *medianFlag {
		fmt.Println("Median:", numbers.CalculateMedian())
	}
	if *modeFlag {
		fmt.Println("Mode:", numbers.CalculateMode())
	}
	if *sdFlag {
		fmt.Println("SD:", numbers.CalculateSD())
	}
}
