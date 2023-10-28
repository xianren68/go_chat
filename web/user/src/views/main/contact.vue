<template>
  <div class="contact">
  <div class="left">
    <!--    搜索-->
    <div class="top">
      <div class="search" @click="()=>{isShow = !isShow}">
        <svg class="icon">
          <use xlink:href="#icon-sousuo"></use>
        </svg>
      </div >
      <input type="text" v-show="isShow" class="search-input" placeholder="查看用户/群组">
      <div class="cancel" v-show="isShow">
        <svg class="icon">
          <use xlink:href="#icon-quxiao"></use>
        </svg>
      </div>
    </div>
    <!--    切换群组/好友-->
    <div class="switch">
      <div class="friend" :class="{select:switchCt}" @click="switchFriend">好友</div>
      <div class="group" :class="{select:!switchCt}" @click="switchGroup">群组</div>
    </div>
    <!--    列表-->
    <List :switchCt="switchCt" v-model="selectIndex"></List>
  </div>
    <div class="right">
      <UserInfo v-if="selectIndex >=0" :information="contactPerson[selectIndex]"></UserInfo>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref} from 'vue'
import {useContactStore} from "../../store"
import {onBeforeMount} from "vue"
import List from '@/components/contact/list.vue'
import UserInfo from "@/components/contact/userInfo.vue";
// 控制搜索框是否显示
const isShow = ref(false)
// 控制选择好友还是群组
const switchCt = ref(true)
// 选中的好友
const selectIndex = ref(-1)
// 渲染列表
const {contactPerson,getFriendList} = useContactStore()
// 切换到用户列表
const switchFriend = ()=>{
  switchCt.value = true
}
// 切换到群组列表
const switchGroup = ()=>{
  switchCt.value = false
}
// 挂载之前，获取用户列表
onBeforeMount(()=>{
  getFriendList()
})
</script>

<style scoped lang="scss">
.contact {
  display: flex;
  align-items: center;
  height: 100%;
  flex-grow:1;
  .left{
    height: 100%;
    width: 20%;
    .top{
      margin-top: 2px;
      margin-left: 10px;
      height:40px;
      display: flex;
      justify-content: start;
      .search{
        width: 20px;
        height: 20px;
        border-radius: 50%;
        background-color: #f5f4fa;
        box-shadow: rgba(50, 50, 93, 0.25) 0px 30px 60px -12px, rgba(0, 0, 0, 0.3) 0px 18px 36px -18px;
        display: flex;
        justify-content: center;
        align-items: center;
        .icon{
          width: 15px;
          height: 15px;
          fill: #bbb;
        }
      }
      .search-input{
        margin-left: 10px;
        padding-left: 10px;
        height: 20px;
        width: 90%;
        font-size: 12px;
        border: none;
        border: #ccc 1px solid;
        border-radius: 5px;
        background-color: transparent;
        outline: none;
        transition: .5s;
      }
      .cancel{
        position: relative;
        height: 20px;
        width: 20px;
        background-color: #f5f4fa;
        border-radius: 50%;
        left:-18px;
        top:2px;
        transition: .5s;
        .icon {
          height: 15px;
          width: 15px;
          color: #ccc;
        }
      }
    }
    .switch{
      display: flex;
      justify-content: space-around;
      align-items: center;
      margin-left: 20px;
      margin-right: 20px;
      font-size: 14px;
      height: 30px;
      font-weight: bolder;
      div{
        width: 30%;
        height: 100%;
        text-align: center;
        transition: .5s;
      }
      .select{
        color: #008efe;
        border-bottom: #008efe 2px solid;
      }

    }
  }
  .right {
    flex-grow: 1;
    height: 100%;
    background-color: #fff;
  }
}
</style>