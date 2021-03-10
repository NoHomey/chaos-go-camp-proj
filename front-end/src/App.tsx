import { BrowserRouter as Router } from 'react-router-dom'
import CssBaseline from '@material-ui/core/CssBaseline';
import PrivateRoutes from './routes/Public'

function App() {
    return (
        <Router>
            <CssBaseline />
            <PrivateRoutes />
        </Router>
    );
}

export default App;
