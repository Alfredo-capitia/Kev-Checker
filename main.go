package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"regexp"
	"time"
)

const catalogURL = "https://www.cisa.gov/sites/default/files/feeds/known_exploited_vulnerabilities.json"

type Vulnerability struct {
	CVEID         string `json:"cveID"`
	VendorProject string `json:"vendorProject"`
	Product       string `json:"product"`
	DateAdded     string `json:"dateAdded"`
}

type Catalog struct {
	CatalogVersion  string          `json:"catalogVersion"`
	Count           int             `json:"count"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Uso: go run . <CVE-ID>")
		return
	}

	targetCVE := os.Args[1]

	if !validateCVE(targetCVE) {
		fmt.Println("Erro: formato de CVE inválido.")
		fmt.Println("Exemplo: CVE-2021-44228")
		return
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	response, err := client.Get(catalogURL)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Println("Erro: o pedido à CISA excedeu o limite de 10 segundos.")
			return
		}

		fmt.Println("Erro ao consultar a CISA:", err)
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
	fmt.Println("Vulnerabilidades:", catalog.Count)

	vuln, found := findCVE(catalog.Vulnerabilities, targetCVE)

	if !found {
		fmt.Println("\nCVE não encontrada no catálogo:", targetCVE)
		return
	}

	fmt.Println("\n[ENCONTRADA NO KEV]")
	fmt.Println("CVE:", vuln.CVEID)
	fmt.Println("Fornecedor:", vuln.VendorProject)
	fmt.Println("Produto:", vuln.Product)
	fmt.Println("Data de inclusão:", vuln.DateAdded)
}

func findCVE(vulnerabilities []Vulnerability, targetCVE string) (Vulnerability, bool) {
	for _, vuln := range vulnerabilities {
		if vuln.CVEID == targetCVE {
			return vuln, true
		}
	}

	return Vulnerability{}, false
}
func validateCVE(cveID string) bool {
	pattern := `^CVE-[0-9]{4}-[0-9]{4,}$`

	matched, err := regexp.MatchString(pattern, cveID)
	if err != nil {
		return false
	}

	return matched
}
