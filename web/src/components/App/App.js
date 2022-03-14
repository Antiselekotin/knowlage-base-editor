import "./App.css"
import Header from "../Header";
import {getArticles, getTags, init} from "../../api/basic";
import {useEffect, useState} from "react";
import Spinner from "../Spinner";
import ScrollableBlock from "../ScrollableBlock";
import SelectableWrapper from "../SelectableWrapper";

const App = () => {
    const [tags, setTags] = useState(null)
    const [articles, setArticles] = useState(null)
    const [isInitError, setIsInitError] = useState(false)
    const [isLoading, setIsLoading] = useState(true)
    const [selectedArticle, setSelectedArticle] = useState(null)
    const [selectedTag, setSelectedTag] = useState(null)

    const selectTag = (tag) => () => {
        setSelectedTag(tag)
        setSelectedArticle(null)
    }
    const selectArticle = (article) => () => {
        setSelectedArticle(article)
        setSelectedTag(null)
    }
    const onLoadSuccess = (data) => {
        setTags(data.tags)
        setArticles(data.articles)
        setIsLoading(false)
    }
    const onLoadFailure = () => {
        setIsInitError(true)
        setIsLoading(false)
    }
    useEffect(() => {
        init(onLoadSuccess, onLoadFailure)
    }, [])
    const articlesList = isLoading ? <Spinner/> :
        <ScrollableBlock items={articles.map((article) => <SelectableWrapper children={article.title} onClick={selectArticle(article)}/>)}
                         maxHeight="calc(50% - 5rem)"/>;
    const tagsList = isLoading ? <Spinner/> :
        <ScrollableBlock items={tags.map((tag) => <SelectableWrapper children={tag.title} onClick={selectTag(tag)}/>)}
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
                {selectedTag?.title} | {selectedArticle?.title}
            </div>
        </div>
    );
}

export default App;
