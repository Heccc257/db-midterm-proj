<template>
    <div class="mx-auto flex min-h-screen flex-col">
        <h1 class="mx-auto mt-[15vh] mb-5 text-center text-2xl font-bold">登录页</h1>
        <div class="container mx-auto">
            <div
                class="mx-auto flex h-[400px] w-1/2 items-center justify-center rounded-md border-[1px] border-slate-200 bg-white px-40 transition-all duration-100 ease-in-out"
            >
                <div class="my-6 max-h-[256px] w-full">
                    <n-tabs size="large" animated :default-value="id_cases[0].name">
                        <n-tab-pane v-for="login in id_cases" :name="login.name" :tab="login.name" :key = login.name>
                            <div v-if="login.name === '登录'">
                                <div class="mb-3 flex h-9 items-center justify-center">
                                    <h2 class="w-15 mr-5 flex-shrink-0 font-bold text-neutral-400">电话号码</h2>
                                    <n-input v-model:value="phoneNumber" type="text" placeholder="" class="w-5"></n-input>
                                </div>
                                <div class="flex h-9 items-center justify-center">
                                    <h2 class="w-15 mr-5 flex-shrink-0 font-bold text-neutral-400">密码</h2>
                                    <n-input v-model:value="password" type="password" placeholder="" class="w-5"></n-input>
                                </div>
                            </div>
                            <div v-else>
                                <div class="mb-3 flex h-9 items-center justify-center">
                                    <h2 class="w-15 mr-5 flex-shrink-0 font-bold text-neutral-400">姓名</h2>
                                    <n-input v-model:value="username" type="text" placeholder="" class="w-5"></n-input>
                                </div>
                                <div class="mb-3 flex h-9 items-center justify-center">
                                    <h2 class="w-15 mr-5 flex-shrink-0 font-bold text-neutral-400">昵称</h2>
                                    <n-input v-model:value="nickname" type="text" placeholder="" class="w-5"></n-input>
                                </div>
                                <div class="mb-3 flex h-9 items-center justify-center">
                                    <h2 class="w-15 mr-5 flex-shrink-0 font-bold text-neutral-400">电话号码</h2>
                                    <n-input v-model:value="phoneNumber" type="text" placeholder="" class="w-5"></n-input>
                                </div>
                                <div class="flex h-9 items-center justify-center">
                                    <h2 class="w-15 mr-5 flex-shrink-0 font-bold text-neutral-400">密码</h2>
                                    <n-input v-model:value="password" type="text" placeholder="" class="w-5"></n-input>
                                </div>
                            </div>
                            <div class="mt-4 flex items-center justify-center">
                                <n-button
                                    strong
                                    secondary
                                    type="success"
                                    @click="executeLogin(login.name)"
                                    class="!rounded-md !px-5"
                                    >
                                        {{login.name}}
                                    </n-button
                                >
                            </div>
                        </n-tab-pane>
                    </n-tabs>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import { NTabs, NTabPane, NInput, NButton } from 'naive-ui';
import { NNotificationProvider, useNotification } from 'naive-ui';
import store from '../store/index'
import API from '../api/CommonData'

import loginAPI from '../api/login';
export default {
    components: {
        NTabs,
        NTabPane,
        NInput,
        NButton,
    },
    data() {
        return {
            id_cases: [
                {
                    name: '登录',
                },
                {
                    name: '注册',
                },
            ],
            username: '',
            nickname: '',
            password: '',
            phoneNumber: '',
        };
    },
    methods: {
        async executeLogin(type) {
            try {
                var res;
                if(type === '注册') {
                    var formdata = new FormData();
                    formdata.append("nick_name", this.nickname);
                    formdata.append("full_name", this.username);
                    formdata.append("phone_number", this.phoneNumber);
                    formdata.append("password_hash", this.password);
                    res = await loginAPI.register(formdata);
                }
                else {
                    var formdata = new FormData();
                    formdata.append("phone_number", this.phoneNumber);
                    formdata.append("password", this.password);
                    res = await loginAPI.login(formdata)
                }
                
                if (res.status === 200) {
                    // convert role from camel case to snake case
                    $notification.success({
                        title: type + '成功',
                        content: '',
                    });
                    if(type === '登录') {
                        store.commit('setToken', res.data.token)
                        store.commit('setUid', res.data.uid)
                        
                        res = await API.getForm('/location_list')
                        // get location list
                        if(res.status === 200) {
                            let locations = {}
                            for(let item of res.data) {
                                locations[item.location_id] = item.building + item.floor
                            }

                            store.commit('setLocations', locations)
                        }
                        else {
                            $notification.error({
                                title: 'location_list请求失败',
                                message: res.message,
                            });
                        }

                        if(res.status == 200) this.$router.push({name: 'UserHome'});
                    }
                }
                else {
                    $notification.error({
                        title: type + '失败',
                        message: res.message,
                    });
                }
            } catch (err) {
                console.log(err);
            }
        },
    },
    setup() {
        window.$notification = useNotification();
    },
};
</script>
