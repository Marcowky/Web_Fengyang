
<template>
  <!-- 顶部导航栏 -->
  <el-card style="border-radius: 20px;margin-left: 10rem;margin-right: 10rem;margin-top: 2rem">
    <el-row :gutter="5">
      <el-col :span="10" :offset="1">
        <div class="grid-content ep-bg-purple" />
        <h>
          最新资讯
        </h>
        <el-card shadow="hover" class="part1">
          <el-button class="page" @click="loadArticles">换一批</el-button>
          <el-scrollbar height="60%" ref="carousel" class="news">
            <el-row v-for="(article, index) in qList" class="news_item">
              <div class="news_mode" @click="showdetail(article)">
                <el-image :src="serverUrl + article.head_image" class="news_img" />
                <span class="news_content">
                  {{ article.title }}
                </span>
                <div class="news_time">发布时间：{{ article.created_at }}</div>
              </div>
            </el-row>
          </el-scrollbar>
        </el-card>
      </el-col>
      <el-col :span="10" :offset="2">
        <div class="grid-content ep-bg-purple" />
        <h>
          景区公告
        </h>
        <el-card shadow="hover" class="part1">
          <el-scrollbar max-height="80%">
            <el-row v-for="(article, index) in warningList">
              <div class="list" @click="showdetail(article)">
                {{ article.title }}
              </div>
              <span style="position:absolute;right: 1.5rem">>></span>
            </el-row>
          </el-scrollbar>
        </el-card>

      </el-col>
    </el-row>
    <el-row style="margin-top: 2rem;margin-bottom: 8px;margin-left: 1.3rem">
      <el-divider>
        <span style="font-size: 1.5rem;font-weight:bold;font-family: 华文楷体">丰阳 欢迎您 !</span>
      </el-divider>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="10" :offset="1">
        <div class="grid-content ep-bg-purple" />
        <h>
          三下乡
        </h>
        <el-card shadow="hover" class="part2">
          <el-card v-for="(article, index) in adviceList" class="modeCard" @click="showdetail(article)">
            <el-image :src="serverUrl + article.head_image" style="width: 100%;height: inherit" />
            <span class="part3-title">{{ article.title }}</span>
          </el-card>
        </el-card>
      </el-col>
      <el-col :span="10" :offset="2">
        <div class="grid-content ep-bg-purple" />
        <h>
          假日时节
        </h>
        <el-card shadow="hover" class="part2">
          <img src="../../assets/event.jpg" style="width: 95%;height: inherit;margin-left: 1rem;" />
          <div data-tockify-component="calendar" data-tockify-calendar="weigp"></div>
        </el-card>
      </el-col>
    </el-row>
  </el-card>
</template>


<script setup>
import { ref, reactive, inject, onMounted, onUpdated } from 'vue'
import warning from '../../config/warning.json'
import news from '../../config/news.json'
import { useRoute, useRouter, onBeforeRouteUpdate } from 'vue-router'
import { articleListOut } from "../../api/article.js";

// 网络请求
const serverUrl = inject("serverUrl")
const axios = inject("axios")
//首先在setup中定义
const router = useRouter()
// 变量初始化
let number = '0'
let qList = ref([])
let newsList = ref([])
let warningList = ref([])
let adviceList = ref([])
let timeStart = 0 //截取第几组的结束
let timeEnd = 1 //默认为0组
let group = 0 //组数
let num = 4//一页展示list数量
let clickNum = 0//点击次数
let pageInfo1 = ({
  pageNum: 1,
  pageSize: 4,
  pageCount: 0,
  count: 0,
  pageArticleType: 'infoArticle',
  categoryId: 1
})
let pageInfo2 = ({
  pageArticleType: 'infoArticle',
  categoryId: 2
})
let pageInfo3 = ({
  pageArticleType: 'infoArticle',
  categoryId: 3
})
// 挂载页面时触发
onMounted(() => {
  loadArticles()
})
// 按条件加载文章列表
const loadArticles = async () => {
  articleListOut(pageInfo1).then(result => {
    if (result != null) {
      qList.value = result.data.data.article
      pageInfo1.count = result.data.data.count
      pageInfo1.pageCount = parseInt(pageInfo1.count / pageInfo1.pageSize) + (pageInfo1.count % pageInfo1.pageSize > 0 ? 1 : 0)
      console.log(qList.value)
      pageInfo1.pageNum++
      if (pageInfo1.pageNum === pageInfo1.pageCount + 1) {
        pageInfo1.pageNum = 1
      }
    }
  })
  articleListOut(pageInfo2).then(result => {
    if (result != null) {
      warningList.value = result.data.data.article
    }
  })
  articleListOut(pageInfo3).then(result => {
    if (result != null) {
      adviceList.value = result.data.data.article
    }
  })
}
//const arrowClick = (val) => {
//  if(val === 'right') {
//    carousel.value.next()
//  } else {
//    carousel.value.prev()
//  }
//}
const showdetail = (article) => {
  router.push({
    path: "/article/detail",
    query: {
      category: 'infoArticle',
      id: article.id
    }
  })
}

