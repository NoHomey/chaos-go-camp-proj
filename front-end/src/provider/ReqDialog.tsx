import * as React from "react"
import { Provider, Value } from "../context/ReqDialog"
import Dialog from "@material-ui/core/Dialog"
import DialogActions from "@material-ui/core/DialogActions"
import DialogContent from "@material-ui/core/DialogContent"
import Button from "@material-ui/core/Button"

enum Kind { Info, Warn, Err, Succ }

export interface State {
    show: React.ReactNode
    kind: Kind
    onOK: null | (() => void)
}

const ReqDialog: React.FC<{}> = ({ children }) => {
    const initial: State = { show: null, kind: Kind.Info, onOK: null }
    const [state, setState] = React.useState<State>(initial)
    const close = () => setState(initial)
    const value: Value = {
        show: node => setState({ show: node, kind: Kind.Info, onOK: null }),
        showFail: () => setState({ show: failText, kind: Kind.Warn, onOK: close }),
        showError: (node, onOK) => setState({ show: node, kind: Kind.Err, onOK: onOK }),
        showResult: (node, onOK) => setState({ show: node, kind: Kind.Succ, onOK: onOK }),
        close
    }
    
    return (
        <Provider value={value}>
            {children}
            <Dialog open={!!state.show} disableEscapeKeyDown maxWidth="sm" fullWidth>
                <DialogContent>
                    {state.show}
                </DialogContent>
                {state.onOK !== null &&
                <DialogActions>
                    <Button color="primary" onClick={state.onOK}>
                        OK
                    </Button>
                </DialogActions>}
            </Dialog>
        </Provider>
    )
}

export default ReqDialog

const failText = "Could not make request to the server. Please ensure you have a stable internet connection"