import * as React from "react"
import { User } from "../service/User"

export enum State { Init, Sign, User }

export interface Value {
    user: null | User
    state: State
    setUser: (user: null | User) => void
    setStateToSign: () => void
}

const Ctx = React.createContext<Value>({
    user: {
        name: "",
        email: "",
    },
    state: State.Init,
    setUser: error,
    setStateToSign: error
})

export const Provider = Ctx.Provider

export function useSetUser() {
    const ctx = React.useContext(Ctx)
    return ctx.setUser
}

export function useUser() {
    const ctx = React.useContext(Ctx)
    return ctx.user
}

export function useState() {
    const ctx = React.useContext(Ctx)
    return {
        state: ctx.state,
        setUser: ctx.setUser,
        setStateToSign: ctx.setStateToSign
    }
}

export default Provider

function error() {
    throw(new Error("User Context with invalid value"))
}