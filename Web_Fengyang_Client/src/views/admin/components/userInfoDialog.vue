<template>
    <el-dialog model-value="showModal" :title="dialogTitle" width="30%" show-close=false @closed='handleClose'>

        <el-form ref="formRef" :rules="rules" :model="user" label-width="100px" :label-position="labelPosition"
            style="min-width: 250px">
            <el-form-item label="ID" prop="ID" v-if="dialogTitle=='编辑用户'" >
                <el-input v-model="user.ID" disabled/>
            </el-form-item>
            <el-form-item label="用户名" prop="UserName">
                <el-input v-model="user.UserName" />
            </el-form-item>
            <el-form-item label="手机号" prop="PhoneNumber">
                <el-input v-model="user.PhoneNumber" />
            </el-form-item>
            <el-form-item label="密码" v-if="dialogTitle=='添加用户'" prop="Password">
                <el-input v-model="user.Password" />
            </el-form-item>
            <el-form-item label="用户类型" prop="UserType">
                <el-input v-model="user.UserType" />
            </el-form-item>
            <el-form-item label="创建时间" prop="CreatedAt" v-if="dialogTitle=='编辑用户'" >
                <el-input v-model="user.CreatedAt" disabled/>
            </el-form-item>
            <el-form-item label="状态" prop="Status">
                <el-input v-model="user.Status" />
            </el-form-item>
        </el-form>

        <el-button @click="submitForm(formRef)" class="button3" type="primary"
            style="left: auto; right: auto; text-align: center; margin-top: 10px;">
            确认
        </el-button>
    </el-dialog>
</template>
  
<script setup>

import { ref, reactive, inject, defineProps, defineEmits, watch } from 'vue';

const props = defineProps({
    dialogTitle: {
        type: String,
        default: '',
        required: true
    },
    dialogTableValue: {
        tyoe: Object,
        default: () => {}
    }
})

const formRef = ref(null)
const user = ref({
    ID: "",
    UserName: "",
    PhoneNumber: "",
    CreatedAt:"",
    Password: "",
    UserType: "",
    Status: ""
})

watch(()=>props.dialogTableValue, () => (
    console.log(props.dialogTableValue),
    user.value = props.dialogTableValue,
    console.log(user.value),
    console.log(user.ID)
), {deep: true, immediate: true})

const emits = defineEmits(['update:modelValue', 'updateUserList'])

const handleClose = () => {
    emits('update:modelValue',false)
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
    console.log("okkkkk")
    if (!formEl) return
    
    await formEl.validate((valid, fields) => {
        if (valid) {
            if(props.dialogTitle=="添加用户"){
                console.log("添加用户")
                register()
            }
            else if(props.dialogTitle=="编辑用户"){
                console.log("编辑用户")
                updateUser()
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
import { ElMessage } from 'element-plus'
const axios = inject("axios")
// 注册
const register = async () => {
    let res = await axios.post("/user/register", {
        userName: user.value.UserName,
        phoneNumber: user.value.PhoneNumber,
        password: user.value.Password,
        userType: "client"
    })

    if (res.data.code == 200) {
        ElMessage({
            message: '添加成功',
            type: 'success',
            offset: 80
        })
        emits('updateUserList')
        handleClose()
    } else {
        ElMessage({
            message: res.data.msg,
            type: 'error',
            offset: 80
        })
    }
}

const updateUser = async () => {
    console.log(user.value.PhoneNumber)
    let res = await axios.put("/user/update", {
        ID: user.value.ID,
        userName: user.value.UserName,
        phoneNumber: user.value.PhoneNumber,
        userType: "client",
        Status: user.value.Status
    })
    if (res.data.code == 200) {
        ElMessage({
            message: res.data.msg,
            type: 'success',
            offset: 80
        })
        emits('updateUserList')
        handleClose()
    } else {
        ElMessage({
            message: res.data.msg,
            type: 'error',
            offset: 80
        })
    }
}

</script>
  
<style scoped></style>
  