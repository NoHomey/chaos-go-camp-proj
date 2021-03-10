import { Switch, Route, Redirect } from 'react-router-dom'
import SignIn from '../page/signIn'
import SignUp from '../page/signUp'
import routes from "./map"

export default function Routes() {
    return (
        <Switch>
            <Route exact path={routes.signIn} component={SignIn}/>
            <Route exact path={routes.signUp} component={SignUp}/>
            <Route exact path="/">
                <Redirect to={routes.signIn} />
            </Route>
        </Switch>
    )
}