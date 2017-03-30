package main

import (
	"log"
	"net/http"
	"os"
	"bufio"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	MigrateDatabase()
	checkIfUsersExist()
}

func startServer() {
	router := NewRouter()
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func checkIfUsersExist() {
	if result := GetUsers(); reflect.ValueOf(result).Elem().Len() == 0 {
		promptCreateUser()
	} else {
		startServer()
	}
}

func promptCreateUser() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Create a user account first")
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')

	user := User{Username: strings.TrimSpace(username), Password: strings.TrimSpace(password)}
	user.Create()

	checkIfUsersExist()
}