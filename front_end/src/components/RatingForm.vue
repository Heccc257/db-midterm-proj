<template>
    <div>
        <n-modal v-model:show="formVisible" preset="dialog"
            :style="bodyStyle"
            title="评价"
        >
            <div class="my-8"></div>
            <n-form
                :label-width="100"
                :model="formValue"
                size="large"
                label-placement="left"
                require-mark-placement="right-hanging"
            >
              <n-form-item label="评价内容" path="comment">
                <n-input v-model:value="formValue.comment" placeholder="输入评价内容" 
                type="textarea" :autosize="{
                    minRows: 3,
                    maxRows: 5
                  }" />
              </n-form-item>
              <n-form-item label="评价内容" path="comment">
                <n-rate allow-half v-model:value="formValue.rating"/>
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
import { NModal, NForm, NFormItem, NSelect, NCard, NInput, NButton, NRate} from 'naive-ui';
import Config from '../config/index'

export default {
    components: {
        NModal,
        NForm,
        NFormItem,
        NSelect,
        NCard,
        NInput,
        NButton,
        NRate
    },
    data() {
        return {
            formValue: {
                comment: null,
                rating: null,
            },
            bodyStyle: {
                width: '600px'
            },
            formVisible: false,
        }
    },
    mounted(){
    },
    methods: {
        setVisible(val) {
            console.log('set show');
            this.formValue = {
                comment: null,
                rating: null,
            },
            this.formVisible = val
        },
        handleSubmit(type) {
            this.formVisible = false
            if(type === 0) return;
            this.formValue.rating = this.formValue.rating * 2 
            this.$emit('submitRating', this.formValue)
        }
    }, 
};
</script>

