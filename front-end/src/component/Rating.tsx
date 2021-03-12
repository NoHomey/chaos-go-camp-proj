import InfoBox, { BoxProps } from "./InfoBox"
import RatingInput from "@material-ui/lab/Rating"

export interface Props extends BoxProps {
    value: number
    onValueChange: (val: number) => void
}

const Rating: React.FC<Props> = props => {
    const { value, onValueChange, ...rest } = props
    return (
    <InfoBox
        info="Rating:"
        boxProps={rest}>
        <RatingInput
            value={value}
            onChange={(e, val) => onValueChange(val!)}
            max={15}
            size="large" />
    </InfoBox>
    )
}

export default Rating