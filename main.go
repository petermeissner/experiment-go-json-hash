package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// declaring a struct
type User struct {
	Name     string
	Password string
	Role     string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func file_exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		// exists
		return true
	}
	return false
}

func read_users_json(path string) []User {
	var (
		err  error
		user []User
	)

	// JSON array to be decoded
	var data []byte
	if file_exists(path) {
		data, err = os.ReadFile(path)
		if err != nil {
			panic(err)
		}
	} else {
		data = []byte(`[]`)
	}

	// decoding JSON array to array
	err = json.Unmarshal(data, &user)

	//
	return user
}

// main function
func main() {
	user := read_users_json("users.json")

	// printing decoded array
	for i := range user {
		fmt.Println(user[i].Name + " - " + user[i].Password +
			" - " + user[i].Role)
	}
}
