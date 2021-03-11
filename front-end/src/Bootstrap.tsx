import { BrowserRouter as Router } from "react-router-dom"
import CssBaseline from "@material-ui/core/CssBaseline";
import ServiceProvider from "./provider/Service"
import ReqDialog from "./provider/ReqDialog"
import User from "./provider/User"
import App from "./App"

export default function Bootstrap() {
    return (
        <Router>
            <ServiceProvider>
                <User>
                    <ReqDialog>
                        <CssBaseline />
                        <App />
                    </ReqDialog>
                </User>
            </ServiceProvider>
        </Router>
    );
}
