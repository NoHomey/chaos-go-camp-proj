import { useService } from "../../context/Service"
import { useSetUser } from "../../context/User"
import { useReqDialog } from "../../context/ReqDialog"
import { SignInData } from "../../service/User"

export default function useSignIn() {
    const user = useService("user")
    const setUser = useSetUser()
    const dialog = useReqDialog()
    return function(data: SignInData) {
        const res = user.SignIn(data)
            .OnFail(() => dialog.showFail())
            .OnError(err => {
                const text = JSON.stringify(err, null, 4)
                dialog.showError(text, dialog.close)
            })
            .OnResult(usr => {
                setUser(usr)
                dialog.close()
            })
        dialog.showLoading("Signing in")
        setTimeout(res.Handle.bind(res), 1500)
    }
}