<template>
    <div class="message">
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
        <div class="msg">

        </div>
    </div>
</template>

<script setup lang="ts">
import {useMsgStore} from "@/store"
import {onBeforeMount,ref} from "vue"
import {messageInt} from "@/type"
import {useRoute} from "vue-router"

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
    width: 100%;
    display: flex;

    .list {
        margin-top: 20px;
        display: flex;
        height: 100%;
        width: 20%;
        overflow: scroll;
        flex-direction: column;
        align-items: center;

        .item {
            height: 50px;
            width: 90%;
            display: flex;
            padding: 0 5px;
            margin-top: 10px;
            align-items: center;
            background: linear-gradient(to right, #d3cce3, #e9e4f0);
            border-radius: 10px;
            box-shadow: rgba(149, 157, 165, 0.2) 0px 8px 24px;

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
            background: linear-gradient(to right, #bdc3c7, #2c3e50);
            box-shadow: rgba(0, 0, 0, 0.25) 0px 54px 55px, rgba(0, 0, 0, 0.12) 0px -12px 30px, rgba(0, 0, 0, 0.12) 0px 4px 6px, rgba(0, 0, 0, 0.17) 0px 12px 13px, rgba(0, 0, 0, 0.09) 0px -3px 5px;
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
}
</style>