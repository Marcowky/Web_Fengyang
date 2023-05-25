<template>
  <!-- 顶部导航栏 -->
  <el-container>
    <el-card style="border-radius: 20px;margin-left: 120px;margin-right: 120px;margin-top: 60px">
      <el-row :gutter="5" style="margin-left: 40px;margin-right: 40px">
        <el-col :span="9" :offset="2" ><div class="grid-content ep-bg-purple" />
          <el-container>
            <el-header>
              <h style="color:#f4932d; font-size: 26px;font-weight: 700;margin-bottom: 20px">
                最新资讯
              </h>
            </el-header>
            <el-main>
              <el-carousel indicator-position="none" type="card" arrow="never" height="400px" direction="vertical" :autoplay="false" ref="carousel">
                <el-carousel-item v-for="v in newslist" :key="v.value" style="width: 350px;height: 330px">
                  <img v-bind:src= "v.res" alt="">
                  <div style="font-size: 15px;font-weight: 500">
                    {{v.content}}
                    <router-link :to="{path: '/info/page'}"
                                 tag="button">>></router-link>
                  </div>
                </el-carousel-item>

              </el-carousel>
            </el-main>
            <div  style="position:absolute; top: 10px; left: 400px; height: 20px; width: 10px; z-index: 1">
              <h1 @click="arrowClick('left')">↑</h1>
              <h2 @click="arrowClick('right')">↓</h2>
            </div>
          </el-container>
        </el-col>

        <el-col :span="9" :offset="4"><div class="grid-content ep-bg-purple" />
          <el-row>
            <h style="color:#e67d6b; font-size: 20px;font-weight: 500;margin-bottom: 20px">
              温馨提醒
            </h>
          </el-row>
          <el-card style="width:300px">
            <el-scrollbar height="200px">
              <el-row v-for="v in warninlist" :key="v.value">
                <div style="line-height: 30px">
                  {{v.content}}
                  <router-link :to="{path: '/info/page'}"
                               tag="button">>></router-link>
                </div>
              </el-row>
            </el-scrollbar>
          </el-card>

        </el-col>
      </el-row>
      <el-row :gutter="20" style="margin-left: 30px;margin-right: 30px">
        <el-col :span="9" :offset="2"><div class="grid-content ep-bg-purple" />
          <el-row>
            <h style="color:#f4932d; font-size: 26px;font-weight: 700;margin-bottom: 20px">
              旅游攻略
            </h>
          </el-row>
          <el-row>
            <el-card class="modeCard">
              <img
                  src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                  style="margin-left:10px;width: 90%;height: 90%"
              />
              <div style="margin-left: 10px">
                <span>攻略1</span>
              </div>
            </el-card>
            <el-card class="modeCard">
              <img
                  src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                  style="margin-left:10px;width: 90%;height: 90%"
              />
              <div style="margin-left: 10px">
                <span>攻略2</span>
              </div>
            </el-card>
          </el-row>
          <el-row>
            <el-card class="modeCard">
              <img
                  src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                  style="margin-left:10px;width: 90%;height: 90%"
              />
              <div style="margin-left: 10px">
                <span>攻略3</span>
              </div>
            </el-card>
            <el-card class="modeCard">
              <img
                  src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                  style="margin-left:10px;width: 90%;height: 90%"
              />
              <div style="margin-left: 10px">
                <span>攻略4</span>
              </div>
            </el-card>
          </el-row>

        </el-col>
        <el-col :span="9" :offset="4"><div class="grid-content ep-bg-purple" />
          <el-row>
            <h style="color:#f4932d; font-size: 20px;font-weight: 500;margin-bottom: 20px">
              重大时事
            </h>
          </el-row>
          <el-row>
            <img
                src="../../assets/pic1.jpg"
                style="width: 500px;height: 300px;margin-right: 10px"
            />
          </el-row>

        </el-col>
      </el-row>
    </el-card>
  </el-container>

</template>

<script setup>
// 导入顶部栏
import TopBar from "../../components/TopBar.vue"
import FooterBar from "../../components/FooterBar.vue";
import {ref} from "vue";
import news from '../../config/news.json';
import warning from '../../config/warning.json'

let carousel = ref(null)
let newslist = news.news_items
let warninlist =warning.warning_items

// 侧边栏和导航栏的点击触发函// 数
const handleSelect = (index) => {    // 这里可以定义导航栏点击函数

};

const arrowClick = (val) => {
  if(val === 'right') {
    carousel.value.next()
  } else {
    carousel.value.prev()
  }
}
</script>

<style lang="scss" scoped>
img{
  height: 200px;
  width: 100%
}
h1{
  font-size: 30px;
  color: coral;
  margin-top: 200px;
}
h2{
  font-size: 30px;
  color: coral;
  margin-top: 60px;
}
.grid-content {
  border-radius: 20px;
  width: 800px;
}
.modeCard{
  width:220px;
  height: 220px;
  margin-right: 10px;
  margin-bottom: 10px;
}
.modeCard:hover{ //鼠标悬停时激活
transform: scale(1.05); //放大倍数
}

.el-carousel__item:nth-child(2n) {
  background-color: rgba(255, 255, 255, 0.7);
}
.el-carousel__item:nth-child(2n + 1) {
  background-color: rgba(255, 255, 255, 0.7);

}

</style>
