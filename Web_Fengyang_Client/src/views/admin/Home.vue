<template>
  <!-- 登录注册弹框 -->
  <LoginDialog ref="loginDialogRef" />
  

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