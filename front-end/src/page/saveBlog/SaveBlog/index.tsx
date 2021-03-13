import Page from "./Page"
import Blog from "../../../data/Blog";
import { Blog as BlogData } from "../../../service/Blog"
import { useService } from "../../../context/Service"
import { useReqDialog } from "../../../context/ReqDialog"
import { RespError } from "../../../response"
import Level from "../../../data/Level";

export interface Props {
    blog: Blog
}

const Comp: React.FC<Props> =({blog}) => {
    const blogService = useService("blog")
    const dialog = useReqDialog()
    const onSave = (blog: Blog) => {
        const res = blogService.Save(normalize(blog))
            .OnFail(() => dialog.showFail())
            .OnError(err => {
                dialog.showError(errorMsg(err), dialog.close)
            })
            .OnResult(id => {
                dialog.showResult(id, dialog.close)
            })
        dialog.showLoading("Saving blog")
        setTimeout(res.Handle.bind(res), minShow)
    }
    return <Page blog={blog} onSaveBlog={onSave} />
}

export default Comp

function normalize(blog: Blog): BlogData {
    return {
        ...blog,
        level: blog.level ?? Level.NotSelected,
        rating: blog.rating ?? 0
    }
}

function errorMsg(err: RespError): string {
    return JSON.stringify(err, null, 4)
}

const minShow = 500