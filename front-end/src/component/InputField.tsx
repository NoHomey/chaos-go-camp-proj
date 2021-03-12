import * as React from 'react'
import TextField, { TextFieldProps } from "@material-ui/core/TextField"
import { Result, errorMsg } from '../validation/Result'
import useShowValidation from "../hook/useShowValidation"

type PickProps
= 'value'
| 'label'
| 'type'
| 'required'
| 'autoComplete'

type InheritProps = Pick<TextFieldProps, PickProps>

export interface Props extends InheritProps {
    validation: Result
    forceError: boolean
    onValueChange: (val: string) => void
}

export const InputField: React.FC<Props> = props => {
    const { validation, onValueChange, forceError, ...rest } = props
    const {
        showError,
        debounceValidation,
        hideValidation,
    } = useShowValidation(forceError)
    return (
        <TextField
            {...rest}
            variant="outlined"
            margin="normal"
            fullWidth
            onChange={e => {
                hideValidation()
                onValueChange(e.target.value)
                debounceValidation()
            }}
            onFocus={hideValidation}
            error={showError && !validation.valid}
            helperText={showError && errorMsg(validation)} />
    )
}

export const MemoInput = React.memo(InputField, (prev, next) => {
    return next.value === prev.value && prev.forceError === next.forceError
})

export default MemoInput