import InfoBox, { BoxProps } from "./InfoBox"
import ToggleButtonGroup from "@material-ui/lab/ToggleButtonGroup"
import ToggleButton from "@material-ui/lab/ToggleButton"

export type Props = BoxProps

const Level: React.FC<Props> = props => {
    return (
    <InfoBox
        info="Level:"
        boxProps={props}>
        <ToggleButtonGroup exclusive value="advanced">
            <ToggleButton value="beginner">
                Beginner
            </ToggleButton>
            <ToggleButton value="intermediate">
                Intermediate
            </ToggleButton>
            <ToggleButton value="advanced">
                Advanced
            </ToggleButton>
            <ToggleButton value="master">
                Master
            </ToggleButton>
        </ToggleButtonGroup>
    </InfoBox>
    )
}

export default Level