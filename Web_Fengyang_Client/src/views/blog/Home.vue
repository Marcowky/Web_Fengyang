<template>
    <!-- 顶部导航栏 -->
    <TopBar :selected="selectedItemIndex" @select="handleSelect" />

    <!-- 侧边栏 -->
    <SideBar :items="menuItems" :selected="selectedItemIndex" @select="handleSelect" />

    <!-- 搜索栏 -->
    <div class="searchBar">
        <el-input class="searchBox" v-model="pageInfo.keyword" placeholder="请输入关键字" />
        <el-button class="searchButton" :icon="Search" @click="loadArticles(0)" circle />
    </div>

    <div class="mainContent">
        <!-- 文章卡片 -->
        <div v-for="(article, index) in articleList" style="margin:15px">
            <!-- 若有封面图 -->
            <el-card v-if="article.head_image" @click="toDetail(article)" style="cursor: pointer;" hoverable shadow="hover">
                <el-image style="width: 200px; float: left" :src="serverUrl + article.head_image" :fit="fit" />
                <template #header>
                    <div>
                        <text style="font-weight:bold; font-size: 20px;">{{ article.title }}</text>
                    </div>
                </template>
                <div style="position: absolute; left: 240px; width: 690px;">
                  <div v-html=article.content></div>
<!--                  <p>{{ article.content + "..." }}</p>-->
                    <div style="position: absolute; margin-top: 10px;">发布时间：{{ article.created_at }}
                    </div>
                </div>
            </el-card>
            <!-- 若无封面图 -->
            <el-card v-else @click="toDetail(article)" style="cursor: pointer;" hoverable shadow="hover">
                <template #header>
                    <div>
                        <text style="font-weight:bold; font-size: 20px;">{{ article.title }}</text>
                    </div>
                </template>
                <div style="height: 80px; ">
                  <div v-html=article.content></div>
<!--                    <p>{{ article.content + "..." }}</p>-->
                    <div style="position: absolute; margin-top: 10px;">发布时间：{{ article.created_at }}
                    </div>
                </div>
            </el-card>
        </div>
        <!-- 页面切换 -->
        <el-pagination class="pageSlider" :small="small" :background="background" layout="prev, pager, next"
            :page-count="pageInfo.pageCount" @current-change="loadArticles" />


    </div>
    <el-button class="publishButton" type="primary" :icon="Edit" size="large" circle @click="goPublish" />
</template>

<script setup>

import { ref, reactive, inject, onMounted } from 'vue'
// icons
import {
    Search,
    Edit
} from '@element-plus/icons-vue'

// 导入路由
import { useRouter } from 'vue-router'
const router = useRouter()

// 网络请求
const serverUrl = inject("serverUrl")
const axios = inject("axios")

// 变量初始化
const articleList = ref([])
const pageInfo = reactive({
    pageNum: 1,
    pageSize: 5,
    pageCount: 0,
    count: 0,
    keyword: "",
    categoryId: window.location.href.slice(-1) // 设置文章类别为地址最后一位
})

// 页面中侧边栏与导航栏的设置
// 1.导入侧边栏和顶部栏
import TopBar from "../../components/TopBar.vue"
import SideBar from "../../components/SideBar.vue"
// 2.导入导航栏路由函数
import { BarRouteGoto } from '../../components/BarRouteFunc.js'
// 3.导入菜单选项配置文件
import config from '../../config/config.json';
// 4.设置侧边栏目录项
const menuItems = config.menuItems.filter(item => item.index.startsWith("5-"));
// 5.导航栏和侧边栏已选选项
const selectedItemIndex = ref("5-" + window.location.href.slice(-1));
// 6.侧边栏和导航栏的点击触发函数
const handleSelect = (index) => {    // 这里可以触发路由跳转或其他操作
    selectedItemIndex.value = index
    BarRouteGoto(router, index) // 设置路由跳转
    pageInfo.categoryId = index[2] // 设置文章种类
    loadArticles(0)
};

// 挂载页面时触发
onMounted(() => {
    loadArticles()
})

// 按条件加载文章列表
const loadArticles = async (pageNum = 0) => {
    if (pageNum != 0) {
        pageInfo.pageNum = pageNum;
    }
    let res = await axios.post(`/article/list?keyword=${pageInfo.keyword}&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}&categoryId=${pageInfo.categoryId}`)
    if (res.data.code == 200) {
        articleList.value = res.data.data.article
    }
    pageInfo.count = res.data.data.count;
    pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
}

const goPublish = async () => {
    try {
        let resUser = await axios.get("user/info")
        if (resUser.data.code == 200) {
            router.push("/blog/publish")
        }
    } catch (err) {
        if (err.response.status === 401) {
            router.push("/user/login")
        }
    }
}

// 前往详情页
const toDetail = (article) => {
    router.push({
        path: "/blog/detail",
        query: {
            id: article.id,
            category: selectedItemIndex.value
        }
    })
}

</script>

<style lang="scss" scoped>
.searchBar {
    position: fixed;
    top: 13px;
    left: 35%;
    transform: translate(-50%, 0%);
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 5px;
    z-index: 9999;
}

.searchBox {
    width: 500px;
}

.mainContent {
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

.pageSlider {
    display: flex;
    justify-content: center;
}

.publishButton {
    position: fixed;
    right: 5%;
    bottom: 5%;
    font-size: 24px;
    box-shadow: 2px 2px 6px #D3D4D8;
}
</style>
