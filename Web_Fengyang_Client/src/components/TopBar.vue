<template>
    <el-menu :default-active="selected" class="topbar" mode="horizontal" :ellipsis="false" @select="handleSelect">
        <el-menu-item index="0">四色丰阳</el-menu-item>
        <div class="弹性盒子" :style="{ flexGrow: 1 }" />
        <template v-for="item in menuItems">
            <template v-if="!item.hasSub&&!item.index.includes('-')">
                <el-menu-item :index="item.index">{{ item.label }}</el-menu-item>
            </template>
            <template v-if="item.hasSub">
                <el-sub-menu :index="item.index">
                    <template #title>{{ item.label }}</template>
                    <template v-for="subItem in menuItems">
                        <el-menu-item :index="subItem.index" v-if="subItem.index.startsWith(item.index + '-')">{{ subItem.label
                        }}</el-menu-item>
                    </template>
                </el-sub-menu>
            </template>
        </template>
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
    methods: {
        handleSelect(index: string) {
            this.$emit('select', index);
        },
    },
    data() {
        return {
            menuItems: [
                { index: '1', label: '游记记录', hasSub: false },
                { index: '2', label: '咨询', hasSub: true },
                { index: '2-1', label: '绿色', hasSub: false },
                { index: '2-2', label: '黄色', hasSub: false },
                { index: '2-3', label: '红色', hasSub: false },
                { index: '3', label: '景点', hasSub: true },
                { index: '3-1', label: '绿色', hasSub: false },
                { index: '3-2', label: '黄色', hasSub: false },
                { index: '3-3', label: '红色', hasSub: false },
                { index: '4', label: '消费', hasSub: true },
                { index: '4-1', label: '绿色', hasSub: false },
                { index: '4-2', label: '黄色', hasSub: false },
                { index: '4-3', label: '红色', hasSub: false },
                { index: '5', label: '论坛', hasSub: true },
                { index: '5-1', label: '游记记录', hasSub: false },
                { index: '5-2', label: '摄影投稿', hasSub: false },
                { index: '5-3', label: '美食', hasSub: false },
                { index: '5-4', label: '景点', hasSub: false },
                { index: '5-5', label: '住宿', hasSub: false },
                { index: '5-6', label: '吐槽投诉', hasSub: false },
                { index: '6', label: '住宿交通', hasSub: false }
            ]
        }
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