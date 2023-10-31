<template>
  <div class="login">
    <p class="title">LOGIN</p>
    <div class="info">
      <div class="name item">
        <svg class="icon">
          <use xlink:href="#icon-user"></use>
        </svg>
        <input type="text" placeholder="用户名" v-model="name">
        <span class="verify" v-show="verName">用户名不能为空</span>
      </div>
      <div class="pwd item">
        <svg class="icon">
          <use xlink:href="#icon-pwd"></use>
        </svg>
        <input type="password" placeholder="密码" v-model="pwd">
        <span class="verify" v-show="verPwd">密码不能为空</span>
      </div>
      <div class="protocol">
        <span class="agree">
          <input type="checkbox" v-model="agree">
          <span>
            同意协议
          </span>
        </span>
        <span class="forget">
          忘记密码
        </span>
      </div>
      <span class="verify" v-show="verAgree">请同意协议</span>
      <button @click="Login">LOGIN</button>
      <div class="register">
        <span>没有账号？ </span>
        <span class="link">去注册</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import {login} from "../../api";
import {useRouter} from "vue-router";
import {ElMessage} from "element-plus"
const router = useRouter()
const name = ref("")
const pwd = ref("")
const agree = ref(false)
// 提示框是否出现
const verName = ref(false)
const verPwd = ref(false)
const verAgree = ref(false)
// 验证
const verify = ():boolean=>{
    let flag = false
    if (name.value == ""){
      verName.value = true
      flag = true
    }
    if (pwd.value == ""){
      verPwd.value = true
      flag = true
    }
    if (agree.value == false){
      verAgree.value = true
      flag = true
    }
    return flag
}
// 用户登录
const Login = async ()=>{
  if(verify()){
    return
  }
  const {data} = await login({name:name.value,password:pwd.value})
  if(data.code == -1){
    ElMessage.error(data.msg)
    return
  }
  localStorage.setItem("userInfo",JSON.stringify(data.userInfo))

  // 登录成功
  ElMessage.success("登录成功")
  router.push('/')
}
</script>

<style lang="scss" scoped>
  .login{
    position: absolute;
    left: 0;
    right: 0;
    top:0;
    bottom: 0;
    margin: auto;
    height: 360px;
    width: 300px;
    background-color: #fff;
    box-shadow: rgba(0, 0, 0, 0.1) 0px 10px 15px -3px, rgba(0, 0, 0, 0.05) 0px 4px 6px -2px;
    .title{
       margin-top: 20px;
        font-size: 20px;
      text-align: center;
      font-weight: bolder;
    }
    .info{
      display: flex;
      flex-direction: column;
      align-items: center;
      line-height: 16px;
      .item {
        border-bottom: #ccc 2px solid;
        height: 30px;
        width: 200px;
        .icon {
          display: inline-block;
          width: 20px;
          height: 20px;
        }
      }
      .name {
        margin-top: 60px;
      }
      .pwd {
        margin-top: 35px;
      }
      input {
        display: inline-block;
        height: 30px;
        width: 160px;
        margin-left: 10px;
        border: none;
        outline: none;
        background-color: transparent;
      }
      .protocol{
        margin-top: 20px;
        width: 200px;
        font-size: 10px;
        color:#ccc;
        input {
          transform: translateY(2px);
          height: 12px;
          width: 10px;
        }
        .agree{
          display: inline-block;
          width: 75px;
          border-right: #ccc 1px solid;
        }
        .forget{
          margin-left: 10px;
          color: #8cc63e;
        }
      }
      button {
        margin-top: 20px;
        background-color: #0090fc;
        font-weight: bolder;
        border-radius: 8px;
        box-shadow: rgba(0, 0, 0, 0.24) 0px 3px 8px;
        width: 200px;
        height: 40px;
        color: #fff;
        border:none;
      }
      .register{
        margin-top: 23px;
        font-size: 12px;
        color: #ccc;
        .link{
          color: #8cc63e;
        }
      }
      .verify{
        display: inline-block;
        width: 200px;
        color: red;
        font-size: 10px;
      }
    }
  }
</style>