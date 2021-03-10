export type RespError = {
    message: string
    name: string
    data: any
}

export type CallBack<T> = (val: T) => void

export interface Response<T> {
    OnResult(cb: CallBack<T>): Response<T>
    OnError(cb: CallBack<RespError>): Response<T>
    OnFail(cb: CallBack<Error>): Response<T>
    Handle(): void
}

export type Wrapped = Promise<globalThis.Response>

export type Consumeer<T> = (r: any) => T

export class Wrapper<T> implements Response<T> {
    private wrapped: Wrapped
    private onResult: null | CallBack<T>
    private onError: null | CallBack<RespError>
    private onFail: null | CallBack<Error>
    private consume: Consumeer<T>

    constructor(wrapped: Wrapped, consume: Consumeer<T>) {
        this.wrapped = wrapped
        this.onResult = null
        this.onError = null
        this.onFail = null
        this.consume = consume
    }

    public OnResult(cb: CallBack<T>): Response<T> {
        this.onResult = cb
        return this
    }

    public OnError(cb: CallBack<RespError>): Response<T> {
        this.onError = cb
        return this
    }

    public OnFail(cb: CallBack<Error>): Response<T> {
        this.onFail = cb
        return this
    }

    public Handle(): void {
        this.wrapped
            .then(resp => resp.json())
            .then(data => {
                if(typeof data.error === 'object') {
                    this.onError!(data.error as RespError)
                } else {
                    this.onResult!(this.consume(data.result!))
                }
            }).catch(this.onFail!)
    }
}

export function Wrap<T>(wrapped: Wrapped): Response<T> {
    return new Wrapper<T>(wrapped, cast)
}

export function Consume<R>(wrapped: Wrapped, consume: Consumeer<R>): Response<R> {
    return new Wrapper(wrapped, consume)
}

export function Any(wrapped: Wrapped): Response<any> {
    return new Wrapper<any>(wrapped, cast)
}

export function Make(wrapped: Wrapped): Response<void> {
    return new Wrapper<void>(wrapped, cast)
}

function cast<T>(res: any): T {
    return res
}

export default Wrap