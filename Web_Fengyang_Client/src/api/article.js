import { axios } from '../main.js'
import { showMessage } from '../components/Message.js'


export async function articleListOut(pageInfo) {
    try {
        let arg = `/article/list?articleType=${pageInfo.pageArticleType}`
        if(pageInfo.sortKey != null) {
            arg = arg + `&order=${pageInfo.sortKey}`
        }
        if(pageInfo.keyword != null){
            arg = arg +`&keyword=${pageInfo.keyword}`
        }
        if(pageInfo.pageNum != null){
            arg = arg +`&pageNum=${pageInfo.pageNum}`
        }
        if(pageInfo.pageSize != null) {
            arg = arg +`&pageSize=${pageInfo.pageSize}`
        }
        if(pageInfo.categoryId != null) {
            arg = arg +`&categoryId=${pageInfo.categoryId}`
        }
        let res = await axios.get(arg)
        return res // 返回响应数据
    } catch (error) {
        return null
    }
}