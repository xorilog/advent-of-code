package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sum := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		var fuelModule int
		_, err = fmt.Sscanf(s.Text(), "%d", &fuelModule)
		if err != nil {
			log.Fatalf("could not read value: %s, err", s.Text(), err )
		}
		divided := fuelModule / 3 - 2
		sum = sum + divided
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}