import * as React from "react"
import createUserService, { Service as UserService } from "../service/User" 
import createReqService, { Service as ReqService } from "../service/Request"
import createFeedService, { Service as FeedService } from "../service/Feed"

export interface Service {
    user: UserService
    request: ReqService
    feed: FeedService
}

const userService = createUserService()
const reqService = createReqService(userService)

export const init: Service = {
    user: userService,
    request: reqService,
    feed: createFeedService(reqService)
}

const Ctx = React.createContext(init)

export const Provider = Ctx.Provider

export function useAll() {
    return React.useContext(Ctx)
}

export function useService<K extends keyof Service>(name: K): Service[K] {
    const ctx = useAll()
    return ctx[name]
}

export default Provider