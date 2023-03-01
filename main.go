package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Started scanning...")
	for {
		genUm := Gen(7)
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://fb.blooket.com/c/firebase/id?id="+genUm, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Origin", "https://play.blooket.com")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Cookie", "bsid=MTY3NzY3ODc3NXxfY3ZhNmE2RHNfOXlXdTc0eFNmRTVPRkZJSkFGTWl6QzJmcjVzeGJWRHhlb1l1RWNUcUFxSGVwcE1SVmdnU1F5NWRQazZjR0FGSXdhVUFKaVM3TlFSVmNtd2ZaWnd3WVB8Lcjt5-fdaFshBozFi92qnDyZDg1sTeilN8jjI3hbGD4=")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-site")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		a, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var res map[string]interface{}
		json.Unmarshal(a, &res)
		if res["success"] != false {
			fmt.Printf("[+]Found Game: %s\n", genUm)
		}
	}
}

func Gen(x int) string {
	finalStr := ""
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 9
	for i := 1; i < x+1; i++ {
		finalStr += strconv.Itoa(rand.Intn(max-min+1) + min)
	}
	return finalStr
}
