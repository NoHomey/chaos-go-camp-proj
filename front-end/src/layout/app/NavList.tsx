import List from "@material-ui/core/List"
import ListItem from "@material-ui/core/ListItem"
import ListItemText from "@material-ui/core/ListItemText"
import Divider from "@material-ui/core/Divider"
import { createStyles, makeStyles } from "@material-ui/core/styles";
import { useHistory } from "react-router-dom";

const useStyles = makeStyles(() =>
    createStyles({
        list: {
            width: 150
        }
    }),
);

export interface Props {
    close: () => void
}

const NavList: React.FC<Props> = ({close}) => {
    const cls = useStyles()
    const history = useHistory()
    const onClick = (to: string) => () => {
        history.push(to)
        close()
    }
    return (
        <>
        <List component="nav" className={cls.list}>
            <ListItem button onClick={onClick("/")}>
                <ListItemText>Read blog</ListItemText>
            </ListItem>
            <ListItem button onClick={onClick("/")}>
                <ListItemText>Save blog</ListItemText>
            </ListItem>
        </List>
        <Divider />
        <List component="nav" className={cls.list}>
            <ListItem button onClick={onClick("/")}>
                <ListItemText>Read blogs</ListItemText>
            </ListItem>
            <ListItem button onClick={onClick("/")}>
                <ListItemText>Unread blogs</ListItemText>
            </ListItem>
            <ListItem button onClick={onClick("/")}>
                <ListItemText>All blogs</ListItemText>
            </ListItem>
        </List>
        <Divider />
        <List component="nav" className={cls.list}>
            <ListItem button onClick={onClick("/")}>
                <ListItemText>Posts</ListItemText>
            </ListItem>
        </List>
        </>
    )
}

export default NavList