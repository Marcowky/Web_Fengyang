<template>
    <!-- 功能栏 -->
    <el-menu :default-active="activeIndex" class="choiceBar" @select="handleSelect">
        <el-menu-item style="color: #409EFF;" v-if="self" index="2">修改</el-menu-item>
        <el-menu-item style="color: #F56C6C;" v-if="self" index="3">删除</el-menu-item>
        <el-menu-item index="1">返回</el-menu-item>
    </el-menu>

    <div class="content">
        <div style="margin:15px">
            <!-- title -->
            <!-- <h1>{{ articleInfo.title }}</h1> -->
            <div style="height: 5px;"></div>
            <text style="font-size: 30px;">{{ articleInfo.title }}</text>

            <div style="height: 60px;  background-color: #FCFAF7;">
                <!-- 作者 -->
                <text style="position: absolute; left: 10px; top: 5px; color: #808080;">作者：{{ articleInfo.username }}
                </text>
                <!-- 发布时间 -->
                <text style="position: absolute; left: 10px; top: 30px; color: #808080;">发布时间：{{ articleInfo.created_at }}
                </text>
                <!-- 分类信息 -->
                <div style="position: absolute; right: 50px; top: 15px; color: #808080;">
                    文章分类：
                    <el-tag class="ml-2" type="success">{{ categoryName }}</el-tag>
                </div>
            </div>
            <!-- 分割线 -->
            <el-divider />
            <!-- 文章内容 -->
            <div :class="$style.article" class="article-content" v-html="articleInfo.content"></div>
            <div style="height: 20px;"></div>
        </div>

    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { userInfo, userBriefInfo } from '../../api/user';
import { articleDelete, articleDetail, articleCategories } from '../../api/article'
import { useRouter, useRoute } from 'vue-router' // 导入路由
import { ElMessageBox } from 'element-plus'

const router = useRouter()
const route = useRoute()
const articleInfo = ref({})
const categoryName = ref("")
const user = ref({})
const self = ref(false)
const activeIndex = ref('0')

// 加载文章种类
const categoryOptions = ref([])
const loadCategories = async () => {
    categoryOptions.value = articleCategories('/' + route.query.category.substring(0, route.query.category.length - 7))
}

// 加载文章
const loadArticle = async () => {
    articleDetail(route.query.category, route.query.id).then(result => {
        if (result != null) {
            articleInfo.value = result.data.data.article
            let label = categoryOptions.value.find((item) => item.value.endsWith(result.data.data.article.category_id)).label// 获取分类
            categoryName.value = label
            getUserInfo()
        }
    })
}

const getUserInfo = async () => {
    userBriefInfo(articleInfo.value.user_id).then(result => {// 获取作者信息
        if (result != null) articleInfo.value.username = result.data.data.userName
    })
    userInfo().then(result => {// 获取用户信息，判断用户是否是作者
        if (result != null) {
            user.value = result.data.data
            if (user.value.id == articleInfo.value.user_id) {
                self.value = true
            }
        }
    })
}

// 前往更新
const toUpdate = () => {
    router.push({
        path: "/article/update",
        query: {
            category: 'blogArticle',
            id: articleInfo.value.id,
        }
    })
}

// 删除文章
const toDelete = async () => {
    ElMessageBox.confirm(
        '是否要删除?',
        '警告',
        {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
        }
    )
        .then(async () => {
            articleDelete("blogArticle", articleInfo.value.id).then(result => {
                if (result == null) {
                    router.go(-1)
                }
            })
        })
        .catch()
}

const handleSelect = (index) => {
    switch (index) {
        case "1":
            router.go(-1)
            break;
        case "2":
            toUpdate()
            break;
        case "3":
            toDelete()
            break;
        default:
            break;
    }
}

// 挂载页面时触发
onMounted(() => {
    loadArticle()
    loadCategories()
})

</script>

<style lang="scss" scoped>
.content {
    position: relative;
    margin: auto;
    width: 70%;
    height: auto;
    background: white;
    box-shadow: 0 0 6px rgba(50, 50, 50, 0.26);
    color:black;
    border-radius: 10px;
    z-index: 99;
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

<style module>
.article img {
    max-width: 100%;
}
</style>
