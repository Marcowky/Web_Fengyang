<template>
    <div class="top-rectangle"></div>
    <el-menu class="topbar" router :default-active="this.$route.fullPath" mode="horizontal" :ellipsis="false">
        <el-menu-item style="position: absolute; left: 10px; height: 53px;" index="0"><el-image style="width: 180px; height: auto;" src="/title.jpg"/></el-menu-item>
        <div class="弹性盒子" :style="{ flexGrow: 1 }" />
        <template v-for="item in menuItems">
            <template v-if="item.mainMenu === 'noSub'">
                <el-menu-item :index="item.index">{{ item.label }}</el-menu-item>
            </template>
            <template v-if="item.mainMenu === 'hasSub'">
                <el-sub-menu :index="item.index">
                    <template #title>{{ item.label }}</template>
                    <template v-for="subItem in menuItems">
                        <el-menu-item :index="item.index + '?category=' + subItem.index"
                            v-if="subItem.mainMenu === item.index && subItem.mainMenu != '/hotel'">{{
                                subItem.label
                            }}</el-menu-item>
                        <el-menu-item :index="subItem.index"
                            v-if="subItem.mainMenu === item.index && subItem.mainMenu == '/hotel'">{{
                                subItem.label
                            }}</el-menu-item>
                    </template>
                </el-sub-menu>
            </template>
        </template>
        <div class="弹性盒子" :style="{ flexGrow: 1 }" />
    </el-menu>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
// 导入菜单选项配置文件
import config from '../config/config.json';

interface MenuItem {
    index: string;
    label: string;
    mainMenu: string;
}

export default defineComponent({
    name: 'TopBar',
    data() {
        return {
            menuItems: config.menuItems as MenuItem[]
        }
    },
});
</script>

<style lang="scss" scoped>
.top-rectangle {
    position: fixed;
    top: 0;
    height: 50px;
    width: 100%;
    background: -webkit-linear-gradient(top, hsla(207, 63%, 70%, 0.37), rgba(255, 255, 255,0)) no-repeat;
    z-index: 998;
    /* 确保黄色长方形位于其他内容之上 */
}

.topbar {
    position: fixed;
    top: 20px;
    height: 60px;
    border-radius: 30px;
    left: 10%;
    right: 10%;
    z-index: 999;
    box-shadow: 0 0 6px rgba(50, 50, 50, 0.26);
}
</style>