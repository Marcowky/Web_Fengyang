<template>
  <!-- 登录注册弹框 -->
  <LoginDialog ref="loginDialogRef" />
  <div class="common-layout">
    <el-container>
      <el-menu class="topbar" router :default-active="this.$route.fullPath" mode="horizontal" :ellipsis="false">
        <el-menu-item style="position: absolute; left: 0px;" index="0">四色丰阳后台</el-menu-item>
        <div class="弹性盒子" :style="{ flexGrow: 1 }" />
      </el-menu>
      <el-container>
        <el-aside v-if="show" width="200px">
          <!-- 侧边栏 -->
          <el-menu class="sidebar" router :default-active="this.$route.fullPath" unique-opened=true>
            <el-sub-menu index="/admin/user">
              <template #title>用户管理</template>
              <el-menu-item index="/admin/user?category=client">客户</el-menu-item>
              <el-menu-item index="/admin/user?category=admin">管理员</el-menu-item>
            </el-sub-menu>
            <el-sub-menu index="/admin/article">
              <template #title>文章修改</template>
              <template v-for="item in menuItems">
                <template v-if="item.index.includes('/') && item.index != '/home' && item.mainMenu != '/hotel'">
                  <el-menu-item :index="'/admin/article?category=' + item.index.substring(1) + 'Article'">{{ item.label
                  }}</el-menu-item>
                </template>
              </template>
            </el-sub-menu>
<!--            <el-menu-item index="/admin/image">图片更换</el-menu-item>-->
            <el-sub-menu index="/admin/carousel">
              <template #title>轮播图管理</template>
<!--              <el-menu-item index="/admin/carousel?category=home">首页</el-menu-item>-->
              <el-menu-item index="/admin/carousel?category=consumpAttraction">消费景点</el-menu-item>
            </el-sub-menu>

          </el-menu>
        </el-aside>
        <el-main class="mainContent">
          <!-- 内容 -->
          <router-view v-if="showRouteView" />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import LoginDialog from '../../components/LoginAndRegister.vue' // 导入登录注册弹框
import { UserStore } from '../../stores/UserStore'
import { toRef } from 'vue';
import config from '../../config/config.json'

const router = useRouter()
const menuItems = config.menuItems;
const loginDialogRef = ref(null)
const userStore = UserStore()
const showRouteView = ref(userStore.token != '')
const tokenRef = toRef(userStore, 'token');
const show = ref(true);

// 挂载页面时触发
onMounted(() => {
  loginDialogRef.value.showDialog()
  loginDialogRef.value.user.userType = "admin"
})

watch(tokenRef, (newToken) => {
  showRouteView.value = newToken !== '';
});

router.beforeEach((to, from) => {
  if (to.path == '/admin/article/publish' || to.path == '/admin/article/update') {
    show.value = false;
  } else {
    show.value = true;
  }
  if (to != from) {
    window.scrollTo(0, 0);
  }
})
</script>

<style scoped>
.topbar {
  position: fixed;
  top: 0px;
  height: 60px;
  width: 100%;
  z-index: 999;
  box-shadow: 0 0 6px rgba(50, 50, 50, 0.26);
}

.sidebar {
  position: fixed;
  top: 20%;
  width: 150px;
  z-index: 999;
  box-shadow: 0 0 6px rgba(50, 50, 50, 0.26);
  border-radius: 0 10px 10px 0;
}

.mainContent {
  margin-top: 100px;
  padding-left: 0;
}
</style>