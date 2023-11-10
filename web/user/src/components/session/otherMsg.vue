<template>
    <div class="other">
        <div class="avatar">
            <img :src="data.avatar" alt="">
        </div>
        <div class="content">
            <span class="name" v-if="data.type != 1">
                {{data.send_name}}
            </span>
            <span class="msg" v-if="!data.md">
                {{data.content}}
            </span>
            <div class="md" v-else>
                <previewmd  :text="data.content"></previewmd>
            </div>
            
        </div>
        
    </div>
</template>

<script setup lang="ts">
import previewmd from './previewmd.vue'
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
        .content {
            max-width: 40%;
            margin-left: 5px;
            .name {
                display: block;
                font-size:10px ;
                color: #aaa;
                margin-bottom: 3px;
            }
            .msg{
                display: inline-block;
                max-width: 100%;
                min-height: 30px;
                margin-right: 5px;
                border-radius: 5px;
                line-height: 30px;
                padding: 0 4px;
                background-color:#ccc;
                font-family: fira code;
                font-size: 16px;
                color: #000;
    
            }
            .md{
                max-width: 100%;
            }
        }

    }
</style>