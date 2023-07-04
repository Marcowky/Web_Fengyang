
<template>
  <!-- 顶部导航栏 -->
    <el-card style="border-radius: 20px;margin-left: 10rem;margin-right: 10rem;margin-top: 2rem">
      <el-row :gutter="5">
        <el-col :span="9" :offset="1" ><div class="grid-content ep-bg-purple" />
              <h>
                最新资讯
              </h>
          <div class="arrow">
          </div>
          <el-card shadow ="hover" class ="part">
            <el-scrollbar height="60%" ref="carousel" class = "news">
              <el-row v-for="(article, index) in newsList" class = "news_item">
                <div class="news_mode">
                  <img v-bind:src="article.res" alt=" " class="news_img"/>
                  <span class="news_content" @click="showdetail(article)">
                  {{ article.content }}
                </span>
                </div>
              </el-row>
            </el-scrollbar>


          </el-card>
        </el-col>

        <el-col :span="9" :offset="3"><div class="grid-content ep-bg-purple" />
            <h>
              景区公告
            </h>
          <el-card shadow ="hover" class="part">
            <el-scrollbar max-height="80%">
              <el-row v-for="(article, index) in warningList">
                <div class="list">
                  {{article.title}}
                  <span style="position:absolute;right: 1.5rem" @click="showdetail(article)">>></span>
                </div>
              </el-row>
            </el-scrollbar>
            <div class="flex-align" @click="change()">
              <span>换一批</span>
            </div>
          </el-card>

        </el-col>
      </el-row>
      <el-row style="margin-top: 20px;margin-bottom: 8px;margin-left: 1.1rem">
        <el-divider>
          <span style="font-size: 1.5rem;font-weight:bold;font-family: 华文楷体">丰阳 欢迎您!</span>
        </el-divider>
      </el-row>

      <el-row :gutter="20">
        <el-col :span="9" :offset="1"><div class="grid-content ep-bg-purple" />
            <h>
              旅游攻略
            </h>
          <el-card shadow ="hover" class="part">
              <el-card v-for="(article, index) in adviceList" class="modeCard">
                  <img
                      src="article.head_image"
                      style="width: 100%;height: inherit"
                  />
                  <span class= "part3-title">{{article.title}}</span>
              </el-card>
          </el-card>
        </el-col>
        <el-col :span="9" :offset="3"><div class="grid-content ep-bg-purple" />
            <h>
              假日时节
            </h>
          <el-card shadow ="hover" class="part">
            <img
                src="../../assets/pic1.jpg"
                style="width: 100%;height: inherit;margin-right: 10px"
            />
            <div data-tockify-component="calendar" data-tockify-calendar="weigp"></div>
          </el-card>

        </el-col>
      </el-row>
    </el-card>

</template>


<script setup>
import { ref, reactive, inject, onMounted,onUpdated } from 'vue'
import warning from '../../config/warning.json'
import news from '../../config/news.json'
import { useRoute, useRouter, onBeforeRouteUpdate } from 'vue-router'

// 网络请求
const serverUrl = inject("serverUrl")
const axios = inject("axios")
//首先在setup中定义
const router = useRouter()
// 变量初始化
let number = '0'
let qList = news.news_items
let newsList = ref({})
let warningList = ref({})
let adviceList = ref({})
//
let timeStart = 0 //截取第几组的结束
let timeEnd = 1 //默认为0组
let group =  0 //组数
let num = 4//一页展示list数量
let clickNum =  0//点击次数

// 挂载页面时触发
onMounted(() => {
  loadArticles()
})

