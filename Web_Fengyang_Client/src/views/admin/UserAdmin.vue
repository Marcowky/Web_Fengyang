<template>
    <el-card>
        <el-table ref="tableRef" :data="tableData" style="width: 100%" border>
            <template v-for="item in user">
                <el-table-column v-if="item.sortable == 'true'" sortable :width="item.width" :prop="item.prop"
                    :label="item.label">
                    <template #default="scope">
                        <el-switch v-if="item.prop == 'status'" v-model="scope.row.status"/>
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
import { ref } from 'vue'
import { user } from "../../config/adminConfig.json"




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
  