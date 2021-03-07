import Tag from "../data/Tag"
import Box from "@material-ui/core/Box"
import Grid from "@material-ui/core/Grid"
import Chip from "@material-ui/core/Chip"
import Typography from "@material-ui/core/Typography"
import Button from "@material-ui/core/Button"
import { PriorityToColor } from "../muiUtil/Tag"
import AddIcon from "@material-ui/icons/Add"

import { makeStyles } from "@material-ui/core/styles"

export interface Props {
    tags: Tag[],
    marginY: number
}

const addIcon = <AddIcon />

const useStyles = makeStyles(theme => ({
    addTagBtn: {
        margin: theme.spacing(2),
    },
    tags: {
        marginBottom: theme.spacing(2),
        listStyle: "none",
        paddingInlineStart: 0
    }
}))

const Tags: React.FC<Props> = ({ tags, marginY }) => {
    const classes = useStyles()
    return (
        <>
            <Box display="flex" alignItems="center" my={marginY}>
                <Typography variant="h6" component="span">
                    Tags:
                </Typography>
                <Button
                    variant="outlined"
                    color="primary"
                    startIcon={addIcon}
                    className={classes.addTagBtn}
                >
                    Add tag
                    </Button>
            </Box>
        <Grid container spacing={2} component="ul" className={classes.tags}>
            {tags.map(({value, priority}) => (
                <Grid item component="li" key={value}>
                    <Chip
                        color={PriorityToColor(priority)}
                        label={value}
                        onDelete={() => console.log(`deleting tag: ${value}`) } />
                </Grid>
            ))}
        </Grid>
    </>
    )
}

export default Tags