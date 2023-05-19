<template>
    <!-- 顶部导航栏 -->
    <TopBar :selected="selectedItemIndex" @select="handleSelect"/>
    
    <div class="tabs">
        <div style="margin:15px">
            <!-- title -->
            <h1>{{ articleInfo.title }}</h1>
            <div style="height: 60px;  background-color: #FCFAF7;">
                <!-- 作者 -->
                <text style="position: absolute; left: 40px; top: 90px; color: #808080;">作者：{{ articleInfo.username }}
                </text>
                <!-- 发布时间 -->
                <text style="position: absolute; left: 40px; top: 115px; color: #808080;">发布时间：{{ articleInfo.created_at }}
                </text>
                <!-- 分类信息 -->
                <div style="position: absolute; right: 50px; top: 95px; color: #808080;">
                    文章分类：
                    <el-tag class="ml-2" type="success">categoryName</el-tag>
                </div>
                <!-- 若为作者，则可以更新或删除文章 -->
                <el-button v-if="self" @click="toUpdate" ghost style="bottom: 45px; left: 805px;"
                    color="#7B3DE0">修改</el-button>
                <el-button v-if="self" @click="toDelete" ghost style="bottom: 45px; left: 815px;"
                    color="#7B3DE0">删除</el-button>
            </div>
            <!-- 分割线 -->
            <el-divider />

            <!-- 文章内容 -->
            <div class="article-content">
                <div v-html="articleInfo.content"></div>
            </div>
        </div>

    </div>
</template>

<script setup>
import { ref, inject, onMounted } from 'vue'

// 导入顶部栏
import TopBar from "../../components/TopBar.vue"

import {BarRouteGoto} from '../../components/BarRouteFunc.js'

// 导入路由
import { useRouter, useRoute } from 'vue-router'
const router = useRouter()
const route = useRoute()

// 网络请求
const axios = inject("axios")
const message = inject("message")
const dialog = inject("dialog")

// 定义变量
const articleInfo = ref({})
const categoryName = ref("")
const user = ref({})
const self = ref(false)

// 导入导航栏路由函数
import {BarRouteGoto} from '../../components/BarRouteFunc.js'
// 导航栏和侧边栏已选选项
const selectedItemIndex = ref("5-"+window.location.href.slice(-1));
// 侧边栏和导航栏的点击触发函数
const handleSelect = (index) => {    // 这里可以触发路由跳转或其他操作
    BarRouteGoto(router, index)
};

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
                message.info(res.data.msg)
                goback()
            } else {
                message.error(res.data.msg)
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
    top: 100px;
    left: 0;
    right: 0;
    margin: auto;
    width: 1000px;
    height: auto;
    background: white;
    box-shadow: 2px 2px 6px #D3D4D8;
    border-radius: 10px;
    z-index: 99;
}

.article-content img {
    max-width: 100% !important;
}</style>
