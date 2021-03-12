import Tag from "../../data/Tag"
import Box from "@material-ui/core/Box"
import Grid from "@material-ui/core/Grid"
import Chip from "@material-ui/core/Chip"
import Typography from "@material-ui/core/Typography"
import Button from "@material-ui/core/Button"
import { PriorityToColor } from "../../muiUtil/Tag"
import AddIcon from "@material-ui/icons/Add"

import { makeStyles } from "@material-ui/core/styles"

export interface Props {
    marginY: number,
    tags: Tag[],
    onAddTag: () => void
    onRemoveTag: (idx: number) => void 
}

const addIcon = <AddIcon />

const useStyles = makeStyles(theme => ({
    addTagBtn: (props: Props) => ({
        margin: theme.spacing(props.marginY),
    }),
    tags: (props: Props) => ({
        marginBottom: theme.spacing(props.marginY),
        listStyle: "none",
        paddingInlineStart: 0
    })
}))

const Tags: React.FC<Props> = props => {
    const classes = useStyles(props)
    return (
        <>
            <Box display="flex" alignItems="center" my={props.marginY}>
                <Typography variant="h6" component="span">
                    Tags:
                </Typography>
                <Button
                    variant="outlined"
                    color="primary"
                    startIcon={addIcon}
                    className={classes.addTagBtn}
                    onClick={props.onAddTag}
                >
                    Add tag
                    </Button>
            </Box>
        <Grid container spacing={2} component="ul" className={classes.tags}>
            {props.tags.map(({value, priority}, idx) => (
                <Grid item component="li" key={value}>
                    <Chip
                        color={PriorityToColor(priority)}
                        label={value}
                        onDelete={() => props.onRemoveTag(idx)} />
                </Grid>
            ))}
        </Grid>
    </>
    )
}

export default Tags