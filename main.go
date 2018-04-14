package main

import (
	"log"
	"os"
	_ "github.com/golang-demos/go-search-rss-demo/matchers"
	"github.com/golang-demos/go-search-rss-demo/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("World")
}
