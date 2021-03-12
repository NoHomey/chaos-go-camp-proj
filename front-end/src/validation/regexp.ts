import { Result, invalid, valid } from "./Result"
import required from "./result/required"

export default function Validator(regexp: RegExp, msg: string): (str: string) => Result {
    return str => {
        const l = str.length
        if(l === 0){
            return required()
        }
        if(!regexp.test(str)) {
            return invalid(msg)
        }
        return valid()
    }
}