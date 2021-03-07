import Layout from "../../layout/SignPage"
import SignInput from "../../component/SignInput";
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Grid from '@material-ui/core/Grid';
import Link from '@material-ui/core/Link';

const links = (
    <Grid container>
        <Grid item xs>
            <Link href="#" variant="body2">
                Forgot password?
            </Link>
        </Grid>
        <Grid item>
            <Link href="#" variant="body2">
                Don't have an account? Sign up
            </Link>
        </Grid>
    </Grid>
)

export default function Page() {
    return (
        <Layout actionButtonLabel="Sign in" link={links}>
            <SignInput
                label="Email address"
                type="email"
                required
                autoComplete="email"
            />
            <SignInput
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