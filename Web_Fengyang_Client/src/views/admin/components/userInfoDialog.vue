<template>
    <el-dialog model-value="showModal" :title="dialogTitle" width="30%" show-close=false @closed='handleClose' center>

        <el-form ref="formRef" :rules="rules" :model="user" label-width="100px" :label-position="labelPosition"
            style="min-width: 250px; padding-right: 40px;">
            <el-form-item label="ID" prop="ID" v-if="dialogTitle == editUserTitle">
                <el-input v-model="user.ID" disabled />
            </el-form-item>
            <el-form-item label="用户名" prop="userName">
                <el-input v-model="user.userName" />
            </el-form-item>
            <el-form-item label="手机号" prop="phoneNumber">
                <el-input v-model="user.phoneNumber" />
            </el-form-item>
            <el-form-item label="密码" v-if="dialogTitle == addUserTitle" prop="password">
                <el-input v-model="user.password" />
            </el-form-item>
            <el-form-item label="创建时间" prop="createdAt" v-if="dialogTitle == editUserTitle">
                <el-input v-model="user.createdAt" disabled />
            </el-form-item>
            <el-form-item label="用户类型" prop="userType">
                <el-tag class="ml-2" type="success">{{ user.userType }}</el-tag>
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

import { ref, reactive, watch } from 'vue';
import { userUpdate } from '../../../api/user';
import { userRegister } from '../../../api/user';
import { showMessage } from '../../../components/Message';

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
const editUserTitle = '编辑用户'
const addUserTitle = '添加用户'
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
    user.value.ID = props.dialogTableValue.ID,
    user.value.userName = props.dialogTableValue.UserName,
    user.value.phoneNumber = props.dialogTableValue.PhoneNumber,
    user.value.createdAt = props.dialogTableValue.CreatedAt,
    user.value.password = props.dialogTableValue.Password,
    user.value.userType = props.dialogTableValue.UserType,
    user.value.status = props.dialogTableValue.Status
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
            if (props.dialogTitle == addUserTitle) {
                register()
            }
            else if (props.dialogTitle == editUserTitle) {
                updateUser()
            }
        } else {
            if (dialogType.value === '登录') {
                showMessage('提交失败', 'error')
            } else if (dialogType.value === '注册') {
                showMessage('注册失败', 'error')
            }
        }
    })
}
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
  