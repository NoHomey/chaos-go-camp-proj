/*import Container from "@material-ui/core/Container"
import Paper from "@material-ui/core/Paper"
import InputField from "../../component/InputField"
import Rating from "../../component/Rating"
import Level from "../../component/Level"
import Button from "@material-ui/core/Button"
import TextField from "@material-ui/core/TextField"
import Accordion from "../../component/Accordion"
import Tags from "../../component/Tags"
import SaveIcon from "@material-ui/icons/Save"

import { makeStyles } from "@material-ui/core/styles"
import { TagPriorty } from "../../data/Tag"

const saveIcon = <SaveIcon />

const useStyles = makeStyles(theme => ({
    form: {
        width: "100%"
    },
    paper: {
        width: "100%",
        marginTop: theme.spacing(6),
        marginBottom: theme.spacing(10),
        padding: theme.spacing(3.5, 5, 3.5, 5)
    },
    saveBtn: {
        marginTop: theme.spacing(4)
    },
    desc: {
        margin: theme.spacing(2, 0, 2, 0)
    }
}))

const shrinkLabel = { shrink: true }

const Page: React.FC<{}> = () => {
    const classes = useStyles()
    return (
        <Container component="main" maxWidth="md">
            <Paper elevation={9} className={classes.paper}>
                <form noValidate className={classes.form}>
                    <InputField InputLabelProps={shrinkLabel} label="Feed URL" value="http://some-url.com" />
                    <InputField InputLabelProps={shrinkLabel} label="Author" value="Some Author" />
                    <InputField InputLabelProps={shrinkLabel} label="Blog title" value="Some Blog Title" />
                    <Accordion
                        className={classes.desc}
                        title="Blog description from the author"
                        body="Some text"
                    />
                    <Rating mt={4} mb={1} />
                    <Level mt={4} mb={1} />
                    <Tags marginY={2} tags={[
                        { value: "go-lang", priority: TagPriorty.Main },
                        { value: "go-blog", priority: TagPriorty.Secondary },
                        { value: "go-videos", priority: TagPriorty.Normal },
                        { value: "go-confs", priority: TagPriorty.Normal },
                    ]} />
                    <TextField
                        fullWidth
                        variant="outlined"
                        margin="normal"
                        label="Quick note"
                        multiline
                        rows={3}
                        rowsMax={12}
                    />
                    <Button
                        variant="outlined"
                        color="primary"
                        endIcon={saveIcon}
                        className={classes.saveBtn}
                    >
                        Save Blog
                    </Button>
                </form>
            </Paper>
        </Container>
    )
}

export default Page*/

export default function Page() { return null }