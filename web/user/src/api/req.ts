import axios from "axios"
const req = axios.create({
    baseURL:"http://127.0.0.1:8080/v1",
    timeout:5000,
})
// 请求拦截器
req.interceptors.request.use(function (config) {
    // 获取本地token
    let token = localStorage.getItem('token')
    if(token){
        config.headers.Authorization = 'Bearer ' + token
    }
    return config
})

// 响应拦截器
req.interceptors.response.use(function (config) {
    // 判断响应头中是否存在token
    let token = config.headers.getAuthorization.toString()
    // 将token存储到本地
    if(token){
        localStorage.setItem('token',token)
    }
    return config
})
export default req