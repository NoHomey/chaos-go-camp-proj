import { Provider, Service } from "../context/Service"
import User from "../service/User"

const service: Service = {
    user: new User()
}

export default function ServiceProvider({ children }: { children: React.ReactNode }) {
    return (
        <Provider value={service}>
            {children}
        </Provider>
    )
}