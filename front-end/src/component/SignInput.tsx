import TextField, { TextFieldProps } from "@material-ui/core/TextField"

export type Props = Omit<TextFieldProps, 'variant' | 'margin' | 'fullWidth'>

const SignInput: React.FC<Props> = props => (
    <TextField
        variant="outlined"
        margin="normal"
        fullWidth
        {...props} />
)

export default SignInput