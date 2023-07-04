<template>
    <TopBar v-if="show" />
    <div class="background"/>
    <div :style="{ marginTop: show ? '100px' : '0' }" class="mainContent">
        <router-view />
    </div>
    
    <FooterBar />
    <RightBar />
</template>
  
<script setup>

import TopBar from "../components/TopBar.vue"
import FooterBar from "../components/FooterBar.vue"
import RightBar from "../components/RightBar.vue"
import { useRoute, useRouter } from 'vue-router'
import { ref } from 'vue';

const router = useRouter();
const route = useRoute()
const show = ref(route.path != "/home");

router.beforeEach((to, from) => {
    if(to !=from) {
      window.scrollTo(0, 0);
    }
    if (to.path.startsWith('/home')) {
        show.value = false;
    } else {
        show.value = true;
    }
})
</script>

<style scoped>
.mainContent {
    margin-top: 100px;
    
}
.background {
    position: fixed;
    top: 0px;
    height: 100%;
    width: 100%;
    background-image: url('../assets/background.jpg');
    background-size: cover;
    background-position: center;
    /* z-index: 10000; */
}
</style>
  
  