import Box, { BoxProps } from "@material-ui/core/Box"
import Typography from "@material-ui/core/Typography"
import RatingInput from "@material-ui/lab/Rating"


import { makeStyles } from "@material-ui/core/styles"

export type Props = Omit<BoxProps, "display" | "alignItems">

const useStyles = makeStyles(theme => ({
    rating: {
        marginLeft: theme.spacing(1)
    }
}))

const Rating: React.FC<Props> = props => {
    const classes = useStyles()
    return (
    <Box display="flex" alignItems="center" {...props}>
        <Typography variant="h6" component="span">
            Rating:
        </Typography>
        <RatingInput
            defaultValue={0}
            max={14}
            size="large"
            className={classes.rating} />
    </Box>
    )
}

export default Rating