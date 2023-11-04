import {defineStore} from 'pinia'
import {sessionInt} from '@/type'
import {reactive} from "vue"
import { useMessageStore } from '.'
const SessionStore = defineStore('session',()=>{
   // 会话列表
   const sessionList:Array<sessionInt> = reactive([])
   // 更新会话（从其他页跳转）
   const updateSessionList = (session:sessionInt) => {
      const messageStore = useMessageStore()
      // 判断与当前会话是否已经存在
      const index = sessionList.findIndex((item)=>item.ID === session.ID && item.type === session.type)
      // 若已存在将其提到数组顶部
      if(index !== -1){
         session = sessionList[index]
         sessionList.splice(index,1)
         messageStore.replaceList(session.ID,session.unReadCount,session.type)
      }else{
         messageStore.messageList.length = 0
      }
      sessionList.unshift(session)
   }
   // 切换会话
   const switchSession = (id:number,type:number) => {
      const index = sessionList.findIndex((item)=> item.type == type && item.ID == id)
      // 切换到头
      const session = sessionList[index]
      sessionList.splice(index,1)
      sessionList.unshift(session)
   }
   return {sessionList,updateSessionList,switchSession}
})
export default SessionStore
