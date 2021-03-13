export enum TagPriorty {
    Normal,
    Secondary,
    Main
}

export interface Tag {
    value: string
    priority: TagPriorty
}

export default Tag