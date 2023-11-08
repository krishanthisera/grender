# Go Gin Web Page Renderer

❗ __Warning: This project is in a very early stage of development.__

This repository contains a simple Go web application built with Gin that renders web pages using headless Chrome through the chromedp library. The rendered HTML content is then returned as the response to the client.

## Prerequisites

Before running the application, make sure you have the following dependencies installed:

- [Go](https://golang.org/dl/)
- [chromedp](https://pkg.go.dev/github.com/chromedp/chromedp)
- [Gin](https://pkg.go.dev/github.com/gin-gonic/gin)

You can install the Go dependencies using `go get`:

```bash
go get -u github.com/chromedp/chromedp
go get -u github.com/gin-gonic/gin
```
