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
                    path: '',
                    component: () => import('@/views/main/session.vue'),
                    children:[
                        {
                            name:'chat',
                            path:'chat',
                            component:() => import('@/views/main/chat.vue')
                        }
                    ]
                },
                {
                    name:"user",
                    path:"user",
                    component:()=>import('@/views/main/homePage.vue')
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
            },
            {
                path:'register',
                name:'register',
                component:()=> import('@/views/ready/register.vue')
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
    if (!token && to.name != 'login' && to.name != 'register') {
        return {name: 'login'}
    }
})

export default router