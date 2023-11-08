<template>
    <div class="info">
      <div class="main">
            <div class="avatar">
                <input type="file" hidden accept="image/*" ref="fileRef" @change="upload">
              <img :src="userstore.userInfo!.Avatar" alt="" title="修改你的头像" @click="fileRef?.click()">
            </div>
            <div class="editImg"></div>
            <p class="line">{{userstore.userInfo!.Line}}</p>
            <div class="contact">
                <svg class="icon">
                  <use xlink:href="#icon-shouji"></use>
                </svg>
                <span>{{userstore.userInfo!.Phone == ""?"未绑定手机号":userstore.userInfo!.Phone}}</span>
              </div>
            <div class="contact">
              <svg class="icon">
                <use xlink:href="#icon-youxiang"></use>
              </svg>
                <span>{{userstore.userInfo!.Email == ""?"未绑定邮箱":userstore.userInfo!.Email}}</span>
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
import {userStore} from '@/store'
import {ref} from 'vue'
import { uploadImage } from '@/api'
import { ElMessage } from 'element-plus';
const userstore = userStore()
const fileRef = ref<HTMLInputElement>()
const upload = async () => {
  const formData = new FormData()
  if (fileRef.value?.files == null) {
    return
  }
  formData.append('file', fileRef.value?.files[0])
  const {data} = await uploadImage(formData)
  if (data.code === 0) {
    ElMessage.success("修改成功")
    // 更新头像
    const info = userStore().userInfo!
    info.Avatar = data.data
    userStore().update(info)
  }else{
    ElMessage.error(data.msg)
  }
}

</script>

<style scoped lang="scss">
  .info {
    flex-grow: 1;
    display: flex;
    height: 100%;
    align-items: center;
    justify-content: center;
    .main {
      display: flex;
      height: 60%;
      width: 60%;
      flex-direction: column;
      border-radius: 10px;
      background: linear-gradient(to right, #005aa7, #fffde4);
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
            fill: #fffde4;
          }
          span {
            margin-left: 20px;
            color: #fffde4;
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