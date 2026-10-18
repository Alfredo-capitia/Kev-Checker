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
	os.Exit(run())
}

func run() int {
	if len(os.Args) != 2 {
		fmt.Println("Uso: go run . <CVE-ID>")
		return 1
	}

	targetCVE := os.Args[1]

	if !validateCVE(targetCVE) {
		fmt.Println("Erro: formato de CVE inválido.")
		fmt.Println("Exemplo: CVE-2021-44228")
		return 1
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Get(catalogURL)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Println("Erro: o pedido à CISA excedeu o limite de 10 segundos.")
			return 1
		}

		fmt.Println("Erro ao consultar a CISA:", err)
		return 1
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		fmt.Println("Erro: a CISA respondeu com status HTTP:", response.Status)
		return 1
	}

	fmt.Println("Status HTTP:", response.Status)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
		return 1
	}

	var catalog Catalog

	if err := json.Unmarshal(body, &catalog); err != nil {
		fmt.Println("Erro ao interpretar o JSON:", err)
		return 1
	}

	fmt.Println("=== CISA KEV Catalog ===")
	fmt.Println("Versão:", catalog.CatalogVersion)
	fmt.Println("Vulnerabilidades:", catalog.Count)

	vuln, found := findCVE(catalog.Vulnerabilities, targetCVE)
	if !found {
		fmt.Println("\nCVE não encontrada no catálogo:", targetCVE)
		return 1
	}

	fmt.Println("\n[ENCONTRADA NO KEV]")
	fmt.Println("CVE:", vuln.CVEID)
	fmt.Println("Fornecedor:", vuln.VendorProject)
	fmt.Println("Produto:", vuln.Product)
	fmt.Println("Data de inclusão:", vuln.DateAdded)

	return 0
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
