import InfoBox, { BoxProps } from "./InfoBox"
import RatingInput from "@material-ui/lab/Rating"

export interface Props extends BoxProps {
    value: null | number
    onValueChange: (val: null | number) => void
}

const Rating: React.FC<Props> = props => {
    const { value, onValueChange, ...rest } = props
    return (
    <InfoBox info="Rating:" {...rest}>
        <RatingInput
            name="rating"
            value={value}
            onChange={(e, val) => onValueChange(val!)}
            max={15}
            size="large" />
    </InfoBox>
    )
}

export default Rating