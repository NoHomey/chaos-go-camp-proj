import * as React from 'react'

export default function useModel() {
    const [email, setEmail] = React.useState("")
    const [password, setPassword] = React.useState("")
    const [remember, setRemember] = React.useState(false)
    return {
        data: {
            email,
            password,
            remember
        },
        event: {
            onEmailChange: setEmail,
            onPasswordChange: setPassword,
            onRememberChange: () => setRemember(!remember)
        }
    }
}