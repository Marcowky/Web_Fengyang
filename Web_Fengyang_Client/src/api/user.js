import { axios } from '../main.js'
import { showMessage } from '../components/Message.js'

export async function userRegister(user) {
    try {
        let res = await axios.post("/user/register", {
            userName: user.userName,
            phoneNumber: user.phoneNumber,
            password: user.password,
            userType: user.userType,
            status: user.status
        })
        showMessage(res.data.msg, 'success')
        return null // 返回响应数据
    } catch (error) {
        showMessage(error.response.data.msg, 'error')
        return error
    }
}

export async function userLogin(user) {
    try {
        let res = await axios.post("/user/login", {
            phoneNumber: user.phoneNumber,
            password: user.password,
            userType: user.userType
        })
        showMessage(res.data.msg, 'success')
        return res // 返回响应数据
    } catch (error) {
        showMessage(error.response.data.msg, 'error')
        return error
    }
}

export async function userDelete(user) {
    try {
        let res = await axios.put("/user/delete", {
            ID: user.ID,
            UserType: user.UserType
        })
        showMessage(res.data.msg, 'success')
        return null // 返回响应数据
    } catch (error) {
        showMessage(error.response.data.msg, 'error')
        return error
    }
}

export async function userUpdate(user) {
    try {
        let res = await axios.put("/user/update", {
            ID: user.ID,
            UserName: user.UserName,
            PhoneNumber: user.PhoneNumber,
            userType: user.UserType,
            Status: user.Status
        })
        showMessage(res.data.msg, 'success')
        return null // 返回响应数据
    } catch (error) {
        showMessage(error.response.data.msg, 'error')
        return error
    }
}

export async function userListOut(pageInfo) {
    try {
        let res = await axios.get(`/user/list?userType=${pageInfo.userType}&keyword=${pageInfo.keyword}&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}&order=${pageInfo.sortKey}`)
        return res // 返回响应数据
    } catch (error) {
        return null
    }
}

export async function userBriefInfo(articleCategory, userID) {
    try {
        let userCategory = (articleCategory == 'blogArticle') ? 'client' : 'admin'
        let res = await axios.get(`user/briefInfo?userType=${userCategory}&id=${userID}`)
        return res // 返回响应数据
    } catch (error) {
        return null
    }
}

export async function userInfo() {
    try {
        let res = await axios.get("user/info")
        return res // 返回响应数据
    } catch (error) {
        return null
    }
}