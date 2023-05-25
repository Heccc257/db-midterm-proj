import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            name: 'Home',
            path: '/',
            redirect: '/user_home',
            component: () => import('../views/HomeView.vue'),
            children: [
                {
                    name: 'UserHome',
                    path: '/user_home',
                    component: () => import('../views/UserHomeView.vue'),
                },
                {
                    name: 'OfferList',
                    path: '/offer_list',
                    component: () => import('../views/OfferListView.vue'),
                },
                {
                    name: 'MyAcceptOffer',
                    path: '/my_accept_offer',
                    component: () => import('../views/MyAcceptOfferView.vue'),
                },
                {
                    name: 'MyOffer',
                    path: '/my_offer',
                    component: () => import('../views/MyOfferView.vue'),
                },
                {
                    name: 'Rank',
                    path: '/rank',
                    component: () => import('../views/RankView.vue'),
                },
            ]
        },
        {
            name: 'Login',
            path: '/login',
            component: () => import('../views/LoginView.vue')
        },
    ]
})

// router.beforeEach((to, from, next) => {
//     if (!localStorage.getItem('token') && to.name != 'Login') {
//         next({name: 'Login'});
//     } else {
//         next();
//     }
// });

export default router;