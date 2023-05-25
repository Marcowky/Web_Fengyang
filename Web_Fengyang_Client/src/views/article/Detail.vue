<template>
    <div class="tabs">
        <n-card>
            <!-- title -->
            <n-h1>{{ articleInfo.title }}</n-h1>
            <div style="height: 75px; background-color: #FCFAF7;">
                <!-- 作者 -->
                <text style="position: absolute; left: 40px; top: 100px; color: #808080;">作者：{{ articleInfo.username }}
                </text>
                <!-- 发布时间 -->
                <text style="position: absolute; left: 40px; top: 125px; color: #808080;">发布时间：{{ articleInfo.created_at }}
                </text>
                <!-- 分类信息 -->
                <div style="position: absolute; right: 50px; top: 97px; color: #808080;">
                    文章分类：
                    <n-tag type="warning">{{ categoryName }}</n-tag>
                </div>
                <!-- 若为作者，则可以更新或删除文章 -->
                <n-button v-if="self" @click="toUpdate" ghost style="bottom: 45px; left: 805px;"
                    color="#7B3DE0">修改</n-button>
                <n-button v-if="self" @click="toDelete" ghost style="bottom: 45px; left: 815px;"
                    color="#7B3DE0">删除</n-button>
            </div>
            <!-- 分割线 -->
            <n-divider />

            <!-- 文章内容 -->
            <div class="article-content">
                <div v-html="articleInfo.content"></div>
            </div>
        </n-card>
    </div>
</template>

<script setup>
import { ref, inject, onMounted } from 'vue'

// 导入路由
import { useRouter, useRoute } from 'vue-router'
const router = useRouter()
const route = useRoute()

// 网络请求
const axios = inject("axios")
const message = inject("message")
import { ElMessage } from 'element-plus'
const dialog = inject("dialog")

// 定义变量
const articleInfo = ref({})
const categoryName = ref("")
const user = ref({})
const self = ref(false)

// 挂载页面时触发
onMounted(() => {
    loadArticle()
})

// 加载文章
const loadArticle = async () => {
    // 获取文章信息
    let resArticle = await axios.get("article/" + route.query.id)
    if (resArticle.data.code == 200) {
        articleInfo.value = resArticle.data.data.article
        // 获取分类徐徐
        let resCategory = await axios.get("article/category/" + resArticle.data.data.article.category_id)
        if (resCategory.data.code == 200) {
            categoryName.value = resCategory.data.data.categoryName
        }
        // 获取作者信息
        let resWriter = await axios.get("user/briefInfo/" + articleInfo.value.user_id)
        articleInfo.value.username = resWriter.data.data.name
        // 获取用户信息，判断用户是否是作者
        let resUser = await axios.get("user/info")
        if (resUser.data.code == 200) {
            user.value = resUser.data.data
            if (user.value.id == articleInfo.value.user_id) {
                self.value = true
            }
        }
    }
}

// 前往更新
const toUpdate = () => {
    router.push({
        path: "/article/update",
        query: {
            id: articleInfo.value.id
        }
    })
}

// 删除文章
const toDelete = async (blog) => {
    dialog.warning({
        title: '警告',
        content: '是否要删除',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: async () => {
            let res = await axios.delete("article/" + articleInfo.value.id)
            if (res.data.code == 200) {
                ElMessage({
                    message: res.data.msg,
                    type: 'warning',
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
        },
        onNegativeClick: () => { }
    })
}

// 回到上级页面
const goback = () => {
    router.go(-1)
}

</script>

<style lang="scss" scoped>
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

.article-content img {
    max-width: 100% !important;
}</style>
