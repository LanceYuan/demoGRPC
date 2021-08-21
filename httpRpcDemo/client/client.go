package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	body := map[string]interface{}{
		"id":     0,
		"params": []string{"lance"},
		"method": "HelloService.Hello",
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	resp, err := http.Post("http://localhost:1234/hello", "application/json", &buf)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res["result"])
}
