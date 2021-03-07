import Layout from "../../layout/SignPage"
import InputField from "../../component/InputField";
import Link from '@material-ui/core/Link';

const link = (
    <Link href="#" variant="body2">
        Have an account? Sign in
    </Link>
)

export default function Page() {
    return (
        <Layout actionButtonLabel="Sign up" link={link}>
            <InputField
                label="Name"
                required
            />
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
        </Layout>
    )
} 