import {createApp} from 'vue'
import App from './App.vue'
import {createPinia} from "pinia";
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import router from "./router";
import DataVVue3 from '@kjgl77/datav-vue3'

import 'normalize.css'
import 'animate.css/animate.min.css'
import 'animate.css/animate.compat.css'
import 'element-plus/theme-chalk/display.css'


const app = createApp(App)
// 使用 Pinia
const pa = createPinia()
app.use(pa)
pa.use(piniaPluginPersistedstate)

app.use(router)
app.use(DataVVue3)
app.mount('#app')
