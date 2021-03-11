import useModel from "./useModel"
import Page from "./Page"

export default function SignIn() {
    const model = useModel()
    return <Page model={model} />
}