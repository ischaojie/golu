package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	APP_KEY    = "218f4a5913a2b64d"
	APP_SECRET = "tn49UOKE8VBy5zezLGWjr8HnusRMcxW8"
	YOUDAO_URL = "https://openapi.youdao.com/api"
)

type Translation struct {
	Translation []*TranNode `json:"translation"`
}

type TranNode struct {
	tr string
}

func HttpGet(url string, data url.Values, ch chan []byte) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		fmt.Println("please enter ctrl+c, request error: ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("please enter ctrl+c, io read error: ", err)
	}

	ch <- body
}

func Truncate(q string) string {
	// 使用rune，解决中文乱码
	s := []rune(q)
	if s == nil {
		return ""
	}
	size := len(s)
	if size <= 20 {
		return string(s)
	} else {
		return string(s[:10]) + strconv.Itoa(size) + string(s[size-10:size])
	}
}

func Trans(q string) {
	// 加盐
	salt := uuid.New().String()
	// 当前时间
	curTime := fmt.Sprintf("%d", time.Now().Unix())
	signStr := APP_KEY + Truncate(q) + salt + curTime + APP_SECRET
	encrypt := sha256.New()
	encrypt.Write([]byte(signStr))
	// sign
	sign := encrypt.Sum(nil)

	urlData := url.Values{
		"from":     {"auto"},
		"to":       {"auto"},
		"signType": {"v3"},
		"curtime":  {curTime},
		"appKey":   {APP_KEY},
		"q":        {q},
		"salt":     {salt},
		"sign":     {fmt.Sprintf("%x", sign)},
	}

	ch := make(chan []byte)

	go HttpGet(YOUDAO_URL, urlData, ch)

	var d interface{}
	if err := json.Unmarshal(<-ch, &d); err != nil {
		log.Fatal("json wrong.")
	}

	data := d.(map[string]interface{})
	fmt.Println("\033[1;31m*******英汉翻译:*******\033[0m")
	for _, trans := range data["translation"].([]interface{}) {
		fmt.Println(trans)
	}

	fmt.Println("\033[1;31m*******网络释义:*******\033[0m")

	if data["web"] == nil {
		return
	}
	for _, web := range data["web"].([]interface{}) {
		webNode := web.(map[string]interface{})
		fmt.Print(webNode["key"], " : ")
		for _, v := range webNode["value"].([]interface{}) {
			fmt.Print(v, " / ")
		}
		fmt.Println()
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("please input the translation word.")
		fmt.Println("example: dict hello")
		os.Exit(0)

	}
	query := strings.Join(args, " ")
	Trans(query)
}
