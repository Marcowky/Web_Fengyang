<template>
    <!-- 顶部导航栏 -->
    <div class="topbar">
        <!-- 返回按钮 -->
        <!-- 这段代码是用来创建一个按钮的，其中 strong 是组件的一个属性，表示按钮是否为粗体。而 quaternary 是指按钮的颜色类型，表示第四种颜色（在这里是紫色）。round 意味着按钮的边角是圆形的。最后，style 属性允许您通过CSS样式来自定义按钮外观。在这个例子中，position: absolute; left: 50px; top: 7px; font-size: 24px; 是一些CSS样式，表示将按钮的位置设置为绝对位置，从左侧和顶部分别偏移50像素和7像素，并将字体大小设置为24像素。 -->
        <n-button @click="goback" strong quaternary round style="position: absolute; left: 50px; top: 7px; font-size: 24px;"
            color="#7B3DE0">
            <n-icon>
                <return-up-back />
            </n-icon>
        </n-button>
        <!-- 标题 -->
        <text style="position:absolute; left: 200px; line-height: 50px; color: #383838">标题</text>
        <!-- 输入框 -->
        <n-input v-model:value="addArticle.title" round placeholder="请输入标题"
            style="position:absolute; left: 265px; top: 8px; width: 1000px; background-color: #F3F0F9;" />
        <!-- 用户头像 -->
        <!-- <n-avatar round size="medium" :src=user.avatarUrl style="position: absolute; right: 190px; top: 8px; " /> -->
        <!-- 发布文章按钮 -->
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

    <!-- 富文本编辑器主体内容区域 -->
    <div class="tabs">
        <n-card>
            <rich-text-editor v-model:modelValue="addArticle.content"></rich-text-editor>
        </n-card>
    </div>

    <!-- 发布文章时弹出的模态框 -->
    <n-modal v-model:show="showModal">
        <div style="width: 400px; height: 450px; background: white;">
            <!-- 封面图片上传 -->
            <n-card title="封面" :bordered="false">
                <div v-if="!newHeadImage" style="width: 300px; height: 150px; margin: 0 auto;">
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
                <div v-else style="width: 230px; margin: 0 auto;">
                    <n-image height="150" :src=serverUrl+addArticle.headImage />
                    <n-button @click="deleteImage" circle style="position: absolute; left: 298px; top: 50px;"
                        color="#383838">
                        <template #icon>
                            <n-icon>
                                <close />
                            </n-icon>
                        </template>
                    </n-button>
                </div>
            </n-card>
            <!-- 分类选择 -->
            <n-card title="分类" :bordered="false">
                <div style="width:300px; margin: 0 auto;">
                    <n-select v-model:value="addArticle.categoryId" :options="categoryOptions" placeholder="请选择分类" />
                </div>
            </n-card>
            <!-- 取消按钮 -->
            <div style="position: absolute; right: 100px; bottom: 30px;">
                <n-button type="default" @click="closeSubmitModal">
                    取消
                </n-button>
            </div>
            <!-- 确认按钮 -->
            <div style="position: absolute; right: 30px; bottom: 30px;">
                <n-button type="primary" @click="submit">
                    确认
                </n-button>
            </div>
        </div>
    </n-modal>
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
import { Send } from "@vicons/ionicons5"
import { ReturnUpBack } from "@vicons/ionicons5"
import { Close } from "@vicons/ionicons5"

// 导入路由
import { useRouter } from 'vue-router'
const router = useRouter()
// const route = useRoute()


// 这也是在Vue.js 3中使用的代码，它使用了inject函数来获取从祖先组件中通过provide提供的三个依赖项。
// serverUrl：这是一个服务器地址的字符串，用于向该地址发送HTTP请求。该值是通过在某个祖先组件中使用provide("serverUrl", serverUrl)提供的。在当前组件中，我们可以使用inject("serverUrl")来访问它。
// axios：这是一个类似于fetch的HTTP客户端工具，用于发送HTTP请求。同样地，这个值也是通过在祖先组件中使用provide("axios", axios)提供的。
// message：这是用于显示用户友好的消息的工具。在祖先组件中，我们可以使用provide("message", message)提供这个值，并在当前组件中使用inject("message")来访问它。
// 通过使用inject和provide，我们可以轻松地实现依赖注入，同时避免了深度嵌套的属性访问和传递。
const serverUrl = inject("serverUrl")
const axios = inject("axios")
const message = inject("message")

// const login = ref(false)
// const user = reactive({
//     avatarUrl: "",
//     id: 0
// })
const categoryOptions = ref([])// 分类列表选项
const addArticle = reactive({// 待发布的文章对象
    id: 0,
    categoryId: 0,
    title: "",
    content: "",
    headImage: "",
})


onMounted(() => {
    // loadAvatar()
    loadCategories()
})

// 加载头像
// const loadAvatar = async () => {
//     let res = await axios.get("/user")
//     console.log(res)
//     if (res.data.code == 200) {
//         user.avatarUrl = serverUrl + res.data.data.avatar
//         user.id = res.data.data.id
//         login.value = true
//     }
// }

// 加载文章种类
const loadCategories = async () => {
    let res = await axios.get("/category")
    console.log(res)
    categoryOptions.value = res.data.data.categories.map((item) => {
        return {
            label: item.name,
            value: item.id
        }
    })
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
const beforeUpload = async (data) => {
    if (data.file.file?.type !== "image/png") {
        message.error("只能上传png格式的图片")
        return false;
    }
    return true;
}

// 控制文章头图
const newHeadImage = ref(false)
const customRequest = async ({ file }) => {
    const formData = new FormData()
    formData.append('file', file.file)
    let res = await axios.post("/image/upload", formData)
    console.log(res)
    addArticle.headImage = res.data.data.filePath
    console.log(addArticle.headImage)
    newHeadImage.value = true
}
const deleteImage = () => {
    addArticle.headImage = ""
    newHeadImage.value = false
}

// 上传文章函数
const submit = async () => {
    let res = await axios.post("/article", {
        category_id: addArticle.categoryId,
        title: addArticle.title,
        content: addArticle.content,
        head_image: addArticle.headImage
    })
    console.log(res)
    if (res.data.code == 200) {
        message.success(res.data.msg)
        goback()
    } else {
        message.error(res.data.msg)
    }
}

// 返回上级页面
const goback = () => {
    router.go(-1)
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
