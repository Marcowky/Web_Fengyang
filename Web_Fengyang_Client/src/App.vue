<template>
  <TopBar @select="handleSelect" />

  <div class="mainBox">
    <router-view />
    <FooterBar />
  </div>

  <RightBar style="position: fixed; top: 450px; right: 20px; z-index: 1;" />
</template>

<script setup>
import TopBar from "./components/TopBar.vue"
import FooterBar from "./components/FooterBar.vue"
import RightBar from "./components/RightBar.vue"

</script>

<style scoped>
.mainBox {
  position: relative;
  min-height: 100%;
  margin-top: 100px;
}
</style>

