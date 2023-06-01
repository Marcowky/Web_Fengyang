
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
              <el-carousel indicator-position="none" type="card" arrow="never" direction="vertical" :autoplay="false" ref="carousel" class="carousel">
                <el-carousel-item v-for="v in newslist" :key="v.value">
                    <img v-bind:src="v.res" alt=""/>
                    <div class="summary" @click="showdetail()">
                      {{v.content}}
                    </div>
                </el-carousel-item>
                <div class="arrow">
                  <link rel="stylesheet" href="http://at.alicdn.com/t/c/font_4099046_8f48c4d9r8b.css">
                  <div style="transform: scale(1.6)">
                    <el-icon class="iconfont icon-jiantouloukong-shang" @click="arrowClick('left')">
                    </el-icon>
                  </div>
                  <div style="transform: scale(1.6);margin-top: 20px">
                    <el-icon class="iconfont icon-jiantouloukong-xia" @click="arrowClick('right')">
                    </el-icon>
                  </div>
                </div>
              </el-carousel>

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
          <span style="font-size: 20px;font-weight:bold;font-family: 华文楷体">丰阳 欢迎您!</span>
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
/*const showarrow = (val)  => {
  if(val === 'enter')
  {
    document.getElementById("arrow").style.display = "block";
  }
  else
  {
    document.getElementById("arrow").style.display = "none";
    arrow.display = false
  }
};*/
</script>

<style lang="scss" scoped>
h{
  color: rgba(244, 147, 45);
  font-size: 26px;
  font-weight: 700;
  margin-bottom: 20px
}
img{
  height: 250px;
  width: 100%
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
.carousel{
  height: 400px;
  width: 500px;
  margin-left: 20px;
}
.arrow{
  display: none;
  position: absolute;
  color: rgba(220, 156, 125, 0.8);
  right:20px;
  bottom:10px;
  z-index: 2
}
.carousel:hover .arrow{
  display: block;
}
.el-carousel__item{
  width: 100%;
  height: inherit;
}
.el-carousel__item:nth-child(2n) {
  background-color: rgba(255, 255, 255, 0.7);
}
.el-carousel__item:nth-child(2n + 1) {
  background-color: rgba(255, 255, 255, 0.7);
}
.summary{
  width: 490px;
  line-height: 40px;
  margin-left: 10px;
  font-size: 18px;
  font-weight: normal;
  color: aliceblue;
  background-color: rgba(1, 1, 2, 0.42);
  overflow: hidden; // 文字超长隐藏
  text-overflow:ellipsis; // 显示...
  white-space: nowrap; // 单行显示
  position:absolute;
  z-index:2;
  bottom: 50px;
}
.summary:hover{
  font-size: 20px;
  font-weight: bold;
}

.list{
  width: 100%;
  margin-left: 30px;
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
