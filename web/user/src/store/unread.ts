import { defineStore } from "pinia"
import { reactive } from "vue"
import {messageInt} from '@/type'
const unread = defineStore('unread',()=>{
    // 未读消息列表
    const unreadmsgList:Array<messageInt> = reactive([])
    // 未读通知列表
    const unreadnoticeList:Array<messageInt> = reactive([])
    // 更新未读消息
    const updateunReadmsg = (data:Array<messageInt>)=>{
        unreadmsgList.push(...data)
    }
    // 更新未读通知
    const updateunReadnotice = (data:Array<messageInt>)=>{
        unreadnoticeList.push(...data)
    }
    return {
        unreadmsgList,
        unreadnoticeList,
        updateunReadmsg,
        updateunReadnotice}
    
})
export default unread