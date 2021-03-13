import { Response, Consume, Make } from "../response";

type Hs = Headers | Record<string, string>

export interface Service {
    Reload(): null | Response<User>
    SignUp(data: SignUpData): Response<void>
    SignIn(data: SignInData): Response<User>
    SignOut(): Response<void>
    Access(): Response<User>
    AugmentHeaders(headers: Hs): Hs
}

export type SignUpData = {
    name: string
    email: string
    password: string
}

export type SignInData = {
    email: string
    password: string
    remember: boolean
}

export type User = {
    name: string
    email: string
}

class Impl implements Service {
    private refreshToken: string
    private accessSyncToken: string
    private timeout: null | number

    public constructor() {
        this.refreshToken = ""
        this.accessSyncToken = ""
        this.timeout = null
    }

    public Reload(): null | Response<User> {
        const refreshToken = localStorage.getItem(refreshTokenKey)
        if(!refreshToken) {
            return null
        }
        this.refreshToken = refreshToken
        return this.Access()
    }

    public SignUp(data: SignUpData): Response<void> {
        return Make(
            fetch(
                url.signUp,
                {
                    method: "POST",
                    cache: "no-cache",
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(data)
                }
            )
        )
    }

    public SignIn(data: SignInData): Response<User> {
        return Consume<User>(
            fetch(
                url.signIn,
                {
                    method: "POST",
                    cache: "no-cache",
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        email: data.email,
                        password: data.password,
                    })
                }
            ),
            res => {
                this.accessSyncToken = res.accessSyncToken
                this.refreshToken = res.refreshToken
                this.refresh(res.accessDuration)
                if(data.remember) {
                    localStorage.setItem(refreshTokenKey, this.refreshToken!)
                }
                return {
                    name: res.name!,
                    email: res.email!
                }
            }
        )
    }

    public SignOut(): Response<void> {
        return Consume<void>(
            fetch(
                url.signOut,
                {
                    method: "POST",
                    cache: "no-cache",
                    headers: {
                        "Authorization": this.authHeader()
                    }
                }
            ),
            () => {
                this.refreshToken = ""
                this.accessSyncToken = ""
                localStorage.removeItem(refreshTokenKey)
            }
        )
    }

    public Access(): Response<User> {
        return Consume<User>(
            fetch(
                url.access,
                {
                    cache: "no-cache",
                    headers: {
                        "Authorization": this.authHeader()
                    }
                }
            ),
            this.consume.bind(this)
        )
    }

    public AugmentHeaders(headers: Hs): Hs {
        if(headers instanceof Headers) {
            headers.set(accessSyncTokenHeader, this.accessSyncToken)
        } else {
            headers[accessSyncTokenHeader] = this.accessSyncToken
        }
        return headers
    }

    private refresh(duration: number) {
        const r = Math.random()
        const time = Math.floor(0.3 * (1 + r) / 2 * duration)
        const timeout = setTimeout(() => {
            this.Access()
                .OnResult(() => console.log("refreshing access"))
                .OnFail(() => this.refresh(retryTime))
                .OnError(err => {
                    //sign out
                    console.log(err)
                })
                .Handle()
        }, time)
    }

    private consume(res: any): User {
        this.accessSyncToken = res.accessSyncToken
        this.refresh(res.accessDuration)
        return {
            name: res.name!,
            email: res.email!
        }
    }

    private authHeader(): string {
        return `PASETO ${this.refreshToken}`
    }
}

export default function create(): Service {
    return new Impl()
}

const url = {
    signUp: "/user/sign-up",
    signIn: "user/sign-in",
    signOut: "/user/sign-out",
    access: "/user/access"
}

const refreshTokenKey = "$__refresh-toke__$"

const accessSyncTokenHeader = "X-Access-Sync-Token"

const retryTime = 500