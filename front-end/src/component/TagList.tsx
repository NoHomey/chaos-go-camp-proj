import Tag from "../data/Tag"
import Grid from "@material-ui/core/Grid"
import Chip from "@material-ui/core/Chip"
import Typography from "@material-ui/core/Typography"
import { PriorityToColor } from "../muiUtil/Tag"
import clsx from "clsx"

import { makeStyles } from "@material-ui/core/styles"

export interface Props {
    tags: Tag[],
    className?: string
}

const useStyles = makeStyles(theme => ({
    list: {
        marginBottom: theme.spacing(2),
        listStyle: "none",
        paddingInlineStart: 0
    }
}))

const TagList: React.FC<Props> = ({ tags, className }) => {
    const classes = useStyles()
    return (
    <Grid container spacing={2} component="ul" className={clsx([classes.list, className])} >
        <Grid item component="li">
            <Typography variant="h6" component="span">
                Tags:
            </Typography>
        </Grid>
        {tags.map(({value, priority}) => (
            <Grid item component="li" key={value}>
                <Chip color={PriorityToColor(priority)} label={value} />
            </Grid>
        ))}
    </Grid>
    )
}

export default TagList