<template>
    <div>
        <n-card title="我的接单" style="margin-bottom: 16px">
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
                                    <div>
                                        <span class="text-[1.2rem] text-slate-400">时限: </span>
                                        <span class="text-[1.1rem] text-slate-700"> {{ item.time_limit }} </span>
                                    </div>
                                </div>
                            </n-col>
                            <n-col :span="10">
                                <div>
                                    <div>
                                        <span class="text-[1.2rem] text-slate-400">下单人id: </span>
                                        <span class="text-[1.1rem] text-slate-700"> {{ item.customer_id }} </span>
                                    </div>
                                </div>
                            </n-col>
                            <n-col :span="2">
                                <div class="h-full w-full"
                                    v-if="tb.name === 'in_progress'"
                                >
                                    <n-button strong secondary type="success" @click="completeOffer(item.offer_id)"> 完成 </n-button>
                                    
                                </div>
                            </n-col>
                        </n-row>
                    </n-card>
                </div>
              </n-tab-pane>
            </n-tabs>
        </n-card>
    </div>
    <n-modal
        v-model:show="showModal"
        preset="dialog"
        title="确认"
        content="确认完成?"
        positive-text="确认"
        negative-text="别急"
        @positive-click="submitComplete"
        @negative-click="onNegativeClick"
    />
</template>

<script>
import { NCard, NTabs, NTabPane, NButton, NRow, NCol, NModal } from 'naive-ui'
import store from '../store/index'
import API from '../api/CommonData'

export default {
    components: {
        NCard, 
        NTabs,
        NTabPane,
        NButton,
        NRow,
        NCol,
        NModal
    },
    data() {
        return {
            tabs: [
                {name: 'in_progress', tab: '进行中'},
                {name: 'completed', tab: '已完成'},
            ],
            tableData: [],
            selectOfferId: null,
            showModal: false,
        }
    },
    methods: {
        async getTableData() {
            let res = await API.getForm('accept_offer_list/' + store.state.UserBasic.uid)
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
        completeOffer(id) {
            this.selectOfferId = id
            this.showModal = true
        },
        onNegativeClick() {
            this.showModal = false
        },
        async submitComplete() {
            let query = {
                uid: store.state.UserBasic.uid,
                offer_id: this.selectOfferId,
            }
            let res = await API.putForm('/complete_offer/' + store.state.UserBasic.token, query)

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
    },
    mounted() {
        this.getTableData()
        
    }
}
</script>