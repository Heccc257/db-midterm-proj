<template>
    <div>
        <n-modal v-model:show="formVisible" preset="dialog"
            :style="bodyStyle"
            title="新增订单"
        >
            <div class="my-8"></div>
            <n-form
                :label-width="150"
                :model="formValue"
                size="large"
                label-placement="left"
                require-mark-placement="right-hanging"
            >
    
                <n-form-item label="任务类别" path="category_name">
                    <n-select v-model:value="formValue.category_name" :options="category" placeholder="选择类别"/>
                </n-form-item>
                <n-form-item label="起始地址" path="pickup_location_id">
                    <n-select v-model:value="formValue.pickup_location_id" :options="locations" placeholder="选择位置"/>
                </n-form-item>
                <n-form-item label="目的地址" path="delivery_location_id">
                    <n-select v-model:value="formValue.delivery_location_id" :options="locations" placeholder="选择位置"/>
                </n-form-item>
                <n-form-item label="时间限制(min)" path="time_limit">
                    <n-input-number v-model:value="formValue.time_limit" clearable />
                </n-form-item>
                <n-form-item label="报酬(元)" path="reward_amount">
                    <n-input-number v-model:value="formValue.reward_amount" clearable />
                </n-form-item>
            </n-form>
            <div class="flex flex-row-reverse ">
                <div>
                    <n-button strong secondary type="success" @click="handleSubmit(1)">
                        确认
                    </n-button>
                </div>
                <div class="mx-4">
                    <n-button strong secondary type="tertiary" @click="handleSubmit(0)">
                        取消
                    </n-button>
                </div>
            </div>
        </n-modal>
    </div>
</template>

<script>
import store from '../store/index'
import { NModal, NForm, NFormItem, NSelect, NCard, NInputNumber, NButton} from 'naive-ui';
import Config from '../config/index'

export default {
    components: {
        NModal,
        NForm,
        NFormItem,
        NSelect,
        NCard,
        NInputNumber,
        NButton
    },
    data() {
        return {
            formValue: {
                category_name: null,
                customer_id: null,
                pickup_location_id: null,
                delivery_location_id: null,
                time_limit: null,
                reward_amount: null,
            },
            bodyStyle: {
                width: '600px'
            },
            category: [],
            locations: [],
            formVisible: false,
        }
    },
    mounted(){
        this.category = Config.category
        this.locations = []
        for(let [index, item] of Object.entries(store.state.LocationList.locations)) {
            this.locations.push({
                label: item,
                value: parseInt(index)
            })
        }
    },
    methods: {
        setVisible(val) {
            console.log('set show');
            this.formValue = {
                category_name: null,
                customer_id: null,
                pickup_location_id: null,
                delivery_location_id: null,
                time_limit: null,
                reward_amount: null,
            },
            this.formVisible = val
        },
        handleSubmit(type) {
            this.formVisible = false
            if(type === 0) return;
            this.$emit('submitOffer', this.formValue)
        }
    }, 
};
</script>

