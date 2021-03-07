export enum TagPriorty {
    Main,
    Secondary,
    Normal
}

export default interface Tag {
    value: string
    priority: TagPriorty
}