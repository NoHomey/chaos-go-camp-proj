import useModel from "./useModel"
import Page from "./Page"

export default function SignUp() {
    const model = useModel()
    return <Page model={model} />
}