import vuex from 'vuex'
import UserBasic from './UserBasic'
import LocationList from './LocationList'

export default new vuex.Store({
    modules: {
        UserBasic,
        LocationList,
    },
})