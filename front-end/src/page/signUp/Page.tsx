import Layout from "../../layout/SignPage"
import InputField from "../../component/InputField"
import Link from "@material-ui/core/Link"
import { Link as RouterLink } from "react-router-dom"
import routes from "../../routes/map"
import email from "../../validation/email"
import password from "../../validation/password"
import { Result, valid, invalid, every } from "../../validation/Result"
import useForceError from "../../hook/useForceError"

const link = (
    <Link to={routes.signIn} component={RouterLink} variant="body2">
        Have an account? Sign in
    </Link>
)

export interface Props {
    model: {
        data: {
            email: string
            password: string
            confirmPassword: string
        }
        event: {
            onEmailChange: (value: string) => void
            onPasswordChange: (value: string) => void
            onConfirmPasswordChange: (value: string) => void
        }
    }
}

const Page: React.FC<Props> = ({model}) => {
    const {data, event} = model
    const emailRes = email(data.email)
    const passwordRes = password(data.password)
    const confirmPasswordRes = match(data.password, data.confirmPassword)
    const valid = every(emailRes, passwordRes, confirmPasswordRes)
    const [forceError, showValidation] = useForceError(valid)
    return (
        <Layout actionButtonLabel="Sign in" link={link} onAction={() => {
            if(!valid) {
                showValidation()
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
            <InputField
                label="Confirm password"
                type="password"
                required
                value={data.confirmPassword}
                validation={confirmPasswordRes}
                forceError={forceError}
                onValueChange={event.onConfirmPasswordChange}
            />
        </Layout>
    )
}

function match(password: string, confirmPassword: string): Result {
    if(password === confirmPassword) {
        return valid()
    }
    return invalid("Confirm password and Password must match")
}

export default Page