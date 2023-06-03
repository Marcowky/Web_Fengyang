
<template>
  <!-- 顶部导航栏 -->
    <el-card style="border-radius: 20px;margin-left: 120px;margin-right: 120px;margin-top: 30px">
      <el-row :gutter="5" style="margin-left: 30px;margin-right: 30px">
        <el-col :span="9" :offset="2" ><div class="grid-content ep-bg-purple" />
              <h>
                最新资讯
              </h>
          <div class="arrow">
          </div>
          <el-card class ="part">
            <el-scrollbar height="400px" ref="carousel" class = "news">
              <el-row v-for="v in newslist"  key= "v.value" class = "news_item">
                <div class="news_mode">
                  <img v-bind:src="v.res" alt=" " class ="news_img"/>
                  <span class="news_content" @click="showdetail()">
                  {{v.content}}
                </span>
                </div>
                <div class="big_mode">
                  <img v-bind:src="v.res" alt=" " class ="big_img"/>
                  <span class="big_content" @click="showdetail()">
                  {{v.content}}
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
            <el-scrollbar height="200px">
              <el-row v-for="v in warninlist" :key="v.value">
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
          <span style="font-size: 22px;font-weight:bold;font-family: 华文楷体">丰阳 欢迎您!</span>
        </el-divider>
      </el-row>

      <el-row :gutter="20" style="margin-left: 30px;margin-right: 30px">
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
import {ref} from "vue";
import news from '../../config/news.json';
import warning from '../../config/warning.json'
import { useRouter } from 'vue-router'

let carousel = ref(null)
let newslist = news.news_items
let warninlist =warning.warning_items
let advicelist = news.news_items

//首先在setup中定义
const router = useRouter()

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
const changelayout = (val) => {
  if(val==="enter")
  {//@mouseover="changelayout('enter')" @mouseleave="changelayout('leave')"
    document.getElementById("news_img").style.width = "100%";//让元素隐藏
    //document.getElementById("news_content").style.position = "absolute";
    document.getElementById("news_content").style.left = "10px";
    document.getElementById("news_content").style.top = "10px";
  }
  else
  {
    document.getElementById("news_img").style.width = "20%";//让元素隐藏
    document.getElementById("news_content").style.left = "120px";
  }

}

</script>

<style lang="scss" scoped>
h{
  color: rgba(244, 147, 45);
  font-size: 26px;
  font-weight: 700;
  margin-bottom: 20px
}

.grid-content {
  border-radius: 20px;
  width: 800px;
}
.part{
  height: 520px;
  width: 600px;
  position: relative;
  border-radius: 10px
}
.news{
  left: 5px;
  top:15px;
}
.news_mode{
  position: relative;
  margin: 5px;
  display: block;
}
.big_mode{
  position: relative;
  margin: 5px;
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
  left: 5px;
  height: 80%;
  width: 80%
}
.news_content{
  width: 70%;
  position: absolute;
  left:150px;
  top:5px;
  font-size: 17px;
  overflow: hidden; // 文字超长隐藏
  text-overflow:ellipsis; // 显示...
  white-space: nowrap; // 单行显示
}
.big_content{
  width: 80%;
  position: absolute;
  left:5px;
  bottom: 55px;
  font-size: 17px;
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
  right:20px;
  bottom:10px;
  z-index: 2
}
.news:hover .arrow{
  display: block;
}

.summary:hover{
  font-size: 20px;
}

.list{
  width: 100%;
  margin-left: 20px;
  margin-bottom: 10px;
  font-size: 18px;
  font-weight: 550;
}
.list:hover{
  transform: scale(1.02);
  color: rgba(0, 0, 255, 0.96);
}

.modeCard{
  float: left;
  width:220px;
  height: 240px;
  margin-left: 40px;
  margin-bottom: 8px;
}
.modeCard:hover{ //鼠标悬停时激活
transform: scale(1.05); //放大倍数
}

</style>
