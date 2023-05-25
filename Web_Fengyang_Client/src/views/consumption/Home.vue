<template>
    <div>
    <!-- 顶部导航栏 -->
    <TopBar @select="handleSelect" />
    
    <div style="position: absolute; top: 100px; left: 0; height: 500px; width: 100%; z-index: 1;">
        <Chart>
        </Chart>
    </div>
    
    <div style="position: absolute; top: 500px; left: 120px; height: 20px; width: 80%; z-index: 1;">
        <SmallBar>
        </SmallBar>
    </div>

    <div style="position: absolute; top: 1000px; left: 120px; height: 20px; width: 80%; z-index: 1;">
        <FooterBar>
        </FooterBar>
    </div>
    
    <div style="position: fixed; top: 450px; right: 20px; z-index: 1;">
        <RightBar>
        </RightBar>
    </div>

    </div>
</template>

<script setup>
// 导入顶部栏
import TopBar from "../../components/TopBar.vue"
import FooterBar from "../../components/FooterBar.vue"
import RightBar from "../../components/RightBar.vue"
import Chart from './Chart.vue'
import SmallBar from './SmallBar.vue'

// 侧边栏和导航栏的点击触发函数
const handleSelect = (index) => {    // 这里可以定义导航栏点击函数

};
</script>