export default {
    state: {
        token: '',
        info: '',
        uid: '',
    },
    mutations:{
       setToken(state, val) {
        state.token = val;
       },
       setUserInfo(state, val) {
        console.log('setUserInfo: ', val);
        state.info = val
       },

       setUid(state, val) {
        state.uid = val
       }
    }
}