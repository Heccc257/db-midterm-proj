<script>
import HomeView from './views/HomeView.vue'
import { NConfigProvider, NNotificationProvider, useNotification } from 'naive-ui';
import store from './store';

export default {
  components: {
    HomeView,
    NConfigProvider,
    NNotificationProvider,
    useNotification
  },
  created(){
        if (sessionStorage.getItem("store")) {
        // console.log('页面重新加载');
            let storet = sessionStorage.getItem("store");
            store.replaceState(Object.assign({}, store.state, JSON.parse(storet == null ? '' : storet)))
        }
        //在页面刷新时将vuex里的信息保存到sessionStorage里
        window.addEventListener("beforeunload", () => {
            // console.log('页面被刷新');
            let state = JSON.stringify(store.state)
            sessionStorage.setItem("store", state == null ? '' : state)
        })
    }
}
</script>

<template>
  <n-notification-provider container-style="margin-top:3rem">
    <n-config-provider preflight-style-disabled> <router-view></router-view> </n-config-provider
  ></n-notification-provider>
</template>