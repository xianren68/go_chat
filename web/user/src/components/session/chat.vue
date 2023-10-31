<template>
    <div class="chat">
        <div class="top">{{sessionInfo.Name}}</div>
        <div class="main">
            <div v-for="item in msgList">
                <meMsg :data="item"></meMsg>
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
import { ref,reactive } from "vue"
import socket from "../../api/socket"
import meMsg from "./meMsg.vue";
// 获取当前聊天人的信息
const props = defineProps(["sessionInfo"])
// 消息数据
const message = ref("")
// 消息列表
const msgList = reactive([])
// 发送消息
const sendMsg = () => {
    const userInfo = JSON.parse(localStorage.getItem("userInfo") as string)
    console.log(userInfo.id);
    
    socket.s?.send(JSON.stringify({from_id:userInfo.ID,target_id:props.sessionInfo.ID,type:1,content:message.value}))
    msgList.push({avatar:"",msg:message.value})
    message.value = ""
}
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
        padding:0 10px;
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
                    margin-left:4px ;
                    fill: #4a499b;
                }
            }

        }

    }
}
</style>