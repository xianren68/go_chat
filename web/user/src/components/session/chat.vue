<template>
    <div class="chat">
        <div class="top">{{ sessionStore.sessionList[0].Name }}</div>
        <div class="main" ref="main">
            <div v-for="item in messageStore.userMessage.get(sessionStore.sessionList[0].ID)">
                <meMsg :data="item" v-if="item.from_id === userstore.getUser()?.ID"></meMsg>
                <otherMsg :data="item" v-else></otherMsg>
            </div>
        </div>
        <!--发送消息-->
        <div class="send">
            <div class="send-main">
                <div class="input">
                    <svg class="icon">
                        <use xlink:href="#icon-huatong"></use>
                    </svg>
                    <input v-model="message">
                    <svg class="icon">
                        <use xlink:href="#icon-biaoqing"></use>
                    </svg>
                    <svg class="icon">
                        <use xlink:href="#icon-attachment"></use>
                    </svg>
                </div>
                <div class="send-icon" @click="sendMsg">
                    <svg class="icon">
                        <use xlink:href="#icon-send"></use>
                    </svg>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, toRaw, onBeforeMount,onMounted } from "vue"
import socket from "../../api/socket"
import meMsg from "./meMsg.vue"
import otherMsg from "./otherMsg.vue"
import { useSessionStore, useMessageStore } from "@/store"
import { putSession, saveMessage } from "@/db"
import { userStore } from "@/store"
import { messageInt } from "@/type"
// 获取会话信息
const sessionStore = useSessionStore()
const userstore = userStore()
const messageStore = useMessageStore()
// 消息数据
const message = ref("")
const main = ref()

// 发送消息
const sendMsg = () => {
    const userInfo = JSON.parse(localStorage.getItem("userInfo") as string)
    // 当前时间
    const nowData = new Date().getTime()
    // 消息
    const send_msg: messageInt = { from_id: userInfo.ID, target_id: sessionStore.sessionList[0].ID, type: 1, content: message.value, send_time: nowData, avatar: userstore.getUser()?.Avatar as string, send_name: userstore.getUser()?.Name as string }
    messageStore.userMessage.get(sessionStore.sessionList[0].ID)?.push(send_msg)
    message.value = ""
    socket.s?.send(JSON.stringify(send_msg))
    sessionStore.sessionList[0].lastMsg = message.value
    sessionStore.sessionList[0].lastMsgTime = nowData
    // 更新会话
    putSession(userstore.db as IDBDatabase, 'usersession', toRaw(sessionStore.sessionList[0]))
    // 存储聊天
    saveMessage(userstore.db as IDBDatabase, send_msg)
}
onBeforeMount(async () => {
    // 获取对应的聊天信息
    const { ID, type, unReadCount } = sessionStore.sessionList[0]
    if (type == 1) {
        // 获取聊天信息
        await messageStore.getuserMessage(ID, unReadCount as number)
    }
    // 减去未读消息
    if (unReadCount == 0){
        return
    }
    // 减去小红点的次数
    userstore.unreadMessage -= unReadCount
    sessionStore.sessionList[0].unReadCount = 0
    // 更新会话
    if(type==1){
        putSession(userstore.db as IDBDatabase,'usersession',toRaw(sessionStore.sessionList[0]))
    }
})
onMounted(() => {
    const observer = new MutationObserver((mutationList) => {
    mutationList.forEach(mutation => {
        let d:any = mutation.addedNodes.item(mutation.addedNodes.length-1)
        d.scrollIntoView()
    });
});
observer.observe(main.value, { childList: true});
})
</script>

<style scoped lang="scss">
.chat {
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;

    .top {
        height: 49px;
        display: flex;
        border-bottom: #ccc 1px solid;
        justify-content: center;
        align-items: center;
        font-family: "Fira Code";
    }

    .main {
        height: 80%;
        padding: 0 10px;
        padding-top: 20px;
        overflow: scroll;
    }

    .main::-webkit-scrollbar {
        display: none;
    }

    .send {
        flex-grow: 1;

        .send-main {
            display: flex;
            align-items: center;

            .input {
                margin-left: 20px;
                height: 40px;
                width: 655px;
                border: 1px solid #ccc;
                border-radius: 13px;
                display: flex;
                align-items: center;

                .icon {
                    height: 20px;
                    width: 20px;
                    margin-left: 8px;
                    fill: #999;
                }

                input {
                    width: 550px;
                    height: 90%;
                    font-size: 16px;
                    border: none;
                    outline: none;
                    margin-left: 10px;
                }
            }

            .send-icon {
                height: 30px;
                width: 30px;
                margin-left: 20px;
                background-color: #ecedf6;
                display: flex;
                align-items: center;
                border-radius: 50%;

                .icon {
                    height: 20px;
                    width: 20px;
                    margin-left: 4px;
                    fill: #4a499b;
                }
            }

        }

    }
}
</style>