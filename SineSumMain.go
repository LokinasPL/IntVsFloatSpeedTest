package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

// N is the number of elements in the slice to be computed, can be changed here to affect the rest of the program easily
const N = 1000000

// ComputeSineSum takes a slice of either ints or floats and returns the sum of the sine of each element as a float
func computeSineSum[T int | float64](data []T) float64 {
	//Initialize sum
	sum := 0.0
	//For each element of the slice in argument
	for _, v := range data {
		//Add the sine of the element to the sum
		sum += math.Sin(float64(v))
	}
	//Return the sum of all the sines
	return sum
}

func main() {
	//Get and parse flag from command line, with default value float
	typ := flag.String("type", "float", "Type of data: float or int")
	flag.Parse()

	//Set seed for random number generation
	rand.Seed(7)

	//Try to run the program with the given type flag, runType returns an error if the flag is invalid, and nil if the program ran successfully
	if err := runType(*typ); err != nil {
		//If the type is invalid, log the error and exit
		log.Fatalf("%v", err)
	}
}

// Function that runs the whole program depending on the flag given in the command line, is separated from the main fonction to be able to test it more easily
func runType(typ string) error {
	//Check the type flag entered
	switch typ {
	//If the type flag is int
	case "int":
		//Create slice of N ints
		arr := make([]int, N)
		//Fill slice with random ints between 0 and 1000
		for i := range arr {
			arr[i] = rand.Intn(1001)
		}
		//Note time before function call
		t0 := time.Now()
		//Call function on slice created
		sum := computeSineSum(arr)
		//Note delay since time before function call
		dt := time.Since(t0)
		//Print type, N, sum, and delay
		fmt.Printf("type=%s N=%d sum=%f elapsed=%s\n", typ, N, sum, dt)
		//Return nil error, since the type was valid and the function ran successfully
		return nil
	//If the type flag is float
	case "float":
		//Create slice of N floats
		arr := make([]float64, N)
		//Fill slice with random floats between 0 and 1
		for i := range arr {
			arr[i] = rand.Float64()
		}
		//Note time before function call
		t0 := time.Now()
		//Call function on slice created
		sum := computeSineSum(arr)
		//Note delay since time before function call
		dt := time.Since(t0)
		//Print type, N, sum, and delay
		fmt.Printf("type=%s N=%d sum=%f elapsed=%s\n", typ, N, sum, dt)
		//Return nil error, since the type was valid and the function ran successfully
		return nil
	//If the type flag is invalid, return an error
	default:
		return fmt.Errorf("unknown type %q: use float or int", typ)
	}
}
