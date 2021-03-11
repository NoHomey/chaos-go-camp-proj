import Layout from "../../layout/SignPage"
import InputField from "../../component/InputField"
import Grid from "@material-ui/core/Grid"
import Link from "@material-ui/core/Link"
import { Link as RouterLink } from "react-router-dom"
import routes from "../../routes/map"
import email from "../../validation/email"
import password from "../../validation/password"
import { LabeledCheckbox } from "../../component/LabeledCheckbox"

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

export interface Props {
    model: {
        data: {
            email: string
            password: string
            remember: boolean
        }
        event: {
            onEmailChange: (value: string) => void
            onPasswordChange: (value: string) => void
            onRememberChange: () => void
        }
    }
}

const Page: React.FC<Props> = ({model}) => {
    const {data, event} = model
    return (
        <Layout actionButtonLabel="Sign in" link={links}>
            <InputField
                label="Email address"
                type="email"
                required
                autoComplete="email"
                value={data.email}
                validation={email(data.email)}
                onValueChange={event.onEmailChange}
            />
            <InputField
                label="Password"
                type="password"
                required
                value={data.password}
                validation={password(data.password)}
                onValueChange={event.onPasswordChange}
            />
            <LabeledCheckbox
                label="Remember me"
                checked={data.remember}
                onToggle={event.onRememberChange} />
        </Layout>
    )
}

export default Page