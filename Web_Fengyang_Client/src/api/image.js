import { axios } from '../main.js'
import { showMessage } from '../components/Message.js'

export async function imageUpload(file) {
    try {
        const formData = new FormData()
        formData.append('file', file.file)
        let res = await axios.post("/image/upload", formData)
        showMessage(res.data.msg, 'success')
        return res 
    } catch (error) {
        showMessage(error.response.data.msg, 'error')
        return null
    }
}

export async function imageCheck(file) {
    const allowedTypes = ['image/jpeg', 'image/png', "image/jpeg"]; // 允许的文件类型
    if (!allowedTypes.includes(file.type)) {
        showMessage("只能上传png/jpg/jpeg格式的图片", 'warning')
        return false;
    }
    return true;
}

export async function imageDelete(imagePath) {
    try {
        let res = await axios.post("/image/delete", { filePath: imagePath })
        showMessage(res.data.msg, 'success')
        return null 
    } catch (error) {
        showMessage(error.response.data.msg, 'error')
        return error
    }
}