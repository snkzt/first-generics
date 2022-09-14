package main

import "fmt"

// Moving the union of int and float (constraint) from the function declaration
// into a new type constraint
// Declare a type constraint: https://go.dev/doc/tutorial/generics
// Number can be number if it's not used from the outside of the package
type Number interface {
  int64 | float64
}

func main() {
  // Initialise a map for int values
  // map[keyType]ValueType
  // This is a map literal which to initialise a map with some data
  ints := map[string]int64{
    "first": 34,
    "second": 12,
  }

  // Initialise a map for float values
  floats := map[string]float64{
    "first": 35.98,
    "second": 26.99,
  }

  fmt.Printf("Non-Generic Sums: %v and %v\n",
    SumInts(ints),
    SumFloat(floats))
 
  fmt.Printf("Generic Sums: %v and %v\n",
    SumIntsOrFloats[string, int64](ints),
    SumIntsOrFloats[string, float64](floats))

  // Implicit type of arguments. Go compiler can infer the types want to use
  // from the types of function arguments: SumIntsOrFloats function declaration.
  fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
  SumIntsOrFloats(ints),
  SumIntsOrFloats(floats))

  // Type arguments omitted as Number, Go compiler infer type argument from other arguments (Number).  
  fmt.Printf("Generic Sums with Constraint: %v and %v\n",
    SumNumbers(ints),
    SumNumbers(floats))
}

// SumInts add together the values of type int64 map and return the total of type int64
func SumInts(m map[string]int64) int64 {
  var sum int64
  for _, value := range m {
    sum += value
  } 
  return sum
}

// SumFloats add together the values of type float64 map and return the total of type float64
func SumFloat(m map[string]float64) float64 {
  var sum float64
  for _, value := range m {
    sum += value
  }
  return sum
}

//SumIntsOrFloats add together a map. It supports int64 and float64 as types of map. 
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
  var sum V
  for _, value:= range m {
    sum += value
  }
  return sum
}

// SumNumbers add together a map regardless the argument of either ints or floats
func SumNumbers[K comparable, V Number](m map[K]V) V {
  var sum V
  for _, value := range m {
    sum += value
  }
  return sum
}
