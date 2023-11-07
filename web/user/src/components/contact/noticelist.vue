<template>
    <div class="notice">
        <div class="title">通知</div>
        <div v-for="item in noticeStore.noticeList" class="item">
            <div class="img">
              <img :src="item.avatar" alt="">
            </div>
            <div class="desc">
              <p class="name">{{item.send_name}}</p>
              <p class="line">{{item.type == 3?"请求添加你为好友":"请求加入群聊:"+item.group_name}}</p>
            </div>
            <div class="cancel" @click="cancel(item)">
                <svg class="icon">
                    <use xlink:href="#icon-quxiao"></use>
                  </svg>
            </div>
            <div class="btn" @click="addFriend(item)">
                同意
            </div>
          </div>
    </div>
</template>

<script setup lang="ts">
import { deleteNotice } from '@/db';
import { useNoticeStore,userStore,useContactStore } from '@/store'
import { messageInt } from '@/type'
import {newFriend,joinCommunity} from '@/api'
import { ElMessage } from 'element-plus';
const noticeStore = useNoticeStore()
const userstore = userStore()
const contactStore = useContactStore()
// 取消通知
const cancel = (item:messageInt) => {
    // 删除消息
    deleteNotice(userstore.db!,item)
    let index = noticeStore.noticeList.indexOf(item)
    noticeStore.noticeList.splice(index,1)
    if(noticeStore.noticeList.length == 0){
        userstore.unreadNotice = false
    }
}
// 发送同意请求
const addFriend = async (item:messageInt) => {
    cancel(item)
    let res:any
    if(item.type == 3){
        // 添加好友
        res = await newFriend({name:item.send_name})
        let data = res.data
        if(data.code !=0){
            ElMessage.error(data.msg)
            return
        }
        ElMessage.success('添加成功')
        // 刷新好友列表
        contactStore.getFriendList()
    }else{
        // 将用户加入群聊
        res = await joinCommunity({name:item.group_name!,id:item.from_id})
        let data = res.data
        if(data.code !=0){
            ElMessage.error(data.msg)
            return
        }
        ElMessage.success('添加成功')
        // 刷新群组列表
        contactStore.getGroupList()

    }
}
</script>

<style lang="scss" scoped>
.notice{
    margin-top: 20px;
    display: flex;
    max-height: 200px;
    overflow: scroll;
    flex-direction: column;
    align-items: center;
    .title{
        height: 20px;
        font-size: 12px;
        font-weight: bold;
        padding-left: 10px;
        width: 100%;
        text-align: left;
        background-color: #f0f4e1;
    }
    .item{
      box-sizing: border-box;
      height: 50px;
      width: 100%;
      display: flex;
      padding:0 10px;
      align-items: center;
      border-bottom:1px solid #eee ;
      border-top:1px solid #eee ;
      margin-bottom: 10px;
      background-color:#f9f9fb;
      .img{
        height:35px;
        width: 35px;
        border-radius: 50%;
        overflow: hidden;
        img{
          height: 100%;
          width: 100%;
        }
      }
      .desc{
        width:108px;
        margin-left: 10px;
        .name{
          font-size: 14px;
          font-weight: normal;
          margin-bottom: 3px;
        }
        .line{
          font-size: 10px;
          color:#aaa;
          white-space: nowrap;

          text-overflow: ellipsis;

          overflow: hidden;
        }
      }
      .btn{
        height: 20px;
        width: 40px;
        font-size: 10px;
        text-align: center;
        line-height: 20px;
        border-radius: 5px;
        border:1px solid #aaa;
      }
      .cancel{
        position: relative;
        top:-18px;
        left:40px;
        .icon{
            height: 15px;
            width:15px;
        }
      }
    }
}
.notice::-webkit-scrollbar{
    display: none;
  }
</style>