import { validate as isEmail } from "email-validator"
import { Result, invalid, valid } from "./Result";

export default function validate(str: string): Result {
    if(isEmail(str)) {
        return valid()
    }
    return invalid("This is not a valid email address")
}