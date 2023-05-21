<template>
  <div>
    <!--顶部导航栏-->
    <el-menu
      router="router"
      :default-active="this.$route.path"
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
      @select="handleSelect"
    >
    <!--注意:在父层级的index中，如果没有添加 “/” 标识的话，router的自动跳转会出问题，
      在父层级跳到其他父层级的子层级下后，再次返回该层级，路径会直接修改后缀，从而导致跳转失败-->
        <!-- <div class="logo">
          <a href="HomePage.vue">
            <img src="../assets/logo.png" alt="">
          </a>
        </div> -->
        <el-menu-item index="/HomePage"><img class="logo" src="../assets/logo.png" alt="" /></el-menu-item>
        <div class = "flex-grow"></div>
        <el-menu-item index="/HomePage">首页</el-menu-item>
        <el-menu-item index="/InfoPage">资讯</el-menu-item>
        <el-menu-item index="/ViewPage">景点</el-menu-item>
        
        <!--这里消费选项菜单有一个问题：如果图标高度增加，消费会往下移，不对齐-->
        <el-sub-menu index="/ConsumptionPage">
          <template #title>消费</template>
          <!--这里包括特产和娱乐，故考虑作下拉子菜单-->
          <el-sub-menu index="/SpcialtyPage">
            <template #title>特产</template>
            <el-menu-item index="2-4-1">特产1</el-menu-item>
            <el-menu-item index="2-4-2">特产2</el-menu-item>
            <el-menu-item index="2-4-3">特产3</el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="RecreationPage">
            <template #title>娱乐</template>
            <el-menu-item index="2-4-4">娱乐1</el-menu-item>
            <el-menu-item index="2-4-5">娱乐2</el-menu-item>
            <el-menu-item index="2-4-6">娱乐3</el-menu-item>
          </el-sub-menu>
        </el-sub-menu>

        <el-menu-item index="/ForumPage">论坛</el-menu-item>

    </el-menu>
  </div>

  <!--轮播图-->
  <div>
    <n-carousel show-arrow autoplay="{delay: 100}">
    <img
      class="carousel-img"
      src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel1.jpeg"
    >
    <img
      class="carousel-img"
      src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel2.jpeg"
    >
    <img
      class="carousel-img"
      src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel3.jpeg"
    >
    <img
      class="carousel-img"
      src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel4.jpeg"
    >
  </n-carousel>
  </div>
  
  <br><br><br>

  <div class="general-title">
    <h3>—景 区 概 况—</h3>
  </div><br>

  <img class="IMG1" src="../assets/spring.jpg" alt="">
  <div class="TXT1">
    <span>{{ spring_mes }}</span>
  </div>

  <img class="IMG2" src="../assets/summer.jpg" alt="">
  <div class="TXT2">
    <span>{{ summer_mes }}</span>
  </div>

  <img class="IMG1" src="../assets/autumn.jpg" alt="">
  <div class="TXT1">
    <span>{{ autumn_mes }}</span>
  </div>

  <img class="IMG2" src="../assets/winter.jpg" alt="">
  <div class="TXT2">
    <span>{{ winter_mes }}</span>
  </div>
  
  <br><br><br>
  <!--footer-->
  <footer class="footer">
    <div class="footer-bottom">
      <div class="footer-container">
        <ul class="footer-nav">
          <!--景区概况-->
          <ul class="css-ul">
            <li>
            <a class="css-main_title" href="/HomePage" >景区概况</a>
            </li>
            <li>
              <a href="/HomePage">四色丰阳</a>
            </li>
            <li>
              <a href="/HomePage">民族风俗</a>
            </li>
            <li>
              <a href="/HomePage">历史人文</a>
            </li>
          </ul>

          <!--景区资讯-->
          <ul class="css-ul">
            <li>
            <a class="css-main_title" href="/InfoPage" >景区资讯</a>
            </li>
            <li>
              <a href="/InfoPage">景区活动</a>
            </li>
            <li>
              <a href="/InfoPage">景区公告</a>
            </li>
            <li>
              <a href="/InfoPage">志愿活动</a>
            </li>
          </ul>

          <!--景区景点-->
          <ul class="css-ul">
            <li>
            <a class="css-main_title" href="#" >景区景点</a>
            </li>
            <li>
              <a href="/InfoPage">热门景点</a>
            </li>
            <li>
              <a href="/InfoPage">游玩攻略</a>
            </li>
            <li>
              <a href="/InfoPage">路线规划</a>
            </li>
          </ul>

          <!--景区消费-->
          <ul class="css-ul">
            <li>
            <a class="css-main_title" href="#" >景区消费</a>
            </li>
            <li>
              <a href="/ConsumptionPage">酒店民宿</a>
            </li>
            <li>
              <a href="/ConsumptionPage">娱乐项目</a>
            </li>
            <li>
              <a href="/ConsumptionPage">美食特产</a>
            </li>
          </ul>

          <!--投诉建议-->
          <ul class="css-ul">
            <li>
            <a class="css-main_title" href="#" >投诉建议</a>
            </li>
            <li>
              <a href="/ForumPage">景区论坛</a>
            </li>
          </ul>

        </ul>
      </div>
    </div>
  </footer>

</template>


<script>

