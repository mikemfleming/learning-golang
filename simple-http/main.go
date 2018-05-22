package main

import "fmt"
import "net/http"
import "io/ioutil"

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	return 
}

