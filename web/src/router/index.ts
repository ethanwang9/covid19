import {createRouter, createWebHistory, RouteRecordRaw} from "vue-router";
import useStore from "../store/index";

const routes: RouteRecordRaw[] = [
    {
        path: "/",
        name: "index",
        meta: {
            title: "首页",
        },
        component: () => import("@/views/index/index.vue"),
    },
    {
        path: "/publicity",
        name: "publicity",
        meta: {
            title: "防疫宣传",
        },
        component: () => import("@/views/index/publicity.vue")
    },
    {
        path: "/login",
        name: "login",
        meta: {
            title: "登录账号",
        },
        component: () => import("@/views/index/login.vue")
    },
    {
        path: "/admin",
        name: "admin",
        component: () => import("@/views/admin/index.vue"),
        meta: {
            auth: true,
        },
        redirect: '/admin/main',
        children: [
            {
                path: "main",
                name: "adminMain",
                meta: {
                    title: "数据看板",
                    auth: true,
                },
                component: () => import("@/views/admin/main/index.vue"),
            },
            {
                path: "query",
                name: "adminQuery",
                meta: {
                    auth: true,
                },
                children: [
                    {
                        path: "danger",
                        name: "adminQueryDanger",
                        meta: {
                            title: "风险地区",
                            auth: true,
                        },
                        component: () => import("@/views/admin/query/danger.vue"),
                    },
                    {
                        path: "travel",
                        name: "adminQueryTravel",
                        meta: {
                            title: "出行政策",
                            auth: true,
                        },
                        component: () => import("@/views/admin/query/travel.vue"),
                    },
                ]
            },
            {
                path: "user",
                name: "adminUser",
                meta: {
                    auth: true,
                },
                children: [
                    {
                        path: "self",
                        name: "adminUserSelf",
                        meta: {
                            title: "个人资料",
                            auth: true,
                        },
                        component: () => import("@/views/admin/user/self.vue"),
                    },
                    {
                        path: "manage",
                        name: "adminUserManage",
                        meta: {
                            title: "管理用户",
                            auth: true,
                        },
                        component: () => import("@/views/admin/user/manage.vue"),
                    },
                ]
            },
            {
                path: "sys",
                name: "adminSys",
                meta: {
                    auth: true,
                },
                children: [
                    {
                        path: "index",
                        name: "adminSysIndex",
                        meta: {
                            title: "系统设置",
                            auth: true,
                        },
                        component: () => import("@/views/admin/sys/index.vue"),
                    },
                    {
                        path: "mp",
                        name: "adminSysMp",
                        meta: {
                            title: "公众号设置",
                            auth: true,
                        },
                        component: () => import("@/views/admin/sys/mp.vue"),
                    },
                ]
            },
        ]
    },
    {
        path: "/status",
        name: "status",
        children: [
            {
                path: "403",
                name: "status403",
                meta: {
                    title: "无法访问该页面",
                },
                component: () => import("@/views/status/403.vue"),
            },
            {
                path: "404",
                name: "status404",
                meta: {
                    title: "404",
                },
                component: () => import("@/views/status/404.vue"),
            },
        ],
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach((to, from, next) => {
    // 动态设置网页标题
    window.document.title = to.meta.title as string || "COVID-19 大数据可视化系统"

    // 404
    if (!router.getRoutes().find((v => v.path === to.path))) {
        next({name: 'status404'})
    }

    // 登录拦截
    const store = useStore().app
    if (to.meta.auth) {
        if (store.isLogin()) {
            next()
        } else {
            next({name: "login"})
        }
    } else {
        if (to.path === "/login" && store.isLogin()) {
            next({name: "admin"})
        }
    }

    next()
})

router.afterEach((to, from, failure) => {

})

export default router