package main

import (
	"fmt"

	"github.com/Jubair-Ahmed-Khan/Go-Learning/auth"
	"github.com/Jubair-Ahmed-Khan/Go-Learning/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("jubair", "1234")
	session := auth.GetSession()

	fmt.Println("session:", session)

	user := user.User{
		Email: "jubair@gmail.com",
		Name:  "Jubair Ahmed Khan",
	}
	// fmt.Println("user email:", user.Email)
	// fmt.Println("user name:", user.Name)
	color.Red(user.Email)
	color.Green(user.Name)
}
