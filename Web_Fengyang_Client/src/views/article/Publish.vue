<template>
    <!-- 富文本编辑器 -->
    <div class="content">
        <div style="padding:15px; width: 100%;">
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
                <el-button style="margin-right: 20px;" type="primary" @click="submit">确认</el-button>
                <el-button type="danger" @click="closeSubmitModal">取消</el-button>
            </div>
        </div>
    </el-dialog>
</template>

<script setup>
import { ref, reactive, inject, onMounted } from 'vue'
import { imageUpload, imageDelete, imageCheck } from '../../api/image'
import RichTextEditor from '../../components/RichTextEditor.vue' // 导入富文本编辑器
import { UploadFilled, Delete } from '@element-plus/icons-vue'
import { useRoute, useRouter } from 'vue-router' // 导入路由
import { articlePost, articleCategories } from '../../api/article'

const router = useRouter()
const route = useRoute()
const serverUrl = inject("serverUrl")
const newHeadImage = ref(false) // 控制文章头图
const categoryOptions = ref([])// 分类列表选项
const showModal = ref(false)
const addArticle = reactive({// 待发布的文章对象
    id: 0,
    categoryId: "",
    title: "",
    content: "",
    headImage: "",
    articleType: route.query.category
})

// 加载文章种类
const loadCategories = async () => {
    categoryOptions.value = articleCategories('/' + addArticle.articleType.substring(0, addArticle.articleType.length - 7))
}

// 控制发布文章时弹窗的显示与隐藏
const showModalModal = () => {
    showModal.value = true
}
const closeSubmitModal = () => {
    showModal.value = false
}

const customRequest = async (file) => {
    imageUpload(file).then(result => {
        if (result != null) {
            addArticle.headImage = result.data.data.filePath
            newHeadImage.value = true
        }
    })
}

const deleteImage = async () => {
    imageDelete(addArticle.headImage).then(result => {
        if (result == null) {
            addArticle.headImage = ""
            newHeadImage.value = false
        }
    })
}

// 上传文章
const submit = async () => {
    articlePost(addArticle).then(result => {
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

onMounted(() => {
    loadCategories()
})
</script>

<style lang="scss" scoped>
.content {
    position: relative;
    margin: auto;
    width: 70%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    background: white;
    box-shadow: 0 0 6px rgba(50, 50, 50, 0.26);
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
    box-shadow: 0 0 6px rgba(50, 50, 50, 0.26);
}

.choiceBar {
    position: fixed;
    top: 25%;
    z-index: 999;
    width: 150px;
    box-shadow: 0 0 6px rgba(50, 50, 50, 0.26);
    border-radius: 0 10px 10px 0;
}
</style>
