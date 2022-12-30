package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User struct fro parsing JSON
type User struct {
	Name     string
	Password string
	Role     string
}

// checks if a file exists via trying to get basic file statistics from os
//
// Note: this is not the only option to check and its based on assumptions
func file_exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		// exists
		return true
	}
	return false
}

// JSON parser specific for array of user infos
//
// - reads in JSON array of users and returns go array of user structs
func read_users_json(path string) []User {
	var (
		err  error
		user []User
	)

	// JSON array to be decoded
	var data []byte
	if file_exists(path) {
		data, err = os.ReadFile(path)
		check_err(err)
	} else {
		data = []byte(`[]`)
	}

	// decoding JSON array to array
	err = json.Unmarshal(data, &user)

	//
	return user
}

// Convenience function for checking errors
//
// - called for side effects
func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

// main function
//
// - read in JSON
// - parse User structs into array
// - print infos to console
// -
func main() {
	users := read_users_json("users.json")

	// printing decoded array
	for i := range users {

		// print infos
		fmt.Println(users[i].Name + " - " + users[i].Password + " - " + users[i].Role)

		// print password hash
		pepper := "UZS75KL"
		hash, salt := pw_hash(users[i].Password, pepper)
		fmt.Println(users[i].Password, hash, salt)

		// check password
		comp := pw_check(salt, pepper, users[i].Password, hash)

		if comp == true {
			fmt.Println("comparison worked for pw/hash - As expected")
		} else {
			fmt.Println("comparison DID NOT work")
		}

		if comp == false {
			fmt.Println("comparison worked")
		} else {
			fmt.Println("comparison DID NOT work for hash/random string - As expected")
		}
	}
}

func str_random(n int) string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!§$%&/()=?+*'#,;.:-_")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func pw_hash(pw string, pepper string) (hash string, salt string) {

	// start measuring time
	start := time.Now()

	// generate random salt
	salt = str_random(12)

	// salt and pepper pw then hash it
	salted_pw := pw + salt + pepper
	my_hash, err := bcrypt.GenerateFromPassword([]byte(salted_pw), bcrypt.DefaultCost+4)
	if err != nil {
		panic(err)
	}

	// reporting execution time
	elapsed := time.Since(start)
	fmt.Printf("Hashing took %s\n", elapsed)

	// return hash and salt used
	return string(my_hash), salt
}

func pw_check(salt string, pepper string, pw string, pw_hash string) bool {

	// start measuring time
	start := time.Now()

	// salt and pepper password
	salted_pw := pw + salt + pepper

	// compare pw and hash
	err := bcrypt.CompareHashAndPassword([]byte(pw_hash), []byte(salted_pw))

	// reporting execution time
	elapsed := time.Since(start)
	fmt.Printf("Check took %s\n", elapsed)

	// return
	if err == nil {
		return true
	} else {
		return false
	}
}
