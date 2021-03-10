import { validate as isEmail } from "email-validator"
import { Result, invalid, valid } from "./Result"
import required from "./result/required"

export default function validate(str: string): Result {
    if(str.length === 0){
        return required()
    }
    if(isEmail(str)) {
        return valid()
    }
    return invalid("This is not a valid email address")
}