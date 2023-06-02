<template>
  <TopBar v-if="showComponent && showFirst"/>
  <router-view />
  <FooterBar v-if="showComponent"/>
  <RightBar v-if="showComponent"/>
</template>

<script setup>

import TopBar from "./components/TopBar.vue"
import FooterBar from "./components/FooterBar.vue"
import RightBar from "./components/RightBar.vue"
import { useRouter } from 'vue-router'
import { ref } from 'vue';

const router = useRouter();
const showComponent = ref(true);
const showFirst = ref(false);

router.beforeEach((to) => {
    if (to.path.startsWith('/admin')) {
        showComponent.value = false;
    } else {
        showComponent.value = true;
    }
    if (to.path.startsWith('/')) {
        showFirst.value = false;
    } else {
        showFirst.value = true;
    }
})


</script>

<style scoped></style>

