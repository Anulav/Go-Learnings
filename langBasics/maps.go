package main

import "fmt"

func main() {
	/* Maps are collections, and we can iterate over them just like we do with arrays and
	   slices. But maps are unordered collections, and there’s no way to predict the order in
	   which the key/value pairs will be returned.

	   The map’s hash table contains a collection of buckets. When we’re storing, removing,
	   or looking up a key/value pair, everything starts with selecting a bucket. This is
	   performed by passing the key—specified in wer map operation—to the map’s hash
	   function. The purpose of the hash function is to generate an index that evenly distributes
	   key/value pairs across all available buckets.

	   The better the distribution, the quicker we can find our key/value pairs as the
	   map grows. If we store 10,000 items in our map, we don’t want to ever look at
	   10,000 key/value pairs to find the one we want. we want to look at the least number
	   of key/value pairs possible. Looking at only 8 key/value pairs in a map of 10,000 items
	   is a good and balanced map. A balanced list of key/value pairs across the right number
	   of buckets makes this possible.

	*/

	//Ways to initialize map
	map1 := map[string]string{"Red": "Stop", "Green": "Go", "Yellow": "Slow down"}
	map2 := make(map[string]string)
	map2["Red"] = "Stop"
	map2["Green"] = "Go"
	fmt.Println("Map1 ", map1)
	fmt.Println("Map2 ", map2)
	//The map key can be a value from any built-in or struct type as long as the value can
	//be used in an expression with the == operator. Slices, functions, and struct types that
	//contain slices can’t be used as map keys. But can use as values.

	//Nil map declaration only
	var map3 map[string]string
	//map3["Red"] = "Stop" This gives error as the map3 is a nil map
	fmt.Print(map3)

	//We get a boolean when we try retrieving a value. That helps to check if a key exists.
	value, exists := map2["Red"]
	if exists {
		fmt.Print("Value @ Red ", value)
	}

	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	//Iterating map using for-range
	for key, value := range colors {
		fmt.Println(key, "->", value)
	}

	delete(colors, "Coral") //delete function used by Map to delete entries of the Map

	//Iterating map using for-range
	for key, value := range colors {
		fmt.Println(key, "->", value)
	}

	//Passing Maps to the function
	//Passing a map between two functions doesn’t make a copy of the map. In fact, you can
	//pass a map to a function and make changes to the map, and the changes will be
	//reflected by all references to the map.
	removeRed(colors)
	fmt.Println(colors)
}

func removeRed(colors map[string]string) {
	delete(colors, "Red")
}
