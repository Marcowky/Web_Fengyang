<template>
    <!-- 上传文章弹框 -->
    <el-dialog v-model="showModal" title="上传文章" width="25%" center>
        <!-- 登录注册跳转按钮 -->
        <el-radio-group v-model="dialogType" size="large">
            <el-radio-button @click="toRegister" label="注册" />
            <el-radio-button @click="toLogin" label="登录" />
        </el-radio-group>
        <el-form ref="formRef" :rules="rules" :model="user">

            <el-form-item v-if="dialogType == '注册'" label="用户名" prop="userName">
                <el-input v-model="user.userName" />
            </el-form-item>

            <el-form-item label="手机号" prop="phoneNumber">
                <el-input v-model="user.phoneNumber" />
            </el-form-item>


            <el-form-item label="密码" prop="password">
                <el-input v-model="user.password" />
            </el-form-item>

            <el-form-item v-if="dialogType == '注册'" label="重新输入密码" prop="repeatPassword">
                <el-input v-model="user.repeatPassword" />
            </el-form-item>

        </el-form>
        <el-checkbox v-model="user.remember">记住密码</el-checkbox>
        <el-button @click="submitForm(formRef)" class="button3" type="primary"
            style="left: auto; right: auto; text-align: center;">
            登录
        </el-button>
    </el-dialog>
</template>


<script setup>
import { ref, reactive, inject } from 'vue'
import { ElMessage } from 'element-plus'
import { UserStore } from '../../stores/UserStore'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()
const formRef = ref(null)
const axios = inject("axios")
const userStore = UserStore()
const user = reactive({
    userName: "",
    phoneNumber: localStorage.getItem("phoneNumber") || route.query.phoneNumber || "",
    password: localStorage.getItem("password") || route.query.password || "",
    rember: localStorage.getItem("rember") == 1 || false,
    repeatPassword: ""
})


const showModal = ref(false)
const dialogType = ref("登录")

function validatePasswordSame(rule, value) {
    return value == user.password;
}

const rules = reactive[{
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
}]

const showDialog = () => {
    showModal.value = true
}

const toRegister = () => {
    dialogType.value = 'register'
}

const toLogin = () => {
    dialogType.value = 'login'
}


const submitForm = async (formEl) => {
    if (!formEl) return
    await formEl.validate((valid, fields) => {
        if (valid) {
            // 执行登录或注册逻辑
            if (dialogType.value === '登录') {
                console.log("login!")
                login()
            } else if (dialogType.value === '注册') {
                console.log("res!")
                register()
            }
            showModal.value = false

        } else {
            console.log("error!")
            if (dialogType.value === '登录') {
                ElMessage({
                    message: "提交失败",
                    type: 'error',
                    offset: 80
                })
                // login()
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
        // router.go(-1)
        ElMessage({
            message: res.data.msg,
            type: 'success',
            offset: 80
        })
    } else {
        ElMessage({
            message: res.data.msg,
            type: 'error',
            offset: 80
        })
    }
}

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
        router.push({
            path: "login",
            query: {
                phoneNumber: user.phoneNumber,
                password: user.password
            }
        })
    } else {
        ElMessage({
            message: res.data.msg,
            type: 'error',
            offset: 80
        })
    }
}

defineExpose({
    showDialog
})
</script>