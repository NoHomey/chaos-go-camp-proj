import * as React from "react"
import { User } from "../service/User"

export interface Value {
    user: User
    setUser: (user: User) => void
}

const Ctx = React.createContext<Value>({
    user: {
        name: "",
        email: "",
    },
    setUser: () => { throw(new Error("User Context with invalid value")) }
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

export default Provider