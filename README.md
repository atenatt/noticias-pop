# 📰 TecPop News - Fique por dentro das novidades! 🚀

Bem-vindo ao **TecPop News**, o projeto mais minimalista e eficaz para você se manter atualizado com as notícias mais importantes! 🎯 

Aqui, você terá um app simples, sem complicações, que coleta as principais manchetes do **Tecmundo** e te apresenta com um visual clean e moderno. A cada 30 segundos uma nova notícia é exibida na tela, com direito a imagem de alta qualidade e um botão "Saiba Mais" para você mergulhar nos detalhes. 📚

## ✨ Funcionalidades
- **Notícias fresquinhas** direto do **Tecmundo**: pop, tecnologia, política e mais!
- **Imagens em alta resolução** (sem aquelas miniaturas de 164x118!).
- **Interface minimalista**: um fundo claro, manchetes centralizadas e tudo o que você precisa saber em poucos cliques.
- **Atualização automática** a cada 30 segundos: a notícia muda, mas o conteúdo continua top!

## 🔧 Como rodar o projeto?

Vamos lá, sem enrolação! Siga os passos abaixo para rodar o **TecPop News** na sua máquina:

### Backend (Golang) 🛠️

1. Clone este repositório:
   ```bash
   git clone https://github.com/seuusuario/tecpop-news.git
   ```

2. Navegue até o diretório `backend`:
   ```bash
   cd tecpop-news/backend
   ```

3. Instale as dependências (via `go mod`):
   ```bash
   go mod tidy
   ```

4. Rode o backend (ele vai coletar as notícias para você):
   ```bash
   go run main.go
   ```

O backend já vai estar rodando na porta `8080`. 🏃‍♂️💨


Agora é só acessar [http://localhost:8080](http://localhost:8080) no navegador e pronto! A mágica vai acontecer. 🎩✨

## 🧠 Tecnologias Utilizadas

- **Golang**: para coletar as notícias do **Tecmundo** de forma rápida e eficiente.
- **Colly**: a biblioteca de scraping que nos ajuda a pegar as notícias fresquinhas.
- **HTML/CSS/JS**: para deixar o frontend bonito e funcional.
- **Imagens em alta resolução**: porque ninguém merece imagens em baixa, né? 🖼️

## 💡 Ideias futuras
- 🎉 Exibir notícias de várias outras fontes.
- 💬 Adicionar filtros de categorias para o usuário escolher.
- 🖌️ Melhorar o design com mais opções de customização.

## 🏗️ Estrutura de diretórios

```
tecpop-news/
│
├── backend/           # Código do backend (Golang + Colly)
│   ├── main.go        # Arquivo principal do backend
│   └── go.mod         # Dependências do Golang
│
└── frontend/          # Arquivos do frontend
    ├── index.html     # Estrutura da página
    ├── script.js      # Lógica de exibição e troca de notícias
    └── styles.css     # Estilo clean do projeto
```

## 🚀 Contribuições

Quer contribuir? Fique à vontade! Basta fazer um fork do projeto, fazer suas melhorias e abrir um pull request. Vamos deixar o **TecPop News** cada vez mais incrível juntos! 💪

---

Se divirta e fique sempre informado com **TecPop News**! 🎉

