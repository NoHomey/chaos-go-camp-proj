import * as React from "react"
import Blog from "../../data/Blog"
import FeedURL from "./FeedURL"
import SaveBlog from "./SaveBlog"

const Page: React.FC<{}> = () => {
    const [blog, setBlog] = React.useState<null | Blog>(null)
    if(blog === null) {
        return <FeedURL setBlogDetails={setBlog} />
    }
    return <SaveBlog blog={blog} />
}

export default Page