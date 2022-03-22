import ArticleEditor from "../ArticleEditor";

const MainPage = ({article, tag, dispatchAction}) => {
    if(article) {
        return <ArticleEditor article={article} dispatchAction={dispatchAction}/>
    }
    if(tag) {
        return <div>{tag?.title}</div>
    }
    return <div>{tag?.title} | {article?.title}</div>
}

export default MainPage;