<template>
    <TopBar v-if="show" />
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

const router = useRouter()
const route = useRoute()
const show = ref(route.path != "/home");

router.beforeEach((to, from) => {
    if (to.path.startsWith('/home')) {
        show.value = false;
    } else {
        show.value = true;
    }
    if (to.path != from.path) {
        console.log("to another path")
        window.scrollTo({
            top: 0
        });
    }
})
</script>
  
<style scoped>
.mainContent {
    margin-top: 100px;
}
</style>
  
  