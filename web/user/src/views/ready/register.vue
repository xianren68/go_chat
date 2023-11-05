<template>
    <div class="register">
        <p class="title">REGISTER</p>
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
            <div class="pwd item">
                <svg class="icon">
                    <use xlink:href="#icon-pwd"></use>
                </svg>
                <input type="password" placeholder="确认密码" v-model="repwd">
                <span class="verify" v-show="verPwd">密码不能为空</span>
            </div>
            <button @click="Login">REGISTER</button>
            <div class="login">
                <span>已有账号？ </span>
                <a class="link" @click="router.push({ name: 'login' })">去登录</a>
            </div>
        </div>
    </div>
</template>
  
<script setup lang="ts">

import { ref } from 'vue'
import { register } from "@/api"
import { useRouter } from "vue-router"
import { ElMessage } from "element-plus"
const router = useRouter()
const name = ref("")
const pwd = ref("")
const repwd = ref("")
// 提示框是否出现
const verName = ref(false)
const verPwd = ref(false)
const verRepwd = ref(false)
// 验证
const verify = (): boolean => {
    let flag = false
    if (name.value == "") {
        verName.value = true
        flag = true
    }
    if (pwd.value == "") {
        verPwd.value = true
        flag = true
    }
    if (repwd.value == "") {
        verRepwd.value = true
        flag = true
    }
    return flag
}
// 用户注册
const Login = async () => {
    if (verify()) {
        return
    }
    // 两次密码不一致
    if (pwd.value != repwd.value) {
        ElMessage.error("两次密码不一致")
        return
    }
    const { data } = await register({ name: name.value, password: pwd.value, identity: repwd.value })
    if (data.code != 0) {
        ElMessage.error(data.msg)
        return
    }


    // 注册成功
    ElMessage.success("注册成功，即将到登录页面")
    router.push({ name: 'login' })
}
</script>
  
<style lang="scss" scoped>
.register {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    margin: auto;
    height: 380px;
    width: 300px;
    background-color: #fff;
    box-shadow: rgba(0, 0, 0, 0.1) 0px 10px 15px -3px, rgba(0, 0, 0, 0.05) 0px 4px 6px -2px;

    .title {
        margin-top: 20px;
        font-size: 20px;
        text-align: center;
        font-weight: bolder;
    }

    .info {
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

        button {
            margin-top: 20px;
            background-color: #0090fc;
            font-weight: bolder;
            border-radius: 8px;
            box-shadow: rgba(0, 0, 0, 0.24) 0px 3px 8px;
            width: 200px;
            height: 40px;
            color: #fff;
            border: none;
        }

        .login {
            margin-top: 23px;
            font-size: 12px;
            color: #ccc;

            .link {
                color: #8cc63e;
            }
        }

        .verify {
            display: inline-block;
            width: 200px;
            color: red;
            font-size: 10px;
        }
    }
}</style>