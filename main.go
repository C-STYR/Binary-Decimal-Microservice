package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

const serverPort = 3333

func main() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s /\n", r.Method)
		})
		mux.HandleFunc("/btod", func(w http.ResponseWriter, r *http.Request) {
			// grab input from body
			// process input with binaryToDecimal(input)
			// return response
		})
		mux.HandleFunc("/dtob", func(w http.ResponseWriter, r *http.Request) {
			// grab input from body
			// process input with decimalToBinary(input)
			// return response
		})
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running server %s\n", err)
			}
		}
	}()
	time.Sleep(100 * time.Millisecond)

	requestUrl := fmt.Sprintf("http://localhost:%d", serverPort)
	res, err := http.Get(requestUrl)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: got response\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	// binaryString := "1001100111"
	// fmt.Println(binaryToDecimal(binaryString))
	// decimal := 615
	// fmt.Println(decimalToBinary(decimal))
	// fmt.Println(binaryString == decimalToBinary(decimal))
}

func binaryToDecimal(input string) (num int) {
	num = 0
	var j float64 = 0
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == 49 {
			num += int(math.Pow(2, j))
		}
		j++
	}
	return
}

func decimalToBinary(num int) (output string) {
	output = ""
	remainders := []int{}
	currentNum := num
	for {
		remainders = append(remainders, currentNum%2)
		quotient := currentNum / 2
		if quotient == 0 {
			break
		}
		currentNum = quotient
	}
	for i := len(remainders) - 1; i >= 0; i-- {
		output += strconv.Itoa(remainders[i])
	}
	return
}
