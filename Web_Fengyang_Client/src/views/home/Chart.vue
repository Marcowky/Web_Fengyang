<template>
    <!--轮播图-->
    <div class="swiper">
      <Swiper 
      :modules="modules" 
      :pagination="{ clickable: true }"
      :autoplay="{delay: 3000, disableOnInteraction: false}"  
      navigation>
        <SwiperSlide>
          <img src="../../assets/pic1.jpg" alt="">
        </SwiperSlide>
        <SwiperSlide>
          <img src="../../assets/pic2.jpg" alt="">
        </SwiperSlide>
        <SwiperSlide>
          <img src="../../assets/pic3.jpg" alt="">
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
      topbarHeight: 60,
    }
  }
}
</script>

<style scoped>
.swiper
img {
  position: relative;
  top: 0px;
  width: 100%;
  
}
</style>