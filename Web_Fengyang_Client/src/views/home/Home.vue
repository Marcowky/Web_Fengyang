<template>
  <!--四色丰阳-->
  <Chart></Chart>
  <TopBar></TopBar>
  <div class="title-1" id="home1"><p>—四 色 丰 阳—</p></div>
  <div class="container-1">
    <el-row>
      <el-col :span="10" :offset="2" >
        <div class="image-container0">
          <div class="image-wrapper">
            <img class="responsive-image" src="../../assets/HomePage/index-pic-1.jpg">
          </div>
        </div>
      </el-col>
      <el-col :span="10"  class="txt-container">
          <h3 class="fengyang_subtitle">红色丰阳</h3>
          <div class="content_txt"><p>{{ red_fengyang }}</p></div>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="10"  :offset="2" class="txt-container">
        <h3 class="fengyang_subtitle">绿色丰阳</h3>
        <div class="content_txt"><p>{{ green_fengyang }}</p></div>
      </el-col>
      <el-col :span="10"  >
        <div class="image-container0">
          <div class="image-wrapper">
            <img class="responsive-image" src="../../assets/HomePage/index-pic-4.jpg">
          </div>
        </div>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="10" :offset="2" >
        <div class="image-container0">
          <div class="image-wrapper">
            <img class="responsive-image" src="../../assets/HomePage/index-pic-5.jpg">
          </div>
        </div>
      </el-col>
      <el-col :span="10"  class="txt-container">
        <h3 class="fengyang_subtitle">金色丰阳</h3>
        <div class="content_txt"><p>{{ gold_fengyang }}</p></div>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="10"  :offset="2" class="txt-container">
        <h3 class="fengyang_subtitle">灰色丰阳</h3>
        <div class="content_txt"><p>{{ gray_fengyang }}</p></div>
      </el-col>
      <el-col :span="10"  >
        <div class="image-container0">
          <div class="image-wrapper">
            <img class="responsive-image" src="../../assets/HomePage/index-pic-8.jpg">
          </div>
        </div>
      </el-col>
    </el-row>
  </div>


  <!--食宿阅论-->
  <div class="title-1" id="home2">
    <p>—食 宿 阅 论—</p>
  </div>
  <div class="container-2">
    <el-row>
      <el-col :span="6">
          <div class="image-container1">
            <el-link href="http://localhost:5173/consumption?category=1">
              <div class="image-wrapper">
                <img class="responsive-image" src="../../assets/HomePage/index-pic-9.jpg">
                <div class="text"><h3>食在丰阳</h3></div>
              </div>
            </el-link>
          </div>
      </el-col>

      <el-col :span="6">
        <div class="image-container1">
          <el-link href="http://localhost:5173/hotel">
            <div class="image-wrapper">
              <img class="responsive-image" src="../../assets/HomePage/index-pic-10.jpg">
              <div class="text"><h3>宿在丰阳</h3></div>
            </div>
          </el-link>
        </div>
      </el-col>

      <el-col :span="6">
        <div class="image-container1">
          <el-link href="http://localhost:5173/info?category=1">
            <div class="image-wrapper">
              <img class="responsive-image" src="../../assets/HomePage/index-pic-11.jpg">
              <div class="text"><h3>阅在丰阳</h3></div>
            </div>
          </el-link>
        </div>
      </el-col>

      <el-col :span="6">
        <div class="image-container1">
          <el-link href="http://localhost:5173/blog?category=1">
            <div class="image-wrapper">
              <img class="responsive-image" src="../../assets/HomePage/index-pic-12.jpg">
              <div class="text"><h3>论在丰阳</h3></div>
            </div>
          </el-link>
        </div>
      </el-col>

    </el-row>
  </div>

  <!--热门景点-->
  <div class="title-1" id="home3">
    <p>—热 门 景 点—</p>
  </div>
  <div class="container-3">
    <el-row :gutter="40">
      <el-col :span="6" :offset="3">
        <div class="image-container2" @click="toArticle('7b145fc0-bd04-4c2a-80ba-a7d098fc02e6')">
          <img class="responsive-image" src="../../assets/HomePage/index-pic-13.jpg" >
          <div class="text2"><h2>{{ row1_title1 }}</h2></div>
        </div>
      </el-col>

      <el-col :span="6" >
        <div class="image-container2" @click="toArticle('ee04916c-585c-417a-aaf5-26c5a4c35488')">
          <img class="responsive-image" src="../../assets/HomePage/index-pic-14.jpg" >
          <div class="text2"><h2>{{ row1_title2 }}</h2></div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="image-container2" @click="toArticle('5c73b37c-07cf-4cab-b6f9-0d8a3b83f816')">
          <img class="responsive-image" src="../../assets/HomePage/index-pic-15.jpg" >
          <div class="text2"><h2>{{ row1_title3 }}</h2></div>
        </div>
      </el-col>
    </el-row>
  </div>

  <div class="white-block"></div>
  <div class="container-4">
    <el-row justify="center" :gutter="10">

      <el-col :span="4">
        
          <div class="image-container3" @click="toArticle('dce2db86-1f2c-429f-823d-220e15bac7bd')">
            <img class="responsive-image" src="../../assets/HomePage/index-pic-16.jpg" >
            <div class="txtbox2">
              <div class="txtbox2_background">
                <p class="txt-title">{{ txt_title_1 }}</p>
              </div>
              <div class="txtbox2-content">
                <p class="txt-content">{{ txt_content_1 }}</p>
              </div>
            </div>
          </div>
        
      </el-col>

      <el-col :span="4">
        
          <div class="image-container3" @click="toArticle('d2a441f4-1472-4ba0-b161-8717e60298e2')">
            <img class="responsive-image" src="../../assets/HomePage/index-pic-17.jpg" >
            <div class="txtbox2">
              <div class="txtbox2_background">
                <p class="txt-title">{{ txt_title_2 }}</p>
              </div>
              <div class="txtbox2-content">
                <p class="txt-content">{{ txt_content_2 }}</p>
              </div>
            </div>
          </div>
        
      </el-col>

      <el-col :span="4">
        
          <div class="image-container3" @click="toArticle('af2b8343-43f9-4b47-98c7-8430b2fe89f4')">
            <img class="responsive-image" src="../../assets/HomePage/index-pic-18.jpg" >
            <div class="txtbox2">
              <div class="txtbox2_background">
                <p class="txt-title">{{ txt_title_3 }}</p>
              </div>
              <div class="txtbox2-content">
                <p class="txt-content">{{ txt_content_3 }}</p>
              </div>
            </div>
          </div>
        
      </el-col>

      <el-col :span="4">
        
          <div class="image-container3" @click="toArticle('19b9deb4-b9ac-4d00-844d-ea4807d9ed5f')">
            <img class="responsive-image" src="../../assets/HomePage/index-pic-19.jpg" >
            <div class="txtbox2">
              <div class="txtbox2_background">
                <p class="txt-title">{{ txt_title_4 }}</p>
              </div>
              <div class="txtbox2-content">
                <p class="txt-content">{{ txt_content_4 }}</p>
              </div>
            </div>
          </div>
        
      </el-col>

      <el-col :span="4">
        
          <div class="image-container3" @click="toArticle('665e238c-dcfd-472b-a7fe-a4c1ac883740')">
            <img class="responsive-image" src="../../assets/HomePage/index-pic-20.jpg" >
            <div class="txtbox2">
              <div class="txtbox2_background">
                <p class="txt-title">{{ txt_title_5 }}</p>
              </div>
              <div class="txtbox2-content">
                <p class="txt-content">{{ txt_content_5 }}</p>
              </div>
            </div>
          </div>
        
      </el-col>

    </el-row>
  </div>


