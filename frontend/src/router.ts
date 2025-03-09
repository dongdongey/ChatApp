import { createRouter, createWebHashHistory } from 'vue-router';

import Login from "./components/Login.vue";
import Chats from "./components/Chats.vue";

const routes = [
    {
        path: '/',
        name: 'home',
        component: Login,
    },
    {
        path: '/chats/:name',
        name: 'chat',
        component: Chats,
    },
];
const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export default router;