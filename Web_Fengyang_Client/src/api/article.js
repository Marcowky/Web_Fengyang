import { axios } from '../main.js'
import { showMessage } from '../components/Message.js'

export async function articleListOut(pageInfo) {
    try {
        let arg = `/article/list?articleType=${pageInfo.pageArticleType}`
        if (pageInfo.sortKey != null) {
            arg = arg + `&order=${pageInfo.sortKey}`
        }
        if (pageInfo.keyword != null) {
            arg = arg + `&keyword=${pageInfo.keyword}`
        }
        if (pageInfo.pageNum != null) {
            arg = arg + `&pageNum=${pageInfo.pageNum}`
        }
        if (pageInfo.pageSize != null) {
            arg = arg + `&pageSize=${pageInfo.pageSize}`
        }
        if (pageInfo.categoryId != null) {
            arg = arg + `&categoryId=${pageInfo.categoryId}`
        }
        let res = await axios.get(arg)
        return res
    } catch (error) {
        return null
    }
}

export async function articlePost(article) {
    try {
        if (article.categoryId == "") {
            showMessage("请选择分类", 'error')
            return false
        }
        if (article.title == "") {
            showMessage("请输入标题", 'error')
            return false
        }
        let res = await axios.post("/article/create", {
            category_id: parseInt(article.categoryId),
            title: article.title,
            content: article.content,
            head_image: article.headImage,
            article_type: article.articleType
        })
        showMessage(res.data.msg, 'success')
        return null
    } catch (error) {
        showMessage(error.response.data.msg, 'error')
        return error
    }
}

export async function articleDelete(articleType, articleID) {
    try {
        let res = await axios.delete(`article/delete?articleType=${articleType}&id=${articleID}`)
        showMessage(res.data.msg, 'success')
        return null
    } catch (error) {
        showMessage(error.response.data.msg, 'error')
        return error
    }
}

export async function articleDetail(articleType, articleID) {
    try {
        let res = await axios.get(`article/detail?articleType=${articleType}&id=${articleID}`)
        return res
    } catch (error) {
        return null
    }
}

export async function articleUpdate(article) {
    try {
        if (article.title == "") {
            showMessage('请输入标题', 'error')
            return false
        }
        if (article.categoryId == "") article.categoryId = article.oldCategoryId.toString()
        let res = await axios.put(`article/update?articleType=${article.articleType}&id=${article.id}`, {
            category_id: parseInt(article.categoryId.slice(-1)),
            title: article.title,
            content: article.content,
            head_image: article.headImage,
            article_type: article.articleType
        })
        showMessage(res.data.msg, 'success')
        return null
    } catch (error) {
        showMessage(error.response.data.msg, 'error')
        return error
    }
}