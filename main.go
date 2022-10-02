package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	ip := os.Args[1]

	rand.Seed(time.Now().Unix())

	flag := true
	successfulCounter := 0
	errorsCounter := 0
	time.AfterFunc(5*time.Minute, func() { flag = false })

	client := &http.Client{}

	for flag == true {
		key := rand.Int()
		value := rand.Int()

		reqPut, err := http.NewRequest(http.MethodPut, "http://"+ip+":58346/storage/"+strconv.Itoa(key), bytes.NewBuffer([]byte(strconv.Itoa(value))))
		reqGet, err := http.NewRequest(http.MethodGet, "http://"+ip+":58346/storage/"+strconv.Itoa(key), nil)

		_, err = client.Do(reqPut)

		if err != nil {
			errorsCounter++
		} else {
			successfulCounter++
		}

		_, err = client.Do(reqGet)

		if err != nil {
			errorsCounter++
		} else {
			successfulCounter++
		}
	}

	fmt.Println("Successful requests: " + strconv.Itoa(successfulCounter))
	fmt.Println("Failured requests: " + strconv.Itoa(errorsCounter))
	fmt.Println("Requests per second: " + fmt.Sprintf("%f", (float64(successfulCounter+errorsCounter)/float64(300))))
}
