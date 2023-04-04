import { createWebHistory, createRouter } from "vue-router";  

import BashView from '../components/BashView.vue'
import ProcessView from '../components/ProcessView.vue'
import LogViews from '../components/LogViews.vue'
import InfoProcess from '../components/InfoProcess.vue'



const routes = [
  {
    path: '/bash',
    name: 'bash',
    component: BashView
  },
  {
    path: '/',
    name: 'main',
    component: ProcessView
  },
  {
    path: '/log',
    name: 'logs',
    component: LogViews
  },
  {
    path: '/proc/:pid',
    name: 'proc',
    component: InfoProcess
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;