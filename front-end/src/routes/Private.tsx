import { Switch, Route, Redirect } from "react-router-dom"
import Page from "../page/saveBlog"
import routes from "./map"

export default function Routes() {
    return (
        <Switch>
            <Route exact path={routes.saveBlog} component={Page}/>
            <Route exact path="/">
                <Redirect to={routes.saveBlog} />
            </Route>
        </Switch>
    )
}