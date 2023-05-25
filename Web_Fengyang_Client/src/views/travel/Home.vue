<template>
  <!-- 顶部导航栏 -->
  <TopBar @select="handleSelect" />
  <div class="wrap_1">
    <!-- 显示框 -->
<!--    <div id="info_window_container">-->
<!--      <div>信息展示：</div>-->
<!--      <div>-->
<!--        <span>地址：</span>-->
<!--        <span>{{placeAddress}}</span>-->
<!--      </div>-->
<!--      <div>-->
<!--        <span>坐标：</span>-->
<!--        <span>{{placeLocation}}</span>-->
<!--      </div>-->
<!--    </div>-->

    <!-- 搜索框 -->
    <div id="search_container">
      <el-input v-model="search_input_content" id="search_input" placeholder="请输入你要查找的关键词" type="text"></el-input>
      <el-button type="primary" id="searchBtn" @click="send_search">search</el-button>
<!--      <el-button type="primary" plain round @click="send()">点击返回丰阳镇中心</el-button>-->
    </div>
  </div>

  <!-- 地图框 -->
  <div id="map_wrap"><div id="container"></div></div>

  <!-- 酒店展示 -->
</template>

<script setup>
import { onMounted,ref,watch } from 'vue';
// 导入顶部栏
import TopBar from "../../components/TopBar.vue"
// 地图加载
import AMapLoader from '@amap/amap-jsapi-loader'


// 安全秘钥
window._AMapSecurityConfig = {
  securityJsCode: '3626fd4fb9d0809a04788c6df4d45953'
}
//***********变量***********
// map地图变量
let map = ref(null);

// search_content是要输入查询的内容
let search_input_content= ref('');
// search_content是要查询的内容
let search_content=ref('');
// 输入查询结果
let placeSearch =ref(null);



//***********函数***********
// 侧边栏和导航栏的点击触发函数
const handleSelect = (index) => {    // 这里可以定义导航栏点击函数
};
// button的点击查询
const send_search= () => {
  search_content.value=search_input_content.value
}
// 根据搜索得到查询结果
const select = (event) => {
  // 在函数内部使用event参数来处理搜索事件
  placeSearch.setCity(event.poi.adcode);
  placeSearch.search(event.poi.name);
  console.log("here");
};

//地图初始化函数
const initMap = () => {
  AMapLoader.load({
    key: 'e0eb946f632c6473f40386d512e056d6', // 申请好的Web端开发者Key，首次调用 load 时必填
    version: '2.0', // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
    plugins: ['AMap.ToolBar', 'AMap.Scale', 'AMap.HawkEye', 'AMap.MapType', 'AMap.Geolocation', 'AMap.AutoComplete', 'AMap.PlaceSearch'] // 需要使用的插件列表，如比例尺'AMap.Scale'等
  }).then((AMap) => {
    map = new AMap.Map('container', {
      viewMode: '3D', // 是否为3D地图模式
      zoom: 18, // 初始化地图级别
      center: [112.283196, 25.043408] // 初始化地图中心点位置
    })
    // 在图面添加工具条控件，工具条控件集成了缩放、平移、定位等功能按钮在内的组合控件
    map.addControl(new AMap.Scale())
    // 在图面添加比例尺控件，展示地图在当前层级和纬度下的比例尺
    map.addControl(new AMap.ToolBar())
    // 在图面添加鹰眼控件，在地图右下角显示地图的缩略图
    map.addControl(new AMap.HawkEye())
    // 在图面添加类别切换控件，实现默认图层与卫星图、实施交通图层之间切换的控制
    map.addControl(new AMap.MapType({defaultType:1,showTraffic:true,showRoad:true}))
    // 在图面添加定位控件，用来获取和展示用户主机所在的经纬度位置
    map.addControl(new AMap.Geolocation())

    // 搜索输入提示
    // search的API属性设置
    const search_autoOptions = {
      input: 'search_input',  //使用提示输入的input的id
      city: '清远'
    }
    // 输入提示
    const search_auto = new AMap.AutoComplete(search_autoOptions)
    // 注册监听，当选中某条记录时会触发
    search_auto.on('select', select)

    // 地图搜索
    // 构造地点查询类,绑定地图
    placeSearch = new AMap.PlaceSearch({
      map: map,
      city: '清远'
    })

  })
}

onMounted(() => {
  initMap();
});

watch(search_content, (newValue) => {
      if (newValue != null) {
        console.log(newValue);
        placeSearch.search(newValue);
        // drawBounds(newValue);
        // map.value.setZoom(16, true, 1);
      }
    }
);
</script>


<style lang="less">
  .wrap_1{
    margin-top: 60px; /* 调整顶部间距，使输入框和按钮位于导航栏下方 */
    display: flex;
    //align-items: center;
    justify-content: center; /* 将子元素水平居中 */
    // > * 是一个 CSS 选择器，用于选择该元素的直接子元素。
    > * {
      margin-right: 10px; /* 添加右侧间距 */
    }

    // 搜索容器
    #search_container{
      display: flex;
      #search_input{
        height: 40px;
        width: 300px;
      }
      #searchBtn {
        height: 40px;
        width: 100px;
      }
    }
  }
  #container {
    padding: 0px;
    margin: 0px;
    width: 100%;
    height: 100%;
  }
  #map_wrap {
    z-index: 10;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    height: 60%;
    width: 60%;
  }
</style>
