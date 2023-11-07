import { messageInt } from '@/type'
import {defineStore} from 'pinia'
import { reactive } from 'vue'
import { getNotice } from '@/db'
import { userStore } from '.'
const noticeStore = defineStore('notice', ()=>{
    const noticeList:Array<messageInt> = reactive([])
    const getNoticeList = async() => {
        const userstore = userStore()
        // 从数据库中获取数据
        noticeList.push(...await getNotice(userstore.db!))
    }
    return {noticeList,getNoticeList}
})
export default noticeStore