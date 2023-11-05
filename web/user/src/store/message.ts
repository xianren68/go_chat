import { getGroupMessage, getGroupUnRead, getUserMessage, getUserUnRead,saveMessage } from '@/db'
import { messageInt } from '@/type'
import {defineStore} from 'pinia'
import { reactive } from 'vue'
import { userStore} from '.'
const messageStore = defineStore('message',()=>{
   // 消息列表
   const messageList:Array<messageInt> = reactive([])
   // 替换消息列表
   const replaceList = async (id:number,unReadCount:number,type:number)=>{
        const db = userStore().db!
        messageList.length = 0
        let unreadList:Array<messageInt> = []
        if(type == 1){
            // 取出已读消息
            const message = await getUserMessage(db,id)
            messageList.push(...(message??[]))
            if(unReadCount > 0){
                // 取出未读消息
                unreadList = await getUserUnRead(db,id)??[]
            }
        }else {
            // 群消息
            // 取出已读消息
            const message = await getGroupMessage(db,id)
            messageList.push(...(message??[]))
            // 取出未读消息
            if(unReadCount > 0){
                unreadList = await getGroupUnRead(db,id)??[]
            }
        }
        // 将未读消息加入数据库和消息列表
        for(let i of unreadList){
            saveMessage(db,i)
        }
        messageList.push(...unreadList)
        
   }
   return {messageList,replaceList}
    
})
export default messageStore