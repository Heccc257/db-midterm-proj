<script>
import NavLink from './NavBar/NavLink.vue';
import store from '../store/index'

export default {
    name: 'Header',
    components: {
        NavLink,
    },
    props: {
        navLinks: {
            type: Array,
            required: true,
            default: () => [],
        },
    },
    mounted(){
    },
    methods: {
        logout(){
            store.commit('setToken', '')
            this.$router.push({name: 'Login'});
        }
    },
};
</script>
<template>
    <div
        id="header"
        class="flex shadow-md h-full w-fit max-w-[300px] flex-col border-r-[1px] border-slate-100 bg-gray-50 py-4 px-6 transition-all duration-300 ease-in-out"
    >
        <div
            class="logo mt-2 mb-4 flex w-fit flex-row justify-center rounded-md bg-blue-50 p-3 text-lg font-bold text-slate-800"
        >
            PKU校园互助系统
        </div>
        <div class="navbar flex flex-col">
            <NavLink v-for="navlink in navLinks" :key="navlink.title" 
                :title="navlink.title" :route="navlink.route" />
        </div>
        <div class="flex flex-grow items-end">
            <button @click="logout">
                <div class="flex items-center justify-center rounded-sm px-3 py-1 transition hover:bg-sky-50">
                    <div class="ml-2 text-slate-600 text-lg opacity-50">{{'退出登录 :)'}}</div>
                </div>
            </button>
        </div>
    </div>
</template>
