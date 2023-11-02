<template>
    <div class="message">
        <!--聊天用户列表-->
        <div class="list">
            <div class="item" v-for="(session, i) in sessionStore.sessionList" :key="i" :class="{ select: show && i == 0 }" @click="switchSession(session.ID,session.type)">
                <div class="img">
                    <img :src="session.Avatar == '' ? 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQFrZh6NXKZ7x0WW0UR2pPf2pXOrCaFcd62Uw&usqp=CAU' : session.Avatar"
                        alt="">
                </div>
                <div class="desc">
                    <p class="name">
                        {{ session.Name }}
                    </p>
                    <p class="line">{{ session?.lastMsg }}</p>
                </div>
                <div class="info">
                    <div class="time">{{session.lastMsgTime?transformTime(session.lastMsgTime):""}}</div>
                    <div class="unread" v-if="session.unReadCount">{{session.unReadCount}}</div>
                </div>
            </div>
        </div>
        <div class="msg">
            <chat v-if="show"></chat>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useSessionStore} from "@/store"
import { onBeforeMount, ref } from "vue"
import { sessionInt } from "@/type"
import { useRoute } from "vue-router"
import chat from "@/components/session/chat.vue"
const sessionStore = useSessionStore()
const route = useRoute()
// 是否有选中的聊天
const show = ref(false)
// 转换时间戳
const transformTime = (time:number):string=>{
    const date = new Date(time)
    return "" + date.getHours() +":" + date.getMinutes()
}
// 切换会话
const switchSession = (id:number,type:number)=>{
    show.value = true
    sessionStore.switchSession(id,type)
}
// 获取路由参数
onBeforeMount(() => {
    const cur: any = route.query
    cur.ID = parseInt(cur.ID)
    cur.type = parseInt(cur.type)
    // 没有传参
    if (isNaN(cur.ID)) return
    cur.unReadCount = 0
    show.value = true
    sessionStore.updateSessionList(cur as sessionInt)
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
            border-bottom: 1px solid #eee;
            background-color: #f9f9fb;

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
                width: 104px;
                height: 80%;
                margin-left: 10px;
                display: flex;
                flex-direction: column;
                .name {
                    font-size: 14px;
                    margin-bottom: 5px;
                    color: #000;
                }

                .line {
                    font-size: 12px;
                    color: #aaa;
                    white-space: nowrap;  

                    text-overflow:ellipsis; 
                   
                    overflow:hidden;
                }

            }

            .info {
                display:flex;
                flex-direction: column;
                height: 80%;
                align-items: center;
                .unread {
                    height: 14px;
                    width: 14px;
                    font-size: 10px;
                    line-height: 14px;
                    color:#fff;
                    text-align: center;
                    border-radius: 50%;
                    background-color: #fa5151;
                }

                .time {
                    color: #ccc;
                    margin-top: 2px;
                    margin-bottom: 5px;
                    font-size: 12px;
                }
            }
        }

    }

    .item.select {
        background: linear-gradient(to right, #f9d9da, #fff, #fff);
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
    height: 100%;
    flex-grow: 1;
}
</style>