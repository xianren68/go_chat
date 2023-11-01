import { defineStore } from "pinia"
import {userInt} from '@/type'
import {ref} from 'vue'
const user = defineStore('user',()=>{
    const unreadNotice = ref(false)
    const unreadMessage = ref(0)
    // db连接
    const db = ref<IDBDatabase>()
    let userInfo:userInt|undefined
    // 登录时
    const update = (data:userInt)=>{
        userInfo = data
        localStorage.setItem('userInfo',JSON.stringify(data))
    }
    // 获取
    const getUser = ()=>{
        if(userInfo){
            return userInfo
        }else{
            userInfo = JSON.parse(localStorage.getItem('userInfo') || '')
            return userInfo
        }
    }
    return {update,getUser,unreadMessage,unreadNotice,db}
})
export default user