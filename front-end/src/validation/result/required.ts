import { Result, invalid } from "../Result"

export default function required(): Result {
    return invalid("This field is required")
}