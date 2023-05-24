<template>
    <!-- 顶部导航栏 -->
    <TopBar @select="handleSelect" />
</template>

<script setup>

// 导入顶部栏
import TopBar from "../../components/TopBar.vue"

// 侧边栏和导航栏的点击触发函数
const handleSelect = (index) => {    // 这里可以定义导航栏点击函数

};
</script>