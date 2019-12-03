package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func fuelCalc(mass int)(fuelMass, leftMass int){
	return mass / 3 - 2, mass % 3
}

func main(){
	f, err := os.Open("input.txt")
	//f, err := os.Open("input-validated.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sum := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		var fuelModule, modulo, fuelMass, fuelSum int
		_, err = fmt.Sscanf(s.Text(), "%d", &fuelModule)
		if err != nil {
			log.Fatalf("could not read value: %s, err", s.Text(), err )
		}

		// Initial calculation of the fuelMass added for the module mass. and the modulo.
		fuelMass, modulo = fuelCalc(fuelModule)

		// Adding the fuelMass to the fuelSum to get the initial mass of fuel for this module
		fuelSum = fuelMass

		// Calculating the amount of fuelMass to add to support itself.
		// While fuelMass is positive and over 0 calculate the required amount of Fuel for it.
		for fuelMass > 0 {
			fuelMass, modulo = fuelCalc(fuelMass)
			// If fuelMass is positive and over 0 Add the fuel Mass to the sum of fuelMass.
			if fuelMass > 0 {
				fuelSum = fuelSum + fuelMass
			}
		}

		// Adding the fuelSum for the module to the shared sum.
		sum = sum + fuelSum
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}