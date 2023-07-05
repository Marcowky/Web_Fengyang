<template>
  <!--轮播图-->
  <div class="swiper">
    <Swiper :modules="modules" :pagination="{ clickable: true }" :autoplay="{ delay: 3000, disableOnInteraction: false }"
      navigation>
      <SwiperSlide>
        <div class="title">
          <div class="title-large">
            <p>四色丰阳</p>
          </div>
          <div class="title-small">
            <p>水流暴冲，湛蓝的水光瞬间爆发，形成了一道喷涌而起的喷泉！</p>
          </div>
        </div>
        <img src="https://www.jiuzhai.com/templates/rt_sienna/custom/images/2022_06/07.jpg" alt="">
      </SwiperSlide>
      <SwiperSlide>
        <div class="title">
          <div class="title-large">
            <p>四色丰阳</p>
          </div>
          <div class="title-small">
            <p>水流暴冲，湛蓝的水光瞬间爆发，形成了一道喷涌而起的喷泉！</p>
          </div>
        </div>
        <img src="https://www.jiuzhai.com/templates/rt_sienna/custom/images/2023_05/06.jpg" alt="">
      </SwiperSlide>
      <SwiperSlide>
        <div class="title">
          <div class="title-large">
            <p>四色丰阳</p>
          </div>
          <div class="title-small">
            <p>水流暴冲，湛蓝的水光瞬间爆发，形成了一道喷涌而起的喷泉！</p>
          </div>
        </div>
        <img src="https://www.jiuzhai.com/templates/rt_sienna/custom/images/2022_06/10.jpg" alt="">
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