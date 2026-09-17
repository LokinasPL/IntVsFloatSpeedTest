package main

import (
	"math"
	"math/rand"
	"testing"
)

func TestComputeSineSum(t *testing.T) {
	//// Test with small int slice
	// Initialize test values
	intSliceSmall := []int{0, 1, 2}
	// Manually calculate the expected output
	expectedInt := math.Sin(0) + math.Sin(1) + math.Sin(2)
	// Calculate the output of the function on this input
	gotInt := computeSineSum(intSliceSmall)
	// Allow small margin of error for floating point comparison
	if math.Abs(gotInt-expectedInt) > 1e-9 {
		// Call error if the ouput is too far from the expected output
		t.Errorf("computeSineSum([]int{0,1,2}) = %f; want %f", gotInt, expectedInt)
	}

	//// Test with large int slice (100k elements)
	// Initialize test values, use for-loop to populate large slice
	intSliceLarge := make([]int, 100000)
	for i := range 100000 {
		intSliceLarge[i] = i
	}
	// Manually calculate the expected output with for-loop
	expectedIntLarge := 0.0
	for _, v := range intSliceLarge {
		expectedIntLarge += math.Sin(float64(v))
	}
	// Calculate the output of the function on this input
	gotIntLarge := computeSineSum(intSliceLarge)
	// Allow small margin of error for floating point comparison
	if math.Abs(gotIntLarge-expectedIntLarge) > 1e-9 {
		// Call error if the ouput is too far from the expected output
		t.Errorf("computeSineSum([]int{0,1,...,99999}) = %f; want %f", gotIntLarge, expectedIntLarge)
	}

	//// Test with small float64 slice
	// Initialize test values
	floatSlice := []float64{0.0, 0.1, 0.5}
	// Manually calculate the expected output
	expectedFloat := math.Sin(0.0) + math.Sin(0.1) + math.Sin(0.5)
	// Calculate the output of the function on this input
	gotFloat := computeSineSum(floatSlice)
	// Allow small margin of error for floating point comparison
	if math.Abs(gotFloat-expectedFloat) > 1e-9 {
		// Call error if the ouput is too far from the expected output
		t.Errorf("computeSineSum([]float64{0.0,0.1,0.5}) = %f; want %f", gotFloat, expectedFloat)
	}

	//// Test with large float64 slice (100k elements)
	floatSliceLarge := make([]float64, 100000)
	// Initialize test values, use for-loop to populate large slice
	for i := range 100000 {
		floatSliceLarge[i] = float64(i) * 0.01
	}
	// Manually calculate the expected output with for-loop
	expectedFloatLarge := 0.0
	for _, v := range floatSliceLarge {
		expectedFloatLarge += math.Sin(v)
	}
	// Calculate the output of the function on this input
	gotFloatLarge := computeSineSum(floatSliceLarge)
	// Allow small margin of error for floating point comparison
	if math.Abs(gotFloatLarge-expectedFloatLarge) > 1e-9 {
		// Call error if the ouput is too far from the expected output
		t.Errorf("computeSineSum([]float64{0.0,0.01,...,999.99}) = %f; want %f", gotFloatLarge, expectedFloatLarge)
	}

	//// Test with empty slice of ints
	// Initialize empty slice of ints
	emptyInt := []int{}
	// Calculate the output of the function on this input
	gotEmpty := computeSineSum(emptyInt)
	// Check if the output is 0.0, which is the expected output for an empty slice
	if gotEmpty != 0.0 {
		// Call error if the ouput is not 0.0
		t.Errorf("computeSineSum([]int{}) = %f; want 0.0", gotEmpty)
	}

	//// Test with empty slice of floats
	// Initialize empty slice of floats
	emptyFloat := []float64{}
	// Calculate the output of the function on this input
	gotEmptyFloat := computeSineSum(emptyFloat)
	// Check if the output is 0.0, which is the expected output for an empty slice
	if gotEmptyFloat != 0.0 {
		//Call error if the ouput is not 0.0
		t.Errorf("computeSineSum([]float64{}) = %f; want 0.0", gotEmptyFloat)
	}
}

