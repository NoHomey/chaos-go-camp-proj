import Container from "@material-ui/core/Container"
import Paper from "@material-ui/core/Paper"
import InputField from "../../../component/InputField"
import Button from "@material-ui/core/Button"
import url from "../../../validation/url"
import useForceError from "../../../hook/useForceError"

import { makeStyles } from "@material-ui/core/styles"

const useStyles = makeStyles(theme => ({
    paper: {
        width: "100%",
        marginTop: theme.spacing(6),
        marginBottom: theme.spacing(10),
        padding: theme.spacing(3.5, 5, 3.5, 5)
    },
    form: {
        width: "100%"
    },
    btn: {
        marginTop: theme.spacing(4)
    },
}))

export interface Props {
    data: {
        feedURL: string
    }
    event: {
        onFeedURLChange: (val: string) => void
        onGetDetails: (feedURL: string) => void
    }
}

const Page: React.FC<Props> = ({data, event}) => {
    const cls = useStyles()
    const { feedURL } = data
    const urlRes = url(feedURL)
    const valid = urlRes.valid
    const [forceError, showValidation] = useForceError(valid)
    return (
        <Container component="main" maxWidth="sm">
            <Paper elevation={9} className={cls.paper}>
                <form noValidate className={cls.form}>
                    <InputField
                        required
                        label="Feed URL"
                        value={feedURL}
                        validation={urlRes}
                        forceError={forceError}
                        onValueChange={event.onFeedURLChange} />
                    <Button
                        variant="outlined"
                        color="primary"
                        className={cls.btn}
                        onClick={() => {
                            if(!valid) {
                                showValidation()
                            } else {
                                event.onGetDetails(feedURL)
                            }
                        }}>
                        Get datails
                    </Button>
                </form>
            </Paper>
        </Container>
    )
}

export default Page