<template>
    <div class="info">
      <div class="main">
            <div class="avatar">
                <input type="file" hidden accept="image/*" ref="fileRef" @change="upload">
              <img :src="userstore.userInfo!.Avatar" alt="" title="修改你的头像" @click="fileRef?.click()">
            </div>
            <div class="editImg"></div>
            <p title="双击编辑"  class="line" @dblclick="LineShow=true" v-if="!LineShow">{{userstore.userInfo!.Line==""?"编辑个性签名":userstore.userInfo!.Line}}</p>
            <div class="edit" v-else>
              <input type="text" placeholder="请输入个性签名" v-model="Line" >
              <svg class="icon" @click="submit(1,Line)">
                <use xlink:href="#icon-icon--duigou"></use>
              </svg>
              <svg class="icon" @click="LineShow=false;Line=''">
                <use xlink:href="#icon-yiquxiao"></use>
              </svg>
            </div>
            <div class="contact">
              <svg class="icon">
                <use xlink:href="#icon-user"></use>
              </svg>
              <span title="双击编辑" v-if="!userNameShow" @dblclick="userNameShow=true">{{userstore.userInfo!.Name}}</span>
              <div class="edit" v-else>
                <input type="text" placeholder="请输入用户名" v-model="userName" >
                <svg class="icon" @click="submit(2,userName)">
                  <use xlink:href="#icon-icon--duigou"></use>
                </svg>
                <svg class="icon" @click="userNameShow =false;userName=''">
                  <use xlink:href="#icon-yiquxiao"></use>
                </svg>
              </div>
            </div>
            <div class="contact">
              <svg class="icon">
                <use xlink:href="#icon-xingbie"></use>
              </svg>
              <span title="双击编辑" v-if="!sexShow" @dblclick="sexShow=true">{{userstore.userInfo!.Name == ""?"未设置":userstore.userInfo!.Gender=="male"?"男":"女"}}</span>
              <div class="edit" v-else>
                <input type="text" placeholder="请输入性别" v-model="sex" >
                <svg class="icon" @click="submit(3,sex)">
                  <use xlink:href="#icon-icon--duigou"></use>
                </svg>
                <svg class="icon" @click="sexShow= false;sex=''">
                  <use xlink:href="#icon-yiquxiao"></use>
                </svg>
              </div>
            </div>
            <div class="contact">
                <svg class="icon">
                  <use xlink:href="#icon-shouji"></use>
                </svg>
                <span title="双击编辑" v-if="!phoneShow" @dblclick="phoneShow=true">{{userstore.userInfo!.Phone == ""?"未绑定手机号":userstore.userInfo!.Phone}}</span>
                <div class="edit" v-else>
                  <input type="tel" placeholder="请输入手机号" v-model="phone" >
                  <svg class="icon" @click="submit(4,phone)">
                    <use xlink:href="#icon-icon--duigou"></use>
                  </svg>
                  <svg class="icon" @click="phoneShow= false;phone=''">
                    <use xlink:href="#icon-yiquxiao"></use>
                  </svg>
                </div>
              </div>
            <div class="contact">
              <svg class="icon">
                <use xlink:href="#icon-youxiang"></use>
              </svg>
                <span title="双击编辑" v-if="!emailShow" @dblclick="emailShow=true">{{userstore.userInfo!.Email == ""?"未绑定邮箱":userstore.userInfo!.Email}}</span>
                <div class="edit" v-else>
                  <input type="tel" placeholder="请输入邮箱" v-model="email" >
                  <svg class="icon" @click="submit(5,email)">
                    <use xlink:href="#icon-icon--duigou"></use>
                  </svg>
                  <svg class="icon" @click="emailShow= false;email=''">
                    <use xlink:href="#icon-yiquxiao"></use>
                  </svg>
                </div>
              </div>
          </div>
      <div class="more" @click="optionShow = !optionShow">
        <svg class="icon">
          <use xlink:href="#icon-gengduo"></use>
        </svg>
      </div>
    </div>
    <div class="option" v-if="optionShow">
      <div class="arrow"></div>
      <div class="item">
        <svg class="icon">
          <use xlink:href="#icon-tuichu"></use>
        </svg>
        <span>退出登录</span>
      </div>
      <div class="item">
        <svg class="icon">
          <use xlink:href="#icon-bianji"></use>
        </svg>
        <span>修改密码</span>
      </div>
    </div>
