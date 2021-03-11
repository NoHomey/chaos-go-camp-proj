import * as React from "react"
import AppLayout from "./layout/app"
import Public from "./routes/Public"
import Private from "./routes/Private"
import { useService } from "./context/Service"
import { useState, State } from "./context/User"
import { useReqDialog } from "./context/ReqDialog"
import { Response } from "./response"
import { User } from "./service/User"

export default function App() {
    const userService = useService("user")
    const dialog = useReqDialog()
    const { state, setUser, setStateToSign } = useState()
    const load = (resp: Response<User>) => {
        const access = resp
            .OnFail(() => {
                dialog.showFail(() => {
                    const access = userService.Reload()
                    dialog.showLoading("Trying again")
                    setTimeout(() => load(access!), 2000)
                })
            })
            .OnError(err => {
                setStateToSign()
                dialog.close()
            })
            .OnResult(usr => {
                setUser(usr)
                dialog.close()
            }).Handle()
    }
    React.useEffect(() => {
        let access = userService.Reload()
        if(access === null) {
            setStateToSign()
        } else {
            dialog.showLoading("Getting access")
            setTimeout(() => load(access!), 1000)
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