let newsData = [];
let currentIndex = 0;

async function fetchNews() {
    try {
        const response = await fetch('/api/news'); // Usando o backend na mesma porta
        if (!response.ok) {
            throw new Error('Erro ao buscar notícias');
        }
        const data = await response.json();
        console.log('Dados recebidos:', data); // Verifique os dados no console
        newsData = data;
        displayNews();
    } catch (error) {
        console.error('Erro ao buscar notícias:', error);
    }
}

function displayNews() {
    if (newsData.length === 0) {
        console.log('Nenhuma notícia disponível');
        return;
    }

    const newsContainer = document.getElementById('news');
    newsContainer.innerHTML = '';

    const article = newsData[currentIndex];
    const newsHTML = `
        <h2>${article.title}</h2>
        <a href="${article.url}" target="_blank">Leia mais</a>
    `;

    newsContainer.innerHTML = newsHTML;

    currentIndex = (currentIndex + 1) % newsData.length;
}

function startNewsRotation() {
    displayNews();
    setInterval(displayNews, 30000); // Atualiza a cada 30 segundos
}

document.addEventListener('DOMContentLoaded', () => {
    fetchNews();  // Buscando as notícias ao carregar a página
    startNewsRotation(); // Iniciando a rotação automática
});
