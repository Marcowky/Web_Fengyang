<template>
    <el-card>
        <el-row :gutter="20">
            <el-col :span="7">
                <el-input placeholder="请输入关键词" clearable v-model="pageInfo.keyword"></el-input>
            </el-col>
            <el-button :icon="Search" @click="loadArticles"> 搜索 </el-button>
            <el-button @click="toPublish()"> 发布 </el-button>
        </el-row>
        <el-table ref="tableRef" :data="articleList" style="width: 100%; margin-top: 20px" border
            @sort-change="handleSortChange">
            <template v-for="item in articleConfig">
                <el-table-column v-if="item.label!='id' && item.sortable == 'true'" sortable="custom" :width="item.width" :prop="item.prop"
                    :label="item.label">
                    <template v-if="item.prop == 'head_image'" #default="scope">
                        <el-tag class="ml-2" v-if="scope.row.head_image!=''" type="success">有</el-tag>
                        <el-tag class="ml-2" v-if="scope.row.head_image==''" type="success">无</el-tag>
                    </template>
                </el-table-column>
                <el-table-column v-if="item.label!='id' && item.sortable == 'false'" :width="item.width" :prop="item.prop" :label="item.label">
                    <template v-if="item.prop == 'option'" #default="scope">
                        <el-button type="primary" @click="toUpdate(scope.row)">修改</el-button>
                        <el-button type="primary" @click="deleteArticle(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </template>
        </el-table>
        <el-pagination style="margin-top: 20px" background layout="prev, pager, next" :page-count="pageInfo.pageCount"
            @current-change="loadArticles" />
    </el-card>
</template>
  
<script  setup>
import { ref, inject, onMounted, reactive } from 'vue'
import { articleConfig } from "../../config/adminConfig.json"
import { ElMessage, ElMessageBox } from 'element-plus'

const axios = inject("axios")
import {
    Search,
} from '@element-plus/icons-vue'
import { useRoute, useRouter, onBeforeRouteUpdate } from 'vue-router'
const pageArticleType = ref('')
const router = useRouter()
const route = useRoute()

// 挂载页面时触发
onMounted(() => {
    pageArticleType.value = route.query.category;
    loadArticles()
})

const handleSortChange = (sort) => {
    // console.log('当前排序字段：', sort.prop);
    // console.log('当前排序方式：', sort.order);
    // 处理排序逻辑...
    switch (sort.prop) {
        case 'ID':
            pageInfo.sortKey = 'id '
            break
        case 'CreatedAt':
            pageInfo.sortKey = 'created_at '
            break
        case 'Status':
            pageInfo.sortKey = 'status '
            break
    }
    switch (sort.order) {
        case 'descending':
            pageInfo.sortKey = pageInfo.sortKey + 'desc'
            break
        case 'ascending':
            pageInfo.sortKey = pageInfo.sortKey + 'asc'
            break
        default:
            pageInfo.sortKey = 'created_at desc'
            break
    }
    // pageInfo.sortKey=sort.prop+' '+sort.order
    // console.log('当前排序query：', pageInfo.sortKey);
    loadArticles()
}

// 4.设置路由守卫
onBeforeRouteUpdate((to, from) => {
    const fromCategory = from.query.category;
    const toCategory = to.query.category;

    if (fromCategory !== toCategory) {
        pageArticleType.value = toCategory
        loadArticles()
    }
});


const articleList = ref([])
const pageInfo = reactive({
    pageNum: 1,
    pageSize: 10,
    pageCount: 0,
    count: 0,
    keyword: "",
    sortKey: 'created_at desc',
    categoryId: ""
})

// 按条件加载文章列表
const loadArticles = async (pageNum = 0) => {
    if (pageNum != 0) {
        pageInfo.pageNum = pageNum;
    }
    if (pageInfo.categoryId == "") {
        pageInfo.categoryId = 0
    }
    let res = await axios.get(`/article/list?articleType=${pageArticleType.value}&keyword=${pageInfo.keyword}&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}&categoryId=${pageInfo.categoryId}`)
    if (res.data.code == 200) {
        articleList.value = res.data.data.article
        // console.log(articleList)
    }
    pageInfo.count = res.data.data.count;
    pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
    console.log(articleList.value)
}

const toPublish = () => {
    router.push({
        path: "/admin/article/publish",
        query: {
            category: pageArticleType.value
        }
    })
}

const toUpdate = (data) => {
    router.push({
        path: "/admin/article/update",
        query: {
            category: pageArticleType.value,
            id: data.id
        }
    })
}


const deleteArticle = async (data) => {

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
            let res = await axios.delete(`article/delete?articleType=${pageArticleType.value}&id=${data.id}`)
            if (res.data.code == 200) {
                ElMessage({
                    message: res.data.msg,
                    offset: 80
                })
                loadArticles()
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

const updateArticle = async (updateUserInfo) => {
    console.log(updateUserInfo.userType)
    let res = await axios.put("/user/update", {
        ID: updateUserInfo.ID,
        UserName: updateUserInfo.UserName,
        PhoneNumber: updateUserInfo.PhoneNumber,
        userType: updateUserInfo.UserType,
        Status: updateUserInfo.Status
    })
    if (res.data.code == 200) {
        ElMessage({
            message: res.data.msg,
            type: 'success',
            offset: 80
        })
    } else {
        ElMessage({
            message: res.data.msg,
            type: 'error',
            offset: 80
        })
    }
}
</script>
  