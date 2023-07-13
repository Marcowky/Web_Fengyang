import { createApp } from "vue";
import App from "./App.vue";
import ElementPlus from 'element-plus'// 导入ElementPlus
import 'element-plus/dist/index.css'
// 导入css文件
import "./style.css";
// import './assets/main.css'
import './assets/base.css'
import './assets/iconfont/iconfont.css'

import { createPinia } from "pinia"; // 引入pinia
import { router } from "./routes/router"; // 引入路由
import axios from "axios"; // 引入axios
import { UserStore } from "./stores/UserStore" // 引入UserStore

axios.defaults.baseURL = "http://localhost:8080"; // 服务端地址全局配置

const app = createApp(App);

app.provide("axios", axios); // 将axios全局放入
app.provide("serverUrl", axios.defaults.baseURL)

app.use(ElementPlus) // 使用Element UI
app.use(createPinia()); // 引入pinia

const userStore = UserStore()

// 拦截器传token
axios.interceptors.request.use((config) => {
    config.headers.authorization = `Bearer ${userStore.token}`
    return config
})

app.use(router); // 引入路由
app.mount("#app");

export {axios, userStore}