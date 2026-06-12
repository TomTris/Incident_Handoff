import { createRouter, createWebHistory } from 'vue-router'
import { isAuthenticated } from '@/api.ts'

const authedRoutes = {

}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/entry',
      name: 'entry',
      component: () => {},
    },
    {
      path: '/incident-detail/:id',
      name: 'incident-detail',
      component: () => import('@/views/IncidentDetailView.vue'),
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue'),
    },
    {
      path: '/incidents',
      name: 'incidents',
      component: () => import('@/views/IncidentsView.vue'),
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'default',
      redirect: "/entry",
    },
  ],
})

router.beforeEach(async (to) => {
  if (to.name == "entry") {
    const ok = await isAuthenticated()
    return ok == true ? "incidents" : "login"
  }
})

export default router
