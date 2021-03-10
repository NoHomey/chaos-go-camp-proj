import { Result, valid } from "./Result";
import min from "./result/minLength"
import max from "./result/maxLength"

export default function Validator(field: string, minLength: number, maxLength: number): (str: string) => Result {
    return str => {
        const l = str.length
        if(l < minLength) {
            return min(field, minLength)
        }
        if(l > maxLength) {
            return max(field, maxLength)
        }
        return valid()
    }
}