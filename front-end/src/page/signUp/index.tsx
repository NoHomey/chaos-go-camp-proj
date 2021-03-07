import Layout from "../../layout/SignPage"
import SignInput from "../../component/SignInput";
import Link from '@material-ui/core/Link';

const link = (
    <Link href="#" variant="body2">
        Have an account? Sign in
    </Link>
)

export default function Page() {
    return (
        <Layout actionButtonLabel="Sign up" link={link}>
            <SignInput
                label="Name"
                required
            />
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
        </Layout>
    )
} 