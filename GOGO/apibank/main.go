package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Albums struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
}

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func main() {
	getAlbums()
	getUsers()
}

func getAlbums()  {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/albums")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var a []Albums
	if err := json.Unmarshal(body, &a); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)
}

func getUsers()  {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var users []Users
	if err := json.Unmarshal(body, &users); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users)

}