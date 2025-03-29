import { createApp } from 'vue';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import App from './App.vue';
import router from './router';

// Create Vue app
const app = createApp(App);

// Use plugins
app.use(ElementPlus);
app.use(router);

// Mount the app
app.mount('#app');