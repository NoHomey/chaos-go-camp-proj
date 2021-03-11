import * as React from "react"
import { Provider } from "../context/User"
import { User } from "../service/User"

export default function UserProvider({ children }: { children: React.ReactNode }) {
    const [user, setUser] = React.useState<User>({ name: "", email: "" })
    return (
        <Provider value={{ user, setUser }}>
            {children}
        </Provider>
    )
}