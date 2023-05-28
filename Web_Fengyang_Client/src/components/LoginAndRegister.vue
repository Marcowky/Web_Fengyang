<template>
    <!-- 上传文章弹框 -->
    <el-dialog v-model="showModal" width="25%" center>
        <div class="box">
            <!-- 登录注册跳转按钮 -->
            <el-radio-group style="margin-bottom: 15px;" v-model="dialogType">
                <el-radio-button @click="toRegister" label="注册" />
                <el-radio-button @click="toLogin" label="登录" />
            </el-radio-group>

            <el-form ref="formRef" :rules="rules" :model="user" label-width="100px" :label-position="labelPosition"
                style="min-width: 250px">
                <!-- 注册时显示用户名输入框 -->
                <el-form-item v-if="dialogType == '注册'" label="用户名" prop="userName">
                    <el-input v-model="user.userName" />
                </el-form-item>
                <!-- 手机号输入框 -->
                <el-form-item label="手机号" prop="phoneNumber">
                    <el-input v-model="user.phoneNumber" />
                </el-form-item>
                <!-- 密码输入框 -->
                <el-form-item label="密码" prop="password">
                    <el-input v-model="user.password" type="password" show-password />
                </el-form-item>
                <!-- 注册时重新输入密码框 -->
                <el-form-item v-if="dialogType == '注册'" label="重新输入密码" prop="repeatPassword">
                    <el-input v-model="user.repeatPassword" type="password" show-password />
                </el-form-item>
            </el-form>

            <!-- 登录时显示记住密码复选框 -->
            <el-checkbox v-if="dialogType == '登录'" v-model="user.rember">记住密码</el-checkbox>

            <!-- 提交按钮 -->
            <el-button @click="submitForm(formRef)" class="button3" type="primary"
                style="left: auto; right: auto; text-align: center; margin-top: 10px;">
                {{ dialogType }}
            </el-button>

        </div>
    </el-dialog>
</template>
  
<script setup>
import { ref, reactive, inject } from 'vue'
import { ElMessage } from 'element-plus'
import { UserStore } from '../stores/UserStore'
import { useRoute } from 'vue-router'

const route = useRoute()
const formRef = ref(null)
const axios = inject("axios")
const userStore = UserStore()
const labelPosition = ref('top')
const showModal = ref(false)
const dialogType = ref("登录")
const user = reactive({
    userName: "",
    phoneNumber: localStorage.getItem("phoneNumber") || route.query.phoneNumber || "",
    password: localStorage.getItem("password") || route.query.password || "",
    rember: localStorage.getItem("rember") == 1 || false,
    repeatPassword: ""
})

// 验证两次输入的密码是否相同
function validatePasswordSame(rule, value) {
    return value == user.password;
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
    ],
    repeatPassword: [
        { required: true, message: "请重新输入密码", trigger: "blur" },
        { validator: validatePasswordSame, message: "两次输入的密码不一致", trigger: "blur" },
    ]
})

// 显示弹框
const showDialog = () => {
    showModal.value = true
}

// 切换为注册模式
const toRegister = () => {
    dialogType.value = '注册'
}

// 切换为登录模式
const toLogin = () => {
    dialogType.value = '登录'
}

// 提交表单
const submitForm = async (formEl) => {
    if (!formEl) return
    await formEl.validate((valid, fields) => {
        if (valid) {
            if (dialogType.value === '登录') {
                login()
            } else if (dialogType.value === '注册') {
                register()
            }
        } else {
            console.log("error!")
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
const login = async () => {
    let res = await axios.post("/user/login", {
        phoneNumber: user.phoneNumber,
        password: user.password
    })

    if (res.data.code == 200) {
        userStore.token = res.data.data.token
        if (user.rember) {
            localStorage.setItem("phoneNumber", user.phoneNumber)
            localStorage.setItem("password", user.password)
            localStorage.setItem("rember", user.rember ? 1 : 0)
        } else {
            localStorage.removeItem("phoneNumber")
            localStorage.removeItem("password")
            localStorage.setItem("rember", user.rember ? 1 : 0)
        }
        ElMessage({
            message: res.data.msg,
            type: 'success',
            offset: 80
        })
        showModal.value = false
    } else {
        ElMessage({
            message: res.data.msg,
            type: 'error',
            offset: 80
        })
    }
}

// 注册
const register = async () => {
    let res = await axios.post("/user/register", {
        userName: user.userName,
        phoneNumber: user.phoneNumber,
        password: user.password
    })

    if (res.data.code == 200) {
        ElMessage({
            message: res.data.msg,
            type: 'success',
            offset: 80
        })
        dialogType.value = '登录'
    } else {
        ElMessage({
            message: res.data.msg,
            type: 'error',
            offset: 80
        })
    }
}

// 暴露方法给组件外部调用
defineExpose({
    showDialog
})
</script>
  
<style scoped>
.box {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}
</style>
  