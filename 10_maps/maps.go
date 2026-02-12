package main

import (
	"fmt"
	"maps"
	"reflect"
)

func main() {
	// maps in Go (hash, object, dictionary)
	// a map is a collection of key-value pairs where each key is unique and maps to a value
	// maps are also known as dictionaries or hash tables in other programming languages
	// syntax for declaring a map in Go
	// var name map[keyType]valueType
	// uninitialized map
	var myMap map[string]int
	// uninitialized maps are nil and cannot be used until they are initialized
	// fmt.Println(myMap["key"]) // this will cause a runtime panic: assignment to entry in nil map
	// we can initialize a map using the make function
	myMap = make(map[string]int)
	// we can also declare and initialize a map in one line
	myMap = map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	fmt.Println(myMap)
	// we can also use short hand declaration for maps
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	fmt.Println(colors)
	// we can access map values using their keys
	fmt.Println(myMap["apple"])
	fmt.Println(colors["green"])
	// we can also modify map values using their keys

	myMap["banana"] = 5
	fmt.Println(myMap)
	// we can also add new key-value pairs to the map
	colors["yellow"] = "#FFFF00"
	fmt.Println(colors)
	// we can also use a for loop to iterate over a map

	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	// we can also use a for loop to iterate over a map without using the key
	for _, value := range colors {
		fmt.Println(value)
	}
	// we can also use a for loop to iterate over a map without using the value
	for key := range myMap {
		fmt.Println("Key:", key)
	}
	// we can also use a for loop to iterate over a map without using the key and value
	for range colors {
		fmt.Println("Iterating over colors map")
	}
	// we can also delete a key-value pair from the map using the delete function
	// syntax for delete function in Go
	// delete(map, key)
	delete(myMap, "orange")
	fmt.Println(myMap)
	// we can also check if a key exists in the map using the comma ok idiom
	value, ok := colors["red"]
	if ok {
		fmt.Println("Value for key 'red':", value)
	} else {
		fmt.Println("Key 'red' does not exist in the colors map")
	}
	// we can also check if a key exists in the map using the value and ok variables
	if value, ok := myMap["banana"]; ok {
		fmt.Println("Value for key 'banana':", value)
	} else {
		fmt.Println("Key 'banana' does not exist in the myMap")
	}
	// we can also use the built-in len function to get the number of key-value pairs in the map
	fmt.Println("Number of key-value pairs in myMap:", len(myMap))
	// we can also use the built-in make function to create a map with a specified initial capacity
	// syntax for make function in Go

	// make(type, initialCapacity)
	myMap = make(map[string]int, 10)
	fmt.Println("myMap with initial capacity of 10:", myMap)

	// maps in Go are implemented as hash tables and have an average time complexity of O(1) for lookups, insertions, and deletions
	// however, the worst-case time complexity can be O(n) in case of hash collisions

	// equality of maps
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"a": 1, "b": 2}
	// maps cannot be compared using the == operator, it will result in a compile-time error
	// fmt.Println(map1 == map2) // this will cause a compile-time error
	// to compare maps, we can use the reflect.DeepEqual function from the reflect package
	fmt.Println("Are map1 and map2 equal?", reflect.DeepEqual(map1, map2))
	fmt.Println(maps.Equal(map1, map2))
	
}