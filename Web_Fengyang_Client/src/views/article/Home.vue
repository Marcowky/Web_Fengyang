<template>
    <!-- 顶栏 -->
    <div class="topbar">
        <!-- 按分类检索 -->
        <n-popselect @update:value="searchByCategory" v-model:value="selectedCategory" :options="categoryOptions" trigger="click">
            <n-button text style="position:absolute; left: 50px; top: 22px; font-size: 18px;">{{categoryName}}</n-button>
        </n-popselect>
        
        <!-- 搜索框 -->
        <n-input v-model:value="pageInfo.keyword" round placeholder="请输入关键字" style="position:absolute; left: 250px; top: 8px; width: 920px; background-color: #F3F0F9;" />

        <!-- 搜索按钮 -->
        <n-button @click="loadArticles(0)" round color="#7B3DE0" style="position:absolute; right: 250px; top: 8px;">
            <template #icon>
                <n-icon>
                    <search />
                </n-icon>
            </template>
        搜索
        </n-button>
    </div>   

    <div class="tabs">
        <n-card>
            <!-- 文章卡片 -->
            <div v-for="(article,index) in articleList" style="margin-bottom:15px">
                <!-- 若有封面图 -->
                <n-card v-if="article.head_image" @click="toDetail(article)" style="cursor: pointer;" hoverable >
                    <n-image width="200" :src=serverUrl+article.head_image style="float: left" />
                    <div style="position: absolute; left: 240px; width: 690px;">
                        <text style="font-weight:bold; font-size: 20px;">{{article.title}}</text>
                        <p>{{article.content+"..."}}</p>
                        <div style="position: absolute; margin-top: 10px;">发布时间：{{article.created_at}}</div>
                    </div>
                </n-card>
                <!-- 若无封面图 -->
                <n-card v-else @click="toDetail(article)" style="cursor: pointer;" hoverable >
                    <div style="height: 140px; ">
                        <text style="font-weight:bold; font-size: 20px;">{{article.title}}</text>
                        <p>{{article.content+"..."}}</p>
                        <div style="position: absolute; margin-top: 10px;">发布时间：{{article.created_at}}</div>                        
                    </div>
                </n-card>
            </div>
            <!-- 换页按钮 -->
            <n-pagination @update:page="loadArticles" v-model:page="pageInfo.pageNum" :page-count="pageInfo.pageCount" />
        </n-card>
    </div>

</template>

<script setup>
import {ref,reactive,inject,onMounted,computed} from 'vue'
// icons
import {Search} from '@vicons/ionicons5'

// 导入路由
import {useRouter } from 'vue-router'
const router = useRouter()

// 网络请求
const serverUrl = inject("serverUrl")
const axios = inject("axios")

// 变量初始化
const selectedCategory = ref(0)
const categoryOptions = ref([])
const articleList = ref([])
const pageInfo = reactive({
  pageNum:1,
  pageSize:5,
  pageCount:0,
  count:0,
  keyword:"",
  categoryId:0
})

// 挂载页面时触发
onMounted(()=>{
    loadArticles()
    loadCategories()
})

// 加载文章列表
const loadArticles = async(pageNum = 0) =>{
    if (pageNum != 0){
        pageInfo.pageNum = pageNum;
    }
    let res = await axios.post(`/article/list?keyword=${pageInfo.keyword}&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}&categoryId=${pageInfo.categoryId}`)
    if (res.data.code == 200) {
        articleList.value = res.data.data.article
    }
    pageInfo.count = res.data.data.count;
    pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
}

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

// 获取分类名字
const categoryName = computed(() => {
    let selectedOption = categoryOptions.value.find((option) => {return option.value == selectedCategory.value})
    return selectedOption ? selectedOption.label : ""
})

// 分类搜索
const searchByCategory = (categoryId) => {
    pageInfo.categoryId = categoryId
    pageInfo.pageNum = 1
    loadArticles()
}

// 前往详情页
const toDetail = (article) => {
    router.push({
        path: "/article/detail",
        query: {
            id: article.id
        }
    }) 
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
    top: 150px;
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
