<template>
  <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick" router style="margin-top:50px;">
    <el-tab-pane label="特色产品" name="first"></el-tab-pane>
    <el-tab-pane label="美食推荐" name="second"></el-tab-pane>
    <el-tab-pane label="团队跟游" name="third"></el-tab-pane>
  </el-tabs>

      <template v-if="showFirst">
        <el-col>
          
          <el-row :span="6">
            <div v-for="(article, index) in articleList" style="margin-left:50px;margin-right:50px;margin-top:30px;">
              <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover" @click="toConsumpution(article)">
                <el-image :src="serverUrl + article.head_image" class="image"/>
                <div class="text-wrapper">
                <div style="font-size: 15px; color: #000;">{{article.title}} </div>
                <div style="font-size: 13px; color: #999;">开放时间 8:00-19:00</div>
                </div>
              </el-card>
            </div>
          </el-row>

        </el-col>
      </template>

  <template v-if="showSecond">
    <el-col>
          
      <el-row :span="6">
        <div v-for="(article, index) in articleList" style="margin-left:50px;margin-right:50px;margin-top:30px">
          <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover" @click="toConsumpution(article)">
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

  <template v-if="showThird">
    <el-col>
          
      <el-row :span="6">
        <div v-for="(article, index) in articleList" style="margin-left:50px;margin-right:50px;margin-top:30px">
          <el-card class="card" :body-style="{ padding: '0px' }" shadow="hover" @click="toConsumpution(article)">
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
const showFirst  = ref(true)
const showSecond = ref(false)
const showThird  = ref(false)

const pageInfo = reactive({
    pageNum: 1,
    pageSize: 5,
    pageCount: 0,
    count: 0,
    keyword: "",
    categoryId: window.location.href.slice(-1) // 设置文章类别为地址最后一位
})

onMounted(() => {
    loadArticles(0,1)
})

// 按条件加载文章列表
const loadArticles = async (pageNum = 0,index = 1) => {
    if (pageNum != 0) {
        pageInfo.pageNum = pageNum;
    }
    let res = await axios.get(`/article/list?articleType=consumptionarticle&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}&categoryId=${index}`)
    if (res.data.code == 200) {
        articleList.value = res.data.data.article
        // console.log(articleList)
    }
    pageInfo.count = res.data.data.count;
    pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
}

const toConsumpution = (article) => {
    router.push({
        path: "/article/detail",
        query: {
            category: "consumptionarticle",
            id: article.id
        }
    })
}

//路由守卫
onBeforeRouteUpdate((to, from) => {
  const toCategory = to.query.category;

  if (toCategory === '1') {
    showFirst.value = true;
    showSecond.value = false;
    showThird.value = false;
    activeName.value = 'first'
    loadArticles(0, 1);
  } else if (toCategory === '2') {
    showFirst.value = false;
    showSecond.value = true;
    showThird.value = false;
    activeName.value = 'second'
    loadArticles(0, 2);
  } else if (toCategory === '3') {
    showFirst.value = false;
    showSecond.value = false;
    showThird.value = true;
    activeName.value = 'third'
    loadArticles(0, 3);
  }

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

  if (tab.paneName === 'third') {
    showThird.value = true
    loadArticles(0,3)
    router.push({ query: { category: '3' } });
  } else {
    showThird.value = false
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
