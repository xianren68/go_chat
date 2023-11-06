import {defineStore} from 'pinia'
import { reactive } from 'vue'
import { getEmojis } from '@/api'
const emojis = defineStore("emojis",()=>{
    const emojilist:Array<string> = reactive([])
    const getEmojiList = async ()=>{
        // emoji列表不存在时才请求
        if(emojilist.length == 0){
            const {data} = await getEmojis() as any
            emojilist.push(...data.data)
        }
    }
    return {getEmojiList,emojilist}
})
export default emojis