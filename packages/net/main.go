package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Data struct {
	User User `json:"user"`
}


type User struct {
	Id int64 `json:"id"`
	Email string `json:"email"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Avatar string `json:"avatar"`
}


func main() {

	req, err := GetRequest("https://reqres.in/api/users/2")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	fmt.Println(string(req))

	body, err := Get("https://reqres.in/api/users/2")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))

}

func GetRequest(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	

	return body, nil
}

func Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}
