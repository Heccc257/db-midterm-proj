<template>
    <div class="my-10 mx-5">
        <div v-for="item in tableData" :key="item.offer_id">
            <n-card title="" class="!rounded-md my-1">
                <n-row>
                    <n-col :span="12">
                        <div class="mb-2 flex flex-col">
                            <div>
                                <span class="text-[1.2rem] text-slate-400">任务: </span>
                                <span class="text-[1.1rem] text-slate-700"> {{ item.category_name }} </span>
                            </div>
                            <div>
                                <span class="text-[1.2rem] text-slate-400">地点: </span>
                                <span class="text-[1.1rem] text-slate-700"> {{ item.from + ' --> ' + item.to }} </span>
                            </div>
                            <div>
                                <span class="text-[1.2rem] text-slate-400">时限: </span>
                                <span class="text-[1.1rem] text-slate-700"> {{ item.time_limit }} </span>
                            </div>
                        </div>
                    </n-col>
                    <n-col :span="10">
                        <h1 class="text-[1.6rem]"> 奖金: </h1>
                        <h1 class="text-[1.6rem] text-center">{{ item.reward_amount + '元' }} </h1>
                    </n-col>
                    <n-col :span="2">
                        <!-- <n-button strong secondary type="warning" size="large">
                                {{'抢!'}}
                        </n-button> -->
                        <div class="h-full w-full bg-amber-200 hover:bg-amber-300 rounded-md text-[2.4rem] text-center hover:cursor-pointer"
                            @click="handleAccept(item)"
                        >
                            <div class="inline-block align-middle h-full w-full py-4">抢!</div>
                            
                        </div>
                    </n-col>
                </n-row>
            </n-card>
        </div>
    </div>
    <n-modal
        v-model:show="showModal"
        preset="dialog"
        title="确认"
        content="确认接单?"
        positive-text="确认"
        negative-text="别急"
        @positive-click="submitAccept"
        @negative-click="onNegativeClick"
    />
</template>

<script>
import { NCard, NRow, NCol, NButton, NModal } from 'naive-ui';
import API from '../api/CommonData'
import store from '../store';

export default {
    components: {
        NCard,
        NRow,
        NCol,
        NButton,
        NModal,
    },
    data() {
        return {
            tableData: [],
            showModal: false,
            selectedOffer: {},
        }
    },
    methods:{
        handleAccept(row) {
            this.selectedOffer = row
            this.showModal = true
        },
        async submitAccept() {
            this.showModal = false
            const token = store.state.UserBasic.token
            let res = await API.postForm('/accept_offer/' + token, {
                uid: store.state.UserBasic.info.user_id,
                offer_id: this.selectedOffer.offer_id
            })

            if(res.status === 200) {
                $notification.success({
                    title: '接单成功',
                    content: res.msg,
                });
                this.getTableData()
            }
            else {
                $notification.error({
                    title: '接单失败',
                    content: res.msg,
                });
            }
        },
        onNegativeClick() {
            this.showModal = false
        },
        async getTableData() {
            let res = await API.getForm('/offer/offer_list');
            if(res.status === 200) {
                const locations = store.state.LocationList.locations
                this.tableData = []
                for(let item of res.data) {
                    item['from'] = locations[item.pickup_location_id]
                    item['to'] = locations[item.delivery_location_id]
                    this.tableData.push(item)
                }
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