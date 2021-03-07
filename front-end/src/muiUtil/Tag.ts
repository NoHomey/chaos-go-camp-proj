import { TagPriorty } from "../data/Tag"
import { PropTypes } from "@material-ui/core"

type Color = Exclude<PropTypes.Color, "inherit">

export function PriorityToColor(priority: TagPriorty): Color {
    switch(priority) {
    case TagPriorty.Main: return "primary"
    case TagPriorty.Secondary: return "secondary"
    case TagPriorty.Normal: return "default"
    }
}