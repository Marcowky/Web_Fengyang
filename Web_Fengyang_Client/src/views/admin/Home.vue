<template>
  <!-- 登录注册弹框 -->
  <LoginDialog ref="loginDialogRef" />
  <el-menu class="topbar" router :default-active="this.$route.fullPath" mode="horizontal" :ellipsis="false">
    <el-menu-item style="position: absolute; left: 0px;" index="0">四色丰阳后台</el-menu-item>
    <div class="弹性盒子" :style="{ flexGrow: 1 }" />
  </el-menu>
  <el-menu class="sidebar" :default-active="this.$route.fullPath" unique-opened="true">
    <el-menu-item>用户管理</el-menu-item>
    
    <el-sub-menu>
      <template #title>文章修改</template>
      
      <template v-for="item in menuItems">
        <template v-if="!item.index.includes('?')">
          <el-menu-item :index="item.index">{{ item.label }}</el-menu-item>
        </template>
      </template>
    </el-sub-menu>
    <el-menu-item>图片更换</el-menu-item>
  </el-menu>
  <router-link to="/admin/article">颜色</router-link>
  <router-view/>
</template>

<script setup>
import { ref, onMounted } from 'vue'

// 导入登录注册弹框
import LoginDialog from '../../components/LoginAndRegister.vue'
const loginDialogRef = ref(null)

// 挂载页面时触发
onMounted(() => {
  loginDialogRef.value.showDialog()
  loginDialogRef.value.userType="admin"
})
import config from '../../config/config.json'

const menuItems = config.menuItems;

</script>

<style scoped>
.topbar {
  position: fixed;
  top: 0px;
  height: 60px;
  width: 100%;
  z-index: 999;
  box-shadow: 0 1px 6px rgba(71, 71, 71, 0.26);
}

.sidebar {
  position: fixed;
  top: 20%;
  width: 150px;
  z-index: 999;
  box-shadow: 2px 0 6px rgba(0, 0, 0, 0.26);
  border-radius: 0 10px 10px 0;
}
</style>