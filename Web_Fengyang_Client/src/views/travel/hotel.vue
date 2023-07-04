INSERT INTO web_fengyang.hotelarticle (id, user_id, category_id, title, content, head_image, article_type, created_at, updated_at) VALUES ('8d449178-3d89-44e9-a05c-2104156cca59', 1, 1, '艾顿酒店 "北京路22号"', 'price: 100
telephone: 123456
website: "www.baidu.com"
center: [112.2896, 25.048]', '../../assets/hotel.png', '2', '2023-06-10 09:56:28', '2023-06-10 09:56:31');

<template>
  <div class="wrap">
    <div class="wrap_1">
      <!-- 搜索框 -->
      <div id="search_container">
        <el-input v-model="search_input_content" id="search_input" placeholder="请输入你要查找的关键词" type="text"></el-input>
        <el-button type="primary" id="searchBtn" @click="send_search">地点查询</el-button>
        <el-button type="primary" id="compute_driving_Btn" @click="compute_driving">路径计算</el-button>
  <!--      <el-button type="primary" plain round @click="send()">点击返回丰阳镇中心</el-button>-->
      </div>
    </div>

    <div class="wrap_2">
    <!-- 地图框 -->
      <div id="map_wrap">
        <div id="container" class="custom_style"></div>
        <div id="panel" class="custom_style"></div>
      </div>
    </div>
  </div>
  <div class="wrap_3">
    <!-- 酒店展示 -->
    <div id="hotel_wrap_title">-----酒店预订------</div>
    <div id="hotel_wrap">
      <el-row>
        <el-col v-for="(item,index) in hotelsData.values()" :key="index" :span="6" :offset="index%3 ==0 ? 2:1">
          <el-card class="custom_style" shadow="always"  :body-style="{ padding: '0px' }" >
            <img src="../../assets/hotel.png" class="hotel_image"/>
            <div class="detail_hotel">
                <h1 class="hotel_title">
                  <span class="hotel_name">{{ item.name }}</span>
                  <span class="hotel_price">{{ item.price }}元/晚起</span>
                </h1>

              <div class="hotel_profile">
                <p>地址: {{ item.placeAddress }}</p>
                <p>电话: {{ item.telephone }}</p>
              </div>

              <div class="hotel_icon" >
                <div class="show_detail">
                  <el-icon :size="20" color="#900" @mouseover="show_icon_detail($event)" @mouseleave="hide_icon_detail($event)" @click="click_show_detail(item.website)"><Document /></el-icon>
                  <span class="icon_detail">显示详情</span>
                </div>
                <div class="hotel_location">
                  <el-icon :size="20" color="#900" @mouseover="show_icon_detail($event)" @mouseleave="hide_icon_detail($event)" @click="click_hotel_location(item.center)"><Position /></el-icon>
                  <span class="icon_detail">地图导航</span>
                </div>
              </div>

            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script setup>
import {ref, onMounted, watch, inject} from 'vue'
// 地图加载
import AMapLoader from '@amap/amap-jsapi-loader'
// element-ui 的图标引入
import { Position, Document } from '@element-plus/icons-vue'
// 安全秘钥
window._AMapSecurityConfig = {
  securityJsCode: '3626fd4fb9d0809a04788c6df4d45953'
};
//***********变量***********
const axios = inject("axios")
// map地图变量
let map = ref(null);
// search_content是要输入查询的内容
let search_input_content= ref('');
// search_content是要查询的内容
let search_content=ref('');
// 输入查询结果
let placeSearch =ref(null);
// 驾车查询
let driving =ref(null);
// 起始点和终点
let start_point =ref(null);
let end_point =ref(null);
var path = [];
// 后端返回的酒店详情变量
var hotelsData = ref([]);

// 没有后台数据时,可用下面的hotels来测试
// const hotels =[
//   {
//     id: 1,
//     name: '酒店1',
//     image: '../../assets/hotel.png',
//     price: 100,
//     website:"www.baidu.com",
//     center: [112.2896, 25.048],
//     placeAddress: '北京路22号',
//     telephone: 123456
//   },
//   {
//     id: 2,
//     name: '酒店2',
//     image: '../../assets/hotel.png',
//     price: 120,
//     website:"www.baidu.com",
//     center: [112.286, 25.0408],
//     placeAddress: '北京路89号',
//     telephone: 123456
//   },
//   {
//     id: 3,
//     name: '酒店3',
//     image: '../../assets/hotel.png',
//     price: 90,
//     website:"www.baidu.com",
//     center: [112.283, 25.0438],
//     placeAddress: '北京路72号',
//     telephone: 123456
//   },
//   {
//     id: 4,
//     name: '酒店3',
//     image: '../../assets/hotel.png',
//     price: 160,
//     website:"www.baidu.com",
//     center: [112.28196, 25.008],
//     placeAddress: '北京路72号',
//     telephone: 123488
//   },
// ]


//***********函数***********

// 侧边栏和导航栏的点击触发函数
const handleSelect = (index) => {    // 这里可以定义导航栏点击函数
};
// button的点击查询
const send_search= () => {
  search_content.value=search_input_content.value
};
// 根据搜索得到查询结果
const select = (event) => {
  // 在函数内部使用event参数来处理搜索事件
  placeSearch.setCity(event.poi.adcode);
  placeSearch.search(event.poi.name);
};

const compute_driving = () => {
  var p1 = start_point.getPosition();
  var p2 = end_point.getPosition();
  driving.search(p1, p2);
  // feedBack('区域搜索', 'error')
};

const show_icon_detail = (event)  => {
  // 通过 event.target 获取触发事件的元素，即被鼠标悬停的元素
  const container = event.target.parentNode.parentNode;
  const detail = container.querySelector('.icon_detail');
  if ( detail ) {
    detail.setAttribute("id", "show");
  }
};

const hide_icon_detail = (event) => {
  const container = event.target.parentNode;
  const detail = container.querySelector('.icon_detail');
  if ( detail ) {
    detail.removeAttribute("id");
  }
};

const click_show_detail= (website) => {
  let fullUrl = website;
  // 检查传递的 website 参数是否以http://或 https://开头。如果不是，我们将前缀设置为 http://，然后将其与 website 参数拼接成完整的 URL。
  if (!/^https?:\/\//i.test(website)) {
    fullUrl = 'http://' + website;
  }
  window.open(fullUrl, '_blank');
};

const click_hotel_location= (location) =>{
  console.log(location);
  // console.log(location[0]);
  map.setZoomAndCenter(18,[location[0], location[1]]);

  // 第一个参数设置为 0，表示将页面滚动到顶部。第二个参数设置为 0，表示在垂直方向上滚动到顶部位置
  // window.scrollTo(0, 0);
  const currentPosition = window.pageYOffset; // 当前滚动位置
  const step = Math.floor(-currentPosition / 20); // 每一步滚动的距离
  const scroll = () => {
    if (window.pageYOffset <= 0) {
      // 已经滚动到顶部
      return;
    }
    window.scrollBy(0, step);
    // 使用 requestAnimationFrame 递归调用滚动函数，实现平滑滚动效果
    requestAnimationFrame(scroll);
  };
  // 调用滚动函数
  scroll();
};

const loadArticles = async() => {
  
  let res = await axios.get(`/article/list?articleType=hotelarticle`)
  if (res.data.code == 200) {
    var articleList = res.data.data.article

    for (var i = 0; i < articleList.length; i++) {
      var rowtitle = articleList[i].title;
      var content = articleList[i].content;
      // 对 obj 进行操作，例如访问属性 obj.property
      // 通过正则表达式匹配键值对的值
      var priceMatch = content.match(/price: (.*?)(\n|$)/);
      var websiteMatch = content.match(/website: "(.*?)"(\n|$)/);
      var centerMatch = content.match(/center: \[(.*?)\](\n|$)/);
      var telephoneMatch = content.match(/telephone: (.*?)(\n|$)/);
      var titleMatch = rowtitle.match(/(.*) "/);
      var placeAddressMatch = rowtitle.match(/\"([^"]+)\"/);
      // 获取匹配结果中的值
      var price = priceMatch ? parseInt(priceMatch[1]) : null;
      var website = websiteMatch ? websiteMatch[1] : null;
      var center = centerMatch ? centerMatch[1].split(',').map(Number) : null; // 分割数字并将其转换为数字类型
      var telephone = telephoneMatch ? parseInt(telephoneMatch[1]) : null;
      var title = titleMatch ? titleMatch[1] : null;
      var placeAddress = placeAddressMatch ? placeAddressMatch[1] : null;
      hotelsData.value.push({
        name: title,
        image: articleList[i].head_image,
        telephone: telephone,
        price: price,
        website: website,
        center: center,
        placeAddress: placeAddress,
      });
    }
    console.log(hotelsData[0])
  }
}
//地图初始化函数
const initMap = () => {
  AMapLoader.load({
    key: 'e0eb946f632c6473f40386d512e056d6', // 申请好的Web端开发者Key，首次调用 load 时必填
    version: '2.0', // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
    plugins: ['AMap.ToolBar', 'AMap.Scale', 'AMap.HawkEye', 'AMap.MapType', 'AMap.Geolocation', 'AMap.AutoComplete', 'AMap.PlaceSearch', 'AMap.Driving', 'AMap.DragRoute', 'AMap.Weather'] // 需要使用的插件列表，如比例尺'AMap.Scale'等
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

    //添加自定义的点标记
    start_point = new AMap.Marker({
      position: new AMap.LngLat(112.283196, 25.043408),
      draggable: true,
      // icon:
      // 在这里设置icon样式，引入assets下面的startpoint.svg文件
      title: 'defineMarkerPlace'
    })
    map.add(start_point)

    //添加自定义的点标记
    end_point = new AMap.Marker({
      position: new AMap.LngLat(112.279502, 25.046681),
      // offset: new AMap.Pixel(-10, -10),
      // icon:
      // 在这里设置icon样式，引入assets下面的endpoint.svg文件
      draggable: true,
      title: 'definePlace',
      zoom: 13
    })
    map.add(end_point)

    driving = new AMap.Driving({
      map: map,
      panel: "panel",
      hideMarkers: true
    });
    compute_driving();

    var weather = new AMap.Weather();
    weather.getForecast('杭州市', function(err, data)
    {
      console.log(err, data);
    });
  })
}

onMounted(() => {
  loadArticles(),
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
  .custom_style{
    border-radius: 20px;
    :hover .hotel_image{
      transform: scale(1.05);
    }
  }
  .wrap{
    border-radius: 20px;
    padding-top: 20px;
    background-color: #eee6e3;
    margin-left: 10%;
    margin-right: 10%;
    .wrap_1{
    //margin-top: 80px; /* 调整顶部间距，使输入框和按钮位于导航栏下方 */
      height: 50px;
      display: flex;
      align-items: center;  //子元素垂直居中
      justify-content: center; /* 将子元素水平居中 */
    // > * 是一个 CSS 选择器，用于选择该元素的直接子元素。
      > * {
        margin-right: 10px; /* 添加右侧间距 */
      }
    // 搜索容器
      #search_container{
        position: absolute;
        display: flex;
        #search_input{
          height: 50px;
          width: 300px;
        }
        #searchBtn {
          height: 50px;
          width: 100px;
        }
        #compute_driving_Btn{
          height: 50px;
          width: 100px;
        }
      }
    }

    .wrap_2{
      height: 800px;
      width: auto;
      margin-left: 5%;
      margin-right: 5%;
      #map_wrap {
        z-index: 10;
      //absolute是相对于最近的父元素来定位
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
      // 使用translate函数将元素在水平和垂直方向上分别向左和向上移动自身宽度和高度的一半，以使其完全居中显示。
        height: 90%;
        width: 100%;
        #container {
          padding: 0px;
          margin: 0px;
          width: 100%;
          height: 100%;
        }
        #panel {
          position: fixed;
          background-color: white;
          max-height: 90%;
          overflow-y: auto;
          top: 10px;
          left: 10px;
          width: 280px;
        }
        #panel .amap-call {
          background-color: #009cf9;
          border-top-left-radius: 4px;
          border-top-right-radius: 4px;
        }
        #panel .amap-lib-driving {
          border-bottom-left-radius: 4px;
          border-bottom-right-radius: 4px;
          overflow: hidden;
        }
      }
    }
  }

  .wrap_3{
    border-radius: 20px;
    height: auto;
    width: auto;
    background-color: #e9e9e0;
    margin-top: 10px ;
    margin-bottom: 20px;
    margin-left: 10%;
    margin-right: 10%;
    #hotel_wrap_title{
      text-align: center;
      font-family: 宋体;
      color: rgb(153, 0, 0);
      font-size: 50px;
    }
    #hotel_wrap {
      .hotel_image{
        max-width: 100%;
        max-height: 100%;
        transition: transform 0.4s;
        /* 样式变化时设置动画时间为0.5秒 */
      }
      .detail_hotel{
        margin: 0px;
        padding: 0px 10px 10px;
        .hotel_title{
          margin: 0px;
          text-align: center;
          .hotel_name{
            max-width: 100%;
            white-space: nowrap;
            text-overflow: ellipsis;
            padding-right: 10px;
            font-size: 28px;
            font-family: 宋体;
            font-weight: bold;
            letter-spacing: -1px;
            color: rgb(200 , 100, 89);
          }
          .hotel_price{
            padding-left: 5px;
            padding-right: 5px;
            width: auto;
            font-size: 18px;
            font-family: 微软雅黑;
            line-height: 16px;
            background: rgb(153, 0, 0);
            color: rgb(255, 255, 255);
          }
        }
        .hotel_profile{
          font-family: 微软雅黑;
          margin-left: 10px;
          font-size: 16px;
          color: #796358;
        }
        .hotel_icon{
          margin-left: 10px;
          display: flex;
        //justify-content: space-between;   //将元素平均分布在容器内，并在它们之间留出空间
          .show_detail{
            display: flex;
            font-family: 微软雅黑;
            color: #796358;
          }
          .hotel_location
          {
            margin-left: 20px;
            display: flex;
            font-family: 微软雅黑;
            color: #796358;
          }
          .icon_detail{
            display: none;
          //opacity: 0;
          //height: 0;
            transition: transform 0.4s;
          }
          .hotel_location :hover .icon_detail{
            display: block;
          }
          #show {
            display: block;
            height: 10px;
          }
        }
      }
    }
  }



</style>
