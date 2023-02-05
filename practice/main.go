package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Println("Network requests")
	resp, err := http.Get(url)
	checkError(err)

	defer resp.Body.Close()
	byte, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	content := string(byte)
	fmt.Println(content)

}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
