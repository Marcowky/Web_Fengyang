
<template>
  <!-- 顶部导航栏 -->
    <el-card style="border-radius: 20px;margin-left: 150px;margin-right: 150px;margin-top: 30px">
      <el-row :gutter="5">
        <el-col :span="9" :offset="2" ><div class="grid-content ep-bg-purple" />
              <h>
                最新资讯
              </h>
          <div class="arrow">
          </div>
          <el-card class ="part">
            <el-scrollbar height="80%" ref="carousel" class = "news">
              <el-row v-for="(article, index) in newsList" class = "news_item">
                <div class="news_mode">
                  <img v-bind:src="article.head_image" alt=" " class ="news_img"/>
                  <span class="news_content" @click="showdetail()">
                  {{article.title}}
                </span>
                </div>
                <div class="big_mode">
                  <img v-bind:src="article.head_image" alt=" " class ="big_img"/>
                  <span class="big_content" @click="showdetail()">
                  {{article.title}}
                </span>
                </div>
              </el-row>
            </el-scrollbar>

          </el-card>
        </el-col>

        <el-col :span="9" :offset="2"><div class="grid-content ep-bg-purple" />
            <h>
              景区公告
            </h>
          <el-card class="part">
            <el-scrollbar height="80%">
              <el-row v-for="v in advicelist" :key="v.value">
                <div class="list">
                  {{v.content}}
                  <router-link :to="{path: '/info/page'}"
                               tag="button">>></router-link>
                </div>
              </el-row>
            </el-scrollbar>
          </el-card>

        </el-col>
      </el-row>
      <el-row style="margin-top: 20px;margin-bottom: 8px;margin-left: 15px">
        <el-divider>
          <span style="font-size: 1.5rem;font-weight:bold;font-family: 华文楷体">丰阳 欢迎您!</span>
        </el-divider>
      </el-row>

      <el-row :gutter="20">
        <el-col :span="9" :offset="2"><div class="grid-content ep-bg-purple" />
            <h>
              旅游攻略
            </h>
          <el-card class="part">
              <el-card v-for="v in advicelist" class="modeCard">
                  <img
                      src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                      style="width: 100%;height: inherit"
                  />
                  <span class= "part3-title">攻略1</span>
              </el-card>
          </el-card>
        </el-col>
        <el-col :span="9" :offset="2"><div class="grid-content ep-bg-purple" />
            <h>
              假日时节
            </h>
          <el-card class="part">
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
import { ref, reactive, inject, onMounted } from 'vue'
import warning from '../../config/warning.json'
import { useRoute, useRouter, onBeforeRouteUpdate } from 'vue-router'

// 网络请求
const serverUrl = inject("serverUrl")
const axios = inject("axios")
// 变量初始化
const newsList = ref([])
const warningList = warning.warning_items
const advicelist = warning.warning_items
//首先在setup中定义
const router = useRouter()
let carousel = ref(null)

// 挂载页面时触发
onMounted(() => {
  loadArticles()
})
// 按条件加载文章列表
const loadArticles = async () => {
  // let res = await axios.get(`/article/list?articleType=infoArticle&keyword=""&pageNum=1&pageSize=5&categoryId=1`)
  let res = await axios.get(`/article/list?articleType=infoArticle&pageNum=1&pageSize=5`)
  if (res.data.code === 200) {
    newsList.value = res.data.data.article
    console.log(newsList)
  }
  else
  {
    console.log("wrong")
  }
}

const arrowClick = (val) => {
  if(val === 'right') {
    carousel.value.next()
  } else {
    carousel.value.prev()
  }
}
const showdetail = (event) => {
 router.push('/info/page')
}


</script>

<style lang="scss" scoped>
h{
  color: rgb(216, 102, 102);
  font-size: 1.5rem;
  font-weight: bold;
  margin-bottom: 1.4rem;
}

.grid-content {
  border-radius: 20px;
  width: 52rem;
}
.part{
  height: 35rem;
  width: inherit;
  position: relative;
  border-radius: 10px;
}
.news{
  left:0.3rem;
  top:1rem;
}
.news_mode{
  position: relative;
  margin: 0.3rem;
  display: block;
}
.big_mode{
  position: relative;
  margin: 0.3rem;
  display: none;
}
.news_item:hover .news_mode{
  display: none;
}
.news_item:hover .big_mode{
  display: block;
}
.news_img{
  height: inherit;
  width: 25%
}
.big_img{
  left: 0.3rem;
  height: 80%;
  width: 80%
}
.news_content{
  width: 70%;
  position: absolute;
  left:2rem;
  top:0.3rem;
  font-size: 17em;
  overflow: hidden; // 文字超长隐藏
  text-overflow:ellipsis; // 显示...
  white-space: nowrap; // 单行显示
}
.big_content{
  width: 80%;
  position: absolute;
  left:0.3rem;
  bottom: 4rem;
  font-size: 1rem;
  font-weight: bold;
  color: aliceblue;
  background-color: rgba(0, 0, 0, 0.5);
  overflow: hidden; // 文字超长隐藏
  text-overflow:ellipsis; // 显示...
  white-space: nowrap; // 单行显示
  z-index: 2
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

.summary:hover{
  font-size: 20px;
}

.list{
  width: 100%;
  margin-left: 1rem;
  margin-bottom: 1rem;
  font-size: 1.2rem;
  font-weight: 550;
}
.list:hover{
  transform: scale(1.02);
  color: rgba(0, 0, 255, 0.96);
}

.modeCard{
  float: left;
  width:13rem;
  height: inherit;
  margin-left: 2.5rem;
  margin-top: 0.5rem;}
.modeCard:hover{ //鼠标悬停时激活
transform: scale(1.05); //放大倍数
}

</style>
