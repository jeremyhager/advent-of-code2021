package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func MethodRequest(method, endpoint string, data []byte) ([]byte, error) {
	url := "https://adventofcode.com/2021/day/" + endpoint + "/input"
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error reaching endpoint: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return body, nil
}

func convertToInt(input string) ([]int, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("[could not read the file:] %v", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var depths []int
	for scanner.Scan() {
		newInt, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, newInt)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("[could not convert to []int:] %v", err)
	}

	return depths, nil

}

func depthCompare(depths []int) int {
	change := 0
	for i, depth := range depths {
		if i == 0 {
			fmt.Println("do nothing")
		} else if (depth - depths[i-1]) > 0 {
			change += 1
		} else {
			fmt.Printf("[did not find change greater than 0:] depths[i-1]: %v, depth: %v\n", depths[i-1], depth)
		}
	}

	return change
}

func main() {
	depths, err := convertToInt("/Users/jeremy.hager/Downloads/input.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("%v", depths)

	change := depthCompare(depths)

	log.Printf("[change:] %v", change)
}
