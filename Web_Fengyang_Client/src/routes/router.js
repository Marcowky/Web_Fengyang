import { createRouter, createWebHashHistory } from "vue-router";

let routes = [
    { path: "/user/login", component: () => import("../views/user/Login.vue") },            // 登录界面
    { path: "/user/register", component: () => import("../views/user/Register.vue") },      // 注册界面
    { path: "/article/publish", component: () => import("../views/article/Publish.vue") },  // 发布界面
    { path: "/article", component: () => import("../views/article/Home.vue") },             // 文章中心
    { path: "/article/detail", component: () => import("../views/article/Detail.vue") },    // 文章详情
    { path: "/article/update", component: () => import("../views/article/Update.vue") },    // 文章更新
    { path: "", component: () => import("../views/Home.vue") },                             // 主页
    { path: "/video", component: () => import("../views/video/Detail.vue") },               // 视频

]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export { router, routes }