// 按条件加载文章列表
const loadArticles = async () => {

  let res = await axios.get(`/article/list?articleType=infoArticle&keyword=""&pageNum=1&pageSize=5&categoryId=2`)
  let res = await axios.get(`/article/list?articleType=infoArticle&pageNum=1&pageSize=12&categoryId=1`)
  if (res.data.code === 200) {
    qList.value = res.data.data.article
  }
  newsList.value = qList.value.slice(
      num * timeStart,
      num * timeEnd
  );
  let res = await axios.get(`/article/list?articleType=infoArticle&pageNum=1&pageSize=5&categoryId=1`)
  if (res.data.code === 200) {
    warningList.value = res.data.data.article
  }
  res = await axios.get(`/article/list?articleType=infoArticle&pageNum=1&pageSize=5&categoryId=3`)
  if (res.data.code === 200) {
    adviceList.value = res.data.data.article
  }


}
const change = async () => {
  if (qList.length > 4 && qList.length > num) {
    //点击的时候获取分为几组
    listlen();
    //每点击一次记录点击次数
    autoIncre();
    clear();
    renderR();
  }
}
// 计算数据的长度，共分为几组，如果不能整除则加1
const listlen = () => {
  let len = qList.length;
  group = len / num;
  if (len % num !== 0) {
    group = parseInt(group) + 1;
  }
}
//每点击一次，记录次数
const autoIncre = () => {
  clickNum++;
  timeStart++;
  timeEnd++;
}
//计算将点击次数和开始截取的参数清空, 如果点击此时大于当前数据的组数，则重新开始计数。
const clear = () =>{
  if (clickNum > group - 1) {
    timeStart = 0;
    timeEnd = 1;
    clickNum = 0;
  }
}
//截取当前每组的数据
const renderR = () =>{
  newsList = qList.slice(
      num * timeStart,
      num * timeEnd
  );
  console.log(newsList);
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
    path: "/info/detail",
    query: {
      id: article.id,
    }
  })
}

</script>

<style lang="scss" scoped>
h{
  color: rgb(216, 102, 102);
  font-size: 1.3rem;
  font-weight: bold;
  margin-bottom: 1.4rem;
}

.grid-content {
  border-radius: 20px;
  width: 30rem;
}
.part{
  height: 31rem;
  width: 32rem;
  position: relative;
  border-radius: 10px;
}

.news_mode{
  position: relative;
  margin: 0.15rem;
}
.news_img{
  transition-delay:9999s;
  height: inherit;
  width: 30%;
  border-radius:10px
}
.news_content{
  transition-delay:9999s;
  position: absolute;
  left:32%;
  right: 10%;
  bottom: 50%;
  font-size: 0.9rem;
  color: black;
  background-color:white;
  border-radius:10px;
  font-weight: normal;
  overflow: hidden; // 文字超长隐藏
  text-overflow:ellipsis; // 显示...
  white-space: nowrap; // 单行显示
  z-index:2;
}

.news{
  left:5%;
  top:1rem;
}
.news:hover .news_img{
  transition: none;
  height: inherit;
  width: 30%
}
.news:hover .news_content{
  transition:none;
  position: absolute;
  left:32%;
  right: 10%;
  bottom: 50%;
  font-size: 0.9rem;
  font-weight: normal;
  color: black;
  background-color:white;
}
.news_item:hover .news_img{
  transition-delay:0s;
  width: 60%;
  height: inherit;
}
.news_item:hover .news_content{
  transition-delay:0s;
  position: absolute;
  left: 0;
  right: 40%;
  bottom: 0;
  font-size: 1rem;
  font-weight: bold;
  color: aliceblue;
  background-color: rgba(0, 0, 0, 0.5);
}

.arrow{
  display: none;
  position: absolute;
  color: rgba(220, 156, 125, 0.8);
  right:1.4rem;
  bottom:0.6rem;
  z-index: 2;
}
.news:hover .arrow{
  display: block;
}

.list{
  width: 100%;
  margin-left: 2rem;
  margin-bottom: 1rem;
  font-size: 0.9rem;
  overflow: hidden; // 文字超长隐藏
  text-overflow:ellipsis; // 显示...
  white-space: nowrap; // 单行显示
}
.list:hover{
  transform: scale(1.02);
  color: rgba(0, 0, 255, 0.96);
  cursor: pointer;
}

.modeCard{
  float: left;
  width:11rem;
  height: inherit;
  margin-left: 2.8rem;
  margin-top: 0.5rem;
}
.modeCard:hover{ //鼠标悬停时激活
transform: scale(1.05); //放大倍数
}

</style>
