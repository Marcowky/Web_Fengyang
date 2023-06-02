<template>
    <userInfoDialog v-model="dialogVisible" />

    <el-card>
        <el-row :gutter="20">
            <el-col :span="7">
                <el-input placeholder="请输入关键词" clearable v-model="pageInfo.keyword"></el-input>
            </el-col>
            <el-button :icon="Search" @click="loadUsers()"> 搜索 </el-button>
            <el-button @click="showDialog"> 添加 </el-button>
        </el-row>
        <el-table ref="tableRef" :data="userList" style="width: 100%; margin-top: 20px" border
            @sort-change="handleSortChange">
            <template v-for="item in userConfig">
                <el-table-column v-if="item.sortable == 'true'" sortable="custom" :width="item.width" :prop="item.prop"
                    :label="item.label">
                    <template #default="scope">
                        <el-switch v-if="item.prop == 'Status'" v-model="scope.row.Status" @change=updateUser(scope.row) />
                    </template>
                </el-table-column>
                <el-table-column v-if="item.sortable == 'false'" :width="item.width" :prop="item.prop" :label="item.label">
                    <template v-if="item.prop == 'option'">
                        <el-button type="primary">Primary</el-button>
                        <el-button type="primary">Primary</el-button>
                    </template>
                </el-table-column>

            </template>

        </el-table>
        <el-pagination style="margin-top: 20px" background layout="prev, pager, next" :page-count="pageInfo.pageCount" @current-change="loadUsers" />
    </el-card>
</template>
  
<script  setup>
import { ref, inject, onMounted, reactive } from 'vue'
import { userConfig } from "../../config/adminConfig.json"
const axios = inject("axios")
import {
    Search,
} from '@element-plus/icons-vue'

// 挂载页面时触发
onMounted(() => {
    loadUsers()
})

const handleSortChange = (sort) => {
    // console.log('当前排序字段：', sort.prop);
    // console.log('当前排序方式：', sort.order);
    // 处理排序逻辑...
    switch(sort.prop){
        case 'ID':
            pageInfo.sortKey='id '
            break
        case 'CreatedAt':
            pageInfo.sortKey='created_at '
            break
        case 'Status':
            pageInfo.sortKey='status '
            break
    }
    switch(sort.order){
        case 'descending':
            pageInfo.sortKey=pageInfo.sortKey+'desc'
            break
        case 'ascending':
            pageInfo.sortKey=pageInfo.sortKey+'asc'
            break
        default:
            pageInfo.sortKey='created_at desc'
            break
    }
    // pageInfo.sortKey=sort.prop+' '+sort.order
    // console.log('当前排序query：', pageInfo.sortKey);
    loadUsers()
}


const userList = ref([])
const pageInfo = reactive({
    pageNum: 1,
    pageSize: 10,
    pageCount: 0,
    count: 0,
    keyword: "",
    sortKey: 'created_at desc'
})

import userInfoDialog from "./components/userInfoDialog.vue"
const dialogVisible = ref(false)

import { ElMessage } from 'element-plus'

// 按条件加载文章列表
const loadUsers = async (pageNum = 0) => {

    if (pageNum != 0) {
        pageInfo.pageNum = pageNum;
    }
    let res = await axios.get(`/user/list?userType=client&keyword=${pageInfo.keyword}&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}&order=${pageInfo.sortKey}`)
    if (res.data.code == 200) {
        userList.value = res.data.data.user
        // console.log(articleList)
    }
    pageInfo.count = res.data.data.count;
    pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
    console.log(userList.value)

}

const showDialog = () => {
    dialogVisible.value = true
}

const updateUser = async (updateUserInfo) => {
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
        showModal.value = false
    } else {
        ElMessage({
            message: res.data.msg,
            type: 'error',
            offset: 80
        })
    }
}
</script>
  