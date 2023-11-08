import { insertData, putSession, saveMessage } from '@/db'
import { useSessionStore, userStore, useMessageStore,useNoticeStore } from '@/store'
import { messageInt} from '@/type'
import { toRaw } from 'vue'
import router from '@/router'
const socket = {
    s: <null | WebSocket>null,
    // 连接socket
    init() {
        if (this.s == null) {
            // 获取token
            let token = localStorage.getItem('token')
            if (!token) {
                router.push('/login')
                return
            }
            this.s = new WebSocket('ws://localhost:9000/v1/chat', [token])
            this.s.onerror = (e) => {
                console.log(e)
            }
        }
    },
    // 接收消息
    receive() {
        const messageStore = useMessageStore()
        const sessionStore = useSessionStore()
        const noticeStore = useNoticeStore();
        this.s!.onmessage = (event: any) => {
            const userstore = userStore()
            const msg: string = event.data
            const val: messageInt = JSON.parse(msg)
            // 通知
            if (val.type > 2) {
                userstore.unreadNotice = true
                insertData(userstore.db!, 'unRead', val)
                if(router.currentRoute.value.name == 'contact'){
                    noticeStore.noticeList.push(val)
                }
                return
            }
            let index: number
            // 正好处于当前会话(加入消息列表，更新会话即可)
            if (router.currentRoute.value.name == 'chat' &&
                ((val.type == 1 && val.from_id == sessionStore.sessionList[0].ID) ||
                    (val.type == 2 && val.group_id == sessionStore.sessionList[0].ID))) {
                // 加入聊天列表
                messageStore.messageList.push(val)
                // 存储进数据库
                saveMessage(userstore.db as IDBDatabase, val)
                // 更新会话
                sessionStore.sessionList[0].lastMsg = val.content
                sessionStore.sessionList[0].lastMsgTime = val.send_time
                index = 0
            } else {
                userstore.unreadMessage++
                insertData(userstore.db!,'unRead',val)
                let session: any
                if (val.type == 1) {
                    // 寻找对应会话
                    index = sessionStore.sessionList.findIndex((item) => item.ID == val.from_id)
                    if (index == -1) {
                        session = { ID: val.from_id, Name: val.send_name, Avatar: val.avatar, lastMsg: val.content, lastMsgTime: val.send_time, type: 1, unReadCount: 1 }
                    }
                } else {
                    // 寻找对应会话
                    index = sessionStore.sessionList.findIndex((item) => item.ID == val.group_id)
                    if (index == -1) {
                        session = { ID: val.group_id!, Name: val.group_name!, Avatar: val.group_avatar!, lastMsg: val.content, lastMsgTime: val.send_time, type: 2, unReadCount: 1 }
                    }
                }
                // 更新会话
                if (index == -1) {
                    sessionStore.sessionList.push(session)
                    index = sessionStore.sessionList.length - 1
                } else {
                    sessionStore.sessionList[index].lastMsg = val.content
                    sessionStore.sessionList[index].lastMsgTime = val.send_time
                    sessionStore.sessionList[index].unReadCount++
                }
            }
            // 防止修改头像和名称
            if(val.type == 1){
                sessionStore.sessionList[index].Name = val.send_name
                sessionStore.sessionList[index].Avatar = val.avatar
            }else{
                sessionStore.sessionList[index].Name = val.group_name!
                sessionStore.sessionList[index].Avatar = val.group_avatar!
            }

            // 存储会话
            putSession(userstore.db as IDBDatabase, val.type == 1 ? 'usersession' : 'groupsession', toRaw(sessionStore.sessionList[index]))

        }
    }
}
export default socket
