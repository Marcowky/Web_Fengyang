<template>
  <div class="carousel-container">
    <userInfoDialog v-model="showDialog" :dialogTitle="dialogTitle" v-if="showDialog" @updateUserList="loadUsers"
                    :dialogTableValue="dialogTableValue" :dialogCarouselType="dialogCarouselType"/>
    <!-- 首页轮播图管理部分 -->
    <el-card class="carousel-card" shadow="hover">
<!--      <h3>首页轮播图</h3>-->
      <el-row :gutter="20">
<!--        <el-col :span="7">-->
<!--&lt;!&ndash;          <el-input&ndash;&gt;-->
<!--&lt;!&ndash;              v-model.trim="homeCarouselPageInfo.keyword"&ndash;&gt;-->
<!--&lt;!&ndash;              placeholder="请输入关键词"&ndash;&gt;-->
<!--&lt;!&ndash;              clearable&ndash;&gt;-->
<!--&lt;!&ndash;              :prefix-icon="searchIcon"&ndash;&gt;-->
<!--&lt;!&ndash;              @keyup.enter.native="loadHomeCarouselList()"&ndash;&gt;-->
<!--&lt;!&ndash;          ></el-input>&ndash;&gt;-->
<!--        </el-col>-->

<!--        <el-button type="primary" icon="el-icon-plus" size="medium" @click="showAddHomeCarouselDialog">-->
        <el-button @click="showAddCarouselDialog">
          添加
        </el-button>
      </el-row>
      <el-table
          :data="homeCarouselList"
          style="width: 100%; margin-top: 20px"
          stripe
          ref="homeCarouselTableRef"
          @sort-change="handleHomeCarouselSortChange"
      >
        <el-table-column prop="id" label="ID" sortable="custom" width="80"></el-table-column>
        <el-table-column prop="order" label="顺序" sortable="custom" width="80"></el-table-column>
        <el-table-column prop="url" label="图片链接">
          <template slot-scope="{ row }">
            <a :href="row.url" target="_blank">{{ row.url }}</a>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150">
          <template slot-scope="{ row }">
            <el-button type="primary" size="small" @click="showEditHomeCarouselDialog(row)">编辑</el-button>
            <el-button type="danger" size="small" @click="deleteHomeCarousel(row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination style="margin-top: 20px" background layout="prev, pager, next"
                     :page-count="homeCarouselPageInfo.pageCount" @current-change="loadHomeCarouselList"
      />
    </el-card>

    <!-- 消费景点轮播图管理部分 -->
    <el-card class="carousel-card" shadow="hover">
      <h3>消费景点轮播图</h3>
      <el-row :gutter="20">
        <el-col :span="7">
          <el-input
              v-model.trim="consumpAttractionCarouselPageInfo.keyword"
              placeholder="请输入关键词"
              clearable
              :prefix-icon="searchIcon"
              @keyup.enter.native="loadConsumpAttractionCarouselList()"
          ></el-input>
        </el-col>
        <el-button type="primary" icon="el-icon-plus" size="medium" @click="showAddConsumpAttractionCarouselDialog">
          添加
        </el-button>
      </el-row>
      <el-table
          :data="consumpAttractionCarouselList"
          style="width: 100%; margin-top: 20px"
          stripe
          ref="consumpAttractionCarouselTableRef"
          @sort-change="handleConsumpAttractionCarouselSortChange"
      >
        <el-table-column prop="id" label="ID" sortable="custom" width="80"></el-table-column>
        <el-table-column prop="order" label="顺序" sortable="custom" width="80"></el-table-column>
        <el-table-column prop="url" label="图片链接">
          <template slot-scope="{ row }">
            <a :href="row.url" target="_blank">{{ row.url }}</a>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150">
          <template slot-scope="{ row }">
            <el-button type="primary" size="small" @click="showEditConsumpAttractionCarouselDialog(row)">编辑</el-button>
            <el-button type="danger" size="small"
                       @click="deleteConsumpAttractionCarousel(row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination style="margin-top: 20px" background layout="prev, pager, next"
                     :page-count="consumpAttractionCarouselPageInfo.pageCount"
                     @current-change="loadConsumpAttractionCarouselList"
      />
    </el-card>

    <!-- 编辑轮播图弹窗 -->
