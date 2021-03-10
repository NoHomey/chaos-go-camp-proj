export interface Result {
    valid: boolean
    message: string
}

export function invalid(msg: string): Result {
    return {
        valid: false,
        message: msg
    }
}

export function valid(): Result {
    return {
        valid: true,
        message: ""
    }
}

export type Validator = () => Result

export function all(vs: Array<Validator>): boolean {
    return vs.every(v => v().valid)
}

export function errorMsg(res: Result): undefined | string {
    return res.valid ? undefined : res.message
}

export default Result