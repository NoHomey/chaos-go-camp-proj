import Layout from "../../layout/SignPage"
import InputField from "../../component/InputField"
import { LabeledCheckbox } from "../../component/LabeledCheckbox"
import Grid from "@material-ui/core/Grid"
import Link from "@material-ui/core/Link"
import { Link as RouterLink } from "react-router-dom"
import routes from "../../routes/map"
import email from "../../validation/email"
import password from "../../validation/password"
import { every } from "../../validation/Result"
import useForceError from "../../hook/useForceError"
import { SignInData } from "../../service/User"

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
    data: {
        email: string
        password: string
        remember: boolean
    }
    event: {
        onEmailChange: (value: string) => void
        onPasswordChange: (value: string) => void
        onRememberChange: () => void
        onSignIn: (data: SignInData) => void
    }
}

const Page: React.FC<Props> = ({data, event}) => {
    const emailRes = email(data.email)
    const passwordRes = password(data.password)
    const valid = every(emailRes, passwordRes)
    const [forceError, showValidation] = useForceError(valid)
    return (
        <Layout actionButtonLabel="Sign in" link={links} onAction={() => {
            if(!valid) {
                showValidation()
            } else {
                event.onSignIn(data)
            }
        }}>
            <InputField
                label="Email address"
                type="email"
                required
                autoComplete="email"
                value={data.email}
                validation={emailRes}
                forceError={forceError}
                onValueChange={event.onEmailChange}
            />
            <InputField
                label="Password"
                type="password"
                required
                value={data.password}
                validation={passwordRes}
                forceError={forceError}
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