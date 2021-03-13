import Page from "./Page"
import Blog from "../../../data/Blog";

export interface Props {
    blog: Blog
}

const Comp: React.FC<Props> =({blog}) => {
    const onSave = (blog: Blog) => console.log(blog)
    return <Page blog={blog} onSaveBlog={onSave} />
}

export default Comp