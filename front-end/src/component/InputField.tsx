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

enum ValidationState { Init, Show, Hide }

export interface Props extends InheritProps {
    validation: Result
    forceError: boolean
    onValueChange: (val: string) => void
}

const waitTime = 700

export const InputField: React.FC<Props> = props => {
    const { validation, onValueChange, forceError, ...rest } = props
    const [showValidation, setShowValidation] = React.useState(ValidationState.Init)
    const debounceValidation = React.useCallback(
        debounce(() => setShowValidation(ValidationState.Show), waitTime),
        []
    )
    const showError =
        showValidation === ValidationState.Show
        ||
        (forceError && (showValidation !== ValidationState.Hide))
    return (
        <TextField
            {...rest}
            variant="outlined"
            margin="normal"
            fullWidth
            onChange={e => {
                setShowValidation(ValidationState.Hide)
                onValueChange(e.target.value)
                debounceValidation()
            }}
            onFocus={() => setShowValidation(ValidationState.Hide)}
            error={showError && !validation.valid}
            helperText={showError && errorMsg(validation)} />
    )
}

export const MemoInput = React.memo(InputField, (prev, next) => {
    return next.value === prev.value && prev.forceError === next.forceError
})

export default MemoInput