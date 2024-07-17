import './assets/main.scss';

import { createApp } from 'vue';
import App from './App.vue';
import router from './router';

import { OhVueIcon, addIcons } from 'oh-vue-icons';
import { BiArrowUp, BiArrowDown, BiChevronBarUp, BiChevronBarDown } from 'oh-vue-icons/icons';

addIcons(BiArrowUp, BiArrowDown, BiChevronBarUp, BiChevronBarDown);

const app = createApp(App);

app.use(router);
app.component('v-icon', OhVueIcon);
app.mount('#app');
