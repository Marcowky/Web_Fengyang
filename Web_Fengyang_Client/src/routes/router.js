import { createRouter, createWebHashHistory } from "vue-router";

let routes = [
    { path: "/login", component: () => import("../views/Login.vue") },
    { path: "/register", component: () => import("../views/Register.vue") },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export { router, routes }
