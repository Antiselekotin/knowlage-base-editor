import DefaultWrapper from "../DefaultWrapper";
import "./ScrollableBlock.css"

const ScrollableBlock = ({items, maxHeight = "100%", wrapper = DefaultWrapper}) => {
    return (
        <div style={{maxHeight: maxHeight}} className="scrollable">
             {items.map((item, index) => wrapper({children: item, key: index}))}
        </div>
    )
}

export default ScrollableBlock