package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

var btime string

func main() {
	fmt.Println("Build at " + btime)
	//if len(os.Args) == 2 && os.Args[1] == "--server" {

	seconds := 1
	seconds, err := strconv.Atoi(os.Getenv("SLEEP"))
	if err != nil {
		seconds = 1
	}

	fmt.Printf("Handler will sleep for %d second(s)", seconds)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(time.Duration(seconds) * time.Second)
		w.Write([]byte("Hello"))
		return
	})
	http.ListenAndServe(":80", nil)
	//} //else {
	// 	for {
	// 		time.Sleep(time.Second)
	// 		var wg sync.WaitGroup
	// 		for i := 0; i < 2; i++ {
	// 			wg.Add(1)
	// 			go func() {
	// 				defer wg.Done()
	// 				r, err := http.DefaultClient.Get("http://localhost:8085")
	// 				if err != nil {
	// 					fmt.Printf("Error: %v", err)
	// 				} else {
	// 					fmt.Printf("%s\n", r.Body)
	// 				}
	// 			}()

	// 			wg.Wait()
	// 		}
	// 	}
	// }
}
