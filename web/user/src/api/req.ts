import axios from "axios"
import NProgressE from 'nprogresse'
import 'nprogresse/style.css'
import router from '@/router'
import { ElMessage } from 'element-plus'
//const router = useRouter()
const req = axios.create({
    baseURL:"http://127.0.0.1:9000/v1",
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
    // 判断登录失效的情况
    let code = config.data.code
    // token错误(跳转到登录页)
    if(code >= 1005 && code <= 1008){
        ElMessage({message:'登录过期，即将返回登录页',type:'error'})
        router.push({name:'login'})
    }
    // 判断响应头中是否存在token
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