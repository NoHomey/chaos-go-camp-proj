import * as React from "react"
import { TagPriorty } from "../../../data/Tag"

export default function useData() {
    const defaultPriority = TagPriorty.Normal
    const [tag, setTag] = React.useState("")
    const [priority, setPriory] = React.useState(defaultPriority)
    return {
        data: {
            value: tag,
            priority
        },
        event: {
            onValueChange: setTag,
            onPriorityChange: setPriory
        },
        action: {
            reset: () => {
                setTag("")
                setPriory(defaultPriority)
            }
        }
    }
}