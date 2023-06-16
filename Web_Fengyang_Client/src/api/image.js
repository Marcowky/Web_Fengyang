import { axios } from '../main.js'
import { showMessage } from '../components/Message.js'

export async function imageUpload(file) {
    try {
        const formData = new FormData()
        formData.append('file', file.file)
        let res = await axios.post("/image/upload", formData)
        showMessage(res.data.msg, 'success')
        return res // 返回响应数据
    } catch (error) {
        showMessage(error.response.data.msg, 'error')
        return null
    }
}

export async function imageDelete(imagePath) {
    try {
        let res = await axios.post("/image/delete", { filePath: imagePath })
        showMessage(res.data.msg, 'success')
        return null // 返回响应数据
    } catch (error) {
        showMessage(error.response.data.msg, 'error')
        return error
    }
}