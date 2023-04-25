/* eslint-disable import/order */
import { createRouter, createWebHistory } from 'vue-router';

import AboutView from '@/views/AboutView.vue';
import AdminView from '@/views/AdminView.vue';
import ContentView from '@/views/ContentView.vue';
import CreateView from '@/views/CreateView.vue';
import HomepageView from '@/views/HomepageView.vue';
import LandingView from '@/views/LandingView.vue';
import LoginView from '@/views/LoginView.vue';
import NotFoundView from '@/views/NotFoundView.vue';
import PaymentView from '@/views/PaymentView.vue';
import ProfileView from '@/views/ProfileView.vue';
import ShopView from '@/views/ShopView.vue';

// Route guards for admin:
function guardAdmin(to, from, next) {
  const admin = localStorage.getItem('admin');

  if (admin) {
    next();
  } else {
    next('/404');
  }
}

// Re Route to Admin Panel
function reRouteAdmin(to, from, next) {
  const admin = localStorage.getItem('admin');

  if (admin) {
    next('/admin');
  } else {
    next();
  }
}

// Re Route to Home if Logged In
function reRouteHome(to, from, next) {
  const user = localStorage.getItem('userType');

  if (user) {
    next('/home');
  } else {
    next();
  }
}

const routes = [
  {
    path: '/',
    name: 'home',
    component: LandingView,
    meta: {
      title: 'Главная',
      auth: false,
    },
  },
  {
    path: '/about',
    name: 'about',
    component: AboutView,
    meta: {
      title: 'О нас',
      auth: false,
    },
  },
  {
    path: '/register',
    name: 'payment',
    component: PaymentView,
    meta: {
      title: 'Регистрация',
      auth: false,
    },
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView,
    meta: {
      title: 'Логин',
      auth: false,
    },
    beforeEnter: reRouteHome,
  },
  {
    path: '/home',
    name: 'homepage',
    component: HomepageView,
    meta: {
      title: 'Домашная страница',
      auth: true,
    },
    beforeEnter: reRouteAdmin,
  },
  {
    path: '/profile',
    name: 'profile',
    component: ProfileView,
    meta: {
      title: 'Профиль',
      auth: true,
    },
  },
  {
    path: '/shop',
    name: 'shop',
    component: ShopView,
    meta: {
      title: 'Магазин',
      auth: true,
    },
  },
  {
    // dynamic view routing
    path: '/content/:content_id',
    name: 'content',
    component: ContentView,
    meta: {
      title: 'Чтение',
      auth: true,
    },
  },
  {
    path: '/create',
    name: 'create',
    component: CreateView,
    meta: {
      title: 'Новая книга',
      auth: false,
    },
    beforeEnter: guardAdmin,
  },
  {
    path: '/admin',
    name: 'admin',
    component: AdminView,
    meta: {
      title: 'Админ Панель',
      auth: false,
    },
  },
  {
    path: '/:pathMatch(.*)*',
    name: '404-page',
    component: NotFoundView,
    meta: {
      title: 'Страница не найдена',
      auth: false,
    },
  },
  // route level code-splitting
  // this generates a separate chunk (about.[hash].js) for this route
  // which is lazy-loaded when the route is visited.
  // component: () => import(/* webpackChunkName: "about" */ ),
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Change document titles:
router.beforeEach((to, from, next) => {
  document.title = `${to.meta.title}` || 'Fairytale Platform';
  next();
});

// route guards for auth:
router.beforeEach(async (to, from, next) => {
  if (to.matched.some((res) => res.meta.auth)) {
    const user = localStorage.getItem('admin') || localStorage.getItem('userType') || null;
    if (!user || user === null) {
      next({ name: 'login' });
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
