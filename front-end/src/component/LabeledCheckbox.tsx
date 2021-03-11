import * as React from 'react'
import FormControlLabel from "@material-ui/core/FormControlLabel"
import Checkbox from "@material-ui/core/Checkbox"

export interface Props {
    label: string
    checked: boolean
    onToggle: () => void
}

export const LabeledCheckbox: React.FC<Props> = ({
    label,
    checked,
    onToggle
}) => (
    <FormControlLabel
        label={label}
        control={
        <Checkbox
            color="primary"
            checked={checked}
            onChange={onToggle} />} />
)

const Memo = React.memo(LabeledCheckbox, (prev, next) => {
    return next.checked === prev.checked
})

export default Memo