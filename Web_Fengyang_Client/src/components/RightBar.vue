<template>
  <el-menu default-active="2" class="el-menu-vertical-demo">
    <el-menu-item index="1">
      <template #title>
        <div class="hover-image-container el-menu-item">
          <el-icon class="iconfont icon-weixinxiaochengxu"></el-icon>
          <text>小程序</text>
          <img src="../assets/xiaochengxu.jpg" class="hover-image" />
        </div>
      </template>
    </el-menu-item>

    <el-menu-item index="2">
      <template #title>
        <div class="hover-image-container el-menu-item">
          <el-icon class="iconfont icon-weixin"></el-icon>
          <span>公众号</span>
          <img src="../assets/gongzhonghao.png" class="hover-image" />
        </div>
      </template>
    </el-menu-item>

    <el-menu-item index="3">
      <template #title>
        <div class="hover-image-container el-menu-item">
          <el-icon>
            <service />
          </el-icon>
          <span>景区热线</span>
          <div class="hover-image">
            <p>景区热线:</p>
            <p>40-16438629</p>
            <p>86-78286993</p>
          </div>
        </div>
      </template>
    </el-menu-item>

    <el-menu-item index="4">
      <template #title>
        <div class="hover-image-container el-menu-item">
          <el-icon>
            <ChatDotSquare />
          </el-icon>
          <span>投诉举报</span>
          <img src="../assets/tousujubao.png" class="hover-image" />
        </div>
      </template>
    </el-menu-item>
  </el-menu>
</template>
  
<script lang="ts" setup>
import { ref } from 'vue'
import {
  Menu as IconMenu,
  Service,
  ChatDotSquare,
} from '@element-plus/icons-vue'

</script>
  
<style scoped>
.hover-image-container {
  position: relative;
}

.hover-image {
  display: none;
  position: absolute;
  border-radius: 10px;
  right: 80px;
  top: -15px;
  width: 100px;
  height: 100px;
  background-color: rgb(56, 56, 56);
}

.hover-image-container:hover .hover-image {
  display: block;
}

.el-menu-vertical-demo:not(.el-menu--collapse) {
  position: fixed;
  bottom: 10%;
  right: 2%;
  border-right-width: 0;
  z-index: 999;
  background-color: rgba(65, 65, 65, 0.7);
  min-height: 200px;
  box-shadow: 2px 0 6px rgba(255, 255, 255, 0.26);
  border-radius: 15px;
}

.el-menu-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  line-height: 2;
  width: 70px;
  height: 70px;
  color: rgb(255, 255, 255);
  border-radius: 15px;
}
</style>
  