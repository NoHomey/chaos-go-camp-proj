import { Provider, Service } from "../context/Service"
import User from "../service/User"

const service: Service = {
    user: new User()
}

const ServiceProvider: React.FC<{}> = ({ children }) => {
    return (
        <Provider value={service}>
            {children}
        </Provider>
    )
}

export default ServiceProvider