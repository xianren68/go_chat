import { getUnreadMsg } from '@/api'
import socket from '@/api/socket'
import { openDb, insertData, getSession, isExistNotify, putSession } from '@/db'
import { userStore } from '@/store'
import { useSessionStore } from '@/store'
import { messageInt } from '@/type'
const userstore = userStore()
const sessionStore = useSessionStore()
export async function init() {
  // 建立socket链接
  socket.init()
  // 打开数据库连接
  // if(userstore.db){
  //   return
  // }
  const db = await openDb((userstore.getUser()?.ID) + 'chat')
  userstore.db = db
  // 获取用户会话列表
  const { count: userCount, map: userMap } = await getSession(db, "usersession")
  // 获取群组会话列表
  const { count: groupCount, map: groupMap } = await getSession(db, "groupsession")
  userstore.unreadMessage += userCount + groupCount
  // 从redis获取所有未读消息
  const { data } = await getUnreadMsg()
  const unReadMsgList:Array<string>|null = data.data
  // 存在未读消息
  if (unReadMsgList) {
    // 遍历所有未读消息，将它们添加到对应的数据库中(倒序)
    for (let i = unReadMsgList.length-1;i >= 0;i--) {
      let value: messageInt = JSON.parse(unReadMsgList[i])
      if (value.type == 1) {
        userstore.unreadMessage++
        /* 私聊 */
        let userSession = userMap.get(value.from_id)
        if (userSession) {
          // 当前会话存在(修改会话状态)
          userSession.lastMsg = value.content
          userSession.lastMsgTime = value.send_time
          userSession.unReadCount ? userSession.unReadCount++ : userSession.unReadCount = 1
        } else {
          userSession = {
            ID: value.from_id, Name: value.send_name, Avatar: value.avatar,
            lastMsg: value.content, lastMsgTime: value.send_time, type: 1,unReadCount:1
          }
        }
         // 添加/更新会话到数据库
         putSession(db,'usersession',userSession)
        userMap.set(value.from_id, userSession)
        // 存进数据库
        insertData(db, "unRead", value)
      } else if (value.type == 2) {
        userstore.unreadMessage++
        /* 群聊 */
        let groupSession = userMap.get(value.group_id as number)
        if (groupSession) {
          // 当前会话存在(修改会话状态)
          groupSession.lastMsg = value.content
          groupSession.lastMsgTime = value.send_time
          groupSession.unReadCount ? groupSession.unReadCount++ : groupSession.unReadCount = 1
        } else {
          groupSession = {
            ID: value.group_id as number, Name: value.group_name as string, Avatar: value.group_avatar as string,
            lastMsg: value.content, lastMsgTime: value.send_time, type: 2,unReadCount:1
          }
        }
        // 添加/更新会话到数据库
        putSession(db,'groupsession',groupSession)
        groupMap.set(value.from_id, groupSession)
        // 存进数据库
        insertData(db, "unRead", value)
      } else {
        /* 通知 */
        userstore.unreadNotice = true
        insertData(db, "unRead", value)
      }
    }
  }
  if (!userstore.unreadNotice) {
    userstore.unreadNotice = await isExistNotify(db)
  }
  // 将所有会话统计到会话表中
  sessionStore.sessionList.push(...userMap.values())
  sessionStore.sessionList.push(...groupMap.values())
  // 按时间排序
  sessionStore.sessionList.sort((i, j) => (j.lastMsgTime as number) - (i.lastMsgTime as number))

}