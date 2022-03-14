import "./Header.css"

const Header = ({children, className}) => {
    let cN = className ? className : "";
    cN = "header " + cN.split(" ").filter(w => w).join(" ")
    return <div className={cN}>{children}</div>
}

export default Header