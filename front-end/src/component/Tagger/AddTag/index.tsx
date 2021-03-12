import Dialog from "./Dialog"
import useData from "./useData"
import Tag from "../../../data/Tag"

export interface Props {
    state: {
        open: boolean
    }
    event: {
        onAddTag: (tag: Tag) => void
        onClose: () => void
    }
}

const AddTag: React.FC<Props> = props => {
    const res = useData()
    const data = res.data
    const event = {
        ...res.event,
        onClose: props.event.onClose,
        onAddTag: (tag: Tag) => {
            props.event.onAddTag(tag)
            res.action.reset()
        }

    }
    const state = props.state
    return <Dialog state={state} data={data} event={event} />
}

export default AddTag