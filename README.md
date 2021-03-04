<p align="center">
  <img src="/logo.jpg" height="400">
</p>

# CryptoScraper

[![Go Report Card](https://goreportcard.com/badge/github.com/uyennguyen16900/crypto-scraper)](https://goreportcard.com/report/github.com/uyennguyen16900/crypto-scraper)

CryptoScraper scrapes today's top 100 crypto coins data and return a list of all cryptocurrencies whose prices are below a given price in ascending order using Go and Colly. 

### 📚 Table of Contents

1. [Project Structure](#project-structure)
2. [Getting Started](#getting-started)
3. [Deliverables](#deliverables)
4. [Resources](#resources)
## Project Structure

```bash
📂 crypto-scraper
├── README.md
├── scraper.go
└── scraper_test.go

```

## Getting Started
To execute the Go program, write ```go run``` followed by the name of the file.
```bash
$ go run scraper.go

```
To execute the test, run
```bash
$ go test
```
To execute all benchmarks within the package, run
```bash
$ go test -bench=.
```

## Resources
- [**Colly** - Docs](http://go-colly.org/docs/)
- [**Go** - Docs](https://golang.org/doc/)
- [Go Examples](http://go-colly.org/docs/examples/basic/)
