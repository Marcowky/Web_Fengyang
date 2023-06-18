import { ElMessage } from 'element-plus'

export function showMessage(message, type) {
    ElMessage({
        message: message,
        type: type,
        offset: 80
    });
}