<template>
    <div class="chat">
        <div class="top">{{ sessionStore.sessionList[0].Name }}</div>
        <div class="main" ref="main">
            <div v-for="item in messageStore.messageList">
                <meMsg :data="item" v-if="item.from_id === userstore.userInfo!.ID"></meMsg>
                <otherMsg :data="item" v-else></otherMsg>
            </div>
        </div>
        <!--发送消息-->
        <div class="send">
            <div class="send-main">
                <div class="input" v-if="!md">
                    <svg class="icon" @click="()=>{md = true;message = ''}">
                        <use xlink:href="#icon-markdown"></use>
                    </svg>
                    <input v-model="message">
                    <svg class="icon" @click="emojiShow = !emojiShow" :class="{select_icon:emojiShow}">
                        <use xlink:href="#icon-biaoqing"></use>
                    </svg>
                    <svg class="icon">
                        <use xlink:href="#icon-attachment"></use>
                    </svg>
                </div>
                <MdEditor v-else v-model="message"></MdEditor>
                <div class="send-icon" @click="sendMsg">
                    <svg class="icon">
                        <use xlink:href="#icon-send"></use>
                    </svg>
                </div>
            </div>
        </div>
    </div>
    <!-- emoji组件 -->
    <Emoji v-if="emojiShow" @insertEmoji="insertEmoji"></Emoji>
    <!-- 控制回到普通聊天框 -->
    <div v-if="md" class="normal" @click="()=>{md = false;message=''}">
        <svg class="icon">
            <use xlink:href="#icon-chat"></use>
        </svg>
    </div>
</template>

<script setup lang="ts">
import { ref, toRaw, onMounted,watch} from "vue"
import socket from "@/api/socket"
import meMsg from "@/components/session/meMsg.vue"
import otherMsg from "@/components/session/otherMsg.vue"
import { useSessionStore, useMessageStore } from "@/store"
import { putSession, saveMessage } from "@/db"
import { userStore } from "@/store"
import { messageInt } from "@/type"
import Emoji from "@/components/emoji.vue"
import  {MdEditor}  from "md-editor-v3"
import 'md-editor-v3/lib/style.css'
// 获取会话信息
const sessionStore = useSessionStore()
const userstore = userStore()
const messageStore = useMessageStore()
// 控制表情是否展开
const emojiShow = ref(false)
// 控制markdown编辑器是否出现
const md = ref(false)
// 插入表情
const insertEmoji = (arg:string)=>{
    message.value += arg
}
watch(messageStore.messageList,()=>{
    md.value =false
    emojiShow.value = false
    message.value = ""
})
// 消息数据
const message = ref("")
const main = ref()
// 发送消息
const sendMsg = () => {
    const userInfo = JSON.parse(localStorage.getItem("userInfo") as string)
    // 当前时间
    const nowData = new Date().getTime()
    // 消息
    let send_msg: messageInt
    if(sessionStore.sessionList[0].type  == 1){
        // 私聊
         send_msg = { from_id: userInfo.ID, target_id: sessionStore.sessionList[0].ID, type: 1, content: message.value, send_time: nowData, avatar: userstore.userInfo!.Avatar, send_name: userstore.userInfo!.Name,md:md.value }
    }else {
        // 群聊
        send_msg = { from_id: userInfo.ID, target_id: 0, type: 2, content: message.value, send_time: nowData, avatar: userstore.userInfo!.Avatar, send_name: userstore.userInfo!.Name,group_id:sessionStore.sessionList[0].ID,group_avatar:sessionStore.sessionList[0].Avatar,group_name:sessionStore.sessionList[0].Name,md:md.value}
    }
    
    messageStore.messageList.push(send_msg)
    message.value = ""
    socket.s?.send(JSON.stringify(send_msg))
    sessionStore.sessionList[0].lastMsg = send_msg.content
    sessionStore.sessionList[0].lastMsgTime = nowData
    // 更新会话
    putSession(userstore.db!, 'usersession', toRaw(sessionStore.sessionList[0]))
    // 存储聊天
    saveMessage(userstore.db!, send_msg)
}

onMounted(() => {
    const observer = new MutationObserver((mutationList) => {
        // 监听子元素变化，以便让聊天随时处于可视区域内
        mutationList.forEach(mutation => {
            let d: any = mutation.addedNodes.item(mutation.addedNodes.length - 1)
            d?.scrollIntoView()
        })
    })
    observer.observe(main.value, { childList: true });
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
        border-bottom: #ccc 1px solid;
        text-align: center;
        line-height: 49px;
        font-family: "Fira Code";
    }

    .main {
        padding: 0 10px;
        flex-grow: 1;
        padding-top: 20px;
        overflow: scroll;
    }

    .main::-webkit-scrollbar {
        display: none;
    }

    .send {
        .send-main {
            margin-top: 15px;
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
                .select_icon{
                    fill:var(--icon-active-color);
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
            .md-editor{
                height: 250px;
                width:655px;
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
.normal {
    position: relative;
    top:-22px;
    left:60px;
    .icon {
        height: 20px;
        width:20px;
    }
}
</style>