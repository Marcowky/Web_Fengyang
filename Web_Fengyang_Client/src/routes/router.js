import { createRouter, createWebHistory } from "vue-router";

let routes = [
    {
        path: '/',
        redirect: '/home',
        component: () => import("../views/clientLayout.vue"),
        children: [
            {
                path: '/home',
                component: () => import("../views/home/Home.vue")
            },
            {
                path: '/info',
                component: () => import("../views/info/Home.vue")
            },
            {
                path: "/attraction",
                component: () => import("../views/attraction/Home.vue")
            },
            {
                path: "/consumption",
                component: () => import("../views/consumption/Home.vue")
            },
            {
                path: "/hotel",
                component: () => import("../views/travel/hotel.vue")
            },
            {
                path: "/traffic",
                component: () => import("../views/travel/traffic.vue")
            },
            {
                path: '/blog',
                component: () => import("../views/blog/Home.vue")
            },
            {
                path: '/article/publish',
                component: () => import("../views/article/Publish.vue")
            },
            {
                path: "/article/update",
                component: () => import("../views/article/Update.vue")
            },
            {
                path: "/article/detail",
                component: () => import("../views/article/Detail.vue")
            },
        ]
    },

    {
        path: "/admin",
        redirect: '/admin/user?category=client',
        component: () => import("../views/admin/adminLayout.vue"),
        children: [
            {
                path: "article",
                component: () => import("../views/admin/ArticleAdmin.vue")
            },
            {
                path: '/admin/article/publish',
                component: () => import("../views/article/Publish.vue")
            },
            {
                path: '/admin/article/update',
                component: () => import("../views/article/Update.vue")
            },
            {
                path: "user",
                component: () => import("../views/admin/UserAdmin.vue")
            },
            {
                path: "image",
                component: () => import("../views/admin/ImageAdmin.vue")
            },
        ]

    },

    // { path: "/login", component: () => import("../views/user/login.vue") },               // 视频
    // { path: "/", component: () => import("../views/home/Home.vue") },                        // 主页
    // { path: "/info", component: () => import("../views/info/Home.vue") }, // 咨询
    // { path: "/info/page", component: () => import("../views/info/News.vue") },
    // { path: "/attraction", component: () => import("../views/attraction/Home.vue") },         // 景点
    // { path: "/consumption", component: () => import("../views/consumption/Home.vue") },     // 消费
    // { path: "/travel", component: () => import("../views/travel/Home.vue") },                // 住宿、旅游
    // { path: "/blog", component: () => import("../views/blog/Home.vue") },                   // 论坛主页
    // { path: "/blog/detail", component: () => import("../views/blog/Detail.vue") },          // 论坛文章详情
    // { path: "/blog/publish", component: () => import("../views/blog/Publish.vue") },  // 发布界面
    // { path: "/blog/update", component: () => import("../views/blog/Update.vue") },    // 文章更新
    // { path: "/admin/article", component: () => import("../views/admin/ArticleAdmin.vue") },    // admin
]

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export { router, routes }
