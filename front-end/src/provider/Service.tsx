import { Provider, init } from "../context/Service"

const ServiceProvider: React.FC<{}> = ({ children }) => {
    return (
        <Provider value={init}>
            {children}
        </Provider>
    )
}

export default ServiceProvider