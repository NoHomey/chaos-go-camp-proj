import * as React from "react"
import AppLayout from "./layout/app"
import Public from "./routes/Public"
import Private from "./routes/Private"
import { useService } from "./context/Service"
import { useState, State } from "./context/User"
import { useReqDialog } from "./context/ReqDialog"

export default function App() {
    const userService = useService("user")
    const dialog = useReqDialog()
    const { state, setUser, setStateToSign } = useState()
    React.useEffect(() => {
        let access = userService.Reload()
        if(access === null) {
            setStateToSign()
        } else {
            access = access
                .OnFail(() => {
                    dialog.showFail()
                    setStateToSign()
                })
                .OnError(err => {
                    setStateToSign()
                    dialog.close()
                })
                .OnResult(usr => {
                    setUser(usr)
                    dialog.close()
                })
            dialog.showLoading("Getting access")
            setTimeout(access.Handle.bind(access), 1000)
        }
    }, [])

    const onSignOut = () => {
        const res = userService.SignOut()
            .OnFail(() => dialog.showFail())
            .OnError(err => {
                dialog.showError("Failed to sign out", dialog.close)
            })
            .OnResult(() => {
                setUser(null)
                dialog.close()
            })
        dialog.showLoading("Signing out")
        setTimeout(res.Handle.bind(res), 1000)
    }
    switch(state) {
    case State.Init:
        return null
    case State.Sign:
        return <Public />
    case State.User:
        return (
            <AppLayout onSignOut={onSignOut}>
                <Private />
            </AppLayout>
        )
    }
}