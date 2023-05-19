<template>
    <!-- 顶部导航栏 -->
    <TopBar :selected="selectedItemIndex" @select="handleSelect"/>
</template>

<script setup>

// 导入路由
import { useRouter } from 'vue-router'
const router = useRouter()

// 导入顶部栏
import TopBar from "../../components/TopBar.vue"

// 导入导航栏路由函数
import {BarRouteGoto} from '../../components/BarRouteFunc.js'
// 侧边栏和导航栏的点击触发函数
const handleSelect = (index) => {    // 这里可以触发路由跳转或其他操作
    BarRouteGoto(router, index)
};
</script>