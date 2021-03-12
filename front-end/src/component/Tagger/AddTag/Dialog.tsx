import Dialog from "@material-ui/core/Dialog"
import DialogTitle from "@material-ui/core/DialogTitle"
import DialogContent from "@material-ui/core/DialogContent"
import DialogActions from "@material-ui/core/DialogActions"
import IconButton from "@material-ui/core/IconButton"
import CloseIcon from "@material-ui/icons/Close"
import Typography from "@material-ui/core/Typography"
import InputField from "../../../component/InputField"
import SelectComp, { Props as SelectProps } from "../../../component/Select"
import MenuItem from "@material-ui/core/MenuItem"
import Button from "@material-ui/core/Button"
import Grid from "@material-ui/core/Grid"
import tag from "../../../validation/tag"
import { Tag, TagPriorty } from "../../../data/Tag"
import { makeStyles } from "@material-ui/core/styles"
import useForceError from "../../../hook/useForceError"

const useStyles = makeStyles(theme => ({
    title: {
        margin: 0,
        padding: theme.spacing(2)
    },
    closeBtn: {
        position: "absolute",
        right: theme.spacing(1),
        top: theme.spacing(1),
        color: theme.palette.grey[500]
    },
    form: {
        width: "100%"
    }
}))

const Select = SelectComp as React.FC<SelectProps<TagPriorty>>

export interface Props {
    state: {
        open: boolean
    }
    data: {
        value: string
        priority: TagPriorty
    }
    event: {
        onAddTag: (tag: Tag) => void
        onClose: () => void
        onValueChange: (val: string) => void
        onPriorityChange: (priority: TagPriorty) => void
    }
}

const Comp: React.FC<Props> = props => {
    const tagRes = tag(props.data.value)
    const valid = tagRes.valid
    const cls = useStyles()
    const [forceError, showValidation] = useForceError(valid)
    return (
        <Dialog
            open={props.state.open}
            onClose={props.event.onClose}
            maxWidth="xs"
            fullWidth
            >
            <DialogTitle disableTypography className={cls.title}>
                <Typography variant="h6">
                    Add tag
                </Typography>
                <IconButton onClick={props.event.onClose} className={cls.closeBtn}>
                    <CloseIcon />
                </IconButton>
            </DialogTitle>
            <DialogContent>
                <form noValidate className={cls.form}>
                    <Grid container spacing={2}>
                        <Grid item xs={8}>
                            <InputField
                                label="Tag"
                                required
                                value={props.data.value}
                                validation={tagRes}
                                forceError={forceError}
                                onValueChange={props.event.onValueChange} />
                        </Grid>
                        <Grid item xs={4}> 
                            <Select
                                fullWidth
                                label="Priority"
                                id="tag-priority-select"
                                value={props.data.priority}
                                onValueChange={props.event.onPriorityChange}>
                                <MenuItem value={TagPriorty.Normal}>
                                    Normal
                                </MenuItem>
                                <MenuItem value={TagPriorty.Secondary}>
                                    Secondary
                                </MenuItem>
                                <MenuItem value={TagPriorty.Main}>
                                    Main
                                </MenuItem>
                            </Select>
                        </Grid>
                    </Grid>
                </form>
            </DialogContent>
            <DialogActions>
                <Button color="primary" onClick={() => {
                    if(!valid) {
                        showValidation()
                    } else {
                        props.event.onAddTag({
                            value: props.data.value,
                            priority: props.data.priority
                        })
                    }
                }}>
                    Add tag
                </Button>
            </DialogActions>
        </Dialog>
    )
}

export default Comp

