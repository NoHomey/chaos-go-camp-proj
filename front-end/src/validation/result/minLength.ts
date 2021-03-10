import { Result, invalid } from "../Result"

export default function res(field: string, min: number): Result {
    return invalid(`${field} must be at least ${min} symbols long`)
}