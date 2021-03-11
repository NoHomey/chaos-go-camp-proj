import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        menuButton: {
            marginRight: theme.spacing(2),
        },
        title: {
            flexGrow: 1,
        },
    }),
);

export interface Props {
    username: string
    onMenuOpen: () => void
    onSignOut: () => void
}

const Bar: React.FC<Props> = ({
    username,
    onMenuOpen,
    onSignOut
}) => {
    const classes = useStyles();
    return (
        <AppBar position="static">
            <Toolbar>
                <IconButton
                    edge="start"
                    className={classes.menuButton}
                    color="inherit"
                    onClick={onMenuOpen}>
                    <MenuIcon />
                </IconButton>
                <Typography variant="h6" className={classes.title}>
                    {"Happy reading " + username}
                </Typography>
                <Button color="secondary" variant="outlined" onClick={onSignOut}>
                    Sign out
                </Button>
            </Toolbar>
        </AppBar>
    );
}

export default Bar