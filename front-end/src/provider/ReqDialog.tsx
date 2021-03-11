import * as React from "react"
import { Provider, Value } from "../context/ReqDialog"
import Dialog from "@material-ui/core/Dialog"
import DialogActions from "@material-ui/core/DialogActions"
import DialogContent from "@material-ui/core/DialogContent"
import Button from "@material-ui/core/Button"
import CircularProgress from "@material-ui/core/CircularProgress"
import Box from "@material-ui/core/Box"

enum Kind { Info, Warn, Err, Succ }

export interface State {
    show: React.ReactNode
    kind: Kind
    loading: boolean
    onOK: null | (() => void)
}

const ReqDialog: React.FC<{}> = ({ children }) => {
    const initial: State = {
        show: null,
        kind: Kind.Info,
        loading: false,
        onOK: null
    }
    const [state, setState] = React.useState<State>(initial)
    const close = () => setState(initial)
    const value: Value = {
        showLoading: text  => setState({
            show: text,
            kind: Kind.Info,
            loading: true,
            onOK: null
        }),
        showFail: (action?: () => void) => setState({
            show: failText,
            kind: Kind.Warn,
            loading: false,
            onOK: action ?? close
        }),
        showError: (node, onOK) => setState({
            show: node,
            kind: Kind.Err,
            loading: false,
            onOK: onOK
        }),
        showResult: (node, onOK) => setState({
            show: node,
            kind: Kind.Succ,
            loading: false,
            onOK: onOK
        }),
        close
    }
    const isRetry = state.kind === Kind.Warn && state.onOK !== close
    return (
        <Provider value={value}>
            {children}
            <Dialog open={!!state.show} disableEscapeKeyDown maxWidth="sm" fullWidth>
                <DialogContent>
                    {state.loading
                    ? <Box width={1} component="span">
                        <Box
                            mr={5}
                            component="span"
                            fontSize="h6.fontSize"
                            fontWeight="fontWeightMedium"
                            color="text.secondary">
                            {state.show}
                        </Box>
                        <CircularProgress size={42} thickness={4}/>
                    </Box>
                    : <Content kind={state.kind}>
                        {state.show}
                    </Content>}
                </DialogContent>
                {state.onOK !== null &&
                <DialogActions>
                    <Button color="primary" onClick={state.onOK}>
                        {isRetry ? "Retry" : "OK"}
                    </Button>
                </DialogActions>}
            </Dialog>
        </Provider>
    )
}

export default ReqDialog

const Content: React.FC<{ kind: Kind }> = ({ kind, children }) => {
    const isText = typeof children === "string"
    return isText ?
        (<Box
            component="span"
            fontSize="h6.fontSize"
            fontWeight="fontWeightMedium"
            color={color(kind)}>
            {children}
        </Box>
        ) : <>{children}</>
}

function color(kind: Kind): string {
    switch(kind) {
    case Kind.Info:
        return "text.primary"
    case Kind.Warn:
        return "warning.main"
    case Kind.Err:
        return "error.main"
    case Kind.Succ:
        return "success.main"
    }
}
const failText = "Could not make request to the server. Please ensure you have a stable internet connection"