</script>

<style lang="scss" scoped>
h {
  color: rgb(216, 102, 102);
  font-size: 1.3rem;
  font-weight: bold;
  margin-top: 1rem;
  margin-bottom: 1.6rem;
}

.grid-content {
  border-radius: 20px;
  width: 30rem;
}

.part1 {
  height: 100%;
  width: 110%;
  position: relative;
  border-radius: 10px;
}

.part2 {
  height: 80%;
  width: 110%;
  position: relative;
  border-radius: 10px;
}

.news_mode {
  position: relative;
  margin: 0.15rem;
}

.news_img {
  transition-delay: 9999s;
  height: inherit;
  width: 30%;
  border-radius: 10px
}

.news_content {
  transition-delay: 9999s;
  position: absolute;
  left: 33%;
  bottom: 60%;
  right: 0;
  font-size: 0.9rem;
  color: black;
  background-color: white;
  border-radius: 10px;
  /*overflow: hidden; // 文字超长隐藏
  text-overflow:ellipsis; // 显示...
  white-space: nowrap; // 单行显示
  z-index:2;*/
}

.news_time {
  transition-delay: 9999s;
  position: absolute;
  left: 33%;
  right: 10%;
  bottom: 20%;
  font-size: 0.7rem;
  color: black;
  background-color: white;
  font-weight: normal;
  visibility: visible;
}

.news {
  position: relative;
  top: 1.5rem;
}

.news:hover .news_img {
  transition: none;
  height: inherit;
  width: 30%
}

.news:hover .news_content {
  transition: none;
  position: absolute;
  left: 33%;
  bottom: 70%;
  right: 0;
  font-size: 0.9rem;
  font-weight: normal;
  color: black;
  background-color: white;

}

.news:hover .news_time {
  transition: none;
  position: absolute;
  left: 33%;
  right: 10%;
  bottom: 20%;
  font-size: 0.7rem;
  color: black;
  background-color: white;
  font-weight: normal;
  visibility: visible;
}

.news_item:hover .news_img {
  transition-delay: 0s;
  width: 70%;
  height: inherit;
}

.news_item:hover .news_content {
  transition-delay: 0s;
  position: absolute;
  left: 0;
  right: 30%;
  bottom: 0.5rem;
  font-size: 1rem;
  font-weight: bold;
  color: aliceblue;
  background-color: rgba(0, 0, 0, 0.5);

}

.news_item:hover .news_time {
  transition-delay: 0s;
  visibility: hidden;
}

.page {
  position: absolute;
  color: rgb(103, 151, 213);
  right: 0.8rem;
  z-index: 2;
}

.list {
  width: 85%;
  margin-left: 1rem;
  margin-bottom: 1rem;
  font-size: 0.9rem;
  overflow: hidden; // 文字超长隐藏
  text-overflow: ellipsis; // 显示...
  white-space: nowrap; // 单行显示
}

.list:hover {
  transform: scale(1.02);
  color: rgba(0, 0, 255, 0.96);
  cursor: pointer;
}

.modeCard {
  float: left;
  width: 12rem;
  height: 12rem;
  border-radius: 5px;
  margin-left: 7%;
  margin-top: 10%;
}

.modeCard:hover {
  //鼠标悬停时激活
  transform: scale(1.05); //放大倍数
}

.part3-title {
  bottom: 0;
  color: black;
}

.modeCard:hover .part3-title {
  color: cornflowerblue;
}</style>
