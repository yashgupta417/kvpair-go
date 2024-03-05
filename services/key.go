package services

import (
	"log"
)

var data = make(map[string]string)

func UpsertKeyToMap(key string, value string) {
	log.Printf("Adding key '%v' to dataset.", key)
	data[key] = value
}
 
func GetKeyFromMap(key string) string {
	log.Printf("Trying fetching key '%v' from dataset", key)
	value, ok := data[key]

	if ok {
		log.Printf("Found key '%v' in dataset", key)
		return value
	} else {
		log.Printf("Key '%v' not found in dataset", key)
		return "-"
	}
}

func DeleteKeyFromMap(key string) {
	log.Printf("Deleting key '%v' from dataset.", key)
	delete(data, key)
}