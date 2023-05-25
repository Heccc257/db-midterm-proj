<template>
    <div class="flex h-screen flex-row">
        <div class="w-72">
            <Aside :navLinks="navLinks"></Aside>
        </div>
        <div class="flex flex-col w-full">
            <Content></Content>
        </div>
    </div>
</template>

<script>
import Aside from '../components/Aside.vue';
import Content from './ContentView.vue'
import { NNotificationProvider, useNotification, NButton } from 'naive-ui';

export default {
    components: {
        Aside,
        NNotificationProvider,
        NButton,
        Content
    },
    data() {
        return {
            navLinks: [
                {
                    title: '个人主页',
                    route: '/user_home',
                },
                {
                    title: '订单列表',
                    route: '/offer_list',
                },
                {
                    title: '我的接单',
                    route: '/my_accept_offer',
                },
                {
                    title: '我的发单',
                    route: '/my_offer',
                },
                {
                    title: '用户排名',
                    route: '/rank',
                },
            ],
        };
    },

    setup() {
        window.$notification = useNotification();
    },
};
</script>
