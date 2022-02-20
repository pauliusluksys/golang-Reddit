// import LoginFormComponent from "/components/LoginFormComponent";
import RegisterFormComponent from "../../../components/RegisterFormComponent.vue";
// import Vue from "vue";
import { createWebHistory, createRouter } from "vue-router";
const routes = [
    {
        path: '/login',
        alias: '/login',
        name: 'login',
        component: () => import("../../../components/LoginFormComponent")
    },
    {
        path: '/register',
        alias: '/register',
        name: 'register',
        component: RegisterFormComponent },
];
const router = createRouter({
    // 4. Provide the history implementation to use. We are using the hash history for simplicity here.
    history: createWebHistory(),
    routes // short for `routes: routes`
})

export default router;

// Now the app has started!