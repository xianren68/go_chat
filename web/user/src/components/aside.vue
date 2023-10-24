<template>
    <div class="aside">
        <div class="avatar">
            <img src="../assets/img/avatar.jpg" alt="">
        </div>
        <a class="item" @click="message">
            <svg class="icon">
                <use xlink:href="#icon-message"></use>
            </svg>
        </a>
        <a class="item" @click="contact">
            <svg class="icon">
                <use xlink:href="#icon-contact"></use>
            </svg>
        </a>
        <a class="setting" @click="setting">
            <svg class="icon">
                <use xlink:href="#icon-setting"></use>
            </svg>
        </a>
    </div>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router";
const router = useRouter()
// 被选中的项
let active: HTMLElement
// 跳转到消息页
const message = (e:Event)=>{
    animation(e.currentTarget as HTMLElement)
}
// 跳转到联系人页
const contact = (e:Event)=>{
    animation(e.currentTarget as HTMLElement)
  router.push('contact')

}
// 跳转到设置页
const setting = (e:Event)=>{
    animation(e.currentTarget as HTMLElement)
}
// 点击选项动画
const animation = (e: HTMLElement) => {
    
    // 重复点击
    if (active == e) {
        return
    }
    // 有被选中的
    if (active != undefined) {
        active.classList.remove('select')
    }
    // 添加样式
    active = e
    active.classList.add('select')
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