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
	ipArray := [16]string{"10.1.2.233", "10.1.2.197", "10.1.2.184", "10.1.2.183", "10.1.2.182", "10.1.2.181", "10.1.2.180", "10.1.2.179", "10.1.2.178", "10.1.2.177", "10.1.2.161", "10.1.2.143", "10.1.2.142", "10.1.2.123", "10.1.2.122", "10.1.2.121"}

	rand.Seed(time.Now().Unix())

	nodesCount, _ := strconv.Atoi(os.Args[1])

	flag := true
	successfulCounter := 0
	errorsCounter := 0
	time.AfterFunc(5*time.Minute, func() { flag = false })

	client := &http.Client{}

	for flag == true {
		ip := rand.Intn(nodesCount)
		key := rand.Int()
		value := rand.Int()

		reqPut, err := http.NewRequest(http.MethodPut, "http://"+ipArray[ip]+":58346/storage/"+strconv.Itoa(key), bytes.NewBuffer([]byte(strconv.Itoa(value))))
		reqGet, err := http.NewRequest(http.MethodGet, "http://"+ipArray[ip]+":58346/storage/"+strconv.Itoa(key), nil)

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
