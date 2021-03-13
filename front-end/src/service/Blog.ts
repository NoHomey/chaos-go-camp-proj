import { Response } from "../response";
import { Service as ReqService } from "./Request"
import Tag from "../data/Tag"
import Level from "../data/Level"

export type Blog = {
    feedURL: string
    author: string
    title: string
    description: string
    tags: Tag[]
    quickNote: string
    rating: number
    level: Level
}

export interface Service {
    Save(blog: Blog): Response<string>
}

class Impl implements Service {
    private reqService: ReqService

    public constructor(reqService: ReqService) {
        this.reqService = reqService
    }

    public Save(blog: Blog): Response<string> {
        return this.reqService.Post<string>(path, blog)
    }
}

export default function create(reqService: ReqService): Service {
    return new Impl(reqService)
}

const path = "/blog"