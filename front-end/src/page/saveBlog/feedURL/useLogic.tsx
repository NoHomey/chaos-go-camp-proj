import * as React from "react"
import Blog from "../../../data/Blog";

export default function useLogic(cb: (blog: Blog) => void) {
    const [url, setURL] = React.useState("")
    return {
        data: {
            feedURL: url
        },
        event: {
            onFeedURLChange: setURL,
            onGetDetails: (feedURL: string) => {
                cb({
                    feedURL,
                    author: "Author",
                    title: "Titile",
                    description: "Desc",
                    tags: [],
                    rating: null,
                    level: null,
                    quickNote: ""
                })
            }
        }
    }
}
