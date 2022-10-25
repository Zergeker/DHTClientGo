package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	ipArray := [50]string{
		"compute-10-0",
		"compute-10-1",
		"compute-10-2",
		"compute-10-3",
		"compute-10-4",
		"compute-10-5",
		"compute-10-6",
		"compute-10-7",
		"compute-10-8",
		"compute-10-9",
		"compute-10-10",
		"compute-10-11",
		"compute-10-12",
		"compute-10-13",
		"compute-10-14",
		"compute-10-15",
		"compute-10-16",
		"compute-10-17",
		"compute-10-18",
		"compute-10-19",
		"compute-10-20",
		"compute-10-21",
		"compute-10-22",
		"compute-10-23",
		"compute-10-24",
		"compute-10-25",
		"compute-10-26",
		"compute-10-27",
		"compute-10-28",
		"compute-10-29",
		"compute-10-30",
		"compute-10-31",
		"compute-10-32",
		"compute-10-33",
		"compute-10-34",
		"compute-10-35",
		"compute-10-36",
		"compute-10-37",
		"compute-10-38",
		"compute-10-39",
		"compute-10-40",
		"compute-10-41",
		"compute-10-42",
		"compute-10-43",
		"compute-10-44",
		"compute-10-45",
		"compute-10-46",
		"compute-10-47",
		"compute-7-0",
		"compute-7-1"}

	nodesCount, _ := strconv.Atoi(os.Args[1])
	port := os.Args[2]
	joinStart := time.Now()
	for i := 1; i < nodesCount; i++ {
		_, err := http.Post("http://"+ipArray[i]+":"+port+"/join?nprime="+ipArray[i-1]+":"+port, "text/plain", bytes.NewReader([]byte("")))
		if err != nil {
			log.Fatal("Node " + ipArray[i] + " could not join node " + ipArray[i-1])
		}
	}
	joinFinished := time.Since(joinStart)

	leaveStart := time.Now()
	for i := nodesCount - 1; i > nodesCount/2; i-- {
		_, err := http.Post("http://"+ipArray[i]+":"+port+"/leave", "text/plain", bytes.NewReader([]byte("")))
		if err != nil {
			log.Fatal("Error occured during node " + ipArray[i] + " leaving")
		}
	}
	leaveFinished := time.Since(leaveStart)

	fmt.Println("Joining time: " + joinFinished.String())
	fmt.Println("Shrinking time: " + leaveFinished.String())
}
