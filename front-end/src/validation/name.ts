import length from "./length"

export const minLength = 3

export const maxLength = 64

const validate = length("Name", minLength, maxLength)

export default validate