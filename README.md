<p align="center">
  <img src="/logo.jpg" height="300">
</p>

# CryptoScraper

[![Go Report Card](https://goreportcard.com/badge/github.com/uyennguyen16900/crypto-scraper)](https://goreportcard.com/report/github.com/uyennguyen16900/crypto-scraper)

CryptoScraper scrapes today's top 100 crypto coins data and return a list of all cryptocurrencies whose prices are below a given price in ascending order using Go and Colly. 

[![asciicast](https://asciinema.org/a/JM3SjIr5nqGSDbJhFHjGglkhD.svg)](https://asciinema.org/a/JM3SjIr5nqGSDbJhFHjGglkhD)
### 📚 Table of Contents

1. [Project Structure](#project-structure)
2. [Getting Started](#getting-started)
3. [Built With](#built-with)
4. [Resources](#resources)
## Project Structure

```bash
📂 crypto-scraper
├── README.md
├── cryptocoinmarketcap.csv
├── output.json
├── scraper.go
└── scraper_test.go

```

## 🚀 Getting Started
Run each command line-by-line to set up the project
```bash
$ git clone git@github.com:uyennguyen16900/crypto-scraper.git
$ cd crypto-scraper
$ git remote rm origin
$ git remote add origin git@github.com:YOUR_GITHUB_USERNAME/crypto-scraper.git
```

To execute the Go program, write ```go run``` followed by the name of the file.
```bash
$ go run scraper.go
```
To execute the test, run
```bash
$ go tests
```
To execute all benchmarks within the package, run
```bash
$ go test -bench=.
```

## Resources
### Built With
- [**Colly** - Docs](http://go-colly.org/docs/)
- [**Go** - Docs](https://golang.org/doc/)

### Example Code
- [Go Examples](http://go-colly.org/docs/examples/basic/)
- [GoByExample - JSON](https://gobyexample.com/json)
