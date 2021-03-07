import Container from "@material-ui/core/Container"
import Button from '@material-ui/core/Button'

import { makeStyles } from '@material-ui/core/styles'

const useStyles = makeStyles(theme => ({
    submit: {
        margin: theme.spacing(2, 0, 2, 0)
    },
    form: {
        width: "100%"
    },
    wrap: {
        width: "100%",
        marginTop: theme.spacing(15)
    }
}))

export interface Props {
    children: JSX.Element[]
    link: React.ReactNode
    actionButtonLabel: string
}

const Page: React.FC<Props> = ({
    children,
    link,
    actionButtonLabel
}) => {
    const classes = useStyles()
    return (
        <Container component="main" maxWidth="sm">
            <div className={classes.wrap}>
                <form noValidate className={classes.form}>
                    {children}
                    <Button
                        fullWidth
                        variant="contained"
                        color="primary"
                        className={classes.submit}
                    >
                        {actionButtonLabel}
                    </Button>
                    {link}
                </form>
            </div>
        </Container>
    )
}

export default Page