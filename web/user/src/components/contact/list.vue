<template>
  <div class="list">
    <div v-for="(item,i) in List" :key="i" class="item-list" @click="$emit('update:modelValue',i)" :class="{select:i==modelValue}">
      <div class="img">
        <img :src="item.Avatar == ''?'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQFrZh6NXKZ7x0WW0UR2pPf2pXOrCaFcd62Uw&usqp=CAU':item.Avatar" alt="">
      </div>
      <div class="desc">
        <p class="name">{{item.Name}}</p>
        <p class="line">{{item.Line}}</p>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import {useContactStore} from "@/store"
defineProps(['List','modelValue'])
defineEmits()
</script>
<style scoped lang="scss">
.list{
  margin-top: 20px;
  display: flex;
  height: 448px;
  overflow: scroll;
  flex-direction: column;
  align-items: center;
  .item-list{
    box-sizing: border-box;
    height: 50px;
    width: 100%;
    display: flex;
    padding:0 10px;
    align-items: center;
    border-bottom:1px solid #eee ;
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
      flex-grow: 1;
      margin-left: 10px;
      .name{
        font-size: 14px;
        font-weight: normal;
        margin-bottom: 3px;
      }
      .line{
        font-size: 10px;
        color:#aaa;
      }
    }
  }
  .item-list.select{
    background: linear-gradient(to right, #f9d9da, #fff,#fff);
    transition: .5s;
  }
}
.list::-webkit-scrollbar{
  display: none;
}
</style>