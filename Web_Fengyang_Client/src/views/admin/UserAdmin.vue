<template>
    <userInfoDialog v-model="showDialog" :dialogTitle="dialogTitle" v-if="showDialog" @updateUserList="loadUsers"
        :dialogTableValue="dialogTableValue" :dialogUserType="dialogUserType" />
    <el-card class="mainTable">
        <el-row :gutter="20">
            <el-col :span="7">
                <el-input :input="loadUsers()" placeholder="请输入关键词" clearable v-model="pageInfo.keyword"
                    :prefix-icon="Search"></el-input>
            </el-col>
            <el-button @click="showAddDialog()"> 添加 </el-button>
        </el-row>
        <el-table ref="tableRef" :data="userList" style="width: 100%; margin-top: 20px" stripe
            @sort-change="handleSortChange">
            <template v-for="item in userConfig">
                <el-table-column v-if="item.sortable == 'true'" sortable="custom" :width="item.width" :prop="item.prop"
                    :label="item.label">
                    <template #default="scope">
                        <el-switch v-if="item.prop == 'Status'" v-model="scope.row.Status" @change=userUpdate(scope.row) />
                    </template>
                </el-table-column>
                <el-table-column v-if="item.sortable == 'false'" :width="item.width" :prop="item.prop" :label="item.label">
                    <template v-if="item.prop == 'option'" #default="scope">
                        <el-button type="warning" @click="showAddDialog(scope.row)">修改</el-button>
                        <el-button type="danger" @click="deleteUser(scope.row)">删除</el-button>
                    </template>
                    <template #default="scope">
                        <el-tag class="ml-2" v-if="item.prop == 'UserType'" type="success">{{ scope.row.UserType }}</el-tag>
                    </template>
                </el-table-column>
            </template>
        </el-table>
        <el-pagination style="margin-top: 20px" background layout="prev, pager, next" :page-count="pageInfo.pageCount"
            @current-change="loadUsers" />
    </el-card>
</template>
  
<script  setup>
import { ref, onMounted, reactive } from 'vue'
import { userConfig } from "../../config/adminConfig.json"
import { ElMessageBox } from 'element-plus'
import { userDelete, userUpdate, userListOut } from '../../api/user';
import { Search } from '@element-plus/icons-vue'
import { useRoute, onBeforeRouteUpdate } from 'vue-router'
import userInfoDialog from "./components/userInfoDialog.vue"

const route = useRoute()
const userList = ref([])
const pageInfo = reactive({
    pageNum: 1,
    pageSize: 10,
    pageCount: 0,
    count: 0,
    keyword: "",
    sortKey: 'created_at desc',
    userType: ''
})
const dialogTableValue = ref({})
const dialogTitle = ref('')
const dialogUserType = ref('')
const showDialog = ref(false)

const handleSortChange = (sort) => {
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
    loadUsers()
}

// 按条件加载文章列表
const loadUsers = async (pageNum = 0) => {
    if (pageNum != 0) {
        pageInfo.pageNum = pageNum;
    }
    userListOut(pageInfo).then(result => {
        if (result != null) {
            userList.value = result.data.data.user
            pageInfo.count = result.data.data.count;
        }
    })
    pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
    // console.log(userList.value)
}

const showAddDialog = (data) => {
    if (!data) {
        dialogTitle.value = "添加用户"
        dialogTableValue.value = {}
    }
    else {
        dialogTitle.value = "编辑用户"
        dialogTableValue.value = JSON.parse(JSON.stringify(data))
    }
    showDialog.value = true
    dialogUserType.value = pageInfo.userType
}

const deleteUser = async (data) => {
    ElMessageBox.confirm(
        '是否要删除?',
        '警告',
        {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
        }
    ).then(async () => {
        userDelete(data).then(result => {
            if (result == null) loadUsers()
        })
    }).catch()
}

// 挂载页面时触发
onMounted(() => {
    pageInfo.userType = route.query.category
})

// 4.设置路由守卫
onBeforeRouteUpdate((to, from) => {
    const fromCategory = from.query.category;
    const toCategory = to.query.category;

    if (fromCategory !== toCategory) {
        pageInfo.userType = toCategory
        loadUsers()
    }
})
</script>
  
<style scoped>
.mainTable {
    box-shadow: 0 0 6px rgba(50, 50, 50, 0.26);
    border-radius: 10px;
}
</style>