func TestRunType(t *testing.T) {
	//// Test with invalid type flag
	// Specify the type that will be tested
	invalidType := "badType"
	// Try to run the program with the invalid type flag, runType should return an error
	err := runType(invalidType)
	if err == nil {
		// If no error is returned, call error since we expected an error for an invalid type flag
		t.Fatalf("expected error for type %q, got nil", invalidType)
	}
	// Check if the error message is the expected one
	want := "unknown type \"badType\": use float or int"
	if err.Error() != want {
		// If the error message is not the expected one, call error with the details of the failure
		t.Fatalf("runType(%q) error = %q; want %q", invalidType, err.Error(), want)
	}

	//// Test with valid type flag "int"
	// Try to run the program with the valid type flag, runType should return nil error
	err = runType("int")
	if err != nil {
		// If an error is returned, call error since we expected no error for a valid type flag
		t.Fatalf("unexpected error for type %q: %v", "int", err)
	}

	//// Test with valid type flag "float"
	// Try to run the program with the valid type flag, runType should return nil error
	err = runType("float")
	if err != nil {
		// If an error is returned, call error since we expected no error for a valid type flag
		t.Fatalf("unexpected error for type %q: %v", "float", err)
	}

	//// Test with valid type flag "INT" (case sensitivity test)
	// Try to run the program with the valid type flag in an invalid case, runType should return an error since the type flag is case sensitive
	err = runType("INT")
	if err == nil {
		// If no error is returned, call error since we expected an error for an invalid type flag
		t.Fatalf("expected error for type %q, got nil", "INT")
	}
	// Check if the error message is the expected one
	want = "unknown type \"INT\": use float or int"
	if err.Error() != want {
		// If the error message is not the expected one, call error with the details of the failure
		t.Fatalf("runType(%q) error = %q; want %q", "INT", err.Error(), want)
	}
}

// Slice structure of the percentages of data to be benchmarked
var percents = []struct {
	name    string
	percent int
}{
	{"1%", 1},
	{"10%", 10},
	{"20%", 20},
	{"30%", 30},
	{"40%", 40},
	{"50%", 50},
	{"60%", 60},
	{"70%", 70},
	{"80%", 80},
	{"90%", 90},
	{"100%", 100},
}

// // Benchmark test for computeSineSum with int slices of varying sizes
func BenchmarkSineSumInts(b *testing.B) {
	// Set seed for random number generation
	rand.Seed(7)
	// Initialize slice to compute on
	intData := make([]int, N)
	// Fill slice with random ints between 0 and 1000
	for i := range intData {
		intData[i] = rand.Intn(1001)
	}
	// For each percentage of data to be benchmarked
	for _, pct := range percents {
		// Run the test for the percentage with the associated name
		b.Run(pct.name, func(b *testing.B) {
			// Run the fonction b.N times (value adjusted dynamically)
			for i := 0; i < b.N; i++ {
				// Run the function on the first N*pct.percent/100 elements of the slice, aka (pct.percent)% of the data
				computeSineSum(intData[:N*pct.percent/100])
			}
		})
	}
}

// // Benchmark test for computeSineSum with float64 slices of varying sizes
func BenchmarkSineSumFloats(b *testing.B) {
	// Set seed for random number generation
	rand.Seed(7)
	// Initialize slice to compute on
	floatData := make([]float64, N)
	// Fill slice with random floats between 0 and 1
	for i := range floatData {
		floatData[i] = rand.Float64()
	}
	// For each percentage of data to be benchmarked
	for _, pct := range percents {
		// Run the test for the percentage with the associated name
		b.Run(pct.name, func(b *testing.B) {
			// Run the fonction b.N times to get the average computing time
			for i := 0; i < b.N; i++ {
				// Run the function on the first N*pct.percent/100 elements of the slice, aka (pct.percent)% of the data
				computeSineSum(floatData[:N*pct.percent/100])
			}
		})
	}
}
