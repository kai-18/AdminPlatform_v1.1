import { createRouter, createWebHistory } from 'vue-router';
import MainLayout from '../layouts/MainLayout.vue';
import IndexPage from '../pages/IndexPage.vue';
import EmployeesPage from '../pages/EmployeesPage.vue';
import UsersPage from '../pages/UsersPage.vue';

const routes = [
  {
    path: '/',
    component: MainLayout,
    children: [
      {
        path: '/',
        component: IndexPage,
      },
      {
        path: 'employees',
        component: EmployeesPage,
      },
      {
        path: 'users',
        component: UsersPage,
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
