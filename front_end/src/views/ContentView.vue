<template>
    <div class="flex flex-col">
        <div></div>
        <div>
            <RouterView class="flex-grow bg-slate-50"></RouterView>
        </div>
    </div>
</template>

<script>

import { NButton } from 'naive-ui'

export default{
    components: {
        NButton,
    },
    data(){
        return{

        }
    },
    method: {
        handleLogout(){

        }
    }
}
</script>