<template>
    <el-dialog model-value="showModal" :title="dialogTitle" width="30%" show-close=false @closed='handleClose'>

        <el-form ref="formRef" :rules="rules" :model="user" label-width="100px" :label-position="labelPosition"
            style="min-width: 250px">
            <el-form-item label="ID" prop="ID" v-if="dialogTitle == '编辑用户'">
                <el-input v-model="user.ID" disabled />
            </el-form-item>
            <el-form-item label="用户名" prop="userName">
                <el-input v-model="user.userName" />
            </el-form-item>
            <el-form-item label="手机号" prop="phoneNumber">
                <el-input v-model="user.phoneNumber" />
            </el-form-item>
            <el-form-item label="密码" v-if="dialogTitle == '添加用户'" prop="password">
                <el-input v-model="user.password" />
            </el-form-item>
            <el-form-item label="用户类型" prop="userType">
                <el-tag class="ml-2" type="success">{{ user.userType }}</el-tag>
            </el-form-item>
            <el-form-item label="创建时间" prop="createdAt" v-if="dialogTitle == '编辑用户'">
                <el-input v-model="user.createdAt" disabled />
            </el-form-item>
            <el-form-item label="状态" prop="status">
                <el-switch v-model="user.status" @change=updateUser(scope.row) />
            </el-form-item>
        </el-form>

        <el-button @click="submitForm(formRef)" class="button3" type="primary"
            style="left: auto; right: auto; text-align: center; margin-top: 10px;">
            确认
        </el-button>
    </el-dialog>
</template>
  
<script setup>

import { ref, reactive, inject, watch } from 'vue';
import { userUpdate } from '../../../api/user';
import { userRegister } from '../../../api/user';

const props = defineProps({
    dialogTitle: {
        type: String,
        default: '',
        required: true
    },
    dialogTableValue: {
        tyoe: Object,
        default: () => { }
    },
    dialogUserType: {
        type: String,
        default: '',
        required: true
    }
})

const formRef = ref(null)
const user = ref({
    ID: "",
    userName: "",
    phoneNumber: "",
    createdAt: "",
    password: "",
    userType: "",
    status: ""
})

watch(() => props.dialogTableValue, () => (
    user.value = props.dialogTableValue
), { deep: true, immediate: true })

watch(() => props.dialogUserType, () => (
    user.value.userType = props.dialogUserType
), { deep: true, immediate: true })

const emits = defineEmits(['update:modelValue', 'updateUserList'])

const handleClose = () => {
    emits('update:modelValue', false)
}

// 表单验证规则
const rules = reactive({
    userName: [
        { required: true, message: "请输入用户名", trigger: "blur" },
        { min: 3, max: 20, message: "用户名长度在 3 到 20 个字符", trigger: "blur" },
    ],
    phoneNumber: [
        { required: true, message: "请输入手机号", trigger: "blur" },
        { min: 11, max: 11, message: "手机号为 11 位", trigger: "blur" },
    ],
    password: [
        { required: true, message: "请输入密码", trigger: "blur" },
        { min: 6, max: 20, message: "密码长度在 6 到 20 个字符", trigger: "blur" },
    ]
})


// 提交表单
const submitForm = async (formEl) => {
    if (!formEl) return

    await formEl.validate((valid, fields) => {
        if (valid) {
            if (props.dialogTitle == "添加用户") {
                register()
            }
            else if (props.dialogTitle == "编辑用户") {
                updateUser()
            }
        } else {
            if (dialogType.value === '登录') {
                ElMessage({
                    message: "提交失败",
                    type: 'error',
                    offset: 80
                })
            } else if (dialogType.value === '注册') {
                ElMessage({
                    message: "注册失败",
                    type: 'error',
                    offset: 80
                })
            }
        }
    })
}

// 登录
import { ElMessage } from 'element-plus'
// 注册
const register = async () => {


    userRegister(user.value).then(result => {
        if (result == null) {
            emits('updateUserList')
            handleClose()
        }
    })

}

const updateUser = async () => {
    userUpdate(user.value).then(result => {
        if (result == null) {
            emits('updateUserList')
            handleClose()
        }
    })
}

</script>
  
<style scoped></style>
  