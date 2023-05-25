<script>
export default {
    props: {
        title: {
            type: String,
            default: 'Text',
        },
        route: {
            type: String,
            default: '',
        },
        icon: {
            type: Object,
            default: null,
        },
        children: {
            type: Array,
            default: () => [],
        }
    },
};
</script>
<script setup>
import { RouterLink } from 'vue-router';
</script>
<template>
    <div class="group flex">
        <div
            class="relative mx-2 my-2 self-center px-1 py-2 text-lg opacity-50 transition-all after:absolute after:right-full after:bottom-[-2px] after:h-[2px] after:w-0 after:bg-sky-500 after:content-[''] group-hover:font-bold group-hover:text-sky-800 group-hover:opacity-80 group-hover:after:right-0 group-hover:after:w-full group-hover:after:transition-all group-hover:after:duration-500"
        >
            <RouterLink :to="route">{{ title }}</RouterLink>
        </div>
    </div>
</template>
