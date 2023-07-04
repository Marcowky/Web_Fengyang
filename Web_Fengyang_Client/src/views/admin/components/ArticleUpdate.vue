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
            <el-upload drag :before-upload="imageCheck" :http-request="customRequest" multiple>
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
            <el-image :src="serverUrl + updateArticle.headImage"></el-image>
            <el-button class="delete-button" size="large" @click="deleteImage" type="danger" :icon="Delete" circle />
        </div>

        <!-- 杂项 -->
        <div class="other-box">
            <!-- 分类选择 -->
            <el-select style="margin-top: 10px; width: 80%;" v-model="updateArticle.categoryId" class="m-2"
                :placeholder="updateArticle.oldCategory" size="large">
                <el-option v-for="item in categoryOptions" :key="item.value" :label="item.label"
                    :value="item.value"></el-option>
            </el-select>
            <!-- 标题输入 -->
            <el-input style="margin-top: 15px; width: 80%;" v-model="updateArticle.title" placeholder="请输入标题"
                size="large" />
            <!-- 按钮 -->
            <div style="margin-top: 17px;">
                <el-button style="margin-right: 20px;" type="primary" @click="submit">确认</el-button>
                <el-button type="danger" @click="closeSubmitModal">取消</el-button>
            </div>
        </div>

    </el-dialog>
</template>

<script setup>
// icons
import {
    Delete,
    UploadFilled
} from '@element-plus/icons-vue'

import { ref, reactive, inject, onMounted } from 'vue'
// 富文本编辑器
import RichTextEditor from '../../../components/RichTextEditor.vue'
import { imageUpload, imageDelete, imageCheck } from '../../../api/image'
import { articleDetail, articleUpdate } from '../../../api/article'
// 导入路由
import { useRouter, useRoute } from 'vue-router'
const router = useRouter()
const route = useRoute()

// 网络请求
const serverUrl = inject("serverUrl")
// 变量初始化
const loadOk = ref(false)
const categoryOptions = ref([])
const updateArticle = reactive({
    id: route.query.id,
    categoryId: "",
    title: "",
    content: "",
    headImage: "",
    oldCategory: "",
    oldCategoryId: "",
    articleType: route.query.category
})
const showModal = ref(false)
const newHeadImage = ref(true)

// 挂载页面时触发
onMounted(() => {
    loadCategories()
    loadArticle()
})

import config from '../../../config/config.json';
// 加载文章种类
const loadCategories = async () => {
    categoryOptions.value = config.menuItems.filter(item => item.mainMenu == '/' + updateArticle.articleType.substring(0, updateArticle.articleType.length - 7)).map((item) => {
        return {
            label: item.label,
            value: item.index
        }
    })
}

// 加载文章
const loadArticle = async () => {
    articleDetail(updateArticle.articleType, updateArticle.id).then(result => {
        if (result != null) {
            let label = categoryOptions.value.find((item) => item.value.endsWith(result.data.data.article.category_id)).label
            updateArticle.oldCategoryId = result.data.data.article.category_id
            updateArticle.oldCategory = label
            updateArticle.title = result.data.data.article.title
            updateArticle.content = result.data.data.article.content
            updateArticle.headImage = result.data.data.article.head_image
            newHeadImage.value = updateArticle.headImage ? true : false
            loadOk.value = true
        }
    })
}

// 显示发布弹窗
const showModalModal = () => {
    showModal.value = true
}


// 关闭发布弹窗
const closeSubmitModal = () => {
    showModal.value = false
}

// 上传封面
const customRequest = async (file) => {

    imageUpload(file).then(result => {
        if (result != null) {
            updateArticle.headImage = result.data.data.filePath
            newHeadImage.value = true
        }
    })
}

// 删除封面
const deleteImage = async () => {

    imageDelete(updateArticle.headImage).then(result => {
        if (result == null) {
            updateArticle.headImage = ""
            newHeadImage.value = false
        }
    })
}

// 提交文章
const submit = async () => {
    articleUpdate(updateArticle).then(result => {
        if (result == null) {
            router.go(-1)
        }
    })
}

const handleSelect = (index) => {
    switch (index) {
        case "1":
            showModalModal()
            break;
        case "2":
            router.go(-1)
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