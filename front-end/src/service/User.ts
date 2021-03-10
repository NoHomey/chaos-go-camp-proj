import { Response, Consume, Make } from "../response";

export interface Service {
    Reload(): null | Response<User>
    SignUp(data: SignUpData): Response<void>
    SignIn(data: SignInData): Response<User>
    SignOut(): Response<void>
    Access(): Response<User>
}

export type SignUpData = {
    name: string
    email: string
    password: string
}

export type SignInData = {
    email: string
    password: string
}

export type User = {
    name: string
    email: string
}

class Impl implements Service {
    private refreshToken: null | string
    private accessSyncToken: null | string

    public constructor() {
        this.refreshToken = null
        this.accessSyncToken = null
    }

    public Reload(): null | Response<User> {
        this.refreshToken = localStorage.getItem(refreshTokenKey)
        if(!this.refreshToken) {
            return null
        }
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
                    body: JSON.stringify(data)
                }
            ),
            res => {
                this.accessSyncToken = res.accessSyncToken
                this.refreshToken = res.refreshToken
                localStorage.setItem(refreshTokenKey, this.refreshToken!)
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
                this.refreshToken = null
                this.accessSyncToken = null
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

    private consume(res: any): User {
        this.accessSyncToken = res.accessSyncToken
        return {
            name: res.name!,
            email: res.email!
        }
    }

    private authHeader(): string {
        return `PASETO ${this.refreshToken}`
    }
}

export default Impl

const url = {
    signUp: "/user/sign-up",
    signIn: "user/sign-in",
    signOut: "/user/sign-out",
    access: "/user/access"
}

const refreshTokenKey = "$__refresh-toke__$"