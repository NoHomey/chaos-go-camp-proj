import { Result, valid, invalid } from "./Result"

export default function validate(str: string): Result {
    try {
        const _ = new URL(str)
        return valid()
    } catch(err) {
        return invalid("This is not a valid URL")
    }
}