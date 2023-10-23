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
// 被选中的项
let active: HTMLElement
// 跳转到消息页
const message = (e:Event)=>{
    animation(e.currentTarget as HTMLElement)
}
// 跳转到联系人页
const contact = (e:Event)=>{
    animation(e.currentTarget as HTMLElement)
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
    width: 4%;
    border-top-left-radius: 15px;
    border-bottom-left-radius: 15px;
    background: linear-gradient(to bottom, #c0c0aa, #1cefff);

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