import {defineStore} from "pinia"
import {getPersonList} from "../api"
import {reactive} from "vue";
interface userInt {
    Name:string,
    ID:number,
    Avatar:string,
    Phone:string,
    Email:string,
    Gender:string,
    Line:string
}
// 联系人/群组列表
const contactStore = defineStore('contact',()=>{
    // 联系人列表
    const contactPerson:Array<userInt> = reactive([])
    // 获取联系人列表
    const getFriendList = async ()=>{
        const{data} = await getPersonList()
        Object.assign(contactPerson,data.data)
    }
    return {contactPerson,getFriendList}
})
export default contactStore