<template>
    <!-- 富文本编辑器 -->
    <div class="content">
        <div style="margin:15px">
            <rich-text-editor v-if="loadOk" v-model:modelValue="updateArticle.content"></rich-text-editor>
        </div>
    </div>

    <!-- 功能栏 -->
    <el-menu class="choiceBar" @select="handleSelect">
        <el-menu-item style="color: #409EFF;" index="1">确认</el-menu-item>
        <el-menu-item style="color: #F56C6C;" index="2">取消</el-menu-item>
    </el-menu>

    <!-- 上传文章弹框 -->
    <el-dialog v-model="showModal" title="修改文章" width="25%" center>
        <!-- 无封面时 -->
        <div v-if="!newHeadImage" style="width: 80%; margin: auto;">
            <n-upload multiple directory-dnd :max="1" @before-upload="beforeUpload" :custom-request="customRequest">
                <n-upload-dragger>
                    <div style="margin-bottom: 12px">
                        <n-icon size="48" :depth="3">
                            <archive-icon />
                        </n-icon>
                    </div>
                    <n-text style="font-size: 16px">
                        点击或者拖动图片到此处
                    </n-text>
                </n-upload-dragger>
            </n-upload>
        </div>
        <!-- 有封面时 -->
        <div v-else style="width: 80%; margin: auto;">
            <el-image :src="serverUrl + updateArticle.headImage"></el-image>
            <el-button class="delete-button" size="large" @click="deleteImage" type="danger" :icon="Delete" circle />
        </div>

        <!-- 杂项 -->
        <div class="other-box">
            <!-- 分类选择 -->
            <el-select style="margin-top: 10px; width: 80%;" v-model="updateArticle.categoryId" class="m-2"
                placeholder="请选择分类" size="large">
                <el-option v-for="item in categoryOptions" :key="item.value" :label="item.label"
                    :value="item.value"></el-option>
            </el-select>
            <!-- 标题输入 -->
            <el-input style="margin-top: 15px; width: 80%;" v-model="updateArticle.title" placeholder="请输入标题"
                size="large" />
            <!-- 按钮 -->
            <div style="margin-top: 17px;">
                <el-button style="margin-right: 20px;" type="danger" @click="closeSubmitModal">取消</el-button>
                <el-button type="primary" @click="submit">确认</el-button>
            </div>
        </div>

    </el-dialog>
</template>

<script setup>
// icons
import { ArchiveOutline as ArchiveIcon } from "@vicons/ionicons5"
import {
    Delete,
} from '@element-plus/icons-vue'

import { ref, reactive, inject, onMounted } from 'vue'
// 富文本编辑器
import RichTextEditor from '../../components/RichTextEditor.vue'
// 导入路由
import { useRouter, useRoute } from 'vue-router'
const router = useRouter()
const route = useRoute()

// 网络请求
const serverUrl = inject("serverUrl")
const axios = inject("axios")
import { ElMessage } from 'element-plus'
// 变量初始化
const loadOk = ref(false)
const categoryOptions = ref([])
const updateArticle = reactive({
    id: 0,
    categoryId: "未选择",
    title: "",
    content: "",
    headImage: "",
})
const showModal = ref(false)
const newHeadImage = ref(true)

// 挂载页面时触发
onMounted(() => {
    loadCategories()
    loadArticle()
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

// 加载文章
const loadArticle = async () => {
    let res = await axios.get(`article/detail?articleType=blogArticle&id=${route.query.id}`)

    if (res.data.code == 200) {
        let label = categoryOptions.value.find((item) => item.value.endsWith(res.data.data.article.category_id)).label
        updateArticle.categoryId = label,
            updateArticle.title = res.data.data.article.title,
            updateArticle.content = res.data.data.article.content,
            updateArticle.headImage = res.data.data.article.head_image,
            newHeadImage.value = updateArticle.headImage ? true : false
        loadOk.value = true
    }

}

// 显示发布弹窗
const showModalModal = () => {
    showModal.value = true
}


// 关闭发布弹窗
const closeSubmitModal = () => {
    showModal.value = false
}

// 图像判断
const beforeUpload = async (data) => {
    if (data.file.file?.type !== "image/png") {
        ElMessage({
            message: "只能上传png格式的图片",
            type: 'error',
            offset: 80
        })
        return false;
    }
    return true;
}

// 上传封面
const customRequest = async ({ file }) => {
    const formData = new FormData()
    formData.append('file', file.file)
    let res = await axios.post("/image/upload", formData)

    updateArticle.headImage = res.data.data.filePath
    newHeadImage.value = true
}

// 删除封面
const deleteImage = () => {
    updateArticle.headImage = ""
    newHeadImage.value = false
}

// 提交文章
const submit = async () => {
    console.log(updateArticle)
    let res = await axios.put(`article/update?articleType=blogArticle&id=${route.query.id}`, {
        category_id: parseInt(updateArticle.categoryId.slice(-1)),
        title: updateArticle.title,
        content: updateArticle.content,
        head_image: updateArticle.headImage,
        article_type: "blogArticle"
    })

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

// 返回上一级
const goback = () => {
    router.go(-2)
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