import {useState} from "react";
import "./Input.css";

const Input = ({value, setValue, isBold = false}) => {
    const [isFocused, setIsFocused] = useState(false);
    let className = "input";
    if (isBold) {
        className += " input--bold";
    }
    if (isFocused) {
        className += " input--focused";
    }
    return (
        <div className={className}>
            <input type="text" value={value}
                   onInput={(e) => setValue(e.target.value)}
                   onFocus={() => setIsFocused(true)}
                   onBlur={() => setIsFocused(false)}/>
        </div>
    )
}

export default Input;