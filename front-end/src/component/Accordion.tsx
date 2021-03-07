import MuiAccordion from "@material-ui/core/Accordion"
import AccordionSummary from "@material-ui/core/AccordionSummary"
import AccordionDetails from "@material-ui/core/AccordionDetails"
import Typography from "@material-ui/core/Typography"
import ExpandMoreIcon from "@material-ui/icons/ExpandMore"

export interface Props {
    title: string
    body: string
    className?: string
}

const expandMoreIcon = <ExpandMoreIcon />

const Accordion: React.FC<Props> = ({ title, body, className }) => (
    <MuiAccordion className={className}>
        <AccordionSummary expandIcon={expandMoreIcon}>
            <Typography variant="body1" component="span">
                {title}
            </Typography>
        </AccordionSummary>
        <AccordionDetails>
            {body}
        </AccordionDetails>
    </MuiAccordion>
)

export default Accordion