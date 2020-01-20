package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(uri string) string {
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func HttpPost(uri string, data string) string {

	resp, err := http.Post(uri, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}

func HttpPostJson(uri string, data []byte) string {

	reader := bytes.NewReader(data)

	resp, err := http.Post(uri, "application/json;charset=UTF-8", reader)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}
