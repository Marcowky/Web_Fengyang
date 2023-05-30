<template>
    <!-- 登录注册弹框 -->
    <LoginDialog ref="loginDialogRef" />
    <!-- 功能栏 -->
    <div class="funcBar">
        <!-- 侧边栏 -->
        <SideBar class="sidebar" :items="menuItems" />
        <el-menu class="choiceBar" :default-active="activeIndex" @select="goPublish" text-color="#409EFF">
            <el-menu-item style="color: #409EFF;" index="1">我要发布!</el-menu-item>
        </el-menu>
    </div>
    <!-- 搜索栏 -->
    <div class="searchButton">
        <el-popover placement="bottom" title="搜索文章" :width="200" trigger="click">
            <template #reference>
                <el-button @click="loadArticles(0)" text bg size="large">
                    <el-icon style="position: relative; top: -2px;" :size="20">
                        <Search />
                    </el-icon>
                </el-button>
            </template>
            <el-input class="searchBox" v-model="pageInfo.keyword" placeholder="请输入关键字" />
        </el-popover>
    </div>

    <!-- 文章列表 -->
    <div class="content">
        <!-- 文章卡片 -->
        <div v-for="(article, index) in articleList" style="margin:15px">
            <!-- 若有封面图 -->
            <el-card class="articleCard" v-if="article.head_image" @click="toDetail(article)" hoverable shadow="hover">
                <el-image style="height: 150px; float: right; margin-bottom: 20px; margin-left: 15px;" :src="serverUrl + article.head_image" />
                <div style="position: relative; height: 150px;">
                    <div style=" margin-bottom: 10px; font-weight:bold; font-size: 20px;">{{ article.title }}</div>
                    <div v-html="article.content"></div>
                    <div style=" position: absolute; bottom: 0px; color: gray;">发布时间：{{ article.created_at }}</div>
                </div>
            </el-card>
            <!-- 若无封面图 -->
            <el-card class="articleCard" v-else @click="toDetail(article)" hoverable shadow="hover">
                <div style="position: relative; height: 120px;">
                    <div style=" margin-bottom: 10px; font-weight:bold; font-size: 20px;">{{ article.title }}</div>
                    <div v-html="article.content"></div>
                    <div style=" position: absolute; bottom: 0px; color: gray;">发布时间：{{ article.created_at }}</div>
                </div>
            </el-card>
        </div>
        <!-- 页面切换 -->
        <el-pagination class="pageSlider" small layout="prev, pager, next" :page-count="pageInfo.pageCount"
            @current-change="loadArticles" />
    </div>
</template>




<script setup>
import { ref, reactive, inject, onMounted, watch } from 'vue'
// icons
import {
    Search
} from '@element-plus/icons-vue'

// 导入路由
import { useRoute, useRouter } from 'vue-router'
const router = useRouter()
const route = useRoute()

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
const activeIndex = ref('0')

// 页面中侧边栏与导航栏的设置
// 1.导入侧边栏
import SideBar from "../../components/SideBar.vue"
// 2.导入菜单选项配置文件
import config from '../../config/config.json';
// 3.设置侧边栏目录项
const menuItems = config.menuItems.filter(item => item.index.startsWith("/blog?category=5-"));
// 4.设置路由跳转监听
watch(
    () => route.fullPath,
    (newValue, oldValue) => {
        pageInfo.categoryId = newValue.charAt(newValue.length - 1) // 设置文章种类
        loadArticles(0) // 加载文章
    }
)

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

// 导入登录注册弹框
import LoginDialog from '../../components/LoginAndRegister.vue'
const loginDialogRef = ref(null)

const goPublish = async () => {
    try {
        let resUser = await axios.get("user/info")
        if (resUser.data.code == 200) {
            router.push("/blog/publish")
        }
    } catch (err) {
        if (err.response.status === 401) {
            // 显示登录注册弹框
            loginDialogRef.value.showDialog()
        }
    }
}

// 前往详情页
const toDetail = (article) => {
    router.push({
        path: "/blog/detail",
        query: {
            id: article.id,
        }
    })
}

</script>

<style lang="scss" scoped>
.searchButton {
    position: fixed;
    display: block;
    left: 50%;
    transform: translateX(-50%);
    top: 10px;
    z-index: 999;
}

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

.articleCard {
    cursor: pointer;
    width: 950px;
}

.pageSlider {
    display: flex;
    justify-content: center;
}

.choiceBar {
    width: 150px;
    box-shadow: 2px 0 6px rgba(0, 0, 0, 0.26);
    border-radius: 0 10px 10px 0;
    margin-top: 10%;
}

.sidebar {
    position: relative;
    top: 0%;
}

.funcBar {
    position: fixed;
    top: 25%;
    z-index: 999;
}
</style>
