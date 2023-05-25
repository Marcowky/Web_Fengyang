<template>
    <el-menu class="topbar" router=true :default-active="this.$route.fullPath" mode="horizontal" :ellipsis="false"
        @select="handleSelect">
        <el-menu-item index="0">四色丰阳</el-menu-item>
        <div class="弹性盒子" :style="{ flexGrow: 1 }" />
        <template v-for="item in menuItems">
            <template v-if="!item.hasSub && !item.index.includes('-')">
                <el-menu-item :index="item.index">{{ item.label }}</el-menu-item>
            </template>
            <template v-if="item.hasSub">
                <el-sub-menu :index="item.index">
                    <template #title>{{ item.label }}</template>
                    <template v-for="subItem in menuItems">
                        <el-menu-item :index="subItem.index" v-if="subItem.index.startsWith(item.index + '?')">{{
                            subItem.label
                        }}</el-menu-item>
                    </template>
                </el-sub-menu>
            </template>
        </template>
    </el-menu>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
// 导入菜单选项配置文件
import config from '../config/config.json';

interface MenuItem {
    index: string;
    label: string;
    mainMenu: boolean;
    hasSub: boolean;
}

export default defineComponent({
    name: 'TopBar',
    emits: ['select'],
    methods: {
        handleSelect(index: string) {
            this.$emit('select', index);
        },
    },
    data() {
        return {
            menuItems: config.menuItems as MenuItem[]
        }
    },
});
</script>

<style lang="scss" scoped>
.topbar {
    position: fixed;
    top: 0px;
    height: 60px;
    width: 100%;
    z-index: 999;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.26);
}
</style>