</template>
const link = "http://localhost:5173/#/consumption"

<script setup>
import Chart from "./Chart.vue";
import TopBar from "./TopBar.vue";
import { useRoute, useRouter, onBeforeRouteUpdate } from 'vue-router'
import { ref, reactive, inject, onMounted } from 'vue'
const router = useRouter()
const route = useRoute()
const red_fengyang = ref("丰阳古镇位于连州市西北部，距离现今有一千多年的历史，一幢古庙、二条古街、三间祠堂、四座门楼，这里包揽了丰阳古村的全部内涵")
const green_fengyang = ref("古镇里的一山一水、一草一木同时也铸造了一道道迷人的风景线，穿梭其中，屋舍俨然，古朴且惬意，美景不断，恬静闲适的气息迎面而至")
const gold_fengyang = ref("这里山水依偎，低调安谧，闲适悠然，丰阳镇里还有一条始建于明朝成化年间的秀美清雅的畔水古村，村内现存明、清时期的历史传统建筑，多达47000平方米")
const gray_fengyang = ref("在村内至今还遗存着丰富的宋代人文景观和完好的明、清建筑，每每走过斑驳的石板古街、重檐飞翘的古宅都能感受浓浓的古村落气息")
const row1_title1 = ref("高山仰止")
const row1_title2 = ref("神凤来兮")
const row1_title3 = ref("风雨虹桥")
const txt_title_1 = ref("奇山峻岭")
const txt_content_1 = ref("如果说丰阳是一个栖息灵魂的“童话世界”，那么扎如就是一座翻飞激情的生态花园。扎如沟是丰阳的第四条沟，位于景区东北方向，全长22.6公里，景区面积80.1平方公里。扎如沟是丰阳风景区中唯一不以水景为特点的景区，沟内环境宁静素美、空气清晰，大气负氧离子浓度18400个/cm3，动植物资源十分丰富，还有一个扎如鹿园，徒步时，说不定路边就有小鹿对你打招呼哦。")
const txt_title_2 = ref("东陂泛舟")
const txt_content_2 = ref("东陂河是丰阳的母亲河，她依着城墙缓缓流淌，世世代代哺育着古城儿女。坐上乌篷船，听着舩公的号子，看着两岸已有百年历史的吊脚楼，别有一番韵味。顺水而下，穿过虹桥一幅水乡的画卷展现于眼前：万寿宫、万名塔、夺翠楼.....一种远离尘世的感觉悠然而生。")
const txt_title_3 = ref("东门城楼")
const txt_content_3 = ref("东门城楼位于城东，紧靠沱江，原名“升恒门”，为凤凰古城四大城门之一。东门城楼始建于清康熙五十四年（1715年），城门下部由紫红砂岩砌成，上部城楼则用古砖砌筑。登上城楼，抚摸着凹凸不平的城墙，时间仿佛就会在这一刻定格，恍惚中，依稀看到它如同像那有着湘西人刚毅的面庞的血性男儿一样，守护着沱江两岸，守护着这一弯流水。东陂河的南岸是古城墙，用紫红沙石砌成，典雅不失雄伟。城墙有东、北两座城楼，久经沧桑，依然壮观。东陂河水清澈，城墙边的河道很浅，水流悠游缓和，可以看到柔波里招摇水草，可以撑一支长篙漫溯。")
const txt_title_4 = ref("神凤文化景区")
const txt_content_4 = ref("一座青山抱古城，一湾沱水绕城过”。 横亘古城的南华山被誉为凤凰的父亲山，碧水逶迤的沱江是凤凰的母亲河。南华山国家森林公园位于凤凰古城南部，临沱江之畔，呈月牙形环抱凤凰古城，神凤文化景区就位于南华山国家森林公园境内，景区入口毗邻虹桥，依山傍水，俯瞰整个凤凰古城。")
const txt_title_5 = ref("虹桥风雨楼")
const txt_content_5 = ref("虹桥又名“卧虹桥”或“风雨桥”，位于凤凰古城东门外迴龙阁沱江河之上。始建于明洪武年间，明嘉靖三十五年（公元1556年）、清康熙九年（1670年）、五十四年（1715年）进行加固维修；民国三年（1914年）湘西镇守使田应诏首倡集资重加修葺，更名为“虹桥”。该桥南北走向，桥型为半圆拱联拱石桥，两墩三孔，为红砂石砌筑。1955年修建成公路，2000年修复为虹桥风雨楼，具有极高的观赏价值和文物保护价值。2011年被确立为湖南省第九批省级文物保护单位。")

