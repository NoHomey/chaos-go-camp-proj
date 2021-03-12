export enum TagPriorty {
    Main,
    Secondary,
    Normal
}

export interface Tag {
    value: string
    priority: TagPriorty
}

export default Tag