import InfoBox, { BoxProps } from "./InfoBox"
import ToggleButtonGroup from "@material-ui/lab/ToggleButtonGroup"
import ToggleButton from "@material-ui/lab/ToggleButton"
import Level from "../data/Level"

export interface Props extends BoxProps {
    value: null | Level
    onValueChange: (val: null | Level) => void
}

const Comp: React.FC<Props> = props => {
    const { value, onValueChange, ...rest } = props
    return (
    <InfoBox info="Level:" {...rest}>
        <ToggleButtonGroup
            exclusive
            value={value}
            onChange={(e, val) => onValueChange(val)}>
            <ToggleButton value={Level.Beginner}>
                Beginner
            </ToggleButton>
            <ToggleButton value={Level.Intermediate}>
                Intermediate
            </ToggleButton>
            <ToggleButton value={Level.Advanced}>
                Advanced
            </ToggleButton>
            <ToggleButton value={Level.Master}>
                Master
            </ToggleButton>
        </ToggleButtonGroup>
    </InfoBox>
    )
}

export default Comp