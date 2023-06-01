<template>
    <!-- 富文本编辑器 -->
    <div class="content">
        <div style="margin:15px">
            <rich-text-editor v-model:modelValue="addArticle.content"></rich-text-editor>
        </div>
    </div>

    <!-- 功能栏 -->
    <el-menu class="choiceBar" @select="handleSelect">
        <el-menu-item style="color: #409EFF;" index="1">发布</el-menu-item>
        <el-menu-item style="color: #F56C6C;" index="2">取消</el-menu-item>
    </el-menu>

    <!-- 上传文章弹框 -->
    <el-dialog v-model="showModal" title="上传文章" width="25%" center>
        <!-- 无封面时 -->
        <div v-if="!newHeadImage" style="width: 80%; margin: auto;">
            <el-upload drag :before-upload="beforeUpload" :http-request="customRequest" multiple>
                <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                <div class="el-upload__text">
                    拖动文件 或 <em>点击上传</em>
                </div>
                <template #tip>
                    <div class="el-upload__tip">
                        仅支持png/jpg/jpeg格式的图片
                    </div>
                </template>
            </el-upload>
        </div>
        <!-- 有封面时 -->
        <div v-else style="width: 80%; margin: auto;">
            <el-image :src="serverUrl + addArticle.headImage"></el-image>
            <el-button class="delete-button" size="large" @click="deleteImage" type="danger" :icon="Delete" circle />
        </div>
        <!-- 杂项 -->
        <div class="other-box">
            <!-- 分类选择 -->
            <el-select style="margin-top: 10px; width: 80%;" v-model="addArticle.categoryId" class="m-2" placeholder="请选择分类"
                size="large">
                <el-option v-for="item in categoryOptions" :key="item.value" :label="item.label"
                    :value="item.value"></el-option>
            </el-select>
            <!-- 标题输入 -->
            <el-input style="margin-top: 15px; width: 80%;" v-model="addArticle.title" placeholder="请输入标题" size="large" />
            <!-- 按钮 -->
            <div style="margin-top: 17px;">
                <el-button style="margin-right: 20px;" type="danger" @click="closeSubmitModal">取消</el-button>
                <el-button type="primary" @click="submit">确认</el-button>
            </div>
        </div>
    </el-dialog>
</template>

<script setup>


// ref：创建一个响应式的引用对象，用于包装一个简单数据类型的值。
// reactive：创建一个响应式的对象，用于包装一个复杂数据类型的值，如对象或数组。
// inject：从祖先组件提供的provide注入一个依赖项，使得当前组件可以访问它。
// onMounted：在组件加载完毕后立即执行一个函数。
import { ref, reactive, inject, onMounted } from 'vue'

// 导入富文本编辑器
import RichTextEditor from '../../components/RichTextEditor.vue'

// 导入一些icons
import { ArchiveOutline as ArchiveIcon } from "@vicons/ionicons5"
import { UploadFilled } from '@element-plus/icons-vue'
// icons
import {
    Delete,
} from '@element-plus/icons-vue'

// 导入路由
import { useRouter } from 'vue-router'
const router = useRouter()



// 这也是在Vue.js 3中使用的代码，它使用了inject函数来获取从祖先组件中通过provide提供的三个依赖项。
// serverUrl：这是一个服务器地址的字符串，用于向该地址发送HTTP请求。该值是通过在某个祖先组件中使用provide("serverUrl", serverUrl)提供的。在当前组件中，我们可以使用inject("serverUrl")来访问它。
// axios：这是一个类似于fetch的HTTP客户端工具，用于发送HTTP请求。同样地，这个值也是通过在祖先组件中使用provide("axios", axios)提供的。
// message：这是用于显示用户友好的消息的工具。在祖先组件中，我们可以使用provide("message", message)提供这个值，并在当前组件中使用inject("message")来访问它。
// 通过使用inject和provide，我们可以轻松地实现依赖注入，同时避免了深度嵌套的属性访问和传递。
const serverUrl = inject("serverUrl")
const axios = inject("axios")
import { ElMessage } from 'element-plus'

