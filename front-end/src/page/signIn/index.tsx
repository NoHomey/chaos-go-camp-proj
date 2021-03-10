import Layout from "../../layout/SignPage"
import InputField from "../../component/InputField"
import FormControlLabel from "@material-ui/core/FormControlLabel"
import Checkbox from "@material-ui/core/Checkbox"
import Grid from "@material-ui/core/Grid"
import Link from "@material-ui/core/Link"
import { Link as RouterLink } from "react-router-dom"
import routes from "../../routes/map"

const links = (
    <Grid container>
        <Grid item xs>
            <Link to="/" variant="body2" component={RouterLink}>
                Forgot password?
            </Link>
        </Grid>
        <Grid item>
            <Link to={routes.signUp} variant="body2" component={RouterLink}>
                Don"t have an account? Sign up
            </Link>
        </Grid>
    </Grid>
)

export default function Page() {
    return (
        <Layout actionButtonLabel="Sign in" link={links}>
            <InputField
                label="Email address"
                type="email"
                required
                autoComplete="email"
            />
            <InputField
                label="Password"
                type="password"
                required
            />
            <FormControlLabel
                control={<Checkbox value="remember" color="primary" />}
                label="Remember me"
            />
        </Layout>
    )
} 