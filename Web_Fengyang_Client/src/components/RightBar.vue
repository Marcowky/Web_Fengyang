<template>
    <el-menu
      default-active="2"
      class="el-menu-vertical-demo"
    >
      <el-menu-item index="1">
        <template #title>
          <div class="hover-image-container el-menu-item">
            <el-icon class="iconfont icon-weixinxiaochengxu"></el-icon>
            <span>小程序</span>
            <img src="../assets/pic1.jpg" class="hover-image"/>
          </div>
        </template>
      </el-menu-item>
      
      <el-menu-item index="2">
        <el-icon class="iconfont icon-weixin"></el-icon>
        <template #title>公众号</template>
      </el-menu-item>
      
      <el-menu-item index="3">
        <el-icon><service /></el-icon>
        <template #title>景区热线</template>
      </el-menu-item>
      
      <el-menu-item index="4">
        <el-icon><ChatDotSquare /></el-icon>
        <template #title>投诉举报</template>
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
    right: 100px;
    top: 0px;
    width: 100px;
    height: 100px;
  }
  
  .hover-image-container:hover .hover-image {
    display: block;
  }

  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 120px;
    background-color: rgba(255, 255, 255, 0.5);
    min-height: 200px;
    border-radius: 15px;
  }

  .el-menu-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    text-align: center;
    line-height: 2;
  }
  </style>
  