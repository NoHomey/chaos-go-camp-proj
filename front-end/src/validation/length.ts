import { Result, valid } from "./Result"
import min from "./result/minLength"
import max from "./result/maxLength"
import required from "./result/required"

export default function Validator(field: string, minLength: number, maxLength: number): (str: string) => Result {
    return str => {
        const l = str.length
        if(l === 0){
            return required()
        }
        if(l < minLength) {
            return min(field, minLength)
        }
        if(l > maxLength) {
            return max(field, maxLength)
        }
        return valid()
    }
}