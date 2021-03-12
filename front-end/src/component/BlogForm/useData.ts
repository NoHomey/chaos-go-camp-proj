import * as React from "react"
import Blog from "../../data/Blog"
import Tag from "../../data/Tag"

export default function useData(blog: Blog) {
    const [author, setAuthor] = React.useState(blog.author)
    const [title, setTitle] = React.useState(blog.title)
    const [description, setDesc] = React.useState(blog.description)
    const [tags, setTags] = React.useState(blog.tags)
    const [quickNote, setQuickNote] = React.useState(blog.quickNote)
    const [rating, setRating] = React.useState(blog.rating)
    const [level, setLevel] = React.useState(blog.level)
    return {
        data: {
            feedURL: blog.feedURL,
            author,
            title,
            description,
            tags,
            quickNote,
            rating,
            level
        },
        event: {
            onAuthorChange: setAuthor,
            onTitleChange: setTitle,
            onDescriptionChange: setDesc,
            onQuickNoteChange: setQuickNote,
            onRatingChange: setRating,
            onLevelChange: setLevel,
            onRemoveTag: (idx: number) => setTags(remove(tags, idx)),
            onAddTag: (tag: Tag) => setTags(add(tags, tag))
        }
    }
}

function remove(tags: Tag[], idx: number): Tag[] {
    return [...tags.slice(0, idx), ...tags.slice(idx + 1)]
}

function add(tags: Tag[], tag: Tag): Tag[] {
    const exists = tags.some(t => t.value === tag.value)
    if(exists) {
        return tags
    }
    const n = [...tags, tag]
    n.sort((a, b) => {
        if(a.priority < b.priority) {
            return -1
        }
        if(a.priority > b.priority) {
            return 1
        }
        if(a.value < b.value) {
            return -1
        }
        if(a.value > b.value) {
            return 1
        }
        return 0
    })
    return n
}