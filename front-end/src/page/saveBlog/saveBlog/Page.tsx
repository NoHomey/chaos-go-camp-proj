import Container from "@material-ui/core/Container"
import Paper from "@material-ui/core/Paper"
import BlogForm from "../../../component/BlogForm"
import Blog from "../../../data/Blog";
import { makeStyles } from "@material-ui/core/styles"

const useStyles = makeStyles(theme => ({
    paper: {
        width: "100%",
        marginTop: theme.spacing(6),
        marginBottom: theme.spacing(10),
        padding: theme.spacing(3.5, 5, 3.5, 5)
    }
}))

export interface Props {
    blog: Blog
    onSaveBlog: (blog: Blog) => void
}

const Page: React.FC<Props> = ({
    blog,
    onSaveBlog
}) => {
    const classes = useStyles()
    return (
        <Container component="main" maxWidth="md">
            <Paper elevation={9} className={classes.paper}>
                <BlogForm
                    actionBtnLabel="Save blog"
                    onAction={onSaveBlog}
                    blog={blog}/>
            </Paper>
        </Container>
    )
}

export default Page