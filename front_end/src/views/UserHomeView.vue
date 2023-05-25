<template>
    <div class="my-10 mx-5">
        <n-card title="用户信息" class="!rounded-md my-1">
            <div class="mb-2 flex flex-col">
                <div class="my-2">
                    <span class="text-[1.2rem] text-slate-400">用户id: </span>
                    <span class="text-[1.1rem] text-slate-700"> {{ info.user_id}} </span>
                </div>
                <div class="my-2">
                    <span class="text-[1.2rem] text-slate-400">昵称: </span>
                    <span class="text-[1.1rem] text-slate-700"> {{ info.nick_name }} </span>
                </div>
                <div class="my-2">
                    <span class="text-[1.2rem] text-slate-400">全名: </span>
                    <span class="text-[1.1rem] text-slate-700"> {{ info.full_name }} </span>
                </div>
                <div class="my-2">
                    <span class="text-[1.2rem] text-slate-400">电话号码: </span>
                    <span class="text-[1.1rem] text-slate-700"> {{ info.phone_number }} </span>
                </div>
                <div class="my-2">
                    <span class="text-[1.2rem] text-slate-400">已发出订单数: </span>
                    <span class="text-[1.1rem] text-slate-700"> {{ info.offer_count }} </span>
                </div>
                <div class="my-2">
                    <span class="text-[1.2rem] text-slate-400">已接取订单数: </span>
                    <span class="text-[1.1rem] text-slate-700"> {{ info.task_count }} </span>
                </div>
            </div>
        </n-card>
    </div>
</template>

<script>
import { NCard } from 'naive-ui'
import store from '../store/index'
import API from '../api/CommonData'

export default {
    components: {
        NCard,
    },
    data() {
        return {
            info: {},
        }
    },
    methods: {
        async getUserInfo() {
            var infores = await API.getForm('/user_info/' + store.state.UserBasic.uid)
            // get user_info
            if(infores.status === 200) {
                store.commit('setUserInfo', infores.data)
                this.info = infores.data
            }
            else {
                $notification.error({
                    title: 'user_info请求失败',
                    message: res.message,
                });
            }
        }
    },
    mounted(){
        this.getUserInfo()
    }
}
</script>