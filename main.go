package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetrequest() {
	data, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error occur while getting a request", err)
		return
	}
	defer data.Body.Close()

	var todo Todo
	err1 := json.NewDecoder(data.Body).Decode(&todo)
	if err1 != nil {
		fmt.Println("error occur while decoding", err1)
	}
	fmt.Println(todo)

}
func performPostrequest() {
	todo := Todo{
		UserId:    1234,
		Id:        34,
		Title:     "Post Request",
		Completed: true,
	}
	res, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error occur while Marshalling", err)
	}
	//convert json into string
	// jsonString := string(res)       instead of converting we use bytes.NewBuffer()
	//convert string into reader
	// jsonReader := strings.NewReader(jsonString)
	//url
	myUrl := "https://jsonplaceholder.typicode.com/todos/"
	//post request
	res1, err1 := http.Post(myUrl, "application/json", bytes.NewBuffer(res))
	if err1 != nil {
		fmt.Println("error occur while post request", err)
		return
	}
	defer res1.Body.Close()

	data, err := io.ReadAll(res1.Body)
	if err != nil {
		fmt.Println("error while reading", err)
		return
	}
	fmt.Println("response is:", string(data))
}
func performPutrequest() {
	todo := Todo{
		UserId:    1234,
		Id:        34,
		Title:     "Post Request By Khushboo",
		Completed: true,
	}
	data, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Eror while Marshalling", err)
		return
	}
	myurl := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, myurl, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error while Requesting", err)
	}
	req.Header.Set("Content-type", "application/json")
	//send a request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while Sending a request")
		return
	}
	defer res.Body.Close()

	res1, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading", err)
		return
	}

	fmt.Println("response is :", string(res1))
	fmt.Println("response is :", res.Status)

}

func main() {
	performGetrequest()
	performPostrequest()
	performPutrequest()

}
