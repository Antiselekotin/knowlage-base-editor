import deepcopy from "deepcopy";

const newStore = (articlesArray, tagsArray) => {
    const actionsStore = {
        articleTitleUpdates: {}, // article_id: {...action}
        articleContentUpdates: {}, // article_id: {...action}
        articleAddTag: {}, // article_id: {tag_id: {...action}, tag_id: {..action}}
        articleRemoveTag: {}, // article_id: {tag_id: {...action}, tag_id: {..action}}
        articleAddConnection: {}, // smaller_id: {bigger_id: {..action}}
        articleRemoveConnection: {}, // smaller_id: {bigger_id: {..action}}
    }

    const articlesBase = {}
    const tagsBase = {}

    articlesArray.forEach(article => articlesBase[article.temp_id] = deepcopy(article))
    tagsArray.forEach(tag => tagsBase[tag.temp_id] = deepcopy(tag))

    const articlesCurrent = deepcopy(articlesBase)
    const tagsCurrent = deepcopy(tagsBase)

    const dispatchAction = (action) => {
        switch (action.type) {
            case "ARTICLE_UPDATE_TITLE": {
                articlesCurrent[action.article_id].title = action.title
                if(action.title === articlesBase[action.article_id]) {
                    delete actionsStore.articleTitleUpdates[action.article_id]
                } else {
                    actionsStore.articleTitleUpdates[action.article_id] = action
                }
            } break;
            case "ARTICLE_UPDATE_CONTENT": {
                articlesCurrent[action.article_id].content = action.content
                if(action.content === articlesBase[action.article_id]) {
                    delete actionsStore.articleContentUpdates[action.article_id]
                } else {
                    actionsStore.articleContentUpdates[action.article_id] = action
                }
            } break;
            case "ARTICLE_ADD_TAG": {
                articlesCurrent[action.article_id].tags.push(action.tag_id)
                tagsCurrent[action.tag_id].articles.push(action.article_id)
                const node = actionsStore.articleRemoveTag[action.article_id]
                if(node && node[action.tag_id]) {
                    delete actionsStore.articleRemoveTag[action.article_id][action.tag_id]
                } else if (!articlesBase[action.article_id].tags.contains(action.tag_id)){
                    actionsStore.articleAddTag[action.article_id][action.tag_id] = actionsStore
                }
            } break;
            case "ARTICLE_REMOVE_TAG": {
                articlesCurrent[action.article_id].tags = articlesCurrent[action.article_id].tags
                    .filter((tag_id) => action.tag_id !== tag_id)
                tagsCurrent[action.tag_id].articles = tagsCurrent[action.tag_id].articles
                    .filter(article_id => action.article_id !== article_id)
                const node = actionsStore.articleAddTag[action.article_id]
                if(node && node[action.tag_id]) {
                    delete actionsStore.articleAddTag[action.article_id][action.tag_id]
                } else if (articlesBase[action.article_id].tags.contains(action.tag_id)){
                    actionsStore.articleRemoveTag[action.article_id][action.tag_id] = actionsStore
                }
            } break;
            case "ARTICLE_ADD_CONNECTION": {
                const min_id = Math.min(action.id_1, action.id_2)
                const max_id = Math.max(action.id_1, action.id_2)
                articlesCurrent[min_id].connections.push(max_id)
                articlesCurrent[max_id].connections.push(min_id)
                const node = actionsStore.articleRemoveConnection[min_id]
                if(node && node[max_id]) {
                    delete actionsStore.articleRemoveConnection[min_id][max_id]
                } else if(!articlesBase[action.id_1].contains(action.id_2)) {
                    actionsStore.articleAddConnection[min_id][max_id] = actionsStore
                }
            } break;
            case "ARTICLE_REMOVE_CONNECTION": {
                const min_id = Math.min(action.id_1, action.id_2)
                const max_id = Math.max(action.id_1, action.id_2)
                articlesCurrent[min_id] = articlesCurrent[min_id].connections.filter(id => id !== max_id)
                articlesCurrent[max_id] = articlesCurrent[max_id].connections.filter(id => id != min_id)
                const node = actionsStore.articleAddConnection[min_id]
                if(node && node[max_id]) {
                    delete actionsStore.articleAddConnection[min_id][max_id]
                } else if(articlesBase[action.id_1].contains(action.id_2)) {
                    actionsStore.articleRemoveConnection[min_id][max_id] = actionsStore
                }
            } break;
        }
        return {
            articles: deepcopy(articlesCurrent),
            tags: deepcopy(tagsCurrent)
        }
    }

    const getActionsList = () => {
        const buff = [
            ...getOneLevelDeepActions(actionsStore.articleTitleUpdates),
            ...getOneLevelDeepActions(actionsStore.articleContentUpdates),
            ...getTwoLevelDeepActions(actionsStore.articleRemoveTag),
            ...getTwoLevelDeepActions(actionsStore.articleAddTag),
            ...getTwoLevelDeepActions(actionsStore.articleAddConnection),
            ...getTwoLevelDeepActions(actionsStore.articleRemoveConnection)];
        return buff
    }

    return {
        dispatchAction,
        getActionsList,
        articles: deepcopy(articlesCurrent),
        tags: deepcopy(tagsCurrent)
    }
}

export {
    newStore
}

const getOneLevelDeepActions = (data) => {
    const buff = [];
    for (const actionKey in data) {
        buff.push(data[actionKey])
    }
    return buff
}

const getTwoLevelDeepActions = (data) => {
    const buff = [];
    for (const objKey in data) {
        const obj = data[objKey]
        for (const actionKey in obj) {
            buff.push(obj[actionKey])
        }
    }
    return buff
}

const objToArray = (obj) => {
    const buff = [];
    for(const key in obj) {
        buff.push(deepcopy(obj[key]));
    }
    return buff;
}