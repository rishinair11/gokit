package main

import "fmt"

var data map[string]string

func init() {
	// Initialize the maps
	data = map[string]string{
		"EL": "Greek",
		"EN": "English",
		"ES": "Spanish",
		"FR": "French",
		"HI": "Hindi",
	}
	fmt.Println("Initialized 'data' map")
}

func add(key, val string) bool {
	if _, ok := data[key]; ok {
		fmt.Println("Key already exists!")
		return false
	}
	data[key] = val
	return true
}

func update(key, val string) bool {
	if _, ok := data[key]; !ok {
		fmt.Println("Key doesn't exist!")
		return false
	}
	data[key] = val
	return true
}

func deleteVal(key string) bool {
	if _, ok := data[key]; !ok {
		fmt.Println("Key doesn't exist!")
		return false
	}
	delete(data, key)
	return true
}

func get(key string) string {
	return data[key]
}

func getAll() map[string]string {
	return data
}

func main() {
	fmt.Println("Original map is => ", data)

	//Add operations
	ok := add("DE", "German")
	if ok {
		fmt.Println("Added new key\nMap is => ", data)
	}
	ok = add("HI", "hindi") // Adding a duplicate key
	if ok {
		fmt.Println("Added new key\nMap is => ", data)
	}

	//Update operations
	ok = update("DE", "Deutsch")
	if ok {
		fmt.Println("Updated existing key\nMap is => ", data)
	}
	ok = update("ZH", "Chinese") // Updating an absent key
	if ok {
		fmt.Println("Updated existing key\nMap is => ", data)
	}

	// Delete operations
	ok = deleteVal("EL")
	if ok {
		fmt.Println("Deleted existing key\nMap is => ", data)
	}
	ok = deleteVal("EL") // Deleting an absent key
	if ok {
		fmt.Println("Deleted existing key\nMap is => ", data)
	}

	// Get operations
	fmt.Println("Got 'FR' => ", get("FR"))
	fmt.Println("Got 'ZH' => ", get("ZH"))
	fmt.Println("Got all => ", getAll())
	// fmt.Println("Got a value")

	// fmt.Println("Got all values")
}
