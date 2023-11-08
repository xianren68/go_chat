<template>
    <div class="me">
        <span class="content" v-if="!data.md">
            {{data.content}}
        </span>
        <div class="md" v-else>
            <previewmd  :text="data.content"></previewmd>
        </div>
        <div class="avatar">
            <img :src="userstore.userInfo!.Avatar" alt="">
        </div>
    </div>
</template>

<script setup lang="ts">
import previewmd from './previewmd.vue'
import {userStore} from '@/store'
const userstore = userStore()
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
            border-radius: 5px;
            line-height: 30px;
            padding: 0 4px;
            background-color: #526fff;
            border: 1px solid #eee;
            font-family: fira code;
            font-size: 16px;
            color: #fff;

        }
        .md {
            max-width: 40%;
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