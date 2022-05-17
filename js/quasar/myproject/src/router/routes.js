
const routes = [
  {
    path: '/',
    component: () => import('layouts/UnauthenticatedLayout.vue'),
    children: [
      { path: '', redirect: to => '/auth' },
      { path: '/auth', name: 'auth', component: () => import('pages/Auth.vue') }
    ]
  },
  {
    path: '/panel',
    component: () => import('layouts/AuthenticatedLayout.vue'),
    children: [
      { path: '', redirect: to => '/panel/home' },
      { path: '/panel/home', name: 'home', component: () => import('pages/Home.vue') },
      { path: '/panel/dns', name: 'dns', component: () => import('pages/Dns.vue') }
    ]
  },
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes
