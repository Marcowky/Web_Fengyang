<template>
  <div class="carousel-box">
      <div class="border1"></div>
      <div class="border2"></div>
      <div class="border3"></div>
      <div class="border4"></div>
    <div class="carousel-card">
      <el-carousel :interval="3000" type="card" height="370px" >
        <el-carousel-item v-for="item in 6" :key="item">
          <div class="carousel-item-wrapper">
            <img :src="`../src/assets/pic${item}.jpg`" :alt="'Slide ' + item" />
          </div>
        </el-carousel-item>
      </el-carousel>
    </div>
  </div>
</template>

<style scoped>
.carousel-box {
  display: flex;
  justify-content: center;
  margin-bottom: 30px;
}

.carousel-card {
  margin-top: 30px;
  width: 92%;
}

.el-carousel__item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.carousel-item-wrapper {
  position: relative;
}

.border1,
.border2,
.border3,
.border4 {
  position: absolute;
  width: 60px;
  height: 60px;
  z-index: 500;
  background-size: contain;
}

.border1 {
  top: 0;
  left: 0;
  background-image: url("../src/assets/border1.png");
  /* 添加边框样式 */
}

.border2 {
  top: 0;
  right: 0;
  background-image: url("../src/assets/border2.png");
  /* 添加边框样式 */
}

.border3 {
  bottom: 0;
  left: 0;
  background-image: url("../src/assets/border3.png");
  /* 添加边框样式 */
}

.border4 {
  bottom: 0;
  right: 0;
  background-image: url("../src/assets/border4.png");
  /* 添加边框样式 */
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
