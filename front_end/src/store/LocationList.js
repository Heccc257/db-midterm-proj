export default {
    state: {
        locations: '',
    },
    mutations:{
       setLocations(state, val) {
            console.log('set locations', val);
            state.locations = val;
       }
    }
}