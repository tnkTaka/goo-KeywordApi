package main

import (
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
)

func PostKeyWordApi(title, body string) string {
	jsonStr := `{"app_id":" My Api Token ","title":"`+title+`","body":"`+body+`","max_num":"5"}`

	req, err := http.NewRequest(
		"POST",
		"https://labs.goo.ne.jp/api/keyword",
		bytes.NewBuffer([]byte(jsonStr)),
	)

	if err != nil {
		return "失敗"
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "失敗"
	}
	defer resp.Body.Close()

	result := execute(resp)
	return result
}

func execute(response *http.Response) string{
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	resBody := string(body)

	return resBody
}

func main() {
	title := "GolangでAPIを叩く"
	body := "GolangでAPIサーバを書こうと思い手始めにgooのキーワード抽出のAPIを叩いてみました。"

	responseApiJson := PostKeyWordApi(title,body)

	fmt.Println(responseApiJson)
}