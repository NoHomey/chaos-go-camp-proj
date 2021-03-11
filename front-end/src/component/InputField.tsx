import * as React from 'react'
import TextField, { TextFieldProps } from "@material-ui/core/TextField"
import { Result, errorMsg } from '../validation/Result'
import debounce from 'lodash.debounce'

type PickProps
= 'value'
| 'label'
| 'type'
| 'required'
| 'autoComplete'

type InheritProps = Pick<TextFieldProps, PickProps>

export interface Props extends InheritProps {
    validation: Result
    onValueChange: (val: string) => void
}

const waitTime = 700

export const InputField: React.FC<Props> = props => {
    const { validation, onValueChange, ...rest } = props
    const [showValidation, setShowValidation] = React.useState(false)
    const debounceValidation = React.useCallback(
        debounce(() => setShowValidation(true), waitTime),
        []
    )
    return (
        <TextField
            {...rest}
            variant="outlined"
            margin="normal"
            fullWidth
            onChange={e => {
                setShowValidation(false)
                onValueChange(e.target.value)
                debounceValidation()
            }}
            error={showValidation && !validation.valid}
            helperText={showValidation && errorMsg(validation)} />
    )
}

export const MemoInput = React.memo(InputField, (prev, next) => {
    return next.value === prev.value
})

export default MemoInput