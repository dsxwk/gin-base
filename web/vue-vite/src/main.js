import { createApp } from 'vue';
import App from '../App.vue';

import routers from '@routers/router';
import {createRouter, createWebHistory} from "vue-router";

const routes = createRouter({
    history: createWebHistory(),
    routes: routers
});

const app = createApp(App);

app.use(routes);
app.mount('#app');
