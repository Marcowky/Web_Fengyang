<template>
  <!-- 登录注册弹框 -->
  <LoginDialog ref="loginDialogRef" />
  <div class="admin">
    <div class="sidebar">
      <div class="logo">
        <h1>后台管理系统</h1>
      </div>
      <div class="menu">
        <el-menu default-active="1">
          <el-menu-item index="user">
            <template #title>
              <i class="el-icon-user"></i>
              <span>用户</span>
            </template>
          </el-menu-item>
          <el-menu-item index="admin">
            <template #title>
              <i class="el-icon-s-custom"></i>
              <span>管理员</span>
            </template>
          </el-menu-item>
          <el-menu-item index="consult">
            <template #title>
              <i class="el-icon-message"></i>
              <span>咨询</span>
            </template>
          </el-menu-item>
          <el-menu-item index="spot">
            <template #title>
              <i class="el-icon-location"></i>
              <span>景点</span>
            </template>
          </el-menu-item>
          <el-menu-item index="consume">
            <template #title>
              <i class="el-icon-wallet"></i>
              <span>消费</span>
            </template>
          </el-menu-item>
          <el-menu-item index="hotel">
            <template #title>
              <i class="el-icon-office-building"></i>
              <span>酒店</span>
            </template>
          </el-menu-item>
          <el-menu-item index="forum">
            <template #title>
              <i class="el-icon-chat-line-square"></i>
              <span>论坛</span>
            </template>
          </el-menu-item>
        </el-menu>
      </div>
    </div>
    <div class="content">
      <div class="header">
        <div class="title">后台管理系统</div>
        <div class="user-info">
          <span>{{ adminName }}</span>
          <el-button type="danger" size="small" @click="logout">退出登录</el-button>
        </div>
      </div>
      <div class="main">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import { ref,onMounted } from 'vue'



// 导入登录注册弹框
import LoginDialog from '../../components/LoginAndRegister.vue'
const loginDialogRef = ref(null)

// 挂载页面时触发
onMounted(() => {
  loginDialogRef.value.showDialog()
})

const state = reactive({
  adminName: 'admin',
})

const logout = () => {
  // TODO: 实现退出登录的逻辑
}
</script>

<style scoped>
.admin {
  display: flex;
  height: 100vh;
}

.sidebar {
  width: 200px;
  background-color: #fff;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.logo {
  height: 60px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.menu {
  padding: 16px 0;
}

.content {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  background-color: #fff;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.main {
  flex: 1;
  padding: 16px;
  overflow-y: auto;
}
</style>