import * as React from 'react'

export default function useForceError(valid: boolean): [boolean, () => void] {
    const [force, setForce] = React.useState(false)
    if(force && valid) {
        setForce(false)
    }
    return [force, () => setForce(true)]
}