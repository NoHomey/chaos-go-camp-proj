import useData from "./useData"
import useSignIn from "./useSignIn"
import Page from "./Page"

export default function SignIn() {
    const {data, onChange} = useData()
    const event = {
        ...onChange,
        onSignIn: useSignIn()
    }
    return <Page data={data} event={event} />
}