const getTags = async () => {
    const tagsRequest = await fetch("/api/tags")
    return await tagsRequest.json()
}

const getArticles = async () => {
    const articlesRequest = await fetch("/api/articles")
    return await articlesRequest.json()
}

const initPromise = async () => {
    const tags = await getTags();
    const articles = await getArticles()
    return {
        tags: tags,
        articles: articles
    }
}

const init = (onSuccess, onFailure) => {
    initPromise()
        .then(onSuccess)
        .catch(onFailure)
}

export {
    init,
    getTags,
    getArticles
}