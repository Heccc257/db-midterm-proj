import http from './Myaxios'
import API_config from './config'

async function postForm(url, requestData) {
    console.log('postForm 表单:', requestData);
    var formdata = new FormData()
    for(let item of Object.keys(requestData)) {
        formdata.append(item, requestData[item])
    }

    let res = await http.request({
        url: API_config.BaseURL + url,
        method: 'post',
        data: formdata,
    })
    console.log('postForm 返回: ', res);
    return res.data;
}

async function getForm(url) {
    // console.log('getForm 表单:', url);
    let res = await http.request({
        url: API_config.BaseURL + url,
        method: 'get',
    })
    console.log('getForm 返回: ', res);
    return res.data;
}

async function putForm(url, query) {
    // console.log('getForm 表单:', url);
    let res = await http.request({
        url: API_config.BaseURL + url,
        method: 'put',
        params: query,
    })
    console.log('putForm 返回: ', res);
    return res.data;
}

export default {
    getForm,
    postForm,
    putForm,
}