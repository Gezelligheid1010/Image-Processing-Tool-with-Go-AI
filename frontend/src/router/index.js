import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import SignUp from '../views/SignUp.vue'
import NewCategory from "../views/NewCategory.vue"
// import Draw from "../views/Draw.vue"
const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err);
}
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/login' // 默认重定向到登录页面
  },
  {
    path: '/category',
    name: 'Category',
    component: () => import('../views/Category.vue'),
  },
  {
    path: '/createCategory',
    name: 'NewCategory',
    component: NewCategory,
    meta: { requireAuth: true }
  },
  {
    path: '/categoryDetail/:category_id',
    name: 'CategoryDetail',
    component: () => import('../views/CategoryDetail.vue'),
  },
  {
    path: '/processImage/:category_id',
    name: 'Draw',
    component: () => import('../views/Draw.vue'),
  },
  {
    path: '/login',
    name: "Login",
    component: Login
  },
  {
    path: '/signup',
    name: "SignUp",
    component: SignUp
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router