import http from './Myaxios'
import API_config from './config'

async function register(requestData) {
    let res = await http.request({
        url: API_config.BaseURL + '/register',
        method: 'post',
        data: requestData,
    })
    // console.log('register res: ', res);
    return res.data;
}

async function login(requestData) {
    let res = await http.request({
        url: API_config.BaseURL + '/login',
        method: 'post',
        data: requestData,
    })
    // console.log('login res: ', res);
    return res.data;
}

export default {
    register,
    login,
}