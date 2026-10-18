# KEV Checker

A command-line tool written in Go that checks whether a Common Vulnerabilities and Exposures (CVE) identifier appears in the CISA Known Exploited Vulnerabilities (KEV) Catalog.

The project is designed as a practical learning exercise in Go, cybersecurity automation, software testing, and DevOps practices.

## Features

- **CVE validation:** validates the format of CVE identifiers before making HTTP requests.
- **CISA KEV integration:** retrieves the official KEV catalog from CISA.
- **Vulnerability lookup:** searches for a specific CVE in the catalog.
- **HTTP timeout:** limits requests to 10 seconds.
- **Error handling:** handles network errors, HTTP failures, response-reading errors, and invalid JSON.
- **Automated tests:** tests CVE validation and vulnerability lookup.
- **Continuous integration:** uses GitHub Actions to run code-quality checks, tests, and builds.

## Requirements

- Go 1.27.2 or a compatible newer version.
- Internet access to retrieve the CISA KEV catalog.

## Getting Started

Clone the repository:

```bash
git clone https://github.com/Alfredo-capitia/Kev-Checker.git
cd Kev-Checker
```

Run the application:

```bash
go run . CVE-2021-44228
```

Replace `CVE-2021-44228` with the CVE identifier you want to check.

## Example Output

```text
Status HTTP: 200 OK
=== CISA KEV Catalog ===
Versão: 2026.10.16
Vulnerabilidades: 1748

[ENCONTRADA NO KEV]
CVE: CVE-2021-44228
Fornecedor: Apache
Produto: Log4j2
Data de inclusão: 2021-12-10
```

Output depends on the current CISA catalog and may change over time.

## Usage

Check a CVE:

```bash
go run . CVE-2021-44228
```

Show usage instructions:

```bash
go run .
```

Test invalid input:

```bash
go run . abc
```

Build a standalone executable:

```bash
go build -o kev-checker .
./kev-checker CVE-2021-44228
```

## Exit Codes

| Code | Meaning |
|---|---|
| `0` | The requested CVE was found in the KEV catalog. |
| `1` | Invalid arguments, invalid CVE format, CVE not found, or an execution error. |

The exit code can be used by shell scripts and CI/CD pipelines to evaluate the result.

## Testing

Run the automated tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

Run static analysis:

```bash
go vet ./...
```

Check formatting:

```bash
gofmt -d .
```

Build the project:

```bash
go build ./...
```

## Continuous Integration

GitHub Actions runs automated checks when changes are pushed to `main` or a pull request targets `main`.

The workflow checks Go formatting, runs `go vet`, executes tests, and builds the project.

## Data Source

The application retrieves vulnerability information from the official CISA Known Exploited Vulnerabilities catalog:

https://www.cisa.gov/known-exploited-vulnerabilities-catalog

The tool reports whether a CVE is present in that catalog. It does not determine whether a system is vulnerable, assess exploitability in a particular environment, or replace a comprehensive security assessment.

## Project Goals

This project is being developed to practice:

- Go programming and error handling.
- HTTP clients and JSON processing.
- Cybersecurity automation.
- Unit testing and code quality.
- Git branching and pull requests.
- CI/CD with GitHub Actions.

## Future Improvements

- Test HTTP interactions with Go's `httptest` package.
- Separate catalog retrieval from command-line logic.
- Improve CLI output and error reporting.
- Expand test coverage.

## Author

**Alfredo Capitia**

GitHub: https://github.com/Alfredo-capitia

