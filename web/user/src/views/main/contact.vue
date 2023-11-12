<template>
  <div class="contact">
  <div class="left">
    <!--    搜索-->
    <div class="top">
      <div class="search" @click="()=>{isShow = !isShow;searchShow=false;search=''}">
        <svg class="icon" :class="{select_icon:isShow}">
          <use xlink:href="#icon-jiahao_o"></use>
        </svg>
      </div >
      <input type="text" v-show="isShow" class="search-input" placeholder="添加用户/群组" v-model="search"
      @keyup.enter="searchEnter">
      <div class="cancel" v-show="isShow" @click="search=''">
        <svg class="icon">
          <use xlink:href="#icon-quxiao"></use>
        </svg>
      </div>
    </div>
    <!-- 添加用户/群组 -->
    <findVue v-if="searchShow" :data="searchData" v-model="searchShow"></findVue> 
    <!-- 通知列表 -->
    <noticelist v-if="noticeStore.noticeList.length>0"></noticelist>
    <!--    切换群组/好友-->
    <div class="switch">
      <div class="friend" :class="{select:switchCt}" @click="switchFriend">好友</div>
      <div class="group" :class="{select:!switchCt}" @click="switchGroup">群组</div>
    </div>
    <!--    列表-->
    <List :List="switchCt?contactStore.contactPerson:contactStore.groupList" v-model="selectIndex"></List>
    <div class="newGroup" v-if="!switchCt" title="创建群聊">
      <svg class="icon" @click="dialogFormVisible= true">
        <use xlink:href="#icon-tianjia"></use>
      </svg>
    </div>
    <el-dialog v-model="dialogFormVisible" title="创建群聊">
      <el-form :model="form">
        <el-form-item label="群名称" label-width="140px">
          <el-input v-model="form.name" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogFormVisible = false">取消</el-button>
          <el-button type="primary" @click="newGroup">
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
    <div class="right">
      <router-view></router-view>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref,reactive} from 'vue'
import {useContactStore} from "../../store"
import {onBeforeMount,watch} from "vue"
import List from '@/components/contact/list.vue'
import {useRouter} from 'vue-router'
import findVue from '@/components/contact/find.vue'
import { getUserByName,getCommunityByName,createCommunity } from '@/api'
import { ElMessage } from 'element-plus'
import noticelist from '@/components/contact/noticelist.vue'
import { useNoticeStore } from '../../store'
const router = useRouter()
const noticeStore = useNoticeStore()
// 控制搜索框是否显示
const isShow = ref(false)
// 控制选择好友还是群组
const switchCt = ref(true)
// 选中的索引
const selectIndex = ref(-1)
// 跳转到群聊会话
const jumpSession = (index:number)=>{
  const {ID,Name,Avatar} = contactStore.groupList[index]
  router.push({name:'session',query:{ID,Name,Avatar,type:2}})
}
// 监听索引的变化
watch(selectIndex,(newSelect)=>{
  // 如果在群聊列表，直接跳转
  
  if(!switchCt.value){
    jumpSession(newSelect)
    return
  }
  router.push(`/contact/userinfo/${newSelect}`)
})
// 好友
const contactStore = useContactStore()
// 切换到用户列表
const switchFriend = ()=>{
  switchCt.value = true
  searchShow.value = false
  search.value = ""
}
// 切换到群组列表
const switchGroup = ()=>{
  switchCt.value = false
  if(contactStore.groupList.length == 0){
    contactStore.getGroupList()
  }
  searchShow.value = false
  search.value = ""
}
// 搜索数据
const search = ref("")
// 控制搜索数据的显示与隐藏
const searchShow = ref(false)
// 搜索获取到的数据
const searchData = reactive<any>({})
// 搜索数据
const searchEnter = async ()=>{
  if(search.value == ""){
    return
  }
  let res:any
  if (switchCt.value){
    if (contactStore.contactPerson.find(item=>item.Name==search.value)){
      ElMessage.error("已存在该好友")
      return
    }
    res = await getUserByName(search.value)
  }else {
    if (contactStore.groupList.find(item=>item.Name==search.value)){
      ElMessage.error("已加入该群组")
      return
    }
    res = await getCommunityByName(search.value)
  }
  let data = res.data
  if(data.code != 0){
    ElMessage.error(data.msg)
    return
  }
  Object.assign(searchData,data.data)
  searchData.isUser = switchCt.value
  searchShow.value = true
}
// 控制创建群聊是否显示
const dialogFormVisible = ref(false)
const form = reactive({
  name:""
})
// 创建群聊
const newGroup = async ()=>{
  dialogFormVisible.value = false
  const {data} = await createCommunity({name:form.name})
  if(data.code != 0){
    ElMessage.error(data.msg)
    return
  }
  ElMessage.success("创建成功")
  contactStore.getGroupList()

}
// 挂载之前，获取用户列表
onBeforeMount(()=>{
  selectIndex.value = -1
  if(contactStore.contactPerson.length == 0){
    contactStore.getFriendList()
  }
  // 获取未读通知
  noticeStore.getNoticeList()
})
</script>

<style scoped lang="scss">
.contact {
  position: relative;
  display: flex;
  align-items: center;
  height: 100%;
  flex-grow:1;
  .left{
    height: 100%;
    width: 20%;
    .top{
      margin-top: 2px;
      margin-left: 10px;
      height:40px;
      display: flex;
      justify-content: start;
      .search{
        width: 20px;
        height: 20px;
        display: flex;
        justify-content: center;
        align-items: center;
        .icon{
          width: 15px;
          height: 15px;
          fill: #bbb;
        }
        .select_icon{
          fill:var(--icon-active-color);
        }
      }
      .search-input{
        margin-left: 10px;
        padding-left: 10px;
        height: 20px;
        width: 90%;
        font-size: 12px;
        border: none;
        border: #ccc 1px solid;
        border-radius: 5px;
        background-color: transparent;
        outline: none;
        transition: .5s;
      }
      .cancel{
        position: relative;
        height: 20px;
        width: 20px;
        background-color: #f5f4fa;
        border-radius: 50%;
        left:-18px;
        top:2px;
        transition: .5s;
        .icon {
          height: 15px;
          width: 15px;
          color: #ccc;
        }
      }
    }
    .switch{
      display: flex;
      justify-content: space-around;
      align-items: center;
      margin-left: 20px;
      margin-right: 20px;
      font-size: 14px;
      height: 30px;
      font-weight: bolder;
      div{
        width: 30%;
        height: 100%;
        text-align: center;
        transition: .5s;
      }
      .select{
        color: #008efe;
        border-bottom: #008efe 2px solid;
      }

    }
    .newGroup{
      position: absolute;
      left:130px;
      top:430px;
      .icon {
        height: 40px;
        width: 40px;
        fill:#f9dada;
      }

    }
  }
  .right {
    flex-grow: 1;
    height: 100%;
    background-color: #fff;
  }
}
</style>