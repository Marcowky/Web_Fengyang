<template>
  <router-view ></router-view>
    <!-- 顶部导航栏 -->
    <TopBar @select="handleSelect" />

    <div class="box">
      

      <div style="position: fixed; top: 450px; right: 20px; z-index: 1;">
          <RightBar>
          </RightBar>
      </div>

      <FooterBar />
    </div>
</template>

<script setup>
// 导入顶部栏、底部栏、右边栏
import TopBar from "./components/TopBar.vue"
import FooterBar from "./components/FooterBar.vue"
import RightBar from "./components/RightBar.vue"

</script>

<style scoped>

</style>

