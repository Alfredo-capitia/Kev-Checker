package main

import (
	"fmt"
	"io"
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("erro ao ler a resposta: ", err)
		return
	}

	fmt.Println("tamanho da resposta:", len(body), "bytes")
	fmt.Println("Inicio da resposta:")
	fmt.Println(string(body[:200]))

}