<!--    <el-dialog-->
<!--        :title="dialogTitle"-->
<!--        :visible.sync="dialogVisible"-->
<!--        width="30%"-->
<!--        destroy-on-close-->
<!--        center-->
<!--        append-to-body-->
<!--    >-->
<!--      <el-form :model="currentCarousel" ref="carouselFormRef" label-width="80px">-->
<!--        <el-form-item label="顺序">-->
<!--          <el-input v-model.number.trim="currentCarousel.order"></el-input>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="图片链接">-->
<!--          <el-input v-model.trim="currentCarousel.url"></el-input>-->
<!--        </el-form-item>-->
<!--      </el-form>-->
<!--      <div slot="footer" class="dialog-footer">-->
<!--        <el-button @click="resetCarouselForm">取消</el-button>-->
<!--        <el-button type="primary" @click="saveCarousel">保存</el-button>-->
<!--      </div>-->
<!--    </el-dialog>-->
<!--    &lt;!&ndash; 添加轮播图弹窗 &ndash;&gt;-->
<!--    <el-dialog-->
<!--        title="添加轮播图"-->
<!--        :visible.sync="addDialogVisible"-->
<!--        width="30%"-->
<!--        destroy-on-close-->
<!--        center-->
<!--        append-to-body-->
<!--    >-->
<!--      <el-form :model="newCarousel" ref="addCarouselFormRef" label-width="80px">-->
<!--        <el-form-item label="顺序">-->
<!--          <el-input v-model.number.trim="newCarousel.order"></el-input>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="图片链接">-->
<!--          <el-input v-model.trim="newCarousel.url"></el-input>-->
<!--        </el-form-item>-->
<!--      </el-form>-->
<!--      <div slot="footer" class="dialog-footer">-->
<!--        <el-button @click="resetAddCarouselForm">取消</el-button>-->
<!--        <el-button type="primary" @click="addCarousel">保存</el-button>-->
<!--      </div>-->
<!--    </el-dialog>-->
  </div>
</template>

<script  setup>
import { ref, inject, onMounted, reactive } from 'vue'
import { userConfig } from "../../config/adminConfig.json"
import { ElMessage, ElMessageBox } from 'element-plus'



const axios = inject("axios")
import {
  Search,
} from '@element-plus/icons-vue'
import { useRoute, onBeforeRouteUpdate } from 'vue-router'
const pageCarouselType = ref('')
const route = useRoute()

