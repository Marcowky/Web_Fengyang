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
    <el-input class="searchButton" v-model="pageInfo.keyword" placeholder="请输入关键字" clearable :prefix-icon="Search" />
    <!-- 文章列表 -->
    <div class="content">
        <!-- 文章卡片 -->
        <div v-for="(article, index) in articleList"  class="articleCard">
            <!-- 若有封面图 -->
            <el-card v-if="article.head_image" @click="toDetail(article)" hoverable shadow="hover">
                <el-image class = "cardImage" 
                    :src="serverUrl + article.head_image" :fit="fit"/>
                <div style="position: relative; height: 150px;">
                    <div style=" margin-bottom: 10px; font-weight:bold; font-size: 20px;">{{ article.title }}</div>
                    <text>{{ cutContent(article.content) }}</text>
                    <div style=" position: absolute; bottom: 0px; color: gray;">发布时间：{{ article.created_at }}</div>
                </div>
            </el-card>
            <!-- 若无封面图 -->
            <el-card v-else @click="toDetail(article)" hoverable shadow="hover">
                <div style="position: relative; height: 120px;">
                    <div style=" margin-bottom: 10px; font-weight:bold; font-size: 20px;">{{ article.title }}</div>
                    <text>{{ cutContent(article.content) }}</text>
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
import { userInfo } from '../../api/user';
import { Search } from '@element-plus/icons-vue'// icons
import { useRouter, onBeforeRouteUpdate } from 'vue-router'// 导入路由
import { articleListOut } from '../../api/article'
import SideBar from "../../components/SideBar.vue" // 1.导入侧边栏
import config from '../../config/config.json' // 2.导入菜单选项配置文件
import LoginDialog from '../../components/LoginAndRegister.vue' // 导入登录注册弹框

const serverUrl = inject("serverUrl")// 网络请求
const router = useRouter()
const articleList = ref([])
const pageInfo = reactive({
    pageNum: 1,
    pageSize: 5,
    pageCount: 0,
    count: 0,
    keyword: "",
    categoryId: window.location.href.slice(-1), // 设置文章类别为地址最后一位
    pageArticleType: "blogArticle"
})
const activeIndex = ref('0')
const menuItems = config.menuItems.filter(item => item.mainMenu == "/blog"); // 设置侧边栏目录项
const loginDialogRef = ref(null)

// 按条件加载文章列表
const loadArticles = async (pageNum = 0) => {
    if (pageNum != 0) {
        pageInfo.pageNum = pageNum;
    }

    articleListOut(pageInfo).then(result => {
        if (result != null) {
            articleList.value = result.data.data.article
            pageInfo.count = result.data.data.count;
            pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
        }
    })
}

const cutContent = (content) => {
    if (content.length <= 100) {
        return content;
    } else {
        return content.substring(0, 100) + '...';
    }
}

const goPublish = async () => {
    userInfo().then(result => {
        if (result != null) {
            router.push({
                path: "/article/publish",
                query: {
                    category: pageInfo.pageArticleType
                }
            })
        }
        else {
            loginDialogRef.value.showDialog()
        }
    })
}


const toDetail = (article) => { // 前往详情页
    router.push({
        path: "/article/detail",
        query: {
            category: pageInfo.pageArticleType,
            id: article.id
        }
    })
}

onBeforeRouteUpdate((to, from) => { // 设置路由守卫,文章种类改变时重新加载文章
    const fromCategory = from.query.category;
    const toCategory = to.query.category;

    if (fromCategory !== toCategory) {
        pageInfo.categoryId = to.query.category // 设置文章种类
        loadArticles(0) // 加载文章
    }
});



onMounted(() => { // 挂载页面时加载文章
    loadArticles()
})

watch(() => pageInfo.keyword, () => ( // 搜索关键词改变时加载文章
    loadArticles()
), { deep: true, immediate: true })

</script>

<style lang="scss" scoped>
.searchButton {
    position: fixed;
    right: 13%;
    top: 33px;
    z-index: 999;
    width: 150px;
}

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

.articleCard {
    cursor: pointer;
    width: 100%;
    padding-left: 15px;
    padding-right: 15px;
    padding-top: 15px;
    // padding-bottom: 5px;
}

.pageSlider {
    display: flex;
    justify-content: center;
    padding: 10px;
}

.choiceBar {
    width: 150px;
    box-shadow: 0 0 6px rgba(50, 50, 50, 0.26);
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

.cardImage {
    height: 150px; 
    width: 250px;  
    float: right; 
    margin-bottom: 20px; 
    margin-left: 15px;
    border-radius: 5px 5px 5px 5px;
}
</style>
