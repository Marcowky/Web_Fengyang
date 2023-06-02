<template>
    <TopBar />
    <div class="mainContent">
        <router-view />
    </div>
    <FooterBar />
    <RightBar />
</template>
  
<script setup>

import TopBar from "../components/TopBar.vue"
import FooterBar from "../components/FooterBar.vue"
import RightBar from "../components/RightBar.vue"

</script>
  
<style scoped>
.mainContent {
    margin-top: 100px;
}
</style>
  
  