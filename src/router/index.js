import Vue from 'vue'
import Router from 'vue-router'

import ViewLogin from '../views/login/index'
import ViewRegister from '../views/register/index'
import View404 from '../views/404'
import ViewContainer from '../views/layout/container'
import ViewList from '../views/list/index'
import ViewDetail from '../views/detail/index'

Vue.use(Router)

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
    {
        path: '/login',
        component: ViewLogin,
        hidden: true
    },
    {
        path: '/register',
        component: ViewRegister,
        hidden: true
    },
    {
        path: '/404',
        component: View404,
        hidden: true
    },
    {
        path: '/',
        redirect: '/',
        component: ViewContainer,
        children: [
            {
                path: '',
                component: ViewList
            },
            {
                path: 'detail',
                component: ViewDetail
            }
        ]
    },
    // 404 page must be placed at the end !!!
    { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
    routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
    const newRouter = createRouter()
    router.matcher = newRouter.matcher // reset router
}

export default router
