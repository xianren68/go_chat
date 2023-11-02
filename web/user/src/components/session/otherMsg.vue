<template>
    <div class="other">
        <div class="avatar">
            <img :src="data.Avatar == '' ? 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQFrZh6NXKZ7x0WW0UR2pPf2pXOrCaFcd62Uw&usqp=CAU' : data.Avatar" alt="">
        </div>
        <span class="content">
            {{data.content}}
        </span>
    </div>
</template>

<script setup lang="ts">
defineProps(['data'])
</script>

<style lang="scss" scoped>
    .other {
        margin-top: 20px;
        display: flex;
        justify-content: flex-start;
        align-items: self-start;
        .avatar {
            width: 30px;
            height: 30px;
            border-radius: 50%;
            overflow: hidden;
            img {
                height: 100%;
                width: 100%;

            }
        }
        .content{
            max-width: 40%;
            min-height: 30px;
            margin-right: 5px;
            border-radius: 10px;
            line-height: 30px;
            padding: 0 4px;
            background-color:#fff;
            border: 1px solid #eee;
            font-family: fira code;
            font-size: 16px;
            color: #000;

        }

    }
</style>