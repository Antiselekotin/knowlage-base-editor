const ArticleUpdateTitleAction = (title, id) => {
    return {
        type: "ARTICLE_UPDATE_TITLE",
        article_id: id,
        title: title
    }
}

const ArticleUpdateContentAction = (content, id) => {
    return {
        type: "ARTICLE_UPDATE_CONTENT",
        article_id: id,
        content
    }
}

const ArticleAddTagAction = (tag_id, article_id) => {
    return {
        type: "ARTICLE_ADD_TAG",
        article_id,
        tag_id
    }
}

const ArticleRemoveTagAction = (tag_id, article_id) => {
    return {
        type: "ARTICLE_REMOVE_TAG",
        article_id,
        tag_id
    }
}

const ArticleAddConnectionAction = (id_1, id_2) => {
    return {
        type: "ARTICLE_ADD_CONNECTION",
        id_1, id_2
    }
}

const ArticleRemoveConnectionAction = (id_1, id_2) => {
    return {
        type: "ARTICLE_REMOVE_CONNECTION",
        id_1, id_2
    }
}

export {
    ArticleAddConnectionAction,
    ArticleAddTagAction,
    ArticleRemoveConnectionAction,
    ArticleRemoveTagAction,
    ArticleUpdateContentAction,
    ArticleUpdateTitleAction,
}

