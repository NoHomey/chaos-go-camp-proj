import { Result, valid } from "./Result"
import required from "./result/required"

export default function validate(str: string): Result {
    const l = str.length
    if(l === 0){
        return required()
    }
    return valid()
}