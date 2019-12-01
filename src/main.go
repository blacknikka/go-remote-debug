package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")

	fmt.Println("hello world.")
}

func main() {
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// key := "StringGetSet_Key"
		// // Set
		// err := client.Set(key, "StringGetSet_Val", 0).Err()
		// if err != nil {
		// 	w.WriteHeader(500)
		// 	w.Write([]byte(err.Error()))
		// }

		// if res, ok := res.(int64); ok {
		// 	w.Write([]byte(fmt.Sprintf("counter: %d", res)))
		// } else {
		// 	w.WriteHeader(500)
		// 	w.Write([]byte("unexpected value"))
		// }
	})
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
