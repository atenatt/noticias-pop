package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"

    "github.com/gocolly/colly"
)

type News struct {
    Title     string `json:"title"`
    URL       string `json:"url"`
    Image     string `json:"image"`
}

func scrapeTecmundo() ([]News, error) {
    var newsList []News

    c := colly.NewCollector()

    // Visitar as páginas de novidades do Tecmundo
    baseURL := "https://www.tecmundo.com.br/novidades"
    for page := 1; page <= 5; page++ { // Iterar nas páginas 1 a 5
        url := fmt.Sprintf("%s?page=%d", baseURL, page)
        c.Visit(url)

        // Extrair as notícias de cada página
        c.OnHTML(".tec--list__item", func(e *colly.HTMLElement) {
            title := e.ChildAttr(".tec--card__thumb__link", "title")
            newsURL := e.ChildAttr(".tec--card__thumb__link", "href")
            image := e.ChildAttr(".tec--card__thumb__image", "data-src")

            // Se o `data-src` não estiver presente, tente capturar o `src`
            if image == "" {
                image = e.ChildAttr(".tec--card__thumb__image", "src")
            }

            // Alterar a resolução da imagem para 1024x768
            image = strings.Replace(image, "ims=164x118", "ims=1080x720", 1)

            // Remover o prefixo "Ir para:" do título, se estiver presente
            title = strings.TrimPrefix(title, "Ir para: ")

            newsList = append(newsList, News{
                Title: title,
                URL:   newsURL,
                Image: image,
            })
        })
    }

    // Erro de scraping
    c.OnError(func(_ *colly.Response, err error) {
        fmt.Println("Erro durante o scraping:", err)
    })

    return newsList, nil
}

func getNews(w http.ResponseWriter, r *http.Request) {
    news, err := scrapeTecmundo()
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
