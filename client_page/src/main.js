import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

// 完整导入element组件库
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// 导入naive组件库
import naive from "naive-ui"; // 引入ui框架
import { createDiscreteApi } from "naive-ui"; // 引入createDiscreteApi

// 导入css文件
import './assets/main.css'

const app = createApp(App)

app.use(router)

app.use(ElementPlus)    // 引入element-plus

app.use(naive); // 引入naive-ui框架

app.mount('#app')
