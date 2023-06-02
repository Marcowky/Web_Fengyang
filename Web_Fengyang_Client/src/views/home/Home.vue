<template>
  <!--四色丰阳-->
  <TopBar></TopBar>
  <Chart></Chart>
  <div class="title-1">
    <p>—四 色 丰 阳—</p>
  </div>
  <div class="container-1">
    <ul>
      <li class="list-1">
        <a>
          <img class="pic-1" src="../../assets/HomePage/index-pic-1.jpg" alt="index-pic-1">
        </a>
      </li>
      <li class="list-2">
        <h3>红色丰阳</h3>
        <div class="fengyang_txt">
          <p>{{ red_fengyang }}</p>
        </div>
      </li>
      <li class="list-3">
        <h3>绿色丰阳</h3>
        <div class="fengyang_txt">
          <p>{{ green_fengyang }}</p>
        </div>
      </li>
      <li class="list-4">
        <a>
          <img class="pic-4" src="../../assets/HomePage/index-pic-4.jpg" alt="index-pic-4">
        </a>
      </li>
      <li class="list-5">
        <a>
          <img class="pic-5" src="../../assets/HomePage/index-pic-5.jpg" alt="index-pic-5">
        </a>
      </li>
      <li class="list-6">
        <h3>金色丰阳</h3>
        <div class="fengyang_txt">
          <p>{{ gold_fengyang }}</p>
        </div>
      </li>
      <li class="list-7">
        <h3>灰色丰阳</h3>
        <div class="fengyang_txt">
          <p>{{ gray_fengyang }}</p>
        </div>
      </li>
      <li class="list-8">
        <a>
          <img class="pic-8" src="../../assets/HomePage/index-pic-8.jpg" alt="index-pic-8">
        </a>
      </li>
    </ul>
  </div>

  <!--食宿阅论-->
  <div class="title-1">
    <p>—食 宿 阅 论—</p>
  </div>
  <div class="container-2">
    <ul>
      <li class="block2-li">
        <el-link href="http://localhost:5173/#/consumption">
          <div class="imgBox">
            <img src="../../assets/HomePage/index-pic-9.jpg" alt="index-pic-9">
            <div class="text">
                 <h3>食在丰阳</h3>
            </div>
          </div>
        </el-link>
      </li>
      <li class="block2-li">
        <el-link href="http://localhost:5173/#/travel">
          <div class="imgBox">
            <img src="../../assets/HomePage/index-pic-10.jpg" alt="index-pic-10">
            <div class="text">
                <h3>宿在丰阳</h3>
            </div>
          </div>
        </el-link>
      </li>
      <li class="block2-li">
        <el-link href="http://localhost:5173/#/info">
          <div class="imgBox">
            <img src="../../assets/HomePage/index-pic-11.jpg" alt="index-pic-11">
            <div class="text">
                <h3>阅在丰阳</h3>
            </div>
          </div>
        </el-link>
      </li>
      <li class="block2-li">
        <el-link href="http://localhost:5173/#/blog">
          <div class="imgBox">
            <img src="../../assets/HomePage/index-pic-12.jpg" alt="index-pic-12">
            <div class="text">
                <h3>论在丰阳</h3>
            </div>
          </div>
        </el-link>
      </li>
    </ul>
  </div>

  <!--热门景点-->
  <div class="title-1">
    <p>—热 门 景 点—</p>
  </div>
  <div class="container-3">
    <ul>
      <li class="block2-li">
        <el-link href="http://localhost:5173/#/consumption">
          <div class="imgBox-2">
            <img class="block3-img" width=448px src="../../assets/HomePage/index-pic-13.jpg" alt="index-pic-13">
            <div class="bg-h2">
              <h2>{{ row1_title1 }}</h2>
            </div>
          </div>
        </el-link>

        <el-link href="http://localhost:5173/#/consumption">
          <div class="imgBox-2">
            <img class="block3-img" width=448px src="../../assets/HomePage/index-pic-14.jpg" alt="index-pic-14">
            <div class="bg-h2">
              <h2>{{ row1_title2 }}</h2>
            </div>
          </div>
        </el-link>

        <el-link href="http://localhost:5173/#/consumption">
          <div class="imgBox-2">
            <img class="block3-img" width=448px src="../../assets/HomePage/index-pic-15.jpg" alt="index-pic-15">
            <div class="bg-h2">
              <h2>{{ row1_title3 }}</h2>
            </div>
          </div>
        </el-link>
      </li> 
    </ul>
</div>

