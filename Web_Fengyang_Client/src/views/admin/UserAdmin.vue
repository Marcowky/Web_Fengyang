<template>
    <el-card>
        <el-table ref="tableRef" :data="userList" style="width: 100%" border>
            <template v-for="item in userConfig">
                <el-table-column v-if="item.sortable == 'true'" sortable :width="item.width" :prop="item.prop"
                    :label="item.label">
                    <template #default="scope">
                        <el-switch v-if="item.prop == 'Status'" v-model="scope.row.Status" @change = updateUser(scope.row) />
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

    </el-card>
</template>
  
<script  setup>
import { ref, inject, onMounted } from 'vue'
import { userConfig } from "../../config/adminConfig.json"
const axios = inject("axios")

// 挂载页面时触发
onMounted(() => {
    loadUsers()
})

const userList = ref([])

import { ElMessage } from 'element-plus'

// 按条件加载文章列表
const loadUsers = async () => {
    let res = await axios.get(`/user/list?userType=client`)
    if (res.data.code == 200) {
        userList.value = res.data.data.user
        // console.log(articleList)
    }
    console.log(userList.value)

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


const tableData = [
    {
        id: '1',
        name: 'Tom',
        phone: '12345678901',
        createat: '2023-1-1',
        usertype: 'client',
        status: true
    },
    {
        id: '2',
        name: 'Marco',
        phone: '15984215111',
        createat: '2023-5-9',
        usertype: 'client',
        status: false
    },
    {
        id: '3',
        name: 'Jennie',
        phone: '13815238521',
        createat: '2022-12-19',
        usertype: 'client',
        status: false
    },
    {
        id: '4',
        name: 'Leo',
        phone: '13784291222',
        createat: '2023-9-20',
        usertype: 'client',
        status: true
    },
]


</script>
  