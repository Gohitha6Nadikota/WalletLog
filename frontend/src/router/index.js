import { createRouter, createWebHistory } from 'vue-router'

import LandingPage from '../pages/LandingPage.vue'
import RegisterPage from '../pages/RegisterPage.vue'
import LoginPage from '../pages/LoginPage.vue'
import DashboardPage from '../pages/Dashboard.vue'
import HomePage from '../pages/Home.vue'
import ExpenseForm from '../pages/ExpenseForm.vue'
import Summary from '../pages/Summary.vue'

const routes = [
  { path: '/', component: LandingPage },
  { path: '/register', component: RegisterPage },
  { path: '/login', component: LoginPage },
  { path: '/dashboard', component: DashboardPage ,meta: { requiresAuth: true }},
  { path: '/home', component: HomePage, meta: { requiresAuth: true } },
  { path: '/add', name: 'AddExpense', component: ExpenseForm, meta: { requiresAuth: true } },
  { path: '/edit/:id', name: 'EditExpense', component: ExpenseForm, meta: { requiresAuth: true } },
  { path: '/summary', component: Summary, meta: { requiresAuth: true } },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const isAuthenticated = !!localStorage.getItem('token')

  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login')
  } else {
    next()
  }
})

export default router
