<template>
    <div class="top-rectangle"></div>
    <el-menu class="topbar" router :default-active="this.$route.fullPath" mode="horizontal" :ellipsis="false">
        <div class="弹性盒子" :style="{ flexGrow: 1 }" />
        <template v-for="item in menuItems">
            <template v-if="item.mainMenu == 'noSub'" >
                <el-menu-item :index="item.index" class="fontset">{{ item.label }}</el-menu-item>
            </template>
            <template v-if="item.mainMenu == 'hasSub'">
                <el-sub-menu :index="item.index" >
                    <template #title >
                        <div  class="fontset">
                            {{ item.label }}
                        </div>
                    </template>
                    <template v-for="subItem in menuItems">
                        <el-menu-item :index="item.index+'?category='+subItem.index" v-if="subItem.mainMenu==item.index">{{
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
import config from '../../config/config.json';

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
.topbar {
    position: absolute;
    top: 20px;
    height: 60px;
    border-radius: 30px;
    border: none;
    left: 10%;
    right:10%;
    z-index: 999;
    background-color: rgb(255, 255, 255,0); 
}

.fontset{
    font-family: "Microsoft YaHei", sans-serif; /* 使用微软雅黑字体 */
    font-weight: bold; /* 加粗 */
    color: white; /* 白色字体颜色 */
}

.top-rectangle {
    position: fixed;
    top: 0;
    height: 100px;
    width: 100%;
    background: -webkit-linear-gradient(top, hsla(0, 0%, 24%, 0.727), rgba(255, 255, 255,0)) no-repeat;
    z-index: 998;
    /* 确保黄色长方形位于其他内容之上 */
}

</style>