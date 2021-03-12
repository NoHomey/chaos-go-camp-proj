import Box, { BoxProps as BoxProps_ } from "@material-ui/core/Box"
import Typography from "@material-ui/core/Typography"

import { makeStyles } from "@material-ui/core/styles"

export type BoxProps = Omit<BoxProps_, "display" | "alignItems">

export interface Props extends BoxProps {
    info: string
    children: React.ReactNode
}

const useStyles = makeStyles(theme => ({
    info: {
        marginRight: theme.spacing(2)
    }
}))

const InfoBox: React.FC<Props> = ({
    info,
    children,
    ...boxProps
}) => {
    const classes = useStyles()
    return (
    <Box display="flex" alignItems="center" {...boxProps}>
        <Typography variant="h6" component="span" className={classes.info}>
            {info}
        </Typography>
        {children}
    </Box>
    )
}

export default InfoBox