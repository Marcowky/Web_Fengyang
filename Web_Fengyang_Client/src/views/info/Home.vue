<template>
  <!-- 顶部导航栏 -->
  <el-container>
    <el-card style="border-radius: 20px;margin-left: 120px;margin-right: 120px;margin-top: 60px">
      <el-row :gutter="5" style="margin-left: 40px;margin-right: 40px">
        <el-col :span="9" :offset="3" ><div class="grid-content ep-bg-purple" />
              <h style="color:#f4932d; font-size: 26px;font-weight: 700;margin-bottom: 20px">
                最新资讯
              </h>
          <el-card style="height: 500px;position: relative;border-radius: 10px">
              <el-carousel indicator-position="none" type="card" arrow="never" direction="vertical" :autoplay="false" ref="carousel" class="carousel">
                <el-carousel-item v-for="v in newslist" :key="v.value" style="width: 100%;height: inherit">
                    <img v-bind:src="v.res" alt="" />
                    <div class="summary">
                      {{v.content}}
                    </div>
                </el-carousel-item>
              </el-carousel>
            <div style="position: absolute;color: coral;right:40px;bottom:120px;z-index: 2">
              <h1 @click="arrowClick('left')" style="margin-bottom: 20px">↑</h1>
              <h2 @click="arrowClick('right')">↓</h2>
            </div>
          </el-card>
        </el-col>

        <el-col :span="9" :offset="1"><div class="grid-content ep-bg-purple" />
            <h style="color:#f4932d; font-size: 26px;font-weight: 700;margin-bottom: 20px">
              温馨提醒
            </h>
          <el-card style="height: 500px;border-radius: 10px">
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
      <el-row style="margin-top:10px; margin-left: 40px;margin-right: 80px">
        <el-divider>
          <span style="font-size: 20px">凤阳欢迎你！</span>
        </el-divider>
      </el-row>

      <el-row :gutter="20" style="margin-left: 30px;margin-right: 30px">
        <el-col :span="9" :offset="3"><div class="grid-content ep-bg-purple" />
            <h style="color:#f4932d; font-size: 26px;font-weight: 700;margin-bottom: 20px">
              旅游攻略
            </h>
          <el-card style="height: 500px;border-radius: 10px">
            <el-row>
              <el-card class="modeCard">
                <div style="margin-left: 10px">
                <img
                    src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                    style="width: 100%;height: inherit"
                />
                  <span>攻略1</span>
                </div>
              </el-card>
              <el-card class="modeCard">
                <div style="margin-left: 10px">
                  <img
                      src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                      style="width: 100%;height: inherit"
                  />
                  <span>攻略1</span>
                </div>
              </el-card>
            </el-row>
            <el-row>
              <el-card class="modeCard">
                <div style="margin-left: 10px">
                  <img
                      src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                      style="width: 100%;height: inherit"
                  />
                  <span>攻略1</span>
                </div>
              </el-card>
              <el-card class="modeCard">
                <div style="margin-left: 10px">
                  <img
                      src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                      style="width: 100%;height: inherit"
                  />
                  <span>攻略1</span>
                </div>
              </el-card>
            </el-row>
          </el-card>


        </el-col>
        <el-col :span="9" :offset="1"><div class="grid-content ep-bg-purple" />
            <h style="color:#f4932d; font-size: 26px;font-weight: 700;margin-bottom: 20px">
              重大时事
            </h>
          <el-card style="height: 500px;border-radius: 10px">
            <img
                src="../../assets/pic1.jpg"
                style="width: 500px;height: 300px;margin-right: 10px"
            />
          </el-card>

        </el-col>
      </el-row>
    </el-card>
  </el-container>

</template>

<script setup>
import {ref} from "vue";
import news from '../../config/news.json';
import warning from '../../config/warning.json'

let carousel = ref(null)
let newslist = news.news_items
let warninlist =warning.warning_items

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
  height: 250px;
  width: 100%
}

.grid-content {
  border-radius: 20px;
  width: 800px;
}
.carousel{
  height: 400px;
  width: 500px;
  margin-left: 20px;

}
.el-carousel__item:nth-child(2n) {
  background-color: rgba(255, 255, 255, 0.7);
}
.el-carousel__item:nth-child(2n + 1) {
  background-color: rgba(255, 255, 255, 0.7);

}
.summary{
  width: 500px;
  line-height: 40px;
  margin-left: 10px;
  font-size: 18px;
  font-weight: 500;
  color: aliceblue;
  background-color: rgba(1, 1, 2, 0.42);
  overflow: hidden; // 文字超长隐藏
  text-overflow:ellipsis; // 显示...
  white-space: nowrap; // 单行显示
  position:absolute;
  z-index:2;
  bottom: 50px;
}
.list{
  width: 100%;
  margin-left: 30px;
  margin-bottom: 10px;
  font-size: 18px;
  font-weight: 550;
}
.list:hover{
  transform: scale(1.05);
  background-color: rgba(246,3,12,0.2);
}


.modeCard{
  width:220px;
  height: 220px;
  margin-left: 25px;
  margin-bottom: 8px;
}
.modeCard:hover{ //鼠标悬停时激活
transform: scale(1.05); //放大倍数
}

</style>

