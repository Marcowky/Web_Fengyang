<template>
  <!--轮播图-->
  <Chart style=" position: relative; left: 0;  width: 100%; " />
  <br><br><br>

  <!--概况-->
  <div class="general-title">
    <h3>—景 区 概 况—</h3>
  </div><br>

  <img class="IMG1" src="../../assets/spring.jpg" alt="">
  <div class="TXT1">
    <span>{{ spring_mes }}</span>
  </div>

  <img class="IMG2" src="../../assets/summer.jpg" alt="">
  <div class="TXT2">
    <span>{{ summer_mes }}</span>
  </div>

  <img class="IMG1" src="../../assets/autumn.jpg" alt="">
  <div class="TXT1">
    <span>{{ autumn_mes }}</span>
  </div>

  <img class="IMG2" src="../../assets/winter.jpg" alt="">
  <div class="TXT2">
    <span>{{ winter_mes }}</span>
  </div>
</template>

<script>
// 页面中侧边栏与导航栏的设置
// 导入顶部栏、轮播图、页脚、右侧悬浮栏
import Chart from "../consumption/Chart.vue"

export default {
  components: {   
    Chart,
  },
  data() {
    return {
      spring_mes: "四季在一年中交替出现，一年四季的变化是连绵不断的过程。其也是阴阳转换，此消彼长、量变产生质变的过程。在这个渐变的过程中，立春、立夏、立秋、立冬，并称“四立”，都是季节的开启。春季排四季之首，新的轮回从此开启。春，代表着温暖、生长。春季，阴阳之气开始转变，万物随阳气上升而萌牙生长，春季是万物生长、春耕播种的季节，所谓“一年之计在于春”。",
      summer_mes: "夏季万物至此皆盛，是农作物进入旺季生长的一个季节。温度升高，天气炎热，狂风暴雨频发，万物盛长。夏季是风暴最多的季节，常伴随着大风、暴雨。北半球夏季，大陆受热气压上升形成低气压，海洋恒温温度相对较低形成高气压，根据环流，夏季吹的是东南风。夏季北半球各地的白昼时间全年最长。",
      autumn_mes: "秋季是收获季节，意味着万物开始从繁茂成长趋向萧索成熟。秋季的前两个节气立秋、处暑，还是酷热天气，因太阳所带来的热力未减弱，所谓“热在三伏”，三伏天出现在小暑与处暑之间，是一年中气温最高且又潮湿、闷热的日子。造成三伏天湿度高原因是，三伏天吹东南风，而东南方是太平洋和印度洋，空气潮湿，风的潮湿造成了三伏天湿度大；到了深秋之后则相反，吹西北风，而西北方是干燥的内陆，干燥的西北风造成深秋之后气候干燥。",
      winter_mes: "冬季，阴阳转变，万物由收到藏，植物生气闭蓄。进入立冬后意味着秋季少雨干燥气候渐过去，转为阴雨寒冻的冬季气候特征。北半球的太阳高度小，白昼时间短。所谓“热在三伏”，冷在三九”，冬至为“一九天”，冬至之前还不算很冷，冬至过后强冷空气南下频繁，并越过了南岭，气温骤降，天气寒冷。冬季真正寒冷，是在冬至后。",
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
.general-title {
  font-family: 宋体;
  color: black;
  font-size: 45px;
  font-weight: bolder;
  text-align: center;
}

.IMG1 {
  width: 300px;
  height: 448px;
  float: left;
  /*这句一定要有*/
  display: inline;
  margin-left: 100px;
  margin-right: 20px;
}

.TXT1 {
  height: 448px;
  text-indent: 2em;
  /* 首行缩进 */
  letter-spacing: 1px;
  /* 字间距 */
  line-height: 45px;
  /* 行间距 */
  color: black;
  margin-right: 100px;
}

.TXT1 span {
  top: 100px;
}

.IMG2 {
  width: 300px;
  height: 448px;
  float: right;
  /*这句一定要有*/
  display: inline;
  margin-left: 20px;
  margin-right: 100px;
}

.TXT2 {
  height: 448px;
  text-indent: 2em;
  /* 首行缩进 */
  letter-spacing: 1px;
  /* 字间距 */
  line-height: 45px;
  color: black;
  margin-left: 100px;
}

.TXT2 span {
  top: 110px;
}

</style>