<div class="container-4">
    <ul>
      <li class="block2-li">
        <el-link href="http://localhost:5173/#/consumption">
          <div class="imgBox-3">
            <img class="block3-img-2" width=448px src="../../assets/HomePage/index-pic-16.jpg" alt="index-pic-9">
              <div class="txtbox">
                <div class="bgtxt">
                  <p class="txt-title">{{ txt_title_1 }}</p>
                </div>
                <div class="bgtxt-content">
                  <p class="txt-content">{{ txt_content_1 }}</p>
                </div>
              </div>
          </div>
        </el-link>

        <el-link href="http://localhost:5173/#/consumption">
          <div class="imgBox-3">
            <img class="block3-img-2" width=448px src="../../assets/HomePage/index-pic-17.jpg" alt="index-pic-9">
              <div class="txtbox">
                <div class="bgtxt">
                  <p class="txt-title">{{ txt_title_2 }}</p>
                </div>
                <div class="bgtxt-content">
                  <p class="txt-content">{{ txt_content_2 }}</p>
                </div>
              </div>
          </div>
        </el-link>

        <el-link href="http://localhost:5173/#/consumption">
          <div class="imgBox-3">
            <img class="block3-img-2" width=448px src="../../assets/HomePage/index-pic-18.jpg" alt="index-pic-9">
              <div class="txtbox">
                <div class="bgtxt">
                  <p class="txt-title">{{ txt_title_3 }}</p>
                </div>
                <div class="bgtxt-content">
                  <p class="txt-content">{{ txt_content_3 }}</p>
                </div>
              </div>
          </div>
        </el-link>

        <el-link href="http://localhost:5173/#/consumption">
          <div class="imgBox-3">
            <img class="block3-img-2" width=448px src="../../assets/HomePage/index-pic-19.jpg" alt="index-pic-9">
              <div class="txtbox">
                <div class="bgtxt">
                  <p class="txt-title">{{ txt_title_4 }}</p>
                </div>
                <div class="bgtxt-content">
                  <p class="txt-content">{{ txt_content_4 }}</p>
                </div>
              </div>
          </div>
        </el-link>

        <el-link href="http://localhost:5173/#/consumption">
          <div class="imgBox-3">
            <img class="block3-img-2" width=448px src="../../assets/HomePage/index-pic-20.jpg" alt="index-pic-9">
              <div class="txtbox">
                <div class="bgtxt">
                  <p class="txt-title">{{ txt_title_5 }}</p>
                </div>
                <div class="bgtxt-content">
                  <p class="txt-content">{{ txt_content_5 }}</p>
                </div>
              </div>
          </div>
        </el-link>
      </li>
    </ul>
  </div>

</template>
const link = "http://localhost:5173/#/consumption"

<script>

