<template>
  <router-view />

</template>

<script setup>

import { onMounted } from 'vue'
import { categoryList } from './api/article';

const loadCategorys = () => {
  categoryList().then(result => {
    if(result != null){
      console.log(result.data.data.categorys)
    }
  })
}

onMounted(() => { // 挂载页面时加载文章
    loadCategorys()
})
</script>

<style scoped></style>

