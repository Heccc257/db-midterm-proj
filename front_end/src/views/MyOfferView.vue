<template>
    <div>
        <n-card title="我的发单" style="margin-bottom: 16px">
            <template #header-extra>
            <div class="mr-5 mt-2">
                <n-button strong secondary type="primary" @click="addOffer">
                    <div class="text-lg font-semibold">新增订单</div>
                </n-button>
            </div>
            </template>
            <n-tabs type="line" animated>
              <n-tab-pane v-for="tb of tabs" :key="tb.name"
              :name="tb.name" :tab="tb.tab">

                <div v-for="item in tableData" :key="item.offer_id">
                    <n-card title="" class="!rounded-md my-1" v-if="item.offer_state === tb.name">
                        <n-row>
                            <n-col :span="12">
                                <div class="mb-2 flex flex-col">
                                    <div>
                                        <span class="text-[1.2rem] text-slate-400">订单编号: </span>
                                        <span class="text-[1.1rem] text-slate-700"> {{ item.offer_id }} </span>
                                    </div>
                                    <div>
                                        <span class="text-[1.2rem] text-slate-400">任务: </span>
                                        <span class="text-[1.1rem] text-slate-700"> {{ item.category_name }} </span>
                                    </div>
                                    <div>
                                        <span class="text-[1.2rem] text-slate-400">地点: </span>
                                        <span class="text-[1.1rem] text-slate-700"> {{ item.from + ' --> ' + item.to }} </span>
                                    </div>
                                    <div v-if="tb.name === 'completed'">
                                        <span class="text-[1.2rem] text-slate-400">完成时间: </span>
                                        <span class="text-[1.1rem] text-slate-700"> {{ item.complete_time }} </span>
                                    </div>
                                    <div v-else>
                                        <span class="text-[1.2rem] text-slate-400">时限: </span>
                                        <span class="text-[1.1rem] text-slate-700"> {{ item.time_limit }} </span>
                                    </div>
                                </div>
                            </n-col>
                            <n-col :span="10">
                                <div v-if="tb.name === 'pending'">
                                    <h1 class="text-[1.6rem]"> 奖金: </h1>
                                    <h1 class="text-[1.6rem] text-center">{{ item.reward_amount + '元' }} </h1>
                                </div>
                                <div v-else>
                                    <div>
                                        <span class="text-[1.2rem] text-slate-400">接单人id: </span>
                                        <span class="text-[1.1rem] text-slate-700"> {{ item.worker_id }} </span>
                                    </div>
                                    <!-- <div>
                                        <span class="text-[1.2rem] text-slate-400">接单人电话: </span>
                                        <span class="text-[1.1rem] text-slate-700"> {{  }} </span>
                                    </div> -->
                                </div>
                            </n-col>
                            <n-col :span="2">
                                <div class="h-full w-full"
                                    v-if="tb.name === 'completed'"
                                >
                                    <n-button strong secondary type="info" @click="addRating(item.offer_id)"> 评价 </n-button>
                                    
                                </div>
                            </n-col>
                        </n-row>
                    </n-card>
                </div>
              </n-tab-pane>
            </n-tabs>
        </n-card>
    </div>
    <OfferForm ref="offerForm" @submitOffer="submitOffer"></OfferForm>
    <RatingForm ref="ratingForm" @submitRating="submitRating"></RatingForm>
</template>

<script>
import { NCard, NTabs, NTabPane, NButton, NRow, NCol } from 'naive-ui'
import store from '../store/index'
import API from '../api/CommonData'
import OfferForm from '../components/OfferForm.vue'
import RatingForm from '../components/RatingForm.vue'

export default {
    components: {
        NCard, 
        NTabs,
        NTabPane,
        NButton,
        NRow,
        NCol,
        OfferForm,
        RatingForm

    },
    data() {
        return {
            tabs: [
                {name: 'pending', tab: '待接单'},
                {name: 'in_progress', tab: '进行中'},
                {name: 'completed', tab: '已完成'},
            ],
            tableData: [],
            ratingOfferId: null,
        }
    },
    methods: {
        async getTableData() {
            let res = await API.getForm('/offer/offer_list/' + store.state.UserBasic.uid)
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
                    title: '数据获取失败',
                    content: res.msg,
                });
            }
        },
        addOffer() {
            this.$refs.offerForm.setVisible(true)
        },
        async submitOffer(form) {
            form.customer_id = store.state.UserBasic.uid
            console.log('add Offer: ', form, store.state.UserBasic.token);
            let res = await API.postForm('/offer/post/' + store.state.UserBasic.token, form)

            if(res.status === 200) {
                $notification.success({
                    title: '订单提交成功',
                    content: '',
                });
                this.getTableData()
            }
            else {
                $notification.error({
                    title: '订单提交失败',
                    content: res.msg,
                });
            }
        },
        addRating(id) {
            this.ratingOfferId = id
            this.$refs.ratingForm.setVisible(true)
        },
        async submitRating(form) {
            form.offer_id = this.ratingOfferId
            console.log('add Rating: ', form, store.state.UserBasic.token);
            let res = await API.postForm('/user_rating/' + store.state.UserBasic.uid + '/' + store.state.UserBasic.token, form)

            if(res.status === 200) {
                $notification.success({
                    title: '评价提交成功',
                    content: '',
                });
                this.getTableData()
            }
            else {
                $notification.error({
                    title: '订单提交失败',
                    content: res.msg,
                });
            }
        }
    },
    mounted() {
        this.getTableData()
        
    }
}
</script>