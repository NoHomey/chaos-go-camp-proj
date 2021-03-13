import { Result, valid, invalid } from "./Result"
import required from "./result/required"

export default function validate(str: string): Result {
    const l = str.length
    if(l === 0){
        return required()
    }
    try {
        const _ = new URL(str)
        return valid()
    } catch(err) {
        return invalid("This is not a valid URL")
    }
}