import Chart from "./Chart.vue";
import TopBar from "./TopBar.vue";
export default {
  components: {
    Chart,
    TopBar,
},
  data() {
    return {
      red_fengyang: "丰阳古镇位于连州市西北部，距离现今有一千多年的历史，一幢古庙、二条古街、三间祠堂、四座门楼，这里包揽了丰阳古村的全部内涵",
      green_fengyang: "古镇里的一山一水、一草一木同时也铸造了一道道迷人的风景线，穿梭其中，屋舍俨然，古朴且惬意，美景不断，恬静闲适的气息迎面而至",
      gold_fengyang: "这里山水依偎，低调安谧，闲适悠然，丰阳镇里还有一条始建于明朝成化年间的秀美清雅的畔水古村，村内现存明、清时期的历史传统建筑，多达47000平方米",
      gray_fengyang: "在村内至今还遗存着丰富的宋代人文景观和完好的明、清建筑，每每走过斑驳的石板古街、重檐飞翘的古宅都能感受浓浓的古村落气息",
      row1_title1: "移石瀑",
      row1_title2: "木香宫",
      row1_title3: "碧湖桥",
      txt_title_1: "象鼻山",
      txt_content_1: "依船靠岸，走上台阶，你会看到一座雄伟的会馆，她设计精巧独特，雕琢玲珑剔透，或飞檐翘角，或回廊游转，或卧龙啸空，或奇兽驰骋，可以说是点雨楼台荟萃的建筑艺术大观。万寿宫内则是我国著名画家黄永玉老先生艺术画展，艺术馆内展示了一代画坛鬼才黄永玉不同时期的创作作品，和为家乡创作的《销魂的九夏芙蓉三湘烟水》、《沱江雨岸》、《故乡荷塘》、《永不回来的风景》四副艺术真迹均在馆内展出。在这里你可感受全才鬼才、博学而不群、倔强、睿智风趣、笔耕不戳、重情重义、讲原则的黄永玉.....",
      txt_title_2: "沱江泛舟",
      txt_content_2: "沱江河是凤凰古城的母亲河，她依着城墙缓缓流淌，世世代代哺育着古城儿女。坐上乌篷船，听着舩公的号子，看着两岸已有百年历史的土家吊脚楼，别有一番韵味。顺水而下，穿过虹桥一幅江南水乡的画卷展现于眼前：万寿宫、万名塔、夺翠楼.....一种远离尘世的感觉悠然而生。沱江的南岸是古城墙，用紫红沙石砌成，典雅不失雄伟。城墙有东、北两座城楼，久经沧桑，依然壮观。沱江河水清澈，城墙边的河道很浅，水流悠游缓和，可以看到柔波里招摇水草，可以撑一支长篙漫溯。沿沱江边而建的吊脚楼群在东门虹桥和北门跳岩附近，细脚伶仃的立在沱江里，像一幅永不回来的风景。",
      txt_title_3: "东门城楼",
      txt_content_3: "东门城楼位于城东，紧靠沱江，原名“升恒门”，为凤凰古城四大城门之一。东门城楼始建于清康熙五十四年（1715年），城门下部由紫红砂岩砌成，上部城楼则用古砖砌筑。登上城楼，抚摸着凹凸不平的城墙，时间仿佛就会在这一刻定格，恍惚中，依稀看到它如同像那有着湘西人刚毅的面庞的血性男儿一样，守护着沱江两岸，守护着这一弯流水。",
      txt_title_4: "神凤文化景区",
      txt_content_4: "一座青山抱古城，一湾沱水绕城过”。 横亘古城的南华山被誉为凤凰的父亲山，碧水逶迤的沱江是凤凰的母亲河。南华山国家森林公园位于凤凰古城南部，临沱江之畔，呈月牙形环抱凤凰古城，神凤文化景区就位于南华山国家森林公园境内，景区入口毗邻虹桥，依山傍水，俯瞰整个凤凰古城。",
      txt_title_5: "虹桥风雨楼",
      txt_content_5: "虹桥又名“卧虹桥”或“风雨桥”，位于凤凰古城东门外迴龙阁沱江河之上。始建于明洪武年间，明嘉靖三十五年（公元1556年）、清康熙九年（1670年）、五十四年（1715年）进行加固维修；民国三年（1914年）湘西镇守使田应诏首倡集资重加修葺，更名为“虹桥”。该桥南北走向，桥型为半圆拱联拱石桥，两墩三孔，为红砂石砌筑。1955年修建成公路，2000年修复为虹桥风雨楼，具有极高的观赏价值和文物保护价值。2011年被确立为湖南省第九批省级文物保护单位。",
    }
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
  },
  // mounted() {
  //   // 改变网页背景颜色
  //   document.querySelector('body').setAttribute('style', 'background-color:#d9e0dc');
  // },
  // beforeDestroy() {
  //   document.querySelector('body').removeAttribute('style')
  // },
};
</script>

<style scoped>

.title-1 {
  width: 100%;
  font-size: 80px;
  font-family: STXingkai;
  font-weight: 100px;
  color: #474141;
  margin-top: 0px;
  margin-bottom: 10px;
  display: flex;
  justify-content: center;
}

.container-1 {
  position: relative;
  width: 100%;
  height: 1440px;
  margin-bottom: 100px;
  padding-top: 0;
}

.container-1 .list-1 {
  top: 0;
}

.container-1 .list-1 .pic-1{
  width: 672px;
  height: 360px;
}

.container-1 .list-2 {
  position: absolute;
  top: 0;
  right: 40px;
  height: 360px;
  width: 672px;
  background: url(../../assets/HomePage/index-pic-2.jpg);
}

.container-1 .list-3 {
  position: absolute;
  top: 360px;
  left: 40px;
  height: 360px;
  width: 672px;
  background: url(../../assets/HomePage/index-pic-3.jpg) ;
}

.container-1 .list-4 {
  position: absolute;
  top: 360px;
  right: 40px;
}

.container-1 .list-4 .pic-4{
  top: 0;
  width: 672px;
  height: 360px;
  left: 0px;
}
.container-1 .list-5 {
  position: absolute;
  top: 720px;
}

.container-1 .list-5 .pic-5{
  width: 672px;
  height: 360px;
}

.container-1 .list-6 {
  position: absolute;
  top: 720px;
  right: 40px;
  height: 360px;
  width: 672px;
  background: url(../../assets/HomePage/index-pic-2.jpg) ;
}

.container-1 .list-7 {
  position: absolute;
  top: 1080px;
  left: 40px;
  height: 360px;
  width: 672px;
  background: url(../../assets/HomePage/index-pic-3.jpg) ;
}

