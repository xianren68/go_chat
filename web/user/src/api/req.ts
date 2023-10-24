import axios from "axios"
import NProgressE from 'nprogresse'
import 'nprogresse/style.css'
const req = axios.create({
    baseURL:"http://127.0.0.1:8080/v1",
    timeout:5000,
})
// 开启进度条
NProgressE.configure({showSpinner:true})
// 请求拦截器
req.interceptors.request.use(function (config) {
    // 开启进度条
    NProgressE.start()
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
    console.log(config)
    let token = config.headers['authorization']?.split(" ")[1]
    // 将token存储到本地
    if(token){
        localStorage.setItem('token',token)
    }
    // 关闭进度条
    NProgressE.done(true)
    return config
})
export default req