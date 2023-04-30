import { createRouter, createWebHashHistory } from "vue-router";

let routes = [
    { path: "/user/login", component: () => import("../views/user/Login.vue") },
    { path: "/user/register", component: () => import("../views/user/Register.vue") },
    { path: "/article/publish", component: () => import("../views/article/Publish.vue") },
    { path: "/article", component: () => import("../views/article/Home.vue") },
    { path: "/article/detail", component: () => import("../views/article/Detail.vue") },
    { path: "/article/update", component: () => import("../views/article/Update.vue") },
    { path: "", component: () => import("../views/Home.vue") },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export { router, routes }