.container-1 .list-8 {
  position: absolute;
  top: 1080px;
  right: 40px;
}

.container-1 .list-8 .pic-8{
  width: 672px;
  height: 360px;
}

.li, h3 {
  color: #000;
  font-size: 30px;
  font-family: 宋体;
  font-weight: 550;
  margin-top: 100px;
  margin-left: 285px;
}

ul,
li {
  list-style: none;
}

a {
  text-decoration: none;
}

.fengyang_txt {
  width: 672px;
  height: 200px;
  font-size: 16px;
  font-family: 宋体;
  color: #000;
  padding-top: 20px;
  padding-left: 30px;
  padding-right: 30px;
  display: flex;
  text-align: center;
  word-break: keep-all;
  line-height: 25px;
}

.container-2 {
  position: relative;
  width: 100%;
  height: 500px;
  margin-bottom: 100px;
  overflow:hidden
}
.container-2 .block2-li {
  float: left;
}

.imgBox {
  position: relative;
  width: 333px;
  overflow: hidden;
  height: 500px;
}
img {
  width: 333px;
  height: 500px;
  transition: all .5s;
}

.imgBox .text {
  position: absolute;
  width: 333px;
  height: 500px;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  text-align: center;
  -webkit-backface-visibility: hidden;
  backface-visibility: hidden;
  background: rgba(0, 0, 0, 0.5);
  opacity: 0;
  -webkit-transition: all 0.35s ease-in-out;  
  -moz-transition: all 0.35s ease-in-out;
  transition: all 0.35s ease-in-out;
}
.imgBox .text h3 {
  color: rgb(243, 243, 240);
  font-size: 40px;
  font-family: 宋体;
  width: 333px;
  margin-left: 0px;
  margin-top: 210px;
}

.imgBox:hover .text {
  opacity: 1;   
}

.imgBox:hover img{
  transform: scale(1.1);
}

.container-3 {
  position: relative;
  width: 100%;
  height: 450px;
}

.container-3 .imgBox-2 {
  width: 410x;
  height: 410px;
  border-radius: 50%;
  margin-right: 40px;
  overflow: hidden;
}

.container-3 .imgBox-2 .block3-img {
  width: 410px;
  height: 410px;
  border-radius: 50%;
}

.bg-h2 {
  position: absolute;
  top: 155px;
  width: 410px;
  height: 110px;
  opacity: 0;
  overflow: hidden;
  display: flex;
  justify-content: center;
  background: rgba(0, 0, 0, 0.5);
}
.bg-h2
h2{
  font-family: 宋体;
  line-height: 110px;
  font-weight: 600;
  position: absolute;
  font-size: 40px;

  color: #f5f5f5;
}

.imgBox-2:hover img{
  transform: scale(1.1);
}

.imgBox-2:hover .bg-h2 {
  opacity: 1;
}

.container-4 {
  position: relative;
  width: 100%;
  height: 230px;
}

.container-4 .imgBox-3 {
  width: 261px;
  height: 230px;
  margin-right: 10px;
}

.container-4 .imgBox-3 .block3-img-2 {
  width: 261px;
  height: 230px;
  margin-right: 10px;
}


.container-4 .txtbox {
  position: absolute;
  /* bottom: -180px; */
  top: 0;
  width: 261px;
  height: 230px;
  transition: all .5s ease-in-out;
  /* z-index: 3; */
  /* -webkit-transition: all 0.35s ease-in-out;  
  -moz-transition: all 0.35s ease-in-out;
  transition: all 0.35s ease-in-out; */
}

.container-4 .txtbox .bgtxt {
  position: absolute;
  /* top: 0; */
  bottom: 0;
  left: 0;
  width: 261px;
  height: 50px;
  display: flex;
  justify-content: center;
  opacity: 1;
  z-index: 1;
  background: rgba(0, 0, 0, 0.5);
}

.container-4 .txtbox .bgtxt .txt-title {
  font-family: 宋体;
  font-weight: 600;
  font-size: 23px;
  color: #f5f5f5;
  line-height: 50px;
}

.container-4 .txtbox .bgtxt-content {
  position: absolute;
  top: 230px;
  width: 261px;
  height: 180px;
  background: rgba(0, 0, 0, 0.5);
  opacity: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  transition: all 0.35s ease-in-out;
}

.txt-content {
  top: 100px;
  font-family: 宋体;
  font-size: 16px;
  color: #f5f5f5;
  text-indent: 13%;
}

.imgBox-3:hover .bgtxt-content {
  position: absolute;
  top: 0px;
  left: 0;
  opacity: 1;
}
</style>
