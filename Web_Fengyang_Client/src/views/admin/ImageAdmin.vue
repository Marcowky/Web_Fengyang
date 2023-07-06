<template>
  <div class="carousel-container">
    <carouselInfoDialog v-model="showDialog" :dialogTitle="dialogTitle" v-if="showDialog" @updateCarouselList="loadCarousels"
                    :dialogTableValue="dialogTableValue" :dialogCarouselType="dialogCarouselType"/>
    <!-- 轮播图管理部分 -->
    <el-card class="mainTable">
      <el-row :gutter="20">
        <el-button @click="showAddCarouselDialog()">
          添加
        </el-button>
      </el-row>
      <el-table
          :data="carouselList"
          style="width: 100%; margin-top: 20px"
          stripe
          ref="CarouselTableRef"
          @sort-change="handleSortChange"
      >
        <template v-for="item in carouselConfig">
          <el-table-column v-if="item.sortable == 'true'" sortable="custom" :width="item.width" :prop="item.prop"
                           :label="item.label">
          </el-table-column>
          <el-table-column v-if="item.sortable == 'false'" :width="item.width" :prop="item.prop" :label="item.label">
            <template v-if="item.prop == 'option'" #default="scope">
              <el-button type="warning" @click="showAddCarouselDialog(scope.row)">修改</el-button>
              <el-button type="danger" @click="deleteCarousel(scope.row)">删除</el-button>
            </template>

            <template #default="scope">
              <el-tag class="ml-2" v-if="item.prop == 'Category'" type="success">{{ scope.row.Category }}</el-tag>
            </template>

            <template v-if="item.prop == 'Url'" #default="scope">
              <el-image style="height: 150px; float: right; margin-bottom: 20px; margin-left: 15px;"
                        :src="serverUrl + scope.row.Url"/>
            </template>
          </el-table-column>
        </template>

      </el-table>
      <el-pagination style="margin-top: 20px" background layout="prev, pager, next"
                     :page-count="pageInfo.pageCount" @current-change="loadCarousels"
      />
    </el-card>

  </div>
</template>

<script  setup>
import { ref, inject, onMounted, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { carouselConfig } from "../../config/adminConfig.json"
import { useRoute, onBeforeRouteUpdate } from 'vue-router'


const serverUrl = inject("serverUrl")// 网络请求
// const axios = inject("axios")
const carouselList = ref([])
const route = useRoute()
const pageInfo = reactive({
  pageNum: 1,
  pageSize: 10,
  pageCount: 0,
  count: 0,
  // keyword: "",
  sortKey: 'Iorder desc',
  category: ''
})
const dialogTableValue = ref({})
const dialogTitle = ref('')
const dialogCarouselType = ref('')
const showDialog = ref(false)

// 挂载页面时触发
onMounted(() => {
  pageInfo.category = route.query.category;
  loadCarousels()
})

const handleSortChange = (sort) => {
  // console.log('当前排序字段：', sort.prop);
  // console.log('当前排序方式：', sort.iorder);
  // 处理排序逻辑...
  switch (sort.prop) {
    case 'ID':
      pageInfo.sortKey = 'id '
      break
    case 'Iorder':
      pageInfo.sortKey = 'Iorder '
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
      pageInfo.sortKey = 'Iorder desc'
      break
  }
  // pageInfo.sortKey=sort.prop+' '+sort.iorder
  // console.log('当前排序query：' , pageInfo.sortKey);
  loadCarousels()
}

// 4.设置路由守卫
onBeforeRouteUpdate((to, from) => {
  const fromCategory = from.query.category;
  const toCategory = to.query.category;

  if (fromCategory !== toCategory) {
    pageInfo.category = toCategory
    loadCarousels()
  }
});

import carouselInfoDialog from "./components/carouselInfoDialog.vue"
import {carouselDelete, carouselListOut} from "../../api/image.js";


// 按条件加载文章列表
const loadCarousels = async (pageNum = 0) => {

  if (pageNum != 0) {
    pageInfo.pageNum = pageNum;
  }
  // let res = await axios.get(`/images/list?userType=${pageCarouselType.value}&keyword=${pageInfo.keyword}&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}&iorder=${pageInfo.sortKey}`)
  // if (res.data.code == 200) {
  //   userList.value = res.data.data.user
  //   // console.log(articleList)
  // }
  carouselListOut(pageInfo).then(result => {
    if (result != null) {
      carouselList.value = result.data.data.carousel
      pageInfo.count = result.data.data.count;
    }
  })


  pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
  // console.log(userList.value)
}


const showAddCarouselDialog = (data) => {

  if (!data) {
    dialogTitle.value = "添加轮播图"
    dialogTableValue.value = {}
  }
  else {
    dialogTitle.value = "编辑轮播图"
    dialogTableValue.value = JSON.parse(JSON.stringify(data))
  }
  showDialog.value = true
  dialogCarouselType.value = pageInfo.category
}
const deleteCarousel = async (data) => {

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
        carouselDelete(data).then(result => {
          if (result == null) loadCarousels()
        })
      }).catch()

}

</script>

<style scoped>
.carousel-container {
  padding: 20px;
}

.mainTable {
  box-shadow: 2px 2px 6px #D3D4D8;
  border-radius: 10px;
}
</style>
