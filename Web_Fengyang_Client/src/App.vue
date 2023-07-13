<template>
  <router-view />
</template>

<script setup>

import { onMounted } from 'vue'
import { categoryStore } from './api/category';


onMounted(() => { // 挂载页面时加载文章
  categoryStore()
})
</script>

<style scoped></style>

