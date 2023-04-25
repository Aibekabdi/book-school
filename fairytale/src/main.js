import App from './App.vue';
import { createApp } from 'vue';
import router from './router';

// temporarily host url:
// const host = 'http://10.42.0.1:8080';
const host = 'http://localhost:8080';

const app = createApp(App);

app.use(router);

app.mount('#app');

export default host;
