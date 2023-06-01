<template>
    <el-menu class="sidebar" router :default-active="this.$route.fullPath" >
        <el-menu-item v-for="item in items" :key="item.mainMenu+'?category='+item.index" :index="item.mainMenu+'?category='+item.index">
            <span>{{ item.label }}</span>
        </el-menu-item>
    </el-menu>
    
</template>
  
<script lang="ts">
import { defineComponent, PropType } from 'vue';

export default defineComponent({
    name: 'SideBar',
    props: {
        items: {
            type: Array as PropType<{ index: string; label: string; mainMenu: string}[]>,
            required: true,
        },
    }
});
</script>
  
<style scoped>
.sidebar {
    position: fixed;
    top: 25%;
    width: 150px;
    z-index: 999;
    box-shadow: 2px 0 6px rgba(0, 0, 0, 0.26);
    border-radius: 0 10px 10px 0;
}
</style>

