import useLogic from "./useLogic"
import Page from "./Page"
import Blog from "../../../data/Blog";

export interface Props {
    setBlogDetails: (blog: Blog) => void
}

const Comp: React.FC<Props> = ({setBlogDetails}) => {
    const {data, event} = useLogic(setBlogDetails)
    return <Page data={data} event={event} />
}

export default Comp