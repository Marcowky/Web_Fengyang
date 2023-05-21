import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../views/HomePage.vue'

let routes = [
  { path:"/", redirect: "/HomePage"},   // �ض���(�о���������)����ҳһ�򿪵�Ĭ����ʾ
  { path: "/HomePage", name: "HomePage", component: HomePage },
  { path: "/InfoPage", name: "InfoPage", component: () => import("../views/InfoPage.vue")}
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
});

export default router
