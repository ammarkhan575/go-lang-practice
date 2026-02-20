package main

import (
	"calculator/auth"
	"calculator/cmd"
	"calculator/math"
	"calculator/user"
	"fmt"
	"math/rand"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("This is my packages folder")
	fmt.Println(math.Add(5, 5))
	fmt.Println(math.Multiply(3.23, 4.99))

	// using cmd package 
	fmt.Println(cmd.CmdPrint())

	fmt.Println("This is my auth package")
	fmt.Println(auth.Credential("mohdammar", "password123"))
	sessionID := auth.GetSessionID() 
	fmt.Println("Session ID: ", sessionID)

	user1 := user.CreateUser(rand.Intn(10000), "Mohd Ammar", "ammarkhan575@gmail.com")
	fmt.Println(user1)

	// updateEmail
	fmt.Println("Before Update")
	fmt.Println(user1.Email)
	user1.UpdateUserEmail("ammarkhan757@gmail.com")
	fmt.Println("After update")
	fmt.Println(user1.Email)
	color.Cyan(user1.Name)
	color.Red("ammarkhan575@gmail.com")

	// to install package we use
	// go get github.com/fatih/color

	// go mod tidy:
	// ✅ adds missing dependencies
	// ✅ removes unused dependencies
	// ✅ updates go.sum
	// ✅ keeps go.mod clean

	// Imagine:
	// Your project imports:
	// import "github.com/gin-gonic/gin"
	// But you didn’t run go get.
	// If you run:
	// go mod tidy
	// 👉 Go automatically downloads and adds it.

	// "go mod tidy ensures go.mod matches the actual imports in the project by adding missing dependencies and removing unused ones."
	
}