const categoryOptions = ref([])// 分类列表选项
const addArticle = reactive({// 待发布的文章对象
    id: 0,
    categoryId: "",
    title: "",
    content: "",
    headImage: "",
})

const value = ref('')

onMounted(() => {
    loadCategories()
})


import config from '../../config/config.json';
// 加载文章种类
const loadCategories = async () => {
    categoryOptions.value = config.menuItems.filter(item => item.index.startsWith("/blog?category=")).map((item) => {
        return {
            label: item.label,
            value: item.index.slice(-3)
        }
    })
    // console.log(categoryOptions)
}

// 控制发布文章时弹窗的显示与隐藏
const showModal = ref(false) // 响应式变量，常用于管理Vue.js组件中的状态和显示/隐藏逻辑，例如用于显示或隐藏模态框或图片上传框等
const showModalModal = () => {
    showModal.value = true
}
const closeSubmitModal = () => {
    showModal.value = false
}

// 判断图片的格式是否符合要求
const beforeUpload = async (file) => {
    const allowedTypes = ['image/jpeg', 'image/png', "image/jpeg"]; // 允许的文件类型
    if (!allowedTypes.includes(file.type)) {
        ElMessage({
            message: "只能上传png/jpg/jpeg格式的图片",
            type: 'error',
            offset: 80
        })
        return false;
    }
    return true;
}

// 控制文章头图
const newHeadImage = ref(false)
const customRequest = async (file) => {
    const formData = new FormData()
    formData.append('file', file.file)
    // console.log(formData)
    let res = await axios.post("/image/upload", formData)
    addArticle.headImage = res.data.data.filePath
    newHeadImage.value = true
}
// const deleteImage = () => {
//     addArticle.headImage = ""
//     newHeadImage.value = false
// }

const deleteImage = async () => {
  try {
    console.log(addArticle.headImage)
    const response = await axios.post("/image/delete", { filePath: addArticle.headImage })
    console.log(response.data)
    addArticle.headImage = ""
    newHeadImage.value = false
  } catch (error) {
    console.error(error)
  }
}

// 上传文章
const submit = async () => {
    if (addArticle.categoryId == "") {
        ElMessage({
            message: "请选择分类",
            type: 'error',
            offset: 80
        })
        return
    }
    if (addArticle.title == "") {
        ElMessage({
            message: "请输入标题",
            type: 'error',
            offset: 80
        })
        return
    }
    // console.log(value)
    let res = await axios.post("/article/create", {
        category_id: parseInt(addArticle.categoryId.slice(-1)),
        title: addArticle.title,
        content: addArticle.content,
        head_image: addArticle.headImage,
        article_type: "blogArticle"
    })
    // console.log(addArticle)
    // console.log(res)
    if (res.data.code == 200) {
        ElMessage({
            message: res.data.msg,
            type: 'success',
            offset: 80
        })
        goback()
    } else {
        ElMessage({
            message: res.data.msg,
            type: 'error',
            offset: 80
        })
    }

}


// 返回上级页面
const goback = () => {
    router.go(-1)
}

const handleSelect = (index) => {
    switch (index) {
        case "1":
            showModalModal()
            break;
        case "2":
            goback()
            break;
        default:
            break;
    }
}
</script>

<style lang="scss" scoped>
.content {
    position: relative;
    margin: auto;
    width: 1000px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    background: white;
    box-shadow: 2px 2px 6px #D3D4D8;
    border-radius: 10px;
    z-index: 99;
}

.other-box {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
}

.delete-button {
    position: absolute;
    right: 0px;
    bottom: 0px;
    box-shadow: 2px 2px 6px #D3D4D8;
}

.choiceBar {
    position: fixed;
    top: 25%;
    z-index: 999;
    width: 150px;
    box-shadow: 2px 0 6px rgba(0, 0, 0, 0.26);
    border-radius: 0 10px 10px 0;
}
</style>
