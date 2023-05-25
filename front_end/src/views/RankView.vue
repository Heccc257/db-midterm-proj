<template>
    <div class="my-10 mx-5">
        <div v-for="item in tableData" :key="item.offer_id">
            <n-card title="" class="!rounded-md my-1">
                <div class="mb-2 flex flex-col">
                    <div>
                        <span class="text-[1.2rem] text-slate-400">分数: </span>
                        <span class="text-[1.1rem] text-slate-700"> {{ item.avg_rating }} </span>
                    </div>
                    <div>
                        <span class="text-[1.2rem] text-slate-400">用户id: </span>
                        <span class="text-[1.1rem] text-slate-700"> {{ item.user_id }} </span>
                    </div>
                    <div>
                        <span class="text-[1.2rem] text-slate-400">昵称: </span>
                        <span class="text-[1.1rem] text-slate-700"> {{ item.nick_name }} </span>
                    </div>
                </div>
            </n-card>
        </div>
    </div>
</template>

<script>
import { NCard, NButton, NModal } from 'naive-ui';
import API from '../api/CommonData'

export default {
    components: {
        NCard,
        NButton,
        NModal,
    },
    data() {
        return {
            tableData: [],
        }
    },
    methods:{
        async getTableData() {
            let res = await API.getForm('/best_users')
            if(res.status === 200) {
                this.tableData = res.data
            }
            else {
                $notification.error({
                    title: '数据请求失败',
                    message: res.message,
                });
            }
        }
    },
    mounted() {
        this.getTableData()
    }
}
</script>