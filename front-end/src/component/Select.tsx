import * as React from "react"
import FormControl, { FormControlProps } from "@material-ui/core/FormControl"
import InputLabel from "@material-ui/core/InputLabel"
import Select from "@material-ui/core/Select"

export interface Props<V> extends FormControlProps {
    value: V
    onValueChange: (val: V) => void
    label: string
    required?: boolean
    children: React.ReactNode
    id: string
}

function Comp<V>(props: Props<V>) {
    const {
        value,
        onValueChange,
        label,
        required,
        children,
        id,
        ...rest
    } = props
    const labelID = `${id}-label`
    return (
        <FormControl {...rest}>
            <InputLabel color="primary" id={labelID}>
                {label}
            </InputLabel>
            <Select
                fullWidth
                required={required}
                labelId={labelID}
                id={id}
                value={value}
                onChange={e => onValueChange(e.target.value as V)}>
                {children}
            </Select>
        </FormControl>
    )
}

export const Memo = React.memo(Comp, (prev, next) => {
    return next.value === prev.value
})

export default Memo