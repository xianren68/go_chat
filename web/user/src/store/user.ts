import { defineStore } from "pinia"
import {userInt} from '@/type'
import {ref} from 'vue'
const user = defineStore('user',()=>{
    const unreadNotice = ref(false)
    const unreadMessage = ref(0)
    const db = ref<IDBDatabase>()
    let userInfo = ref<userInt|null>()
    // 登录时
    const update = (data:userInt|null)=>{
        if(data == null){
            data = userInfo.value!
        }else{
            userInfo.value = data
        }
        localStorage.setItem('userInfo',JSON.stringify(data))
    }
    // 获取
    const getUser = ()=>{
        userInfo.value = JSON.parse(localStorage.getItem('userInfo') || 'null')
    }
    return {update,getUser,unreadMessage,unreadNotice,db,userInfo}
})
export default user