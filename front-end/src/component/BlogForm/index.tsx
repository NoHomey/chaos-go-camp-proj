import InputField from "../InputField"
import TextField from "@material-ui/core/TextField"
import Rating from "../Rating"
import Level from "../Level"
import Button from "@material-ui/core/Button"
import Accordion from "../Accordion"
import Tagger from "../Tagger"
import SaveIcon from "@material-ui/icons/Save"
import Blog from "../../data/Blog"
import { makeStyles } from "@material-ui/core/styles"
import useData from "./useData"
import required from "../../validation/required"
import { every } from "../../validation/Result"
import useForceError from "../../hook/useForceError"

const saveIcon = <SaveIcon />

const useStyles = makeStyles(theme => ({
    form: {
        width: "100%"
    },
    saveBtn: {
        marginTop: theme.spacing(4)
    },
    desc: {
        margin: theme.spacing(2, 0, 2, 0)
    }
}))

const shrinkLabel = { shrink: true }

export interface Props {
    blog: Blog
    actionBtnLabel: string
    onAction: (blog: Blog) => void
}

const Page: React.FC<Props> = ({
    blog,
    actionBtnLabel,
    onAction
}) => {
    const classes = useStyles()
    const { data, event } = useData(blog)
    const authorRes = required(data.author)
    const titleRes = required(data.title)
    const valid = every(authorRes, titleRes)
    const [forceError, showValidation] = useForceError(valid)
    return (
        <form noValidate className={classes.form}>
            <TextField
                variant="outlined"
                margin="normal"
                fullWidth
                InputLabelProps={shrinkLabel}
                label="Feed URL"
                value={data.feedURL}
                disabled />
            <InputField
                required
                InputLabelProps={shrinkLabel}
                label="Author"
                value={data.author}
                validation={authorRes}
                forceError={forceError}
                onValueChange={event.onAuthorChange} />
            <InputField
                required
                InputLabelProps={shrinkLabel}
                label="Blog title"
                value={data.title}
                validation={titleRes}
                forceError={forceError}
                onValueChange={event.onTitleChange} />
            <Accordion
                className={classes.desc}
                title="Blog description from the author"
                body={data.description}
            />
            <Rating
                mt={4}
                mb={1}
                value={data.rating}
                onValueChange={event.onRatingChange} />
            <Level
                mt={4}
                mb={1}
                value={data.level}
                onValueChange={event.onLevelChange} />
            <Tagger
                marginY={2}
                tags={data.tags}
                onAddTag={event.onAddTag}
                onRemoveTag={event.onRemoveTag} />
            <TextField
                fullWidth
                variant="outlined"
                margin="normal"
                label="Quick note"
                multiline
                rows={3}
                rowsMax={12}
                value={data.quickNote}
                onChange={e => event.onQuickNoteChange(e.target.value)}
            />
            <Button
                variant="outlined"
                color="primary"
                endIcon={saveIcon}
                className={classes.saveBtn}
                onClick={() => {
                    if(!valid) {
                        showValidation()
                    } else {
                        onAction(data)
                    }
                }}>
                {actionBtnLabel}
            </Button>
        </form>
    )
}

export default Page