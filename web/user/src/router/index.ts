import {createRouter, createWebHistory} from "vue-router";

const routes = [
    {
        path: '/',
        name: 'main',
        // 懒加载
        component: () => import('@/views/main.vue'),
        children:
            [
                {
                    name: 'contact',
                    path: 'contact',
                    component: () => import('@/views/main/contact.vue')
                },
                {
                    name: 'session',
                    path: 'session',
                    component: () => import('@/views/main/session.vue')
                }
            ]

    },
    {
        path: '/ready',
        name: 'ready',
        component: () => import('@/views/ready.vue'),
        children: [
            {
                path: '/',
                name: 'login',
                component: () => import('@/views/ready/login.vue')
            }
        ]
    }
]
const router = createRouter({
    routes,
    history: createWebHistory()
})
// 全局前置守卫
router.beforeEach((to) => {
    let token = localStorage.getItem("token")
    if (!token && to.name != 'login') {
        return {name: 'login'}
    }
})

export default router