import { getMessage, getunRead, saveMessage } from '@/db'
import { messageInt } from '@/type'
import {defineStore} from 'pinia'
import { reactive } from 'vue'
import { userStore} from '.'
const messageStore = defineStore('message',()=>{
    // 私聊消息
    const userMessage = reactive(new Map<number,Array<messageInt>>())
    // // 群聊消息
    // const groupMessage = reactive(new Map<number,Array<messageInt>>())
    // 获取私聊消息
    const getuserMessage = async (from_id:number,unread:number)=>{
        const userstore = userStore()
        // 判断是否已存在
        if(userMessage.get(from_id)){
            return
        }
        // 获取已读消息
       const data = await getMessage(userstore.db as IDBDatabase,from_id)
       const list = [...data]
        // 判断是否有未读消息
        if(unread > 0){
            // 读取未读消息
            const unreadList = await getunRead(userstore.db as IDBDatabase,from_id)
            list.push(...unreadList)
            // 全部存入已读消息
            for(let i of unreadList){
                saveMessage(userstore.db as IDBDatabase,i)
            }
        }
        console.log(list)
        userMessage.set(from_id,list)
    }
    return {userMessage,getuserMessage}
    
})
export default messageStore