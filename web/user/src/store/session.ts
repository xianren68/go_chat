import {defineStore} from 'pinia'
import {sessionInt} from '@/type'
import {reactive} from "vue"
const SessionStore = defineStore('message',()=>{
   // 会话列表
   const sessionList:Array<sessionInt> = reactive([])
   const updateSessionList = (session:sessionInt) => {
      // 判断与当前会话是否已经存在
      const index = sessionList.findIndex((item)=>item.ID === session.ID && item.type === session.type)
      // 若已存在将其提到数组顶部
      if(index !== -1){
         session = sessionList[index]
         sessionList.splice(index,1)
      }
      sessionList.unshift(session)
   }
   return {sessionList,updateSessionList}
})
export default SessionStore
