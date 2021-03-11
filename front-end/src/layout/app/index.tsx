import * as React from "react"
import AppBar from "./Bar"
import NavList from "./NavList"
import Drawer from "@material-ui/core/Drawer"
import { createStyles, makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles(() =>
    createStyles({
        main: {
            width: "100%"
        }
    }),
);

export interface Props {
    onSignOut: () => void
}

const Layout: React.FC<Props> = ({ onSignOut, children }) => {
    const [open, setOpen] = React.useState(false)
    const classes = useStyles()
    const close = () => setOpen(false)
    return (
        <>
            <AppBar
                username="Ivo Statev"
                onMenuOpen={() => setOpen(true)}
                onSignOut={onSignOut} />
            <main className={classes.main}>
                {children}
            </main>
            <Drawer open={open} onClose={onDrawerClose(close)}>
                <NavList close={close} />
            </Drawer>
        </>
    )
}

export default Layout

const onDrawerClose = (cb: () => void) => (ev: React.KeyboardEvent | React.MouseEvent) => {
    if(
        ev.type === "keydown"
        &&
        (
            (ev as React.KeyboardEvent).key === "Tab"
            ||
            (ev as React.KeyboardEvent).key === "Shift"
        )
    ) {
        return;
    }
    cb()
}