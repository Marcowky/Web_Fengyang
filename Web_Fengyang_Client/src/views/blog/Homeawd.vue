<template>
    <div class="box"><!-- 主容器 -->
        <div class="mainContent">
        </div>
        <div class="footer"></div>
    </div>
</template>

<script setup>
import { ref, reactive, inject, onMounted } from 'vue'
// icons
import {
    Search
} from '@element-plus/icons-vue'

// 导入底部栏
import FooterBar from "../../components/FooterBar.vue"

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
const activeIndex = ref('0')

// 页面中侧边栏与导航栏的设置
// 1.导入侧边栏和顶部栏
import TopBar from "../../components/TopBar.vue"
import SideBar from "../../components/SideBar.vue"
// 2.导入菜单选项配置文件
import config from '../../config/config.json';
// 3.设置侧边栏目录项
const menuItems = config.menuItems.filter(item => item.index.startsWith("/blog?category=5-"));
// 4.侧边栏和导航栏的点击触发函数
const handleSelect = (index) => {
    pageInfo.categoryId = index.charAt(index.length - 1) // 设置文章种类
    loadArticles(0) // 加载文章
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
    background-color: blue;
    z-index: 9999;
    height: 100px;
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

.box {
    position: relative;
    min-height: 100%;
}

.footer {
    position: relative;
    background-color: yellow;
    bottom: 0px;
    height: 100px;
    z-index: 1000;
    width: 100%;
}
</style>
