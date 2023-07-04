<template>
    <!--轮播图-->
    <div class="swiper">
      <Swiper 
      :modules="modules" 
      :pagination="{ clickable: true }"
      :autoplay="{delay: 3000, disableOnInteraction: false}"  
      navigation>
        <SwiperSlide>
          <img src="https://www.jiuzhai.com/templates/rt_sienna/custom/images/2022_06/07.jpg" alt="">
        </SwiperSlide>
        <SwiperSlide>
          <img src="https://www.jiuzhai.com/templates/rt_sienna/custom/images/2023_05/06.jpg" alt="">
        </SwiperSlide>
        <SwiperSlide>
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
      modules: [ Pagination, Autoplay, Navigation ], 
    }
  }
}
</script>

<style scoped>
.swiper
img {
  margin:0px;
  position: relative;
  top: 0px;
  width: 100%;
  height: 100%;
  background-size: 100%;
}

.swiper {
  margin:0px;
  position: relative;
  top: 0px;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 100;
}

</style>