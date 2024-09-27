package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
)

type News struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func scrapeNews() ([]News, error) {
	var newsList []News

	c := colly.NewCollector()

	// Extrair as notícias da página do G1, por exemplo
	c.OnHTML(".feed-post-link", func(e *colly.HTMLElement) {
		title := e.Text
		url := e.Attr("href")

		newsList = append(newsList, News{
			Title: title,
			URL:   url,
		})
	})

	// Erro de scraping
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Erro durante o scraping:", err)
	})

	// Visitar o site do G1
	c.Visit("https://g1.globo.com/")

	return newsList, nil
}

func getNews(w http.ResponseWriter, r *http.Request) {
	news, err := scrapeNews()
	if err != nil {
		http.Error(w, "Erro ao fazer scraping das notícias", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(news)
}

func main() {
	// Servir arquivos estáticos da pasta frontend
	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs) // Servir o frontend

	http.HandleFunc("/api/news", getNews) // API de notícias com scraping

	fmt.Println("Servidor iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
