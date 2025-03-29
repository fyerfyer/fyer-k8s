import { createRouter, createWebHistory } from 'vue-router';

// Import views
import Home from '@/views/Home.vue';
import Pods from '@/views/Pods.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: { title: 'Dashboard' }
  },
  {
    path: '/pods',
    name: 'Pods',
    component: Pods,
    meta: { title: 'Kubernetes Pods' }
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

// Update page title based on route
router.beforeEach((to, from, next) => {
  document.title = to.meta.title ? `${to.meta.title} - K8s Dashboard` : 'K8s Dashboard';
  next();
});

export default router;