<template>
    <el-menu :default-active="selected" class="topbar" mode="horizontal" :ellipsis="false" @select="handleSelect">
        <el-menu-item index="0">四色丰阳</el-menu-item>
        <div class="弹性盒子" :style="{ flexGrow: 1 }" />
        <el-menu-item index="1">首页</el-menu-item>
        <el-sub-menu index="2">
            <template #title>咨询</template>
            <el-menu-item index="2-1">绿色</el-menu-item>
            <el-menu-item index="2-2">黄色</el-menu-item>
            <el-menu-item index="2-3">红色</el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="3">
            <template #title>景点</template>
            <el-menu-item index="3-1">绿色</el-menu-item>
            <el-menu-item index="3-2">黄色</el-menu-item>
            <el-menu-item index="3-3">红色</el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="4">
            <template #title>消费</template>
            <el-menu-item index="4-1">绿色</el-menu-item>
            <el-menu-item index="4-2">黄色</el-menu-item>
            <el-menu-item index="4-3">红色</el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="5">
            <template #title>论坛</template>
            <el-menu-item index="5-1">游记记录</el-menu-item>
            <el-menu-item index="5-2">摄影投稿</el-menu-item>
            <el-menu-item index="5-3">美食</el-menu-item>
            <el-menu-item index="5-4">景点</el-menu-item>
            <el-menu-item index="5-5">住宿</el-menu-item>
            <el-menu-item index="5-6">吐槽投诉</el-menu-item>
        </el-sub-menu>
        <el-menu-item index="6">住宿交通</el-menu-item>
    </el-menu>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
    name: 'TopBar',
    props: {
        selected: {
            type: String,
            default: '1'
        },
    },
    emits: ['select'],
    methods:{
        handleSelect(index: string) {
            this.$emit('select', index);
        },
    }
});
</script>

<style lang="scss" scoped>
.topbar {
    position: fixed;
    top: 0px;
    width: 100%;
    z-index: 999;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.26);
}
</style>