<template>
  <div class="carousel-box">
    <div class="carousel-card">
      <el-carousel :interval="3000" type="card" height="370px" >
      <el-carousel-item v-for="item in 6" :key="item">
        <img :src="`../src/assets/pic${item}.jpg`" :alt="'Slide ' + item" />
      </el-carousel-item>
    </el-carousel>
  </div>
    
  </div>
</template>
<script setup>



</script>
  
<style scoped>
.carousel-box {
  display: flex;
  justify-content: center;
  margin-bottom: 30px;
}

.carousel-card {
  width: 100%;
}

.el-carousel__item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.el-carousel__item:nth-child(2n) {
  border-radius: 15px;
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  border-radius: 15px;
  background-color: #d3dce6;
}
</style>
  
  