<template>
    <div class="find">
        <div class="img">
            <img :src="data.Avatar" alt="">
        </div>
        <div class="desc">
            <p class="name">{{ data.Name }}</p>
            <p class="line">{{ data.Line }}</p>
        </div>
        <div class="add" @click="sendRequest">
            {{ data.isUser ? "添加" : "加入" }}
        </div>
    </div>
</template>

<script setup lang="ts">
import socket from '@/api/socket'
import { messageInt } from '@/type'
import { ElMessage } from 'element-plus'
import { userStore } from '@/store'
const userstore = userStore()

const props = defineProps(['data',"modelValue"])
const emits = defineEmits()
// 发送请求
const sendRequest = () => {
    let send_msg: messageInt
    const now = Date.now()
    // 添加好友
    if (props.data.isUser) {
        send_msg = {
            from_id: userstore.userInfo.ID,
            target_id: props.data.ID,
            content: "",
            type: 3,
            avatar: userstore.userInfo.Avatar,
            send_name: userstore.userInfo.Name,
            send_time: now,
        }
    }else{
        // 请求加群
        send_msg = {
            from_id: userstore.userInfo.ID,
            target_id: props.data.OwnerId,
            content: "",
            type: 4,
            avatar: userstore.userInfo.Avatar,
            send_name: userstore.userInfo.Name,
            group_id: props.data.ID,
            group_name:props.data.Name,
            send_time: now,
        }
    }
    // 发送消息
    socket.s?.send(JSON.stringify(send_msg))
    ElMessage.success("请求已发送")
    emits('update:modelValue',false)
}
</script>

<style scoped lang="scss">
.find {
    box-sizing: border-box;
    margin: 10px 0;
    height: 50px;
    width: 100%;
    display: flex;
    padding: 0 10px;
    align-items: center;
    border-bottom: 1px solid #eee;
    border-top: 1px solid #eee;
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
        margin-left: 10px;

        .name {
            font-size: 14px;
            font-weight: normal;
            margin-bottom: 3px;
        }

        .line {
            font-size: 10px;
            color: #aaa;
        }
    }

    .add {
        height: 20px;
        width: 40px;
        font-size: 10px;
        text-align: center;
        line-height: 20px;
        border: 1px #ccc solid;
        border-radius: 5px;
    }

}</style>