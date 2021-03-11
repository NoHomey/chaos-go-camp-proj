import * as React from "react"
import User, { Service as UserService } from "../service/User" 

export interface Service {
    user: UserService
}

const Ctx = React.createContext<Service>({
    user: new User()
})

export const Provider = Ctx.Provider

export function useAll() {
    return React.useContext(Ctx)
}

export function useService<K extends keyof Service>(name: K): Service[K] {
    const ctx = useAll()
    return ctx[name]
}

export default Provider