import * as React from "react"
import debounce from 'lodash.debounce'

enum ValidationState { Init, Show, Hide }

const waitTime = 700

export default function useShowValidation(forceError: boolean) {
    const [show, setShow] = React.useState(ValidationState.Init)
    const showValidation = () => setShow(ValidationState.Show)
    const debounceValidation = React.useCallback(
        debounce(showValidation, waitTime),
        []
    )
    const showError =
        show === ValidationState.Show
        ||
        (forceError && (show !== ValidationState.Hide))
    return {
        showError,
        debounceValidation,
        showValidation,
        hideValidation: () => setShow(ValidationState.Hide)
    }
}