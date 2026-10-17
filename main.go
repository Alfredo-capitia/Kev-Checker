package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Catalog struct {
	CatalogVersion string `json:"catalogVersion"`
	DateReleased   string `json:"dateReleased"`
	Count          int    `json:"count"`
}

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

	var catalog Catalog

	err = json.Unmarshal(body, &catalog)

	if err != nil {
		fmt.Println("Erro ao interpretar o Json:", err)
		return
	}

	fmt.Println("=== CISA KEV Catalog ===")
	fmt.Println("Versão:", catalog.CatalogVersion)
	fmt.Println("Publicado em:", catalog.DateReleased)
	fmt.Println("Vulnerabilidades:", catalog.Count)

}
