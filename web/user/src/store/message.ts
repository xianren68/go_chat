import {defineStore} from 'pinia'
import {messageInt} from '@/type'
import {reactive} from "vue"
const MsgStore = defineStore('message',()=>{
   const msgList:Array<messageInt> = reactive([])
   const updateMsgList = (msg:messageInt) => {
      // 判断与当前用户的聊天是否已经存在
      const index = msgList.findIndex((item)=>item.ID === msg.ID)
      // 若已存在将其提到数组顶部
      if(index !== -1){
         msg = msgList[index]
         msgList.splice(index,1)
      }
      msgList.unshift(msg)
   }
   return {msgList,updateMsgList}
})
export default MsgStore
