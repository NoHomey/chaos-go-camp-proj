import { Result, valid, invalid } from "./Result"
import min from "./result/minLength"
import max from "./result/maxLength"
import required from "./result/required"

export default function Validate(str: string): Result {
    const l = str.length
    if(l === 0){
        return required()
    }
    if(l < minLength) {
        return min("Password", minLength)
    }
    if(l > maxLength) {
        return max("Password", maxLength)
    }
    let lowerCase = false
    let upperCase = false
    let digit = false
    for(const x of str) {
        if(inRange(x, 'a', 'z')) {
            lowerCase = true
            continue
        }
        if(inRange(x, 'A', 'Z')) {
            upperCase = true
            continue
        }
        if(inRange(x, '0', '9')) {
            digit = true
            continue
        }
    }
    if(!lowerCase) {
        return invalid("Password must include at least one lowercased english letter")
    }
    if(!upperCase) {
        return invalid("Password must include at least one uppercased english letter")
    }
    if(!digit) {
        return invalid("Password must include at least one digit")
    }
    return valid()
}

function inRange(x: string, a: string, b: string): boolean {
    return a <= x && x <= b
}

export const minLength = 8

export const maxLength = 32