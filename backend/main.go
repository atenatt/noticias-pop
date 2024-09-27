package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gocolly/colly"
)

type News struct {
    Title     string `json:"title"`
    Subtitle  string `json:"subtitle"`
    URL       string `json:"url"`
    Image     string `json:"image"`
}

func scrapeNews() ([]News, error) {
    var newsList []News

    c := colly.NewCollector()

    // Extrair as notícias, incluindo título, subtítulo, URL e imagem
    c.OnHTML(".feed-post-body", func(e *colly.HTMLElement) {
        title := e.ChildText(".feed-post-link")
        subtitle := e.ChildText(".feed-post-body-resumo p") // Captura do subtítulo corrigida
        url := e.ChildAttr(".feed-post-link", "href")
        image := e.ChildAttr("img", "src")

        newsList = append(newsList, News{
            Title:    title,
            Subtitle: subtitle,
            URL:      url,
            Image:    image,
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

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(news)
}

func main() {
    // Servir arquivos estáticos da pasta frontend
    fs := http.FileServer(http.Dir("../frontend"))
    http.Handle("/", fs)

    // API para as notícias
    http.HandleFunc("/api/news", getNews)

    fmt.Println("Servidor iniciado em http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
