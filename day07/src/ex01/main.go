package main

import (
	"log"
	"moneybag/ex00"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	for i := 0; i < 1000; i++ {
		result := ex00.MinCoins2(10000, []int{1, 5, 10, 25, 50, 100, 500, 1000})
		log.Println(result)
	}
}
