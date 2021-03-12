import * as React from "react"
import debounce from 'lodash.debounce'

enum ValidationState { Init, Show, Hide }

const waitTime = 700

export default function useShowValidation(forceError: boolean) {
    const [showValidation, setShowValidation] = React.useState(ValidationState.Init)
    const debounceValidation = React.useCallback(
        debounce(() => setShowValidation(ValidationState.Show), waitTime),
        []
    )
    const showError =
        showValidation === ValidationState.Show
        ||
        (forceError && (showValidation !== ValidationState.Hide))
    return {
        showError,
        debounceValidation,
        hideValidation: () => setShowValidation(ValidationState.Hide)
    }
}