import { Response, Wrap } from "../response";
import { Service as UserService } from "./User"

export interface Service {
    Get<T>(url: string): Response<T>
    Post<T>(url: string, body: Object): Response<T>  
}

class Impl implements Service {
    private userService: UserService

    public constructor(userService: UserService) {
        this.userService = userService
    }

    public Get<T>(url: string): Response<T> {
        const headers = {}
        this.userService.AugmentHeaders(headers)
        return Wrap<T>(
            fetch(
                url,
                {
                    cache: "no-cache",
                    headers: headers
                }
            )
        )
    }

    public Post<T>(url: string, body: Object): Response<T> {
        const headers = { 'Content-Type': 'application/json' }
        this.userService.AugmentHeaders(headers)
        return Wrap<T>(
            fetch(
                url,
                {
                    method: "POST",
                    cache: "no-cache",
                    headers,
                    body: JSON.stringify(body)
                }
            )
        )
    }
}

export default function create(userService: UserService): Service {
    return new Impl(userService)
}