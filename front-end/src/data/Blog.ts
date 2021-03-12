import Level from "./Level";
import Tag from "./Tag"

export default interface Blog {
    feedURL: string
    author: string
    title: string
    description: string
    tags: Tag[]
    quickNote: string
    rating: null | number
    level: null | Level
}