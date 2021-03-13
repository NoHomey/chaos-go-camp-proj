import * as React from "react"
import Blog from "../../../data/Blog";
import { useService } from "../../../context/Service"
import { useReqDialog } from "../../../context/ReqDialog"
import { Details } from "../../../service/Feed"
import { RespError } from "../../../response"

export default function useLogic(cb: (blog: Blog) => void) {
    const [url, setURL] = React.useState("")
    const feed = useService("feed")
    const dialog = useReqDialog()
    return {
        data: { feedURL: url },
        event: {
            onFeedURLChange: setURL,
            onGetDetails: (feedURL: string) => {
                const res = feed.Details(feedURL)
                    .OnFail(() => dialog.showFail())
                    .OnError(err => {
                        dialog.showError(errorMsg(err), dialog.close)
                    })
                    .OnResult(details => {
                        cb(blog(feedURL, details))
                        dialog.close()
                    })
                dialog.showLoading("Getting feed details")
                setTimeout(res.Handle.bind(res), minShow)
            }
        }
    }
}

function blog(feedURL: string, details: Details): Blog {
    return {
        feedURL,
        author: details.author,
        title: details.title,
        description: details.description,
        tags: [],
        rating: null,
        level: null,
        quickNote: ""
    }
}

function errorMsg(err: RespError): string {
    return JSON.stringify(err, null, 4)
}

const minShow = 1500