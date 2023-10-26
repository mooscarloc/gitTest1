export default [
    {
        path: '/',
        name: 'Index',
        redirect: '/home',
        component: () => (import('../components/main')),
        children: [
            {
                path: '/home',
                name: 'Home',
                component: () => (import('../views/home.vue')),
                meta:{
                    title:'Management System - Home'
                }
            },
            {
                path: '/home/project',
                name: 'Project',
                component: () => (import('../views/project/index.vue')),
                meta:{
                    title:'Management System - Dataset Display'
                }
            },
            {
                path: '/info',
                name: 'Info',
                component: () => (import('../views/project/info.vue')),
                meta:{
                    title:'Management System - Project Config'
                }
            },
            {
                path: '/info/config',
                name: 'Config',
                component: () => (import('../views/project/config.vue')),
                meta:{
                    title:'Management System - Factor Config'
                }
            },
        ]
    },
    {
        path:'/login',
        name:'Login',
        component:()=>import('../views/login'),
        meta:{
            title:'Management System - Login'
        }
    }
]