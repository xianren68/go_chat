<template>
    <div class="message">
        <!--聊天用户列表-->
        <div class="list">
            <div class="item" v-for="(session, i) in sessionStore.sessionList" :key="i"
                :class="{ select: route.name == 'chat' && i == 0 }" @click="switchSession(session.ID, session.type)"
                @mouseup.right="deleteSe($event,i)">
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
                    <div class="time">{{ session.lastMsgTime ? transformTime(session.lastMsgTime) : "" }}</div>
                    <div class="unread" v-if="session.unReadCount">{{ session.unReadCount }}</div>
                </div>
            </div>
        </div>
        <div class="msg">
            <router-view></router-view>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useSessionStore, useMessageStore, userStore } from "@/store"
import { onBeforeMount, toRaw } from "vue"
import { sessionInt } from "@/type"
import { useRouter, useRoute } from "vue-router"
import { putSession, deleteSession } from "@/db";
import { ElMessageBox } from "element-plus";
const sessionStore = useSessionStore()
const messageStore = useMessageStore()
const userstore = userStore()
const route = useRoute()
const router = useRouter()
// 转换时间戳
const transformTime = (time: number): string => {
    const date = new Date(time)
    const minutes = date.getMinutes()
    if (minutes - 10 < 0) {
        return "" + date.getHours() + ":" + "0" + minutes
    }
    return "" + date.getHours() + ":" + minutes

}
// 切换会话
const switchSession = async (id: number, type: number) => {
    sessionStore.switchSession(id, type)
    // 获取对应的聊天信息
    const { unReadCount } = sessionStore.sessionList[0]
    const ID = id
    // 更新聊天列表
    await messageStore.replaceList(ID, unReadCount as number, type)
    // 减去小红点的次数
    userstore.unreadMessage -= unReadCount
    sessionStore.sessionList[0].unReadCount = 0
    // 更新会话
    putSession(userstore.db!, sessionStore.sessionList[0].type == 1 ? "usersession" : "groupsession", toRaw(sessionStore.sessionList[0]))
    router.push({ name: 'chat' })
}
// 删除会话
const deleteSe = (e:Event,index: number) => {
    e.preventDefault()
    ElMessageBox.confirm('确定要删除该会话吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
    }).then(() => {
        if (sessionStore.sessionList[index].type == 1) {
            deleteSession(userstore.db!, 'usersession', sessionStore.sessionList[index].ID)
        } else {
            deleteSession(userstore.db!, 'groupsession', sessionStore.sessionList[index].ID)
        }
        sessionStore.sessionList.splice(index, 1)
        if (index == 0) {
            router.push("/")
        }
    }
    )

}
// 获取路由参数
onBeforeMount(() => {
    const cur: any = route.query
    cur.ID = parseInt(cur.ID)
    cur.type = parseInt(cur.type)
    // 没有传参
    if (isNaN(cur.ID)) return
    cur.unReadCount = 0
    sessionStore.updateSessionList(cur as sessionInt)
    router.push({ name: 'chat' })
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

                    text-overflow: ellipsis;

                    overflow: hidden;
                }

            }

            .info {
                display: flex;
                flex-direction: column;
                height: 80%;
                align-items: center;

                .unread {
                    height: 14px;
                    width: 14px;
                    font-size: 10px;
                    line-height: 14px;
                    color: #fff;
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
}</style>