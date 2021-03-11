import * as React from "react"
import { Provider, State as UserState } from "../context/User"
import { User } from "../service/User"

interface State {
    user: null | User
    state: UserState
}

const UserProvider: React.FC<{}> = ({ children }) => {
    const [state, setState] = React.useState<State>({
        user: null,
        state: UserState.Init
    })
    return (
        <Provider value={{
            user: state.user,
            state: state.state,
            setUser: usr => {
                if(usr === null) {
                    setState({
                        user: null,
                        state: UserState.Sign
                    })
                } else {
                    setState({
                        user: usr,
                        state: UserState.User
                    })
                }
            },
            setStateToSign: () => {
                setState({
                    user: state.user,
                    state: UserState.Sign
                })
            }
        }}>
            {children}
        </Provider>
    )
}

export default UserProvider