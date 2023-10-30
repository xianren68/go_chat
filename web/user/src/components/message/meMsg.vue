<template>
    <div class="me">
        <span class="content">
            {{data.msg}}
        </span>
        <div class="avatar">
            <img src="../../assets/img/avatar.jpg" alt="">
        </div>
    </div>
</template>

<script setup lang="ts">
defineProps(['data'])
</script>

<style lang="scss" scoped>
    .me {
        margin-top: 20px;
        display: flex;
        justify-content: flex-end;
        align-items: self-start;
        .content{
            max-width: 40%;
            min-height: 30px;
            margin-right: 5px;
            border-radius: 10px;
            line-height: 30px;
            padding: 0 4px;
            background-color: #526fff;
            border: 1px solid #eee;
            font-family: fira code;
            font-size: 16px;
            color: #fff;

        }
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

    }
</style>