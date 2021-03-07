import TextField, { TextFieldProps } from "@material-ui/core/TextField"

export type Props = Omit<TextFieldProps, 'variant' | 'margin' | 'fullWidth'>

const InputField: React.FC<Props> = props => (
    <TextField
        variant="outlined"
        margin="normal"
        fullWidth
        {...props} />
)

export default InputField