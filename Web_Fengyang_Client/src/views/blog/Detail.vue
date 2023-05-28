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
                    <el-tag class="ml-2" type="success">{{categoryName}}</el-tag>
                </div>
            </div>
            <!-- 分割线 -->
            <el-divider />
            <!-- 文章内容 -->
            <div class="article-content" v-html="articleInfo.content"></div>
            <div style="height: 20px;"></div>
        </div>

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
import { ElMessage } from 'element-plus'
import { ElMessageBox } from 'element-plus'

// 定义变量
const articleInfo = ref({})
const categoryName = ref("")
const user = ref({})
const self = ref(false)
const activeIndex = ref('0')

// 挂载页面时触发
onMounted(() => {
    loadArticle()
    loadCategories()
})

import config from '../../config/config.json';
// 加载文章种类
const categoryOptions = ref([])
const loadCategories = async () => {
    categoryOptions.value = config.menuItems.filter(item => item.index.startsWith("/blog?category=")).map((item) => {
        return {
            label: item.label,
            value: item.index.slice(-3)
        }
    })
}

// 加载文章
const loadArticle = async () => {
    // 获取文章信息
    let resArticle = await axios.get("article/" + route.query.id)
    if (resArticle.data.code == 200) {
        articleInfo.value = resArticle.data.data.article
        // 获取分类
        let label = categoryOptions.value.find((item) => item.value.endsWith(resArticle.data.data.article.category_id)).label
            categoryName.value = label
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
        path: "/blog/update",
        query: {
            id: articleInfo.value.id,
        }
    })
}

// 删除文章
const toDelete = async (blog) => {

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
            let res = await axios.delete("article/" + articleInfo.value.id)
            if (res.data.code == 200) {
                ElMessage({
                    message: res.data.msg,
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
        })
        .catch()
}

// 回到上级页面
const goback = () => {
    router.go(-1)
}

const handleSelect = (index) => {

    switch (index) {
        case "1":
            goback()
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

</script>

<style lang="scss" scoped>
.content {
    position: relative;
    margin: auto;
    width: 1000px;
    height: auto;
    background: white;
    box-shadow: 2px 2px 6px #D3D4D8;
    border-radius: 10px;
    z-index: 99;
}

.choiceBar {
    position: fixed;
    top: 25%;
    z-index: 999;
    width: 150px;
    box-shadow: 2px 0 6px rgba(0, 0, 0, 0.26);
    border-radius: 0 10px 10px 0;
}
</style>
