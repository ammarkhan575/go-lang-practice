package auth

import "fmt"


func Credential(username string, password string) bool {
	fmt.Println(username, " ", password);
	return true
}