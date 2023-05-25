<template>
    <!-- 顶部导航栏 -->
    <TopBar @select="handleSelect" />

    <div class="box">
        <div style="position: relative; margin-top:100px; width: 100%; z-index: 1;">
            <Chart />
        </div>

        <div style="position: relative; margin: auto; width: 80%; z-index: 1;">
            <SmallBar />
        </div>

            <FooterBar />

    </div>
</template>

<script setup>
// 导入顶部栏
import TopBar from "../../components/TopBar.vue"
import FooterBar from "../../components/FooterBar.vue"
import Chart from './Chart.vue'
import SmallBar from './SmallBar.vue'

// 侧边栏和导航栏的点击触发函数
const handleSelect = (index) => {    // 这里可以定义导航栏点击函数

};
</script>

<style lang="scss" scoped>
.box {
    position: relative;
    min-height: 100%;
}
</style>