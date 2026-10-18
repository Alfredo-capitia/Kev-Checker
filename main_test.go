package main

import "testing"

func TestFindCVE(t *testing.T) {
	vulnerabilities := []Vulnerability{
		{
			CVEID:         "CVE-2021-44228",
			VendorProject: "Apache",
			Product:       "Log4j2",
		},
		{
			CVEID:         "CVE-2023-34362",
			VendorProject: "Progress",
			Product:       "MOVEit Transfer",
		},
	}

	tests := []struct {
		name        string
		targetCVE   string
		wantFound   bool
		wantProduct string
	}{
		{
			name:        "CVE encontrada",
			targetCVE:   "CVE-2021-44228",
			wantFound:   true,
			wantProduct: "Log4j2",
		},
		{
			name:      "CVE não encontrada",
			targetCVE: "CVE-9999-99999",
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := findCVE(vulnerabilities, tt.targetCVE)

			if found != tt.wantFound {
				t.Fatalf("esperava found=%t; obtido %t", tt.wantFound, found)
			}

			if found && got.Product != tt.wantProduct {
				t.Errorf("esperava produto %q; obtido %q", tt.wantProduct, got.Product)
			}
		})
	}
}
