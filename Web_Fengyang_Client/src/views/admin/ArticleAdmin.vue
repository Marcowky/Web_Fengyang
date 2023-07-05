<template>
    <hotelInfoDialog v-model="showDialog" :dialogTitle="dialogTitle" v-if="showDialog" @updateHotelsList="loadArticles"
        :dialogTableValue="dialogTableValue" />
    <el-card class="mainTable">
        <el-row :gutter="20">
            <el-col :span="7">
                <el-input :input="loadArticles()" placeholder="请输入关键词" clearable v-model="pageInfo.keyword"
                    :prefix-icon="Search"></el-input>
            </el-col>
            <div v-if="showAdd">
                <el-button v-if="showButton" @click="toPublish()"> 发布 </el-button>
                <el-button v-else @click="showAddDialog()"> 发布 </el-button>
            </div>
        </el-row>
        <el-table ref="tableRef" :data="articleList" style="width: 100%; margin-top: 20px" stripe
            @sort-change="handleSortChange">
            <template v-for="item in articleConfig">
                <el-table-column v-if="item.label != 'id' && item.sortable == 'true'" sortable="custom" :width="item.width"
                    :prop="item.prop" :label="item.label">
                    <template v-if="item.prop == 'head_image'" #default="scope">
                        <el-tag class="ml-2" v-if="scope.row.head_image != ''"
                            :type="getColor(scope.row.head_image == '')">有</el-tag>
                        <el-tag class="ml-2" v-if="scope.row.head_image == ''"
                            :type="getColor(scope.row.head_image == '')">无</el-tag>
                    </template>
                    <template v-if="item.prop == 'category_id'" #default="scope">
                        <el-tag class="ml-2" :type="getColor(scope.row.category_id)">{{
                            getCategoryName(scope.row.category_id) }}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column v-if="item.label != 'id' && item.sortable == 'false'" :width="item.width" :prop="item.prop"
                    :label="item.label">
                    <template v-if="item.prop == 'option'" #default="scope">
                        <el-button v-if="showButton" type="warning" @click="toUpdate(scope.row)">修改</el-button>
                        <el-button v-else type="warning" @click="showAddDialog(scope.row)">修改</el-button>
                        <el-button type="danger" @click="deleteArticle(scope.row)">删除</el-button>
                    </template>

                </el-table-column>
            </template>
        </el-table>
        <el-pagination style="margin-top: 20px" background layout="prev, pager, next" :page-count="pageInfo.pageCount"
            @current-change="loadArticles" />
    </el-card>
</template>
  
<script  setup>
import { ref, onMounted, reactive } from 'vue'
import { articleConfig } from "../../config/adminConfig.json"
import config from "../../config/config.json"
import { ElMessageBox } from 'element-plus'
import hotelInfoDialog from "./components/hotelInfoDialog.vue"
import { articleDelete } from '../../api/article'
import { Search } from '@element-plus/icons-vue'
import { useRoute, useRouter, onBeforeRouteUpdate } from 'vue-router'
import { articleListOut } from '../../api/article'

const router = useRouter()
const route = useRoute()
const showAdd = ref()
const showButton = ref()
const dialogTitle = ref('')
const dialogTableValue = ref('')
const showDialog = ref(false)

// 这里是酒店增/改弹窗的触发函数,向dialogTablieValue注入值,并通过组件之间的传参给dialog
const showAddDialog = (data = false) => {
    if (!data) {
        dialogTitle.value = "添加酒店信息"
        console.log("添加酒店信息")
        dialogTableValue.value = { "test": 111 }
    }
    else {
        console.log("修改酒店信息")
        console.log(data)
        dialogTitle.value = "编辑酒店信息"
        dialogTableValue.value = JSON.parse(JSON.stringify(data))
    }
    showDialog.value = true
}

const handleSortChange = (sort) => {
    // 处理排序逻辑...
    switch (sort.prop) {
        case 'user_id':
            pageInfo.sortKey = 'user_id '
            break
        case 'category_id':
            pageInfo.sortKey = 'category_id '
            break
        case 'head_image':
            pageInfo.sortKey = 'head_image '
            break
        case 'created_at':
            pageInfo.sortKey = 'created_at '
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
    loadArticles()
}

// 4.设置路由守卫
onBeforeRouteUpdate((to, from) => {
    const fromCategory = from.query.category;
    const toCategory = to.query.category;

    if (fromCategory !== toCategory) {
        pageInfo.pageArticleType = toCategory
        showButton.value = (pageInfo.pageArticleType !== 'hotelarticle')
        showAdd.value = (pageInfo.pageArticleType != 'blogarticle')
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
    categoryId: "",
    pageArticleType: ""
})

const getCategoryName = (categoryId) => {
    if (pageInfo.pageArticleType != 'hotelarticle') {
        return config.menuItems.find(item => item.mainMenu == '/' + pageInfo.pageArticleType.substring(0, pageInfo.pageArticleType.length - 7) && item.index == categoryId).label
    }
    return "酒店记录"
}

const getColor = (categoryId) => {
    switch (categoryId) {
        case 1: return "success"
        case 2: return "primary"
        case 3: return "warning"
        case 4: return "danger"
        case 5: return "warning"
        case 6: return "info"
        case true: return "info"
        case false: return "success"
    }
}

// 按条件加载文章列表
const loadArticles = async (pageNum = 0) => {
    if (pageNum != 0) {
        pageInfo.pageNum = pageNum;
    }
    if (pageInfo.categoryId == "") {
        pageInfo.categoryId = 0
    }

    articleListOut(pageInfo).then(result => {
        if (result != null) {
            articleList.value = result.data.data.article
            pageInfo.count = result.data.data.count;
            pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
        }
    })

}

const toPublish = () => {
    router.push({
        path: "/admin/article/publish",
        query: {
            category: pageInfo.pageArticleType
        }
    })
}

const toUpdate = (data) => {
    router.push({
        path: "/admin/article/update",
        query: {
            category: pageInfo.pageArticleType,
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
            articleDelete(pageInfo.pageArticleType, data.id).then(result => {
                if (result == null) {
                    loadArticles()
                }
            })
        })
        .catch()

}

// 挂载页面时触发
onMounted(() => {
    pageInfo.pageArticleType = route.query.category
    showAdd.value = (pageInfo.pageArticleType != 'blogarticle')
    showButton.value = (pageInfo.pageArticleType != 'hotelarticle')
    loadArticles()
})
</script>

<style scoped>
.mainTable {
    box-shadow: 2px 2px 6px #D3D4D8;
    border-radius: 10px;
}
</style>
  