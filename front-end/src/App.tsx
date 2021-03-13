import * as React from "react"
import AppLayout from "./layout/app"
import Public from "./routes/Public"
import Private from "./routes/Private"
import { useService } from "./context/Service"
import { useState, State } from "./context/User"
import { useReqDialog } from "./context/ReqDialog"
import { Response } from "./response"
import { User } from "./service/User"
import { useHistory } from "react-router-dom"

export default function App() {
    const userService = useService("user")
    const dialog = useReqDialog()
    const history = useHistory()
    const { state, setUser, setStateToSign } = useState()
    const load = (resp: Response<User>) => {
        resp
            .OnFail(() => {
                dialog.showFail(() => {
                    dialog.showLoading("Trying again")
                    const access = userService.Reload()
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
                history.replace("/")
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