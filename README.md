# go-abcc
[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hexoul/go-abcc/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/hexoul/go-abcc)](https://goreportcard.com/report/github.com/hexoul/go-abcc)
[![GoDoc](https://godoc.org/github.com/hexoul/go-abcc?status.svg)](https://godoc.org/github.com/hexoul/go-abcc)

> ABCC API Client written in Golang

## Usage
- As library, start from `abcc.GetInstanceWithKey('YOUR_ACCESS_KEY', 'YOUR_SECRET_KEY')`
- As program, start from `abcc.GetInstance()` after executing `go run -abcc:accesskey=[YOUR_ACCESS_KEY] -abcc:secretkey=[YOUR_SECRET_KEY]`

## Features
| Type        | Endpoint                        | Done |
|-------------|---------------------------------|------|
| Common      | /v1/common/timestamp            | ✔ |
| Common      | /v1/common/markets              | ✔ |
| Exchange    | /v1/exchange/orders             | ✔ |
| Member      | /v1/members/me                  | ✔ |
| Member      | /v1/members/trades              | ✔ |
