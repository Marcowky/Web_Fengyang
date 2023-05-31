<template>
  <TopBar />
  <router-view />
  <FooterBar />
  <RightBar />
</template>

<script setup>
import TopBar from "./components/TopBar.vue"
import FooterBar from "./components/FooterBar.vue"
import RightBar from "./components/RightBar.vue"

</script>

<style scoped></style>

