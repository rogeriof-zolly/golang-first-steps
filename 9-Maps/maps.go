package main

import "fmt"

func main() {

	map1 := map[string]string{
		"key": "value",
	}

	fmt.Println(map1)

	map2 := map[string]map[string]string{
		"key": {
			"mapKey": "mapValue",
		},
	}

	fmt.Println(map2)

	user := map[string]map[string]string{
		"personalData": {
			"name":      "Rogerio",
			"cellphone": "+5554999999999",
		},
		"universityData": {
			"instituition": "University of Santa Cruz do Sul",
			"course":       "Computer Engineering",
		},
	}

	fmt.Println("----------------------------------")
	fmt.Println("Name:", user["personalData"]["name"])
	fmt.Println("Course:", user["universityData"]["course"])

	user["professionalData"] = map[string]string{
		"job": "Quality Assurance Analyst",
	}

	fmt.Println("Job:", user["professionalData"]["job"])
	fmt.Println("----------------------------------")

}
