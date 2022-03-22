import "./App.css"
import Header from "../Header";
import {getArticles, getTags, init} from "../../api/basic";
import {useEffect, useState} from "react";
import Spinner from "../Spinner";
import ScrollableBlock from "../ScrollableBlock";
import SelectableWrapper from "../SelectableWrapper";
import MainPage from "../MainPage";
import {newStore} from "../../api/actions-store";

const App = () => {
    const [tags, setTags] = useState(null)
    const [articles, setArticles] = useState(null)
    const [isInitError, setIsInitError] = useState(false)
    const [isLoading, setIsLoading] = useState(true)
    const [selectedArticleId, setSelectedArticleId] = useState(null)
    const [selectedTagId, setSelectedTagId] = useState(null)
    const [store, setStore] = useState(null)

    const selectTag = (tag_id) => () => {
        setSelectedTagId(tag_id)
        setSelectedArticleId(null)
    }
    const selectArticle = (article_id) => () => {
        setSelectedArticleId(article_id)
        setSelectedTagId(null)
    }
    const onLoadSuccess = (data) => {
        const tmpStore = newStore(data.articles, data.tags)
        setStore(tmpStore)
        setTags(tmpStore.tags)
        setArticles(tmpStore.articles)
        setIsLoading(false)
    }
    const onLoadFailure = () => {
        setIsInitError(true)
        setIsLoading(false)
    }
    useEffect(() => {
        init(onLoadSuccess, onLoadFailure)
    }, [])

    const dispatchAction = (action) => {
        const {tags: newTags, articles: newArticles} = store.dispatchAction(action)
        setTags(newTags)
        setArticles(newArticles)
    }

    if (isLoading) {
        return null
    }
    const articlesList = isLoading ? <Spinner/> :
        <ScrollableBlock items={Object.values(articles).map((article) => <SelectableWrapper children={article.title} onClick={selectArticle(article.temp_id)}/>)}
                         maxHeight="calc(50% - 5rem)"/>;
    const tagsList = isLoading ? <Spinner/> :
        <ScrollableBlock items={Object.values(tags).map((tag) => <SelectableWrapper children={tag.title} onClick={selectTag(tag.temp_id)}/>)}
                                                                maxHeight="calc(50% - 5rem)"/>;
    return (
        <div className="app">
            <div className="app__side-menu">
                <Header>Статьи</Header>
                {articlesList}
                <Header className='mt-3'>Теги</Header>
                {tagsList}
            </div>
            <div className="app__main-window">
                <MainPage article={articles[selectedArticleId]} tag={tags[selectedTagId]} dispatchAction={dispatchAction}/>
            </div>
        </div>
    );
}

export default App;
