<template>
  <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
    <el-tab-pane label="特色产品" name="first"></el-tab-pane>
    <el-tab-pane label="美食推荐" name="second"></el-tab-pane>
    <el-tab-pane label="团队跟游" name="third"></el-tab-pane>
  </el-tabs>

      <template v-if="showFirst">
        <el-row>
          
          <el-col :span="6">
            <div v-for="(article, index) in articleList" style="margin:15px">
              <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover" @click="toDetail(article)">
                <img :src="serverUrl + article.head_image" class="image"/>
                <div v-html="article.content"></div>
                <div class="text-wrapper">
                <div style="font-size: 15px; color: #000;">
                  <span>阳光牧场</span>
                </div>
                <div style="font-size: 13px; color: #999;">
                  <span>开放时间8:00-18:00</span>
                </div>
              </div>
              </el-card>
            </div>
          </el-col>

          <el-col :span="6" :offset="3">
            <a href="http://www.baidu.com">
              <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover">
                <img src="../../assets/pic8.jpg" class="image"/>
                <div class="text-wrapper">
                <div style="padding: 0px; font-size: 15px; color: #000;">
                  <span>青山绿水</span>
                </div>
                <div style="padding: 0px; font-size: 13px; color: #999;">
                  <span>开放时间8:00-18:00</span>
                </div>
              </div>
              </el-card>
            </a>
          </el-col>

          <el-col :span="6" :offset="3">
            <a href="http://www.baidu.com">
              <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover">
                <img src="../../assets/pic9.jpg" class="image"/>
                <div class="text-wrapper">
                  <div style="padding: 0px; font-size: 15px; color: #000;">
                    <span>万里晴空</span>
                  </div>
                  <div style="padding: 0px; font-size: 13px; color: #999;">
                    <span>开放时间8:00-18:00</span>
                  </div>
                </div>
                </el-card>
              </a>
            </el-col>
          </el-row>
  
        <el-row class="card-wrapper">
          <el-col :span="6" >
            <a href="http://www.baidu.com">
            <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover">
              <img src="../../assets/pic10.jpg" class="image"/>
              <div class="text-wrapper">
              <div style="padding: 0px; font-size: 15px; color: #000;">
                <span>阡陌交通</span>
              </div>
              <div style="padding: 0px; font-size: 13px; color: #999;">
                <span>开放时间8:00-18:00</span>
              </div>
            </div>
            </el-card>
            </a>
          </el-col>

          <el-col :span="6" :offset="3">
            <a href="http://www.baidu.com">
            <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover">
              <img src="../../assets/pic11.jpg" class="image"/>
              <div class="text-wrapper">
              <div style="padding: 0px; font-size: 15px; color: #000;">
                <span>夕阳乡巷</span>
              </div>
              <div style="padding: 0px; font-size: 13px; color: #999;">
                <span>开放时间8:00-18:00</span>
              </div>
            </div>
            </el-card>
            </a>
          </el-col>

          <el-col :span="6" :offset="3">
            <a href="http://www.baidu.com">
            <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover">
              <img src="../../assets/pic12.jpg" class="image"/>
              <div class="text-wrapper">
              <div style="font-size: 15px; color: #000;">
                <span>落日余晖</span>
              </div>
              <div style="font-size: 13px; color: #999;">
                <span>开放时间8:00-18:00</span>
              </div>
              </div>
            </el-card>
            </a>
          </el-col>

        </el-row>

      </template>

  <template v-if="showSecond">
    <el-row>
      <el-col
        v-for="(o, index) in 6"
        :key="o"
        :span="6"
        :offset="index % 3 ? 2 : 0"
      >
        <el-card class="card" :body-style="{ padding: '0px' }">
          <img src="../../assets/pic7.jpg" class="image"/>
          <div style="padding: 5px; font-size: 13px; color: #999;">
            <span>开放时间8:00-18:00</span>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </template>
</template>

<script setup>
import { ref, reactive, inject, onMounted } from 'vue'
import { useRoute, useRouter, onBeforeRouteUpdate } from 'vue-router'
const activeName = ref('first')
const articleList = ref([])
const activeIndex = ref('0')
const router = useRouter()
const route = useRoute()
const serverUrl = inject("serverUrl")
const axios = inject("axios")

const pageInfo = reactive({
    pageNum: 1,
    pageSize: 5,
    pageCount: 0,
    count: 0,
    keyword: "",
    categoryId: window.location.href.slice(-1) // 设置文章类别为地址最后一位
})

onMounted(() => {
    loadArticles()
})

// 按条件加载文章列表
const loadArticles = async (pageNum = 0) => {
    if (pageNum != 0) {
        pageInfo.pageNum = pageNum;
    }
    let res = await axios.get(`/article/list?articleType=consumptionarticle&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}&categoryId=${pageInfo.categoryId}`)
    if (res.data.code == 200) {
        articleList.value = res.data.data.article
        // console.log(articleList)
    }
    pageInfo.count = res.data.data.count;
    pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
}

const toDetail = (article) => {
    router.push({
        path: "/blog/detail",
        query: {
            id: article.id,
        }
    })
}

const showFirst  = ref(true)
const showSecond = ref(false)
// import { TabsPaneContext } from 'element-plus'
const handleClick = (tab, event) => {
  console.log(tab, event)
  
  if (tab.paneName === 'first') {
    showFirst.value = true
  } else {
    showFirst.value = false
  }

  if (tab.paneName === 'second') {
    showSecond.value = true
  } else {
    showSecond.value = false
  }
}
</script>

<style>
.demo-tabs > .el-tabs__content {
  padding: 12px;
  color: #6b778c;
}

.card-wrapper {
  margin-top: 20px;
}

.text-wrapper {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: flex-end;
}

.card {
  width: 280px;
  height: 220px;
  border-radius: 15px;
}

.card:hover{
  transform: scale(1.05);
}

.image {
  width: 280px;
  height: 180px;
  object-fit: cover;
}

a{
  text-decoration:none;
}
</style>
