import Container from "@material-ui/core/Container"
import Paper from "@material-ui/core/Paper"
import InputField from "../../component/InputField"
import Rating from "@material-ui/lab/Rating"
import Button from "@material-ui/core/Button"
import Box from "@material-ui/core/Box"
import Grid from "@material-ui/core/Grid"
import Typography from "@material-ui/core/Typography"
import Chip from "@material-ui/core/Chip"
import AddIcon from "@material-ui/icons/Add"
import TextField from "@material-ui/core/TextField"
import Accordion from "@material-ui/core/Accordion"
import AccordionSummary from "@material-ui/core/AccordionSummary"
import AccordionDetails from "@material-ui/core/AccordionDetails"
import ExpandMoreIcon from "@material-ui/icons/ExpandMore"
import SaveIcon from "@material-ui/icons/Save"

import { makeStyles } from "@material-ui/core/styles"

const addIcon = <AddIcon />
const expandMoreIcon = <ExpandMoreIcon />
const saveIcon = <SaveIcon />

const useStyles = makeStyles(theme => ({
    form: {
        width: "100%"
    },
    paper: {
        width: "100%",
        marginTop: theme.spacing(6),
        marginBottom: theme.spacing(10),
        padding: theme.spacing(3)
    },
    importBtn: {
        marginTop: theme.spacing(4)
    },
    inline: {
        margin: theme.spacing(2, 0, 2, 0)
    },
    rating: {
        marginLeft: theme.spacing(1)
    },
    tags: {
        marginBottom: theme.spacing(2),
        listStyle: "none",
        paddingInlineStart: 0
    },
    addTagBtn: {
        margin: theme.spacing(2),
    }
}))

const shrinkLabel = { shrink: true }

const Page: React.FC<{}> = () => {
    const classes = useStyles()
    return (
        <Container component="main" maxWidth="sm">
            <Paper elevation={9} className={classes.paper}>
                <form noValidate className={classes.form}>
                    <InputField InputLabelProps={shrinkLabel} label="Feed URL" value="http://some-url.com" />
                    <InputField InputLabelProps={shrinkLabel} label="Author" value="Some Author" />
                    <InputField InputLabelProps={shrinkLabel} label="Blog title" value="Some Blog Title" />
                    <Accordion className={classes.inline}>
                        <AccordionSummary expandIcon={expandMoreIcon}>
                            <Typography variant="body1" component="span">
                                Blog description from the author
                            </Typography>
                        </AccordionSummary>
                        <AccordionDetails>
                            Some text
                        </AccordionDetails>
                    </Accordion>
                    <Box display="flex" alignItems="center" my={2}>
                        <Typography variant="h6" component="span">
                            Rating:
                        </Typography>
                        <Rating defaultValue={0} max={14} size="large" className={classes.rating} />
                    </Box>
                    <Box display="flex" alignItems="center" my={2}>
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
                        <Grid item component="li">
                            <Chip color="primary" label="go-lang" onDelete={() => { }} />
                        </Grid>
                        <Grid item component="li">
                            <Chip color="primary" label="go-lang" onDelete={() => { }} />
                        </Grid>
                        <Grid item component="li">
                            <Chip color="primary" label="go-blog" onDelete={() => { }} />
                        </Grid>
                        <Grid item component="li">
                            <Chip color="primary" label="videos" onDelete={() => { }} />
                        </Grid>
                        <Grid item component="li">
                            <Chip color="primary" label="go-confs" onDelete={() => { }} />
                        </Grid>
                        <Grid item component="li">
                            <Chip color="secondary" label="go-lang" onDelete={() => { }} />
                        </Grid>
                        <Grid item component="li">
                            <Chip color="secondary" label="go-lang" onDelete={() => { }} />
                        </Grid>
                        <Grid item component="li">
                            <Chip color="secondary" label="go-blog" onDelete={() => { }} />
                        </Grid>
                        <Grid item component="li">
                            <Chip color="secondary" label="videos" onDelete={() => { }} />
                        </Grid>
                        <Grid item component="li">
                            <Chip color="default" label="go-confs" onDelete={() => { }} />
                        </Grid>
                        <Grid item component="li">
                            <Chip color="default" label="go-lang" onDelete={() => { }} />
                        </Grid>
                        <Grid item component="li">
                            <Chip color="default" label="go-lang" onDelete={() => { }} />
                        </Grid>
                        <Grid item component="li">
                            <Chip color="default" label="go-blog" onDelete={() => { }} />
                        </Grid>
                        <Grid item component="li">
                            <Chip color="default" label="videos" onDelete={() => { }} />
                        </Grid>
                        <Grid item component="li">
                            <Chip color="default" label="go-confs" onDelete={() => { }} />
                        </Grid>
                    </Grid>
                    <TextField
                        fullWidth
                        variant="outlined"
                        margin="normal"
                        label="Quick note"
                        multiline
                        rows={3}
                        rowsMax={12}
                    />
                    <Button
                        variant="outlined"
                        color="primary"
                        endIcon={saveIcon}
                        className={classes.importBtn}
                    >
                        Save Blog
                    </Button>
                </form>
            </Paper>
        </Container>
    )
}

export default Page