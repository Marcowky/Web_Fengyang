<template xmlns="http://www.w3.org/1999/html">

  <div id="wrap">
    <h2 class="title">丰阳交通简介</h2>
    <div class="passage_content"> 丰阳镇位于连州市西北部，距市区32公里，总面积168.7公里，丰阳镇毗邻广西和湖南两省，自古便是三省边界最重要的商贸通道，更是南粤古驿道六条线路中肇雷古驿道文化线路中的重要组成部分。
    <div class="passage_content"> 由于乡镇面积较小，人口不多，丰阳镇目前没有直接的火车站，因此我们在下面提供了前往连州站（离丰阳最近的火车站）的铁路交通，以及通过大巴以及自驾直接抵达丰阳的公路交通，供您选择。
    <div class="passage_content"> 无论您选择哪种交通方式，丰阳镇都值得一游。它拥有独特的自然风光和丰富的人文历史，为游客提供了难忘的旅行体验。请根据您的出发地和时间，请选择最适合您的交通方式前往丰阳镇，祝您旅途愉快！</div>
    </div></div>
    <div id="center"><div class="img_box"><img src="../../assets/location.png" class="centered-image"></div></div>
    <div style="text-align: center;">丰阳地理分布图</div>
  </div>



  <div class="train" id="wrap">
    <h2 class="title">铁路交通</h2>
    <div class="passage_content"> 丰阳镇目前没有直接的铁路车站，我们推荐您可以乘坐火车到广东省清远市的清远站，这也是离丰阳镇最近的火车站。
    <div class="passage_content">由于清远位于广东的西北方向，邻近湖南，江西省份，因此有多条高铁途径连州，从那里您可以在下方选择合适的班次前往连州站之后，再选择其他交通方式前往丰阳镇。
      </div><div class="passage_content">下面是清远市的列车时刻表:</div></div>
    <div id="center"><div class="img_box"><img src="../../assets/train1.png" class="centered-image"></div></div>

  </div>




  <div id="wrap">
    <h2 class="title"> 公路交通</h2>
    <div class="passage_content">目前乘坐大巴可以选择连州、韶关、郴州三个市区选择大巴前往丰阳镇，请根据出发地及时间，择优选择乘车地点
    <div class="passage_content">值得一提的是，如果您想通过自驾游前往丰阳镇，请您点击这里，我们为您规划路线</div></div>
    <div class="subtitle">连州方向</div>
    <div class="traffic_wrap_detail">
      <div class="traffic_wrap">
        <div class="passage_content">连州客运站有班车前往丰阳，新南门客运站位于锦江区新南路，发车时间在早上七点，直达沟口；连州客运站位于金牛区西三环路五段，上高速较快，大巴经停漳扎镇，直达丰阳，此路线相对其他两条路线较快</div>
        <div class="passage_content">具体路线为： 连州（茶店子/新南门）——G55高速——茂县——西岸——连州（具体汽车票务请查询：<a target="_blank" href="https://www.checi.cn/qichezhan/476_910/">连州汽车票务网</a>)</div>
        <div class="passage_content">连州方向—路程150公里，耗时1~2小时</div>
      </div>
      <img class="img_style" src="../../assets/bus_station.jpg">
    </div>

    <div class="subtitle">韶关方向</div>
    <div class="traffic_wrap_detail">
      <img class="img_style" src="../../assets/bus_station.jpg">
      <div class="traffic_wrap">
        <div class="passage_content">广东韶关每天都有发往丰阳的班车：丰阳汽车中心站早上7:30，汽车东站早上8:00，汽车南站早上8:45，每天班次不是特别多，请提前一天购票</div>
        <div class="passage_content">具体路线为： 韶关——茂县——临西——丰阳（具体汽车票务请查询：<a target="_blank" href="http://ticket.gdcd.gov.cn/station_more_notice">广东省汽车票务网</a>)</div>
        <div class="passage_content">韶关方向—路程320公里，耗时7~8小时</div>
      </div>
    </div>

    <div class="subtitle">郴州方向</div>
    <div class="traffic_wrap_detail">
      <div class="traffic_wrap">
        <div class="passage_content">郴州位于湖南省，到丰阳的直达车在平政汽车站坐，每天有三班，仅有两班到达丰阳景区，买票的时候一定要问清楚目的地。大巴车途中会休息吃饭，建议准备零食。</div>
        <div class="passage_content">具体路线为： 郴州——德县——蓝山——丰阳（具体汽车票务请查询：<a target="_blank" href="https://www.checi.cn/qichezhan/130_331/">郴州汽车票务网</a>)</div>
        <div class="passage_content">郴州方向—路程440公里，耗时9~10小时</div>
      </div>
      <img class="img_style" src="../../assets/bus_station.jpg">
    </div>

  </div>

</template>
<script setup>

</script>



<style scoped>
#wrap{
  width: 70%;
  left: 15%;
  padding-left: 2%;
  padding-right: 2%;
  padding-bottom: 20px;
  //background-color: #7B3DE0;
  background-color: white;
}
#center{
  display: flex;
  justify-content:center;
}
.traffic_wrap_detail{
  display: flex;
  .traffic_wrap{
  }
  .img_style{
    left_margin: 10px;
    height: 100%;
    scale: 0.9;
  }
}

.img_box{
  width: 820px;
  display: flex;
  //justify-content: center;
  align-items: center;
  .centered-image {
    display: block;
    max-width: 100%;
    max-height: 100%;
  }
}
.title{
  color: rgb(86, 99, 136);
  font-size: 32px;
  top: 10px;
  margin: 20px;
  text-align: center;
}

.subtitle{
  color: rgb(133, 100, 83);
  text-align: left;
  font-weight: bold;
  font-size: 20px;
  margin: 20px;
  //padding-bottom: 80px;
}
.passage_content{
  color: rgb(133, 100, 83);
  font-size: 19px;
  font-family: 等线;
  margin-bottom: 20px;
  text-indent: 2em;
/*这个是首行缩进*/
}
</style>