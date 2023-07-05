<template>
    <el-dialog model-value="showModal" :title="dialogTitle" width="25%" show-close=false @closed='handleClose' center>

        <el-form ref="formRef" :rules="rules" :model="user" label-width="100px" :label-position="labelPosition"
            style="min-width: 250px; padding-right: 40px;">
            <el-form-item label="ID" prop="ID" v-if="dialogTitle == editUserTitle">
                <el-input v-model="user.ID" disabled />
            </el-form-item>
            <el-form-item label="用户名" prop="UserName">
                <el-input v-model="user.UserName" />
            </el-form-item>
            <el-form-item label="手机号" prop="PhoneNumber">
                <el-input v-model="user.PhoneNumber" />
            </el-form-item>
            <el-form-item label="密码" v-if="dialogTitle == addUserTitle" prop="Password">
                <el-input v-model="user.Password" />
            </el-form-item>
            <el-form-item label="创建时间" prop="CreatedAt" v-if="dialogTitle == editUserTitle">
                <el-input v-model="user.CreatedAt" disabled />
            </el-form-item>
            <el-row>
                <el-col :span="11" style="padding-left: 15px;">
                    <el-form-item label="用户类型" prop="UserType">
                        <el-tag class="ml-2" type="success">{{ user.UserType }}</el-tag>
                    </el-form-item>
                </el-col>
                <el-col :span="12">
                    <el-form-item label="状态" prop="Status">
                        <el-switch v-model="user.Status" @change="updateUser(scope.row)" />
                    </el-form-item>
                </el-col>
            </el-row>
        </el-form>

        <div class="button_area">
            <el-button @click="submitForm(formRef)" type="primary" style="margin-right: 30px;">
                确认
            </el-button>
            <el-button @click="handleClose" type="danger">
                取消
            </el-button>
        </div>

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
        type: Object,
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
    UserName: "",
    PhoneNumber: "",
    CreatedAt: "",
    Password: "",
    UserType: "",
    Status: ""
})

watch(() => props.dialogTableValue, () => (
    user.value = props.dialogTableValue
), { deep: true, immediate: true })

watch(() => props.dialogUserType, () => (
    user.value.UserType = props.dialogUserType
), { deep: true, immediate: true })

const emits = defineEmits(['update:modelValue', 'updateUserList'])

const handleClose = () => {
    emits('update:modelValue', false)
}

// 表单验证规则
const rules = reactive({
    UserName: [
        { required: true, message: "请输入用户名", trigger: "blur" },
        { min: 3, max: 20, message: "用户名长度在 3 到 20 个字符", trigger: "blur" },
    ],
    PhoneNumber: [
        { required: true, message: "请输入手机号", trigger: "blur" },
        { min: 11, max: 11, message: "手机号为 11 位", trigger: "blur" },
    ],
    Password: [
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
    let tempUser = ({
        ID: user.value.ID,
        userName: user.value.UserName,
        phoneNumber: user.value.PhoneNumber,
        createdAt: user.value.CreatedAt,
        password: user.value.Password,
        userType: user.value.UserType,
        status: user.value.Status
    })
    userRegister(tempUser).then(result => {
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
  
<style scoped>
.button_area {
    display: flex;
    justify-content: center;
}
</style>
  