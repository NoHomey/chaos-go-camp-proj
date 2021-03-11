import * as React from "react"

export interface Value {
    show: (node: React.ReactNode) => void
    showFail: () => void
    showError: (node: React.ReactNode, onOK: () => void) => void
    showResult: (node: React.ReactNode, onOK: () => void) => void
    close: () => void
}

const Ctx = React.createContext<Value>({
    show: error,
    showFail: error,
    showError: error,
    showResult: error,
    close: error
})

export const Provider = Ctx.Provider

export function useReqDialog() {
    return React.useContext(Ctx)
}

export default Provider

function error() {
    throw(new Error("ReqDialog Context with invalid value"))
}