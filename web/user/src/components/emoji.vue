<template>
    <div class="emoji">
        <div class="list">
            <div v-for="item in emojis.emojilist" class="item" @click="insert(item)">{{item}}</div>
        </div>
        <div class="arrow"></div>
    </div>
    
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useEmojis } from '@/store'
const emits = defineEmits(['insertEmoji'])
const emojis = useEmojis()
// 获取emojis
onMounted(()=>{
    emojis.getEmojiList()
})
// 向父组件传递emoji
const insert = (item:string)=>{
    emits("insertEmoji",item)
}
</script>

<style lang="scss" scoped>
.emoji{
    position: relative;
    bottom:210px;
    left:580px;
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100px;
    .list {
        flex-direction: row;
        flex-wrap: wrap;
        display: flex;
        padding-bottom: 10px;
        background-color: #fff;
        border-radius: 15px;
        border: 1px solid #ccc;
        width: 180px;
        height: 140px;
        overflow: auto;
        box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
        .item {
            height: 30px;
            width: 30px;
            font-size: 16px;
            display:flex;
            justify-content: center;
            align-items: center;
        }
        .item:hover{
            transform: scale(1.2);
            transition: .3s;
        }
    }
    .arrow{
        position: relative;
        top: -2px;
        border: 10px solid transparent;
        border-top:10px solid #fff ;
    }
}
.list::-webkit-scrollbar{
    display: none;
  }
</style>