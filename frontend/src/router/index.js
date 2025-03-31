import { createRouter, createWebHistory } from 'vue-router';

// Import view components
import Home from '@/views/Home.vue';
import Pods from '@/views/Pods.vue';
import PodDetail from '@/views/PodDetail.vue';

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
    meta: { title: 'Kubernetes Pods' },
    children: [
      {
        path: ':namespace/:name',
        name: 'PodDetail',
        component: PodDetail,
        meta: { title: 'Pod Details' }
      }
    ]
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

// Update page title based on the route
router.beforeEach((to, from, next) => {
  // If it's the Pod detail page, set a more specific title
  if (to.name === 'PodDetail' && to.params.name) {
    document.title = `Pod: ${to.params.name} - K8s Dashboard`;
  } else {
    document.title = to.meta.title ? `${to.meta.title} - K8s Dashboard` : 'K8s Dashboard';
  }
  next();
});

export default router;