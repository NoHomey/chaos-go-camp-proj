import InfoBox, { BoxProps } from "./InfoBox"
import RatingInput from "@material-ui/lab/Rating"

export type Props = BoxProps

const Rating: React.FC<Props> = props => {
    return (
    <InfoBox
        info="Rating:"
        boxProps={props}>
        <RatingInput
            defaultValue={0}
            max={15}
            size="large" />
    </InfoBox>
    )
}

export default Rating