// 挂载页面时触发
onMounted(() => {
  pageCarouselType.value = route.query.category;
  loadUsers()
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
  loadUsers()
}

// 4.设置路由守卫
onBeforeRouteUpdate((to, from) => {
  const fromCategory = from.query.category;
  const toCategory = to.query.category;

  if (fromCategory !== toCategory) {
    pageCarouselType.value=toCategory
    loadUsers()
  }
});


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


// 按条件加载文章列表
const loadUsers = async (pageNum = 0) => {

  if (pageNum != 0) {
    pageInfo.pageNum = pageNum;
  }
  let res = await axios.get(`/user/list?userType=${pageCarouselType.value}&keyword=${pageInfo.keyword}&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}&order=${pageInfo.sortKey}`)
  if (res.data.code == 200) {
    userList.value = res.data.data.user
    // console.log(articleList)
  }
  pageInfo.count = res.data.data.count;
  pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
  console.log(userList.value)
}

const dialogTableValue = ref({})
const dialogTitle = ref('')
const dialogCarouselType = ref('')
const showDialog = ref(false)
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
  dialogCarouselType.value = pageCarouselType.value
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
  )
      .then(async () => {
        let res = await axios.put("/user/delete", {
          ID: data.ID,
          UserType: data.UserType
        })
        if (res.data.code == 200) {
          ElMessage({
            message: res.data.msg,
            type: 'success',
            offset: 80
          })
          loadUsers()
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
  } else {
    ElMessage({
      message: res.data.msg,
      type: 'error',
      offset: 80
    })
  }
}
</script>


<!--<script>-->
<!--export default {-->
<!--  name: "CarouselPage",-->
<!--  data() {-->
<!--    return {-->
<!--// 首页轮播图列表-->
<!--      homeCarouselList: [],-->
<!--// 消费景点轮播图列表-->
<!--      consumpAttractionCarouselList: [],-->
<!--// 当前编辑的轮播图-->
<!--      currentCarousel: {},-->
<!--// 添加轮播图表单-->
<!--      newCarousel: {},-->
<!--// 首页轮播图分页信息-->
<!--      homeCarouselPageInfo: {-->
<!--        page: 1,-->
<!--        pageSize: 10,-->
<!--        keyword: "",-->
<!--        totalCount: 0,-->
<!--        pageCount: 0,-->
<!--      },-->
<!--// 消费景点轮播图分页信息-->
<!--      consumpAttractionCarouselPageInfo: {-->
<!--        page: 1,-->
<!--        pageSize: 10,-->
<!--        keyword: "",-->
<!--        totalCount: 0,-->
<!--        pageCount: 0,-->
<!--      },-->
<!--// 编辑/添加轮播图弹窗是否显示-->
<!--      dialogVisible: false,-->
<!--      addDialogVisible: false,-->
<!--// 弹窗标题-->
<!--      dialogTitle: "",-->
<!--// 搜索图标-->
<!--      searchIcon: "el-icon-search",-->
<!--    };-->
<!--  },-->
<!--  methods: {-->
<!--// 加载首页轮播图列表-->
<!--    loadHomeCarouselList(page = this.homeCarouselPageInfo.page) {-->
<!--      const params = {-->
<!--        page,-->
<!--        pageSize: this.homeCarouselPageInfo.pageSize,-->
<!--        keyword: this.homeCarouselPageInfo.keyword,-->
<!--        type: "HOME",-->
<!--      };-->
<!--// 调用接口获取数据-->
<!--      this.homeCarouselList = [];-->
<!--// 更新分页信息-->
<!--      this.homeCarouselPageInfo.totalCount = 0;-->
<!--      this.homeCarouselPageInfo.pageCount = 0;-->
<!--    },-->
<!--// 加载消费景点轮播图列表-->
<!--    loadConsumpAttractionCarouselList(page = this.consumpAttractionCarouselPageInfo.page) {-->
<!--      const params = {-->
<!--        page,-->
<!--        pageSize: this.consumpAttractionCarouselPageInfo.pageSize,-->
<!--        keyword: this.consumpAttractionCarouselPageInfo.keyword,-->
<!--        type: "CONSUMPTION_ATTRACTION",-->
<!--      };-->
<!--// 调用接口获取数据-->
<!--      this.consumpAttractionCarouselList = [];-->
<!--// 更新分页信息-->
<!--      this.consumpAttractionCarouselPageInfo.totalCount = 0;-->
<!--      this.consumpAttractionCarouselPageInfo.pageCount = 0;-->
<!--    },-->
<!--// 显示添加轮播图弹窗-->
<!--    showAddHomeCarouselDialog() {-->
<!--      this.dialogTitle = "添加首页轮播图";-->
<!--      this.newCarousel = {};-->
<!--      this.addDialogVisible = true;-->
<!--    },-->
<!--// 显示编辑轮播图弹窗-->
<!--    showEditHomeCarouselDialog(carousel) {-->
<!--      this.dialogTitle = "编辑首页轮播图";-->
<!--      this.currentCarousel = Object.assign({}, carousel);-->
<!--      this.dialogVisible = true;-->
<!--    },-->
<!--// // 显示添加消费景点轮播图弹窗-->
<!--//     showAddConsumpAttractionCarouselDialog() {-->
<!--//       this.dialogTitle = "添加消费景点轮播图";-->
<!--//       this.newCarousel = {};-->
<!--//       this.addDialogVisible = true;-->
<!--//     },-->
<!--// // 显示编辑消费景点轮播图弹窗-->
<!--//     showEditConsumpAttractionCarouselDialog(carousel) {-->
<!--//       this.dialogTitle = "编辑消费景点轮播图";-->
<!--//       this.currentCarousel = Object.assign({}, carousel);-->
<!--//       this.dialogVisible = true;-->
<!--//     },-->
<!--// 重置编辑/添加轮播图表单-->
<!--    resetCarouselForm() {-->
<!--      this.dialogVisible = false;-->
<!--      this.$nextTick(() => {-->
<!--        this.$refs.carouselFormRef.resetFields();-->
<!--        this.currentCarousel = {};-->
<!--      });-->
<!--    },-->
<!--    resetAddCarouselForm() {-->
<!--      this.addDialogVisible = false;-->
<!--      this.$nextTick(() => {-->
<!--        this.$refs.addCarouselFormRef.resetFields();-->
<!--        this.newCarousel = {};-->
<!--      });-->
<!--    },-->
<!--// 保存编辑/添加轮播图信息-->
<!--    saveCarousel() {-->
<!--      const formName = this.dialogTitle === "添加轮播图" ? "addCarouselFormRef" : "carouselFormRef";-->
<!--      this.$refs[formName].validate(valid => {-->
<!--        if (valid) {-->
<!--          // 调用接口保存数据-->
<!--          if (this.dialogTitle === "添加轮播图") {-->
<!--            // 添加轮播图-->
<!--            this.addCarousel(this.newCarousel);-->
<!--          } else {-->
<!--            // 编辑轮播图-->
<!--            this.editCarousel(this.currentCarousel);-->
<!--          }-->
<!--        } else {-->
<!--          return false;-->
<!--        }-->
<!--      });-->
<!--    },-->
<!--// 添加轮播图-->
<!--    addCarousel(carousel) {-->
<!--      // 调用接口添加数据-->
<!--      this.addDialogVisible = false;-->
<!--      // 重新加载数据-->
<!--      if (this.dialogTitle === "添加首页轮播图") {-->
<!--        this.loadHomeCarouselList();-->
<!--      } else {-->
<!--        this.loadConsumpAttractionCarouselList();-->
<!--      }-->
<!--    },-->
<!--// 编辑轮播图-->
<!--    editCarousel(carousel) {-->
<!--      // 调用接口更新数据-->
<!--      this.dialogVisible = false;-->
<!--      // 重新加载数据-->
<!--      if (this.dialogTitle === "编辑首页轮播图") {-->
<!--        this.loadHomeCarouselList();-->
<!--      } else {-->
<!--        this.loadConsumpAttractionCarouselList();-->
<!--      }-->
<!--    },-->
<!--// 删除轮播图-->
<!--    deleteHomeCarousel(id) {-->
<!--      // 调用接口删除数据-->
<!--      // 重新加载数据-->
<!--      this.loadHomeCarouselList();-->
<!--    },-->
<!--    deleteConsumpAttractionCarousel(id) {-->
<!--      // 调用接口删除数据-->
<!--      // 重新加载数据-->
<!--      this.loadConsumpAttractionCarouselList();-->
<!--    },-->
<!--// 首页轮播图分页信息变化回调方法-->
<!--    handleHomeCarouselSortChange(sort) {-->
<!--      this.homeCarouselPageInfo.page = sort.currentPage;-->
<!--      this.homeCarouselPageInfo.pageSize = sort.pageSize;-->
<!--      this.loadHomeCarouselList();-->
<!--    },-->
<!--// 消费景点轮播图分页信息变化回调方法-->
<!--    handleConsumpAttractionCarouselSortChange(sort) {-->
<!--      this.consumpAttractionCarouselPageInfo.page = sort.currentPage;-->
<!--      this.consumpAttractionCarouselPageInfo.pageSize = sort.pageSize;-->
<!--      this.loadConsumpAttractionCarouselList();-->
<!--    },-->
<!--  },-->
<!--};-->
<!--</script>-->

<style scoped>
.carousel-container {
  padding: 20px;
}

.carousel-card {
  margin-bottom: 20px;
}
</style>
