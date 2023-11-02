<template>
    <div class="info">
      <div class="main">
            <div class="avatar">
              <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQFrZh6NXKZ7x0WW0UR2pPf2pXOrCaFcd62Uw&usqp=CAU" alt="">
            </div>
            <p class="line">{{information.Line}}</p>
            <div class="contact">
                <svg class="icon">
                  <use xlink:href="#icon-shouji"></use>
                </svg>
                <span>{{information.Phone == ""?"未绑定手机号":information.Phone}}</span>
              </div>
            <div class="contact">
              <svg class="icon">
                <use xlink:href="#icon-youxiang"></use>
              </svg>
                <span>{{information.Email == ""?"未绑定手机号":information.Email}}</span>
              </div>
            <div class="connect">
              <svg class="icon" @click="jumpMsg(information.ID,information.Name,information.Avatar)">
                <use xlink:href="#icon-faxiaoxi"></use>
              </svg>
              <svg class="icon">
                <use xlink:href="#icon-dadianhua"></use>
              </svg>
              <svg class="icon">
                <use xlink:href="#icon-shipin"></use>
              </svg>
            </div>
          </div>
      <div class="more">
        <svg class="icon">
          <use xlink:href="#icon-gengduo"></use>
        </svg>
      </div>
    </div>
</template>
<script setup lang="ts">
import {useRouter} from 'vue-router'
// 路由
const router = useRouter()
defineProps(['information'])
// 跳转到聊天页面
const jumpMsg = (ID:number,Name:string,Avatar:string)=>{
  router.push({name:'session',query:{ID,Name,Avatar,type:1}})
}
</script>

<style scoped lang="scss">
  .info {
    display: flex;
    height: 100%;
    width: 100%;
    align-items: center;
    justify-content: center;
    .main {
      display: flex;
      height: 60%;
      width: 60%;
      flex-direction: column;
      border-radius: 10px;
      background: linear-gradient(to right, #43c6ac, #f8ffae);
      box-shadow: rgba(0, 0, 0, 0.25) 0px 14px 28px, rgba(0, 0, 0, 0.22) 0px 10px 10px;
      align-items: center;
      .avatar{
        height: 100px;
        width: 100px;
        margin-top: -60px;
        overflow: hidden;
        border-radius: 50%;
        img{
          height: 100%;
          width: 100%;
        }
      }
      .line{
        margin:18px 0;
        font-weight: bold;
        font-family: "Fira Code Light";
        font-style: italic;
      }
      .contact{
          display: flex;
          width: 260px;
          margin-top: 20px;
          justify-content: start;
          .icon {
            height: 20px;
            width: 20px;
            fill: #008c8c;
          }
          span {
            margin-left: 20px;
            color: #999;
          }

      }
      .connect{
        margin-top: 35px;
        width: 250px;
        display: flex;
        justify-content: space-around;
        .icon {
          height: 30px;
          width: 30px;
        }
      }
    }
    .more{
      position: relative;
      top:-120px;
      left:-30px;
      .icon{
        height: 20px;
        width: 20px;
      }
    }
  }
</style>