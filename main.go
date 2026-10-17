package main

import (
	"fmt"
	"net/http"
)

func main() {

	url := "https://www.cisa.gov/sites/default/files/feeds/known_exploited_vulnerabilities.json"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("erro ao consultar a CISA", err)
		return
	}
	defer response.Body.Close()

	fmt.Println("Status HTTP: ", response.Status)
	fmt.Println("Código Http: ", response.StatusCode)

}
