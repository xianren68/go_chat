import {defineStore} from "pinia"
import {getPersonList} from "../api"
import {reactive} from "vue";
// 联系人/群组列表
const contactStore = defineStore('contact',()=>{
    // 联系人列表
    const contactPerson = reactive([])
    // 获取联系人列表
    const getFriendList = async ()=>{
        const{data} = await getPersonList()
        console.log(data)
        Object.assign(contactPerson,data.data)
        console.log(contactPerson)
    }
    return {contactPerson,getFriendList}
})
export default contactStore