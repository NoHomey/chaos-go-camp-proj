import { BrowserRouter as Router } from "react-router-dom"
import CssBaseline from "@material-ui/core/CssBaseline";
import PrivateRoutes from "./routes/Public"
import ServiceProvider from "./provider/Service"
import ReqDialog from "./provider/ReqDialog"
import User from "./provider/User"

function App() {
    return (
        <Router>
            <CssBaseline />
            <ServiceProvider>
                <User>
                    <ReqDialog>
                        <PrivateRoutes />
                    </ReqDialog>
                </User>
            </ServiceProvider>
        </Router>
    );
}

export default App;
