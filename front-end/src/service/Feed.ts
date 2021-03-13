import { Response } from "../response";
import { Service as ReqService } from "./Request"

export type Details = {
    title: string
    author: string
    description: string
}

export interface Service {
    Details(feedURL: string): Response<Details>
}

class Impl implements Service {
    private reqService: ReqService

    public constructor(reqService: ReqService) {
        this.reqService = reqService
    }

    public Details(feedURL: string): Response<Details> {
        const enc = window.btoa(feedURL)
        const path = `${paretPath}/${enc}`
        return this.reqService.Get<Details>(path)
    }
}

export default function create(reqService: ReqService): Service {
    return new Impl(reqService)
}

const paretPath = "/feed/details"