<template>
    <div class="message">
        <!--聊天用户列表-->
        <div class="list">
            <div class="item" v-for="(msg,i) in msgList" :key="i" :class="{select:show && i== 0}">
                <div class="img">
                    <img
                        :src="msg.Avatar == ''?'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQFrZh6NXKZ7x0WW0UR2pPf2pXOrCaFcd62Uw&usqp=CAU':item.Avatar"
                        alt="">
                </div>
                <div class="desc">
                    <p class="name">{{ msg.Name }}</p>
                    <p class="line">{{ msg?.lastMsg }}</p>
                </div>
            </div>
        </div>
        <div class="msg" v-if="show">
            <chat  :userInfo="msgList[0]"></chat>
        </div>
    </div>
</template>

<script setup lang="ts">
import {useMsgStore} from "@/store"
import {onBeforeMount,ref} from "vue"
import {messageInt} from "@/type"
import {useRoute} from "vue-router"
import chat from "@/components/message/chat.vue"

const {msgList, updateMsgList} = useMsgStore()

const route = useRoute()
// 是否有选中的聊天
const show = ref(false)
// 获取路由参数
onBeforeMount(() => {
    const cur = route.query
    // 没有传参
    if (!cur) return
    show.value = true
    updateMsgList(cur)
})
</script>

<style lang="scss" scoped>
.message {
    height: 100%;
    flex-grow: 1;
    display: flex;

    .list {
        display: flex;
        height: 100%;
        width: 20%;
        overflow: scroll;
        flex-direction: column;
        align-items: center;

        .item {
        box-sizing: border-box;
            height: 50px;
            width: 100%;
            display: flex;
            padding: 0 3px;
            align-items: center;
            border-bottom:1px solid #eee ;
            background-color:#f9f9fb;

            .img {
                height: 35px;
                width: 35px;
                border-radius: 50%;
                overflow: hidden;

                img {
                    height: 100%;
                    width: 100%;
                }
            }

            .desc {
                flex-grow: 1;
                margin-left: 10px;

                .name {
                    font-size: 14px;
                    font-weight: bold;
                    margin-bottom: 3px;
                }

                .line {
                    font-size: 10px;
                    color: #aaa;
                }
            }
        }

        .item.select {
                background: linear-gradient(to right, #f9d9da, #fff,#fff);
                transition: .5s;
        }
    }

    .msg {
        height: 100%;
        flex-grow: 1;
        background-color: #fff;
    }

    .list::-webkit-scrollbar {
        display: none;
    }
    .msg {
        height:100%;
        flex-grow:1;
    }
}
</style>