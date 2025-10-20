 Go REST API Client

A simple "REST API Client built in Go (Golang)" that demonstrates how to perform CRUD operations (`GET`, `POST`, `PUT`, and `DELETE`) using Go’s standard `net/http` package.  
This project interacts with the public [JSONPlaceholder](https://jsonplaceholder.typicode.com/) API — a free online REST service used for testing and prototyping.

OVERVIEW:
This project showcases how to:
- Make HTTP requests in Go
- Handle JSON encoding/decoding
- Work with REST API endpoints
- Manage HTTP methods like GET, POST, PUT, and DELETE

Tech Stack
- Language: Go 1.21+
- Packages Used:
  - `net/http` — HTTP client for requests
  - `encoding/json` — For JSON serialization/deserialization
  - `io` / `ioutil` — Reading API responses
  - `strings` — For creating request bodies

 📁 Project Structure:
├── main.go # Contains all CRUD functions and main() execution
└── README.md # Project documentation


 Features:
-Fetch Todo using **GET** request  
-Create a new Todo using **POST** request  
- Update an existing Todo using **PUT** request  
- Delete a Todo using **DELETE** request  
- JSON marshalling/unmarshalling with `encoding/json`  
-Uses Go’s `http.Client` for manual request handling

 ▶️ How to Run
 1️⃣ Clone the repository
 in bash
git clone https://github.com/<your-username>/go-rest-api-client.git
cd go-rest-api-client
2️⃣ Run the program:
go run main.go

******OUTPUT****
Todo is {1 1 delectus aut autem false}
response is : {
  "userId": 12,
  "id": 23,
  "title": "post Request",
  "completed": true
}
Response is: {
  "userId": 12,
  "id": 23,
  "title": "post Request By Khushboo",
  "completed": true
}
response is : 200 OK