export default{
  data(){
    return{
      spring_mes:"四季在一年中交替出现，一年四季的变化是连绵不断的过程。其也是阴阳转换，此消彼长、量变产生质变的过程。在这个渐变的过程中，立春、立夏、立秋、立冬，并称“四立”，都是季节的开启。春季排四季之首，新的轮回从此开启。春，代表着温暖、生长。春季，阴阳之气开始转变，万物随阳气上升而萌牙生长，春季是万物生长、春耕播种的季节，所谓“一年之计在于春”。",
      summer_mes:"夏季万物至此皆盛，是农作物进入旺季生长的一个季节。温度升高，天气炎热，狂风暴雨频发，万物盛长。夏季是风暴最多的季节，常伴随着大风、暴雨。北半球夏季，大陆受热气压上升形成低气压，海洋恒温温度相对较低形成高气压，根据环流，夏季吹的是东南风。夏季北半球各地的白昼时间全年最长。",
      autumn_mes:"秋季是收获季节，意味着万物开始从繁茂成长趋向萧索成熟。秋季的前两个节气立秋、处暑，还是酷热天气，因太阳所带来的热力未减弱，所谓“热在三伏”，三伏天出现在小暑与处暑之间，是一年中气温最高且又潮湿、闷热的日子。造成三伏天湿度高原因是，三伏天吹东南风，而东南方是太平洋和印度洋，空气潮湿，风的潮湿造成了三伏天湿度大；到了深秋之后则相反，吹西北风，而西北方是干燥的内陆，干燥的西北风造成深秋之后气候干燥。",
      winter_mes:"冬季，阴阳转变，万物由收到藏，植物生气闭蓄。进入立冬后意味着秋季少雨干燥气候渐过去，转为阴雨寒冻的冬季气候特征。北半球的太阳高度小，白昼时间短。所谓“热在三伏”，冷在三九”，冬至为“一九天”，冬至之前还不算很冷，冬至过后强冷空气南下频繁，并越过了南岭，气温骤降，天气寒冷。冬季真正寒冷，是在冬至后。",    
    }
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
  },
  mounted() {
      // 改变网页背景颜色
      document.querySelector('body').setAttribute('style', 'background-color:#d9e0dc');
  },
  beforeDestroy() {
      document.querySelector('body').removeAttribute('style')
  },
};


</script>

<style scoped>

.el-menu-demo{
  background-color:#edefeb;
  height: 70px;
}
/*去掉选中el-menu-item时的下划线*/
.el-menu--horizontal > .el-menu-item {
  border-bottom: none !important;
}
.el-menu--horizontal > .el-menu-item.is-active {
  border-bottom: none !important;
}
/* 设置选中el-menu-item时的状态 */
.el-menu-item .is-active{
  border-bottom-color: none !important;
}
/* 设置选中el-menu-item时的样式 */
.el-menu-item:hover{
    color: #1787FB !important;
    text-decoration: underline;
}

.logo{
  width:150px;
  height:47px;
}
.flex-grow {
  flex-grow: 1;
}
.carousel-img {
  width: 100%;
  height: 360px;
  object-fit: cover;
}
.general-title{
  font-family: 宋体;
  color: black;
  font-size: 45px;
  font-weight: bolder;
  text-align: center;
}
.IMG1{
  width: 300px;
  height: 448px;
  float: left;  /*这句一定要有*/ 
  display: inline;
  margin-right: 20px;
}
.TXT1{
  height: 448px;
  text-indent:2em;    /* 首行缩进 */
  letter-spacing:1px;   /* 字间距 */
  line-height: 45px;  /* 行间距 */
  color: black;
}
.TXT1 span{
  top: 100px;
}

.IMG2{
  width: 300px;
  height: 448px;
  float: right;  /*这句一定要有*/ 
  display: inline;
  margin-left: 20px;
}
.TXT2{
  height: 448px;
  text-indent:2em;    /* 首行缩进 */
  letter-spacing:1px;   /* 字间距 */
  line-height: 45px;
  color: black;
}
.TXT2 span{
  top: 110px;
}

/* footer的css */
.footer{
  position: relative;
}

.footer-bottom{
  background:#353432;
  width: 100%;
  padding: 0;
}

.footer-container{
  position: static;
  padding: 0px 50px;
  margin:10 auto;
}

.footer-nav{
  display: flex;
  flex-wrap: wrap;
  padding: 40px 0px;
}

.css-ul{
  margin-left: 40px;
}

.css-main_title{
  display: inline-block;
  padding: 10px 0;
  font-size: 16px;
  color: #ffffff;
  border-bottom: 1px solid #A0A0A0;   /* 景区概况底下那条横线 */
}

ul, li{
  list-style: none;   /* 去掉列表前面的小圈圈 */
}

/* .css-li{
  display: block;
  font-size: 14px;
  line-height: 30px;
  color: #aaaaaa;
} */

a{    /* 改每个主标题下a的样式，一开始用的是每个a里面加.css-li，后面发现可以直接用这个解决*/
  display: block;
  font-size: 10px;
  line-height: 30px;
  color: #aaaaaa;
}
</style>



<!-- el-submenu子菜单部分  el-submenu的代码不能包含在scoped -->
<style lang="less">
.el-submenu .el-submenu__title {
  font-weight: bold;
  font-size: 16px;
  color: #1787FB;
}
.el-menu--horizontal>.el-submenu .el-submenu__title {
    height: 92px;
    line-height: 92px;
    border-bottom: 2px solid transparent;
    color: #666666;
    i{
      font-size: 16px;
      font-weight: bold;
    }
}
.el-menu .el-menu--popup .el-menu--popup-bottom-start{
  top: 92px;
}

.el-menu--horizontal>.el-submenu .el-submenu__title:hover{
    color: #1787FB;
    /* less语法，实现鼠标悬停时icon变色 */
    i {
        color: #1787FB;
    } 
}
</style>