import "./ArticleEditor.css";
import Input from "../Input";
import {ArticleUpdateTitleAction} from "../../api/article-actions";

const ArticleEditor = ({article, dispatchAction}) => {
    return (
        <Input value={article.title} setValue={(val) => dispatchAction(ArticleUpdateTitleAction(val, article.temp_id))} isBold/>
    )
}

export default ArticleEditor;