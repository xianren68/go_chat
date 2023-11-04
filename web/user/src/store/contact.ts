import {defineStore} from "pinia"
import {getPersonList,getCommunityList} from "../api"
import {reactive} from "vue";
import {groupInt,  userInt} from "@/type";
// 联系人/群组列表
const contactStore = defineStore('contact',()=>{
    // 联系人列表
    const contactPerson:Array<userInt> = reactive([])
    // 群聊列表
    const groupList:Array<groupInt> = reactive([])
    // 获取联系人列表
    const getFriendList = async ()=>{
        const{data} = await getPersonList()
        Object.assign(contactPerson,data.data??[])
    }
    // 获取群组列表
    const getGroupList = async ()=>{
        const {data} = await getCommunityList()
        groupList.push(...data.data??[])
    }
    return {contactPerson,getFriendList,groupList,getGroupList}
})
export default contactStore