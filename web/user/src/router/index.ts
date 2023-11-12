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
                    component: () => import('@/views/main/contact.vue'),
                    children:[
                        {
                            path:'userinfo/:index',
                            name:"userinfo",
                            component: () => import('@/views/main/contact/userInfo.vue'),
                            beforeEnter: (_to, from, next) => {
                                if(from.name === 'contact'){
                                    next()
                                }
                            }
                        }
                    ]
                },
                {
                    name: 'session',
                    path: '',
                    component: () => import('@/views/main/session.vue'),
                    children:[
                        {
                            name:'chat',
                            path:'chat',
                            component:() => import('@/views/main/session/chat.vue')
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
        beforeEnter: (_to, _from, next) => {
            // 路由守卫
            if (localStorage.getItem('token')) {
                return false
            }
            next()
                
        },
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
    },
    {
        path:'/repassword',
        name:"repassword",
        component:()=>import('@/views/rePassword.vue'),
        beforeEnter: (_to, from, next)=> {
            if(from.name != 'login' && from.name != 'user'){
                return false
            }
            next()
        },
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