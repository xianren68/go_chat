<template>
    <div class="re">
        <div class="content">
            <p class="title">CHANGE</p>
            <div class="info">
                <div class="query item" v-if="query">
                    <svg class="icon">
                        <use xlink:href="#icon-youxiang"></use>
                    </svg>
                    <div>{{email}}</div>
                </div>
                <div class="email item" v-else>
                    <svg class="icon">
                        <use xlink:href="#icon-youxiang"></use>
                    </svg>
                    <input type="email" placeholder="您绑定的邮箱" v-model="email">
                    <span class="verify" v-show="emailShow">邮箱输入错误</span>
                </div>
                <div class="code item">
                    <div class="con">
                        <svg class="icon">
                            <use xlink:href="#icon-yanzheng"></use>
                        </svg>
                        <input type="text" placeholder="请输入验证码" v-model="code">
                        <button @click="sendCode">{{times==0?"发送验证码":(times + "s")}}</button>
                    </div>
                    <span class="verify" v-show="verCode">验证码为6位</span>
                </div>
                <div class="pwd item">
                    <svg class="icon">
                        <use xlink:href="#icon-pwd"></use>
                    </svg>
                    <input type="password" placeholder="新密码" v-model="pwd">
                    <span class="verify" v-show="verPwd">密码不能为空</span>
                </div>
                <div class="pwd item">
                    <svg class="icon">
                        <use xlink:href="#icon-pwd"></use>
                    </svg>
                    <input type="password" placeholder="确认密码" v-model="repwd">
                    <span class="verify" v-show="verPwd">密码不能为空</span>
                </div>
                <button @click="submit">SUBMIT</button>
            </div>
        </div>
    </div>
</template>
  
<script setup lang="ts">

import { ref,onMounted } from 'vue'
import { useRouter,useRoute} from "vue-router"
import { ElMessage } from "element-plus"
import { isEmail } from '@/utils/verify'
import {sendVerify,verfiyCode} from '@/api'
const router = useRouter()
const route = useRoute()
const email = ref("")
const code = ref("")
const pwd = ref("")
const repwd = ref("")
// 提示框是否出现
const verCode = ref(false)
const verPwd = ref(false)
const verRepwd = ref(false)
const emailShow = ref(false)
const id = ref("")
// 过期时间
const times = ref(0)
// 是否是修改密码
const query = ref(false)
// 发送验证码
const sendCode = async () => {
    if (times.value > 0){
        ElMessage.info("请不要重复发送")
        return
    }
    if(!isEmail(email.value)){
        emailShow.value = true
    }
    emailShow.value = false
    const {data} = await sendVerify({email:email.value})
    if(data.code == 0){
        ElMessage.success(data.msg)
        id.value = data.Id + ""
        times.value = 300
        const ti = setInterval(()=>{
            times.value--
            if(times.value == 0){
                clearInterval(ti)
            }
        },1000)
        return
    }
    ElMessage.error(data.msg)
}
// 验证
const verify = (): boolean => {
    let flag = false
    if(!isEmail(email.value)){
        emailShow.value = true
        flag = true
    }
    if (code.value == "" || code.value.length != 6) {
        verCode.value = true
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
// 用户修改密码
const submit = async () => {
    if (verify()) {
        return
    }
    // 两次密码不一致
    if (pwd.value != repwd.value) {
        ElMessage.error("两次密码不一致")
        return
    }
    const {data} = await verfiyCode({email:email.value,code:code.value,password:pwd.value,id:id.value})
    if(data.code == 0){
        ElMessage.success(data.msg)
        localStorage.clear()
        router.push('/login')
        return
    }
    ElMessage.error(data.msg)
}
onMounted(() => {
    const qu:any = route.query
    if(!qu){
        return
    }
    query.value = true
    email.value = qu.email

})
</script>
  
<style lang="scss" scoped>
.re {
    height: 100vh;
    width: 100vw;
    background: linear-gradient(to right, #00bf8f, #001510);
    .content {
        position: absolute;
        left: 0;
        right: 0;
        top: 0;
        bottom: 0;
        margin: auto;
        height: 460px;
        width: 400px;
        color:#fff;
        background-color: #202020;
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
                width: 253px;

                .icon {
                    display: inline-block;
                    width: 20px;
                    height: 20px;
                    fill:#fff;
                }
            }
            .email {
                margin-top: 50px;
            }
            .query{
                margin-top: 50px;
                div{
                    margin-left: 5px;
                    display: inline-block;
                }
            }
            .code {
                margin-top: 35px;
                .con{
                    display: flex;
                align-items: center;
                button{
                    margin:0;
                    background-color: transparent;
                    border: 1px solid #fff;
                    border-radius: 0;
                    border-bottom: none;
                    font-weight: normal;
                    font-size: 10px;
                    height: 30px;
                    width: 80px;
                }
                }
            }

            .pwd {
                margin-top: 35px;
            }

            input {
                display: inline-block;
                height: 30px;
                width: 160px;
                color:#fff;
                margin-left: 10px;
                border: none;
                outline: none;
                background-color: transparent;
            }

            button {
                margin-top: 40px;
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
    }


}
</style>