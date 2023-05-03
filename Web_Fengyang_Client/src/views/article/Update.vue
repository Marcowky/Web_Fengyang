<template>
    <div class="topbar">
        <!-- 返回按键 -->
        <n-button @click="goback" strong quaternary round style="position: absolute; left: 50px; top: 7px; font-size: 24px;" color="#7B3DE0">
            <n-icon>
                <return-up-back />
            </n-icon>
        </n-button>
        <!-- 标题栏 -->
        <text style="position:absolute; left: 200px; line-height: 50px; color: #383838">标题</text>
        <n-input v-model:value="updateArticle.title" round placeholder="请输入标题" style="position:absolute; left: 265px; top: 8px; width: 1000px; background-color: #F3F0F9;" />
        <!-- 发布按钮 -->
        <div style="position: absolute; right: 50px; top: 8px">
            <n-button round color="#7B3DE0" @click="showModalModal">
                <template #icon>
                    <n-icon>
                        <send />
                    </n-icon>
                </template>
                发布
            </n-button>
        </div>     
    </div>
    <!-- 富文本编辑器 -->
    <div class="tabs">
        <n-card>
            <rich-text-editor v-if="loadOk" v-model:modelValue="updateArticle.content"></rich-text-editor>
        </n-card>
    </div>
    <!-- 发布弹窗 -->
    <n-modal v-model:show="showModal">
        <div style="width: 400px; height: 450px; background: white;">
            <n-card title="封面" :bordered="false" >
                <!-- 若无封面 -->
                <div v-if="!newHeadImage" style="width: 300px; height: 150px; margin: 0 auto;">
                    <n-upload
                    multiple
                    directory-dnd
                    :max="1"
                    @before-upload="beforeUpload"
                    :custom-request="customRequest"
                    >
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
                <!-- 若有封面 -->
                <div v-else style="width: 230px; margin: 0 auto;">
                    <n-image height="150" :src=serverUrl+updateArticle.headImage />    
                    <n-button @click="deleteImage" circle style="position: absolute; left: 298px; top: 50px;" color="#383838">
                        <template #icon>
                            <n-icon><close /></n-icon>
                        </template>
                    </n-button>                
                </div>
            </n-card>
            <!-- 分类选项 -->
            <n-card title="分类" :bordered="false">
                <div style="width:300px; margin: 0 auto;">
                    <n-select v-model:value="updateArticle.categoryId" :options="categoryOptions" placeholder="请选择分类"/>
                </div>
            </n-card>
            <!-- 取消和确认发布按钮 -->
            <div style="position: absolute; right: 100px; bottom: 30px;">
                <n-button type="default" @click="closeSubmitModal">
                    取消
                </n-button>
            </div>
            <div style="position: absolute; right: 30px; bottom: 30px;">
                <n-button type="primary" @click="submit">
                    确认
                </n-button>
            </div>                
        </div>
    </n-modal>

</template>

<script setup>
// icons
import { ArchiveOutline as ArchiveIcon } from "@vicons/ionicons5"
import { Send } from "@vicons/ionicons5"
import { ReturnUpBack } from "@vicons/ionicons5"
import { Close } from "@vicons/ionicons5"

import {ref,reactive,inject, onMounted} from 'vue'
// 富文本编辑器
import RichTextEditor from '../../components/RichTextEditor.vue'
// 导入路由
import {useRouter, useRoute} from 'vue-router'
const router = useRouter()
const route = useRoute()
// 网络请求
const serverUrl = inject("serverUrl")
const axios = inject("axios")
const message = inject("message")
// 变量初始化
const loadOk = ref(false)
const user = reactive({
    avatarUrl: "",
    id: 0
})
const categoryOptions = ref([])
const updateArticle = reactive({
    id: 0,
    categoryId: 0,
    title:"",
    content:"",
    headImage:"",
})
const showModal = ref(false)
const newHeadImage = ref(false)

// 挂载页面时触发
onMounted(() => {
    loadCategories()
    loadArticle()
})

// 加载分类
const loadCategories = async() =>{
    let res = await axios.get("/article/category")
    categoryOptions.value = res.data.data.categories.map((item)=>{
      return {
        label:item.name,
        value:item.id
      }
    })
}

// 加载文章
const loadArticle = async() => {
    let res = await axios.get("/article/" + route.query.id)
    
    if (res.data.code == 200) {
        updateArticle.categoryId = res.data.data.article.categoryId,
        updateArticle.title = res.data.data.article.title,
        updateArticle.content = res.data.data.article.content,
        updateArticle.headImage = res.data.data.article.headImage,
        newHeadImage.value = updateArticle.headImage? true: false    
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
const beforeUpload = async(data) => {
    if (data.file.file?.type !== "image/png") {
        message.error("只能上传png格式的图片")
        return false;
    }
    return true;
}

// 上传封面
const customRequest = async({file}) => {
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
const submit = async() => {
    let res = await axios.put("/article/" + route.query.id, {
        category_id: updateArticle.categoryId,
        title: updateArticle.title,
        content: updateArticle.content,
        head_image: updateArticle.headImage
    })
    
    if (res.data.code == 200) {
        message.success(res.data.msg) 
        goback()       
    } else {
        message.error(res.data.msg)
    }
}

// 返回上一级
const goback= () => {
    router.go(-2)    
}

</script>

<style lang="scss" scoped>
.topbar {
    position: sticky;
    top: 0;
    height: 50px;
    background: white;
    box-shadow: 0px 1px 5px #D3D4D8;
}
.tabs {
    position: absolute;
    top: 75px;
    left: 0;
    right: 0;
    margin: auto;
    width: 1000px;
    height: auto;
    background: white;  
    box-shadow: 0px 1px 3px #D3D4D8; 
    border-radius: 5px; 
}
</style>