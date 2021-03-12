import * as React from "react"
import Tags from "./Tags"
import AddTag from "./AddTag"
import Tag from "../../data/Tag"

export interface Props {
    marginY: number,
    tags: Tag[],
    onAddTag: (tag: Tag) => void
}

const Tagger: React.FC<Props> = ({
    marginY,
    tags,
    onAddTag
}) => {
    const [open, setOpen] = React.useState(false)
    return (
        <>
            <Tags
                marginY={marginY}
                tags={tags}
                onAddTag={() => setOpen(true)} />
            <AddTag
                state={{ open }}
                event={{
                    onAddTag: tag => {
                        onAddTag(tag)
                        setOpen(false)
                    },
                    onClose: () => setOpen(false)
                }} />
        </>
    )
}

export default Tagger