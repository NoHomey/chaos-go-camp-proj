import { Result, invalid } from "../Result"

export default function res(field: string, max: number): Result {
    return invalid(`${field} must be at most ${max} symbols long`)
}