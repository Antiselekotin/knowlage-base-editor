const Spinner = ({size}) => {
    const sizeText = size ? `${size}px` : '2rem'
    return (
        <div className="spinner-border" role="status" style={{height: sizeText, width: sizeText}} >
        </div>
    )
}

export default Spinner