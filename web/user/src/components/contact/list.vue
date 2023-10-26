<template>
  <div class="list" v-if="switchCt">
    <div v-for="(item,i) in contactPerson" :key="i" class="item-list" @click="$emit('update:modelValue',i)" :class="{select:i==modelValue}">
      <div class="img">
        <img :src="item.Avatar == ''?'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQFrZh6NXKZ7x0WW0UR2pPf2pXOrCaFcd62Uw&usqp=CAU':item.Avatar" alt="">
      </div>
      <div class="desc">
        <p class="name">{{item.Name}}</p>
        <p class="line">{{item.Line}}</p>
      </div>
    </div>
  </div>
  <div class="list" v-if="!switchCt" >
    <div v-for="item in contactPerson" class="item-list">
      <span>{{item.Name}}</span>
    </div>
  </div>
</template>
<script setup lang="ts">
import {useContactStore} from "@/store"
defineProps({
  switchCt:Boolean,
  modelValue:Number,
})
defineEmits()
const contactPerson = useContactStore().contactPerson
</script>
<style scoped lang="scss">
.list{
  margin-top: 20px;
  padding:5px;
  display: flex;
  height: 448px;
  overflow: scroll;
  flex-direction: column;
  align-items: center;
  .item-list{
    height: 50px;
    width: 94%;
    display: flex;
    padding:0 5px;
    margin-top: 10px;
    align-items: center;
    background: linear-gradient(to right, #d3cce3, #e9e4f0);
    border-radius: 10px;
    box-shadow: rgba(149, 157, 165, 0.2) 0px 8px 24px;
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
        font-weight: bold;
        margin-bottom: 3px;
      }
      .line{
        font-size: 10px;
        color:#aaa;
      }
    }
  }
  .item-list.select{
    background: linear-gradient(to right, #bdc3c7, #2c3e50);
    box-shadow: rgba(0, 0, 0, 0.25) 0px 54px 55px, rgba(0, 0, 0, 0.12) 0px -12px 30px, rgba(0, 0, 0, 0.12) 0px 4px 6px, rgba(0, 0, 0, 0.17) 0px 12px 13px, rgba(0, 0, 0, 0.09) 0px -3px 5px;
    transition: .5s;
  }
}
.list::-webkit-scrollbar{
  display: none;
}
</style>