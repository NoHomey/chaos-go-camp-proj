import regexp from "./regexp"

const test = /^\w+(?:(?:-|_)\w+)*$/

const msg = "Valid tags consists of alpha numeric words separated by dash or underscore"

const validate = regexp(test, msg)

export default validate