import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

// ��������element�����
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// ����naive�����
import naive from "naive-ui"; // ����ui���
import { createDiscreteApi } from "naive-ui"; // ����createDiscreteApi

// ����css�ļ�
import './assets/main.css'

const app = createApp(App)

app.use(router)

app.use(ElementPlus)    // ����element-plus

app.use(naive); // ����naive-ui���

app.mount('#app')
