<template>
    <el-tabs v-model="activeName" :tab-position="tabPosition" style="height: 800px" class="demo-tabs" @tab-click="handleClick">
      <el-tab-pane label="热门景点" name="first"></el-tab-pane>
      <el-tab-pane label="游玩攻略" name="second"></el-tab-pane>

    <template v-if="showFirst">
        <el-col>
          
          <el-row :span="6">
            <div v-for="(article, index) in articleList" style="margin-left:40px;margin-right:40px;margin-bottom:30px">
              <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover" @click="toAttraction(article)">
                <el-image :src="serverUrl + article.head_image" class="image"/>
                <div class="text-wrapper">
                <div style="font-size: 15px; color: #000;">{{article.title}} </div>
                <div style="font-size: 13px; color: #999;">{{article.content}}</div>
                </div>
              </el-card>
            </div>
          </el-row>

        </el-col>
      </template>

      <template v-if="showSecond">
        <el-col>
              
          <el-row :span="6">
            <div v-for="(article, index) in articleList" style="margin-left:40px;margin-right:40px;margin-top:30px">
              <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover" @click="toAttraction(article)">
                <el-image :src="serverUrl + article.head_image" class="image"/>
                <div class="text-wrapper">
                <div style="font-size: 15px; color: #000;">{{article.title}} </div>
                <div style="font-size: 13px; color: #999;">{{article.content}}</div>
                </div>
              </el-card>
            </div>
          </el-row>
    
        </el-col>
      </template>
    </el-tabs>
</template>

  <script setup>
  import { ref, reactive, inject, onMounted } from 'vue'
  import { useRoute, useRouter, onBeforeRouteUpdate } from 'vue-router'
  // import type { TabsPaneContext } from 'element-plus'

  const tabPosition = ref('left')
  const activeName = ref('first')
  const articleList = ref([])
  const activeIndex = ref('0')
  const router = useRouter()
  const route = useRoute()
  const serverUrl = inject("serverUrl")
  const axios = inject("axios")
  const showFirst  = ref(true)
  const showSecond = ref(false)

  const pageInfo = reactive({
      pageNum: 1,
      pageSize: 5,
      pageCount: 0,
      count: 0,
      keyword: "",
      categoryId: window.location.href.slice(-1) // 设置文章类别为地址最后一位
  })

  const handleCategoryChange = (category) => {
    showFirst.value = false;
    showSecond.value = false;
    activeName.value = '';

    if (category === '1') {
      showFirst.value = true;
      activeName.value = 'first';
      loadArticles(0, 1);
    } else if (category === '2') {
      showSecond.value = true;
      activeName.value = 'second';
      loadArticles(0, 2);
    }
  };

  onMounted(() => {
    const category = route.query.category;
    handleCategoryChange(category);
  })

  // 按条件加载文章列表
  const loadArticles = async (pageNum = 0,index = 1) => {
      if (pageNum != 0) {
          pageInfo.pageNum = pageNum;
      }
      let res = await axios.get(`/article/list?articleType=attractionArticle&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}&categoryId=${index}`)
      if (res.data.code == 200) {
          articleList.value = res.data.data.article
          // console.log(articleList)
      }
      pageInfo.count = res.data.data.count;
      pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
      console.log(articleList.value)
  }

  const toAttraction = (article) => {
      router.push({
          path: "/article/detail",
          query: {
              category: "attractionarticle",
              id: article.id,
          }
      })
  }

  onBeforeRouteUpdate((to, from) => {
  const toCategory = to.query.category;
  handleCategoryChange(toCategory);
});


// import { TabsPaneContext } from 'element-plus'
const handleClick = (tab, event) => {
  console.log(tab, event)
  
  if (tab.paneName === 'first') {
    showFirst.value = true
    loadArticles(0,1)
    router.push({ query: { category: '1' } });
  } else {
    showFirst.value = false
  }

  if (tab.paneName === 'second') {
    showSecond.value = true
    loadArticles(0,2)
    router.push({ query: { category: '2' } });
  } else {
    showSecond.value = false
  }
}


  </script>

 <style>
  .demo-tabs > .el-tabs__content {
    padding: 32px;
    color: #6b778c;
    font-size: 32px;
    font-weight: 600;
  }
  
  .el-tabs--right .el-tabs__content,
  .el-tabs--left .el-tabs__content {
    height: 100%;
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
  