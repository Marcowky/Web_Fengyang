import { createRouter, createWebHashHistory } from "vue-router";

let routes = [
    { path: "/user/login", component: () => import("../views/user/Login.vue") },            // 登录界面
    { path: "/user/register", component: () => import("../views/user/Register.vue") },      // 注册界面
    // { path: "/article/publish", component: () => import("../views/article/Publish.vue") },  // 发布界面
    // { path: "/article", component: () => import("../views/article/Home.vue") },             // 文章中心
    // { path: "/article/detail", component: () => import("../views/article/Detail.vue") },    // 文章详情
    // { path: "/article/update", component: () => import("../views/article/Update.vue") },    // 文章更新
    // { path: "/oldhome", component: () => import("../views/oldHome.vue") },                  // 主页

    { path: "/video", component: () => import("../views/video/Detail.vue") },               // 视频


    { path: "/", component: () => import("../views/home/Home.vue") },                        // 主页

    { path: "/info", component: () => import("../views/info/Home.vue") }, // 咨询
    { path: "/info/page", component: () => import("../views/info/News.vue") },

    { path: "/attraction", component: () => import("../views/attraction/Home.vue") },         // 景点

    { path: "/consumption", component: () => import("../views/consumption/Home.vue") },     // 消费

    { path: "/travel", component: () => import("../views/travel/Home.vue") },                // 住宿、旅游

    { path: "/blog", component: () => import("../views/blog/Home.vue") },                   // 论坛主页
    { path: "/blog/detail", component: () => import("../views/blog/Detail.vue") },          // 论坛文章详情
    { path: "/blog/publish", component: () => import("../views/blog/Publish.vue") },  // 发布界面
    { path: "/blog/update", component: () => import("../views/blog/Update.vue") },    // 文章更新

]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export { router, routes }