const toArticle = (articleId) =>{
    router.push({
        path: "/article/detail",
        query: {
            category: "attractionArticle",
            id: articleId,
        }
    })
    console.log(articleId)
}
  // mounted() {
  //   // 改变网页背景颜色
  //   document.querySelector('body').setAttribute('style', 'background-color:#d9e0dc');
  // },
  // beforeDestroy() {
  //   document.querySelector('body').removeAttribute('style')
  // },

</script>

<style scoped>
.title-1 {
  width: 100%;
  font-size: 80px;
  font-family: STXingkai;
  font-weight: 100;
  color: #474141;
  margin-top: 100px;
  margin-bottom: 10px;
  display: flex;
  justify-content: center;
}

img{
  transition: all .5s;
}

.container-1 {
  .image-container0 {
    aspect-ratio: 670/360;
    .image-wrapper {
      width: 100%;
      height: 100%;
      overflow: hidden;
      align-items: center;
      justify-content: center;
      .responsive-image {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
    }
  }
  .txt-container {
    background: url(../../assets/HomePage/index-pic-2.jpg);
    .fengyang_subtitle {
      color: rgb(0, 0, 0);
      font-size: 34px;
      top: 6%;
      margin: 20px;
      text-align: center;
      font-family: 宋体;
      
    }
    .content_txt
    {
      top: 8%;
      left: 5%;
      width: 90%;
      text-indent:2em;
      font-size: 18px;
      font-family: 宋体;
      color: rgb(0, 0, 0);
      display: flex;
      text-align: center;
      word-break: keep-all;
      line-height: 30px;
      
    }
  }
}

.image-container0:hover img {
    transform: scale(1.1);
  }

.container-2{
  margin-left: 3%;
  margin-right: 3%;
  .image-container1 {
    aspect-ratio: 300/400;

      .image-wrapper {
        position: relative;
        width: 100%;
        height: 100%;
        overflow: hidden;
        display: flex;
        align-items: center;
        justify-content: center;
        .responsive-image {
          width: 100%;
          height: 100%;
          object-fit: cover;
        }
      }
  }
  .text {
     position: absolute;
     width: 100%;
     height: 100%;
     background: rgba(0, 0, 0, 0.5);
     opacity: 0;
     transition: opacity 0.35s ease-in-out;
     display: flex;
     justify-content: center;/* 水平居中 */
     align-items: center;/* 垂直居中 */
     
   }
  .text h3 {
    
    color: #fff;
    font-size: 40px;
    font-family: 宋体;
    width: 333px;
    writing-mode: vertical-rl;
  }
}
.image-container1:hover .text {
  opacity: 1;
}
.image-container1:hover img {
    transform: scale(1.1);
  }



.container-3 {
  display: flex;
  align-content: center;
  .image-container2{
    aspect-ratio: 1/1;
    position: relative;
    border-radius: 50%;
    /*将容器裁成圆形*/
    overflow: hidden;
    /*在容器内超过容器的元素 都会被隐藏*/
    .responsive-image{
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
    .text2{
      position: absolute;
      top: 50%;
      left: 50%;
      width: 100%;
      height: 20%;
      overflow: hidden;
      transform: translate(-50%, -50%);
      background: rgba(0, 0, 0, 0.5);
      opacity: 0;
      transition: opacity 0.35s ease-in-out;
    }
    .text2 h2{
      padding: 10px;
       text-align: center;
       font-family: 宋体;
       font-size: 40px;
       color: #f5f5f5;
     }
  }
  .image-container2:hover img{
    transform: scale(1.1);
  }
  .image-container2:hover .text2{
    opacity: 1;
  }
}

.white-block
{
  height: 40px;
}


.container-4 .image-container3 {
  aspect-ratio: 275/230;
  position: relative;
  overflow: hidden;
  .responsive-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
  .txtbox2 {
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    transition: all .5s ease-in-out;
    display: flex;
    align-items: center;
    justify-content: center;
    text-align: center;
    .txtbox2_background {
      position: absolute;
      bottom: 0;
      z-index: 1;
      height: 23%;
      width: 100%;
      background: rgba(0, 0, 0, 0.5);
      overflow: hidden;
    }
    .txt-title {
      font-family: 宋体;
      font-weight: 600;
      font-size: 23px;
      color: #f5f5f5;
      line-height: 100%;
    }
    .txt-content {
      top: 47%;
      font-family: 宋体;
      font-size: 16px;
      color: #f5f5f5;
      text-indent: 2em;
    }
  }
}

.container-4 .image-container3 .txtbox2 .txtbox2-content {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 23%;
  background: rgba(0, 0, 0, 0.5);
  opacity: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  transition: all 0.35s ease-in-out;
  transform: translateY(100%);
  /* transition: transform 0.5s ease, opacity 0.5s ease; */
}

.container-4 .image-container3:hover .txtbox2 .txtbox2-content {
  opacity: 1;
  transform: translateY(0);
}
</style>