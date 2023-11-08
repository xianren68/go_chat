<template>
    <div class="aside">
        <div class="avatar" @click="router.push({name:'user'})">
            <img :src="userstore.userInfo!.Avatar" alt="">
        </div>
        <a class="item" @click="message" :class="{select:router.currentRoute.value.name == 'session'||router.currentRoute.value.name == 'chat' }">
            <svg class="icon">
                <use xlink:href="#icon-message"></use>
            </svg>
            <p class="unreadMsg" v-if="userstore.unreadMessage>0">{{userstore.unreadMessage}}</p>
        </a>
        <a class="item" @click="contact" :class="{select:router.currentRoute.value.name == 'contact'}">
            <svg class="icon">
                <use xlink:href="#icon-contact"></use>
            </svg>
            <p class="unreadNotice" v-if="userstore.unreadNotice"></p>
        </a>
        <a class="setting" @click="setting">
            <svg class="icon">
                <use xlink:href="#icon-setting"></use>
            </svg>
        </a>
    </div>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router"
import {userStore} from '@/store'
const router = useRouter()
const userstore = userStore()
// 跳转到消息页
const message = ()=>{
    router.push({name:'session'})
}
// 跳转到联系人页
const contact = ()=>{
  router.push('contact')

}
// 跳转到设置页
const setting = ()=>{
}
</script>

<style lang="scss" scoped>
.aside {
    display: flex;
    position: relative;
    flex-direction: column;
    align-items: center;
    height: 100%;
    width: 60px;
    border-radius: 10px;
    background-color: #484aa1;
  box-shadow: rgba(50, 50, 93, 0.25) 0px 6px 12px -2px, rgba(0, 0, 0, 0.3) 0px 3px 7px -3px;

    .avatar {
        margin-top: 30px;
        margin-bottom: 10px;
        width: 30px;
        height: 30px;
        border-radius: 50%;
        overflow: hidden;

        img {
            height: 100%;
            width: 100%;
        }
    }

    .icon {
        margin-top: 20px;
        width: 20px;
        height: 20px;
    }
    .unreadMsg {
        position: relative;
        text-align: center;
        font-size: 10px;
        top:-29px;
        left:13px;
        color: #fff;
        background-color: #fa5151;
        border-radius: 5px;
    }
    .unreadNotice{
        position: relative;
        width: 6px;
        height: 6px;
        top:-26px;
        left:18px;
        border-radius: 50%;
        background-color: #fa5151;
    }
    p {
        transition: 0.5s;
    }
    .select .icon {
        transform: scale(1.2);
        transition: .5s;
        fill: var(--icon-active-color);
    }

    .setting {
        position: absolute;
        bottom: 20px;

    }
}</style>