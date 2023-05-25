import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router';
import vuex from 'vuex'

export const app = createApp(App)
app.use(router)
app.use(vuex)

const meta = document.createElement('meta');
meta.name = 'naive-ui-style';
document.head.appendChild(meta);

app.mount('#app');
