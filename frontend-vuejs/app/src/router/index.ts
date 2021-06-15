import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import AccountDetails from "@/components/Account.vue";
import LoginPage from "@/components/Login.vue";
import Logout from "@/components/Logout.vue";
import {USERTOKEN} from "@/components/types";
import Tickets from "@/components/Tickets.vue";
import Home from "@/components/Home.vue";
import Forbidden from "@/components/403.vue";

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/login',
    name: 'login',
    component: LoginPage
  },
  {
    path: '/account/:accountid',
    name: 'accountDetails',
    component: AccountDetails
  },
  {
    path: '/logout',
    name: 'logout',
    component: Logout
  },
  {
    path: '/tickets',
    name: 'tickets',
    component: Tickets
  },
  {
    path: '/forbidden',
    name:'forbidden',
    component: Forbidden
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) =>{
  const loggedIn = localStorage.getItem(USERTOKEN);
  const publicPages = ['/login'];
  const authRequired = !publicPages.includes(to.path);
  if (authRequired && !loggedIn) {
    return next('/login');
  }
  next();
})
export default router
