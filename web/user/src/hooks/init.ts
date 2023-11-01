import { getUnreadMsg } from '@/api'
import socket from '@/api/socket'
import { openDb, insertData,getSession, isExistNotify} from '@/db'
import { userStore} from '@/store'
import { useSessionStore } from '@/store'
import { messageInt} from '@/type'
import { log } from 'console'
const userstore = userStore()
const sessionStore = useSessionStore()
export async function init(){
  // 建立socket链接
  socket.init()
  // 打开数据库连接
  const db = await openDb((userstore.getUser()?.ID) + 'chat')
  userstore.db = db
  // 获取用户会话列表
  const {count:userCount,map:userMap} = await getSession(db,"usersession")
  // 获取群组会话列表
  const {count:groupCount,map:groupMap} = await getSession(db,"groupsession")
  userstore.unreadMessage += userCount + groupCount
  // 从redis获取所有未读消息
  const {data} = await getUnreadMsg()
  const unReadMsgList = data.data
  if(!unReadMsgList){
    return
  }
  // 遍历所有未读消息，将它们添加到对应的数据库中
  for(let i  of unReadMsgList){
        let value:messageInt = JSON.parse(i as string)
        if(value.type == 1){
          userstore.unreadMessage ++
          /* 私聊 */
          let userSession = userMap.get(value.from_id)
          if(userSession){
            // 当前会话存在(修改会话状态)
            userSession.lastMsg = value.content
            userSession.lastMsgTime = value.send_time
            userSession.unReadCount?userSession.unReadCount++:userSession.unReadCount = 1
          }else{
            userSession = {ID:value.from_id,Name:value.send_name,Avatar:value.avatar,
            lastMsg:value.content,lastMsgTime:value.send_time,type:1}
          }
          userMap.set(value.from_id,userSession)
          // 存进数据库
          insertData(db,"unread",value)
        }else if (value.type == 2){
          userstore.unreadMessage++
          /* 群聊 */
          let groupSession = userMap.get(value.group_id as number)
          if(groupSession){
            // 当前会话存在(修改会话状态)
            groupSession.lastMsg = value.content
            groupSession.lastMsgTime = value.send_time
            groupSession.unReadCount?groupSession.unReadCount++:groupSession.unReadCount = 1
          }else{
            groupSession = {ID:value.group_id as number,Name:value.group_name as string,Avatar:value.group_avatar as string,
            lastMsg:value.content,lastMsgTime:value.send_time,type:2}
          }
          groupMap.set(value.from_id,groupSession)
          // 存进数据库
          insertData(db,"unread",value)
        }else {
          /* 通知 */
          userstore.unreadNotice = true
          insertData(db,"unread",value) 
        }
  }
  if(!userstore.unreadNotice){
    userstore.unreadNotice = await isExistNotify(db)
  }
  // 将所有会话统计到会话表中
  sessionStore.sessionList.push(...userMap.values())
  sessionStore.sessionList.push(...groupMap.values())
  // 按时间排序
  sessionStore.sessionList.sort((i,j)=>(j.lastMsgTime as number) - (i.lastMsgTime as number))

}