import "./SelectableWrapper.css"

const SelectableWrapper = ({children, onClick = () => {}}) => {
    return (
        <div className="selectable-wrapper" onClick={onClick}>{children}</div>
    )
}

export default SelectableWrapper