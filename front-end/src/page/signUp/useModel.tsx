import * as React from 'react'

export default function useModel() {
    const [email, setEmail] = React.useState("")
    const [password, setPassword] = React.useState("")
    const [confirmPassword, setConfirmPassword] = React.useState("")
    return {
        data: {
            email,
            password,
            confirmPassword
        },
        event: {
            onEmailChange: setEmail,
            onPasswordChange: setPassword,
            onConfirmPasswordChange: setConfirmPassword
        }
    }
}