</template>
<script setup lang="ts">
import {userStore} from '@/store'
import {ref} from 'vue'
import { uploadImage } from '@/api'
import { ElMessage } from 'element-plus'
import {isEmail,isMobile,isSex} from '@/utils/verify'
import {updateUserInfo} from '@/api'
const userstore = userStore()
const fileRef = ref<HTMLInputElement>()
// 修改头像
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
// 个性签名
const Line = ref("")
// 控制个性签名编辑框是否出现
const LineShow = ref(false)
// 用户名
const userName = ref("")
// 控制用户名编辑框是否出现
const userNameShow = ref(false)
// 性别
const sex = ref("")
// 控制性别编辑框是否出现
const sexShow = ref(false)
// 手机号
const phone = ref("")
// 控制手机号编辑框是否出现
const phoneShow = ref(false)
// 邮箱
const email = ref("")
// 控制邮箱编辑框是否出现
const emailShow = ref(false)
// 验证并发送请求
const submit = async (t:number,val:string) => {
  if(val == ""){
    ElMessage.error("不能为空")
    return
  }
  let d:any  = {}
  let res:any
  switch (t){
    // 签名
    case 1:
      d.line = val
      res = await updateUserInfo(d)
      if (res.data.code === 0){
        userstore.userInfo!.Line = val
        Line.value = ''
        LineShow.value = false
      }else {
        ElMessage.error(res.data.msg)
        return
      }
      break
    case 2:
      d.Name = val
      res = await updateUserInfo(d)
      if (res.data.code === 0){
        userstore.userInfo!.Name = val
        userName.value = ''
        userNameShow.value = false
      }else {
        ElMessage.error(res.data.msg)
        return
      }
      break
    case 3: 
      if(!isSex(val)){
        ElMessage.error("性别输入错误")
        return
      }
      val = val=="男"?"male":"female"
      d.sex = val
      res = await updateUserInfo(d)
      if (res.data.code === 0){
        userstore.userInfo!.Gender = val
        sex.value = ''
        sexShow.value = false
      }else {
        ElMessage.error(res.data.msg)
        return
      }
      break
    case 4:
      if(!isMobile(val)){
        ElMessage.error("手机号输入错误")
        return
      }
      d.phone = val
      res = await updateUserInfo(d)
      if (res.data.code === 0){
        userstore.userInfo!.Phone = val
        phone.value = ''
        phoneShow.value = false

      }else {
        ElMessage.error(res.data.msg)
        return
      }
      break
    case 5:
      if(!isEmail(val)){
        ElMessage.error("邮箱输入错误")
        return
      }
      d.email = val
      res = await updateUserInfo(d)
      if (res.data.code === 0){
        userstore.userInfo!.Email = val
        email.value = ''
        emailShow.value = false
      }else {
        ElMessage.error(res.data.msg)
        return
      }
      break
  }
  ElMessage.success("修改成功")
  userstore.update(null)
}
// 操作是否出现
const optionShow = ref(false)
</script>

<style scoped lang="scss">
  .info {
    flex-grow: 1;
    height: 100%;
    position: relative;
    .main {
      position:absolute;
      top:0;
      left: 0;
      right: 0;
      bottom: 0;
      margin:auto;
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
        margin-bottom: 10px;
        overflow: hidden;
        border-radius: 50%;
        img{
          height: 100%;
          width: 100%;
        }
      }
      .line{
        margin:10px 0;
        font-weight: bold;
        font-family: "Fira Code Light";
        font-style: italic;
        color: #fff;
      }
      .contact{
          display: flex;
          width: 260px;
          margin-top: 20px;
          justify-content: start;
          align-items: center;
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
    .edit{
      margin-top:10px;
      display: flex;
      align-items: center;
      width: 260px;
      height: 25px;
      border-bottom: #ccc 1px solid;
      input {
        width: 150px;
        padding-left: 15px;
        color: #fffde4;
        border: none;
        outline: none;
        background-color: transparent;
      }
      .icon {
        margin-left: 10px;
        height: 20px;
        width: 20px;
      }
    }
    .more{
      position: relative;
      top:120px;
      left:720px;
      .icon{
        height: 20px;
        width: 20px;
      }
    }
  }
  .option{
    position: relative;
    height: 50px;
    top:115px;
    left: -186px;
    background-color: #fff;
    padding: 5px 10px;
    .arrow {
      position: absolute;
      top:18px;
      left:-25px;
      border:15px solid rgba(0,0,0,0);
      border-right:15px solid #fff ;
    }
    .item {
      display: flex;
      align-items: center;
      font-size: 12px;
      margin-top: 5px;
      .icon{
        height: 15px;
        width: 15px;
        margin-right: 10px;
      }
    }
  }
</style>