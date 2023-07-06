<template>
  <!--轮播图-->
  <div class="swiper">
    <Swiper :modules="modules" :pagination="{ clickable: true }" :autoplay="{ delay: 3000, disableOnInteraction: false }"
      navigation>
      <SwiperSlide v-for="item in chatImageData">
        <div class="title">
          <div class="title-large">
            <p>{{ item.titleLargeText }}</p>
          </div>
          <div class="title-small">
            <p>{{ item.titleSmallText }}</p>
          </div>
        </div>
        <img :src="item.imagePath" alt="">
      </SwiperSlide>

    </Swiper>
    <!-- 如果需要导航按钮 -->
    <div class="swiper-button-prev"></div>
    <div class="swiper-button-next"></div>
  </div>
</template>

<script>
// 首页必须自己导入
import { Swiper, SwiperSlide } from "swiper/vue"
import "swiper/css"
import { Pagination, Autoplay, Navigation } from "swiper"
import "swiper/css/pagination"
import "swiper/css/navigation"

export default {
  components: {
    Swiper,
    SwiperSlide,
  },
  data() {
    return {
      modules: [Pagination, Autoplay, Navigation],
      chatImageData: [
        {
          imagePath: "/src/assets/HomePage/chart/homeChartPic3.jpg",
          titleLargeText: "一碧千里",
          titleSmallText: "绿意盎然的盛夏，谱写一场夏天的故事。"
        },
        {
          imagePath: "/src/assets/HomePage/chart/homeChartPic2.jpg",
          titleLargeText: "碧水微澜",
          titleSmallText: "湛蓝的水，青翠的树，山中画，画中意，犹如仙境。"
        },
        {
          imagePath: "/src/assets/HomePage/chart/homeChartPic1.jpg",
          titleLargeText: "水帘悬挂",
          titleSmallText: "跌宕奔腾，似玉如银；水清爽而透明，轻盈而飘逸。"
        }
      ]
    }
  }
}
</script>

<style scoped>
.swiper img {
  margin: 0px;
  position: relative;
  top: 0px;
  width: 100%;
  height: 100%;
  background-size: 100%;
}

.swiper {
  margin: 0px;
  position: relative;
  top: 0px;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 100;
}

.title {
  position: absolute;
  font-family: Microsoft YaHei Light;
  top: 150px;
  left: 150px;
  color: #ffffff;
  z-index: 1000;
  text-shadow: 1px 1px 10px rgb(0, 0, 0);
  /* -webkit-text-stroke: 0.2px black; */
}

.title-large {
  font-size: 65px;
}

.title-small {
  margin-top: 5px;
  margin-left: 10px;
  font-size: 30px;
}
</style>