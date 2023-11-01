import {messageInt, sessionInt, unreadData } from "@/type"
import { log } from "console"

// 打开数据库
export function openDb(dbName: string):Promise<IDBDatabase>{
    return new Promise((resolve, reject) => {
        let db: IDBDatabase
        const request = window.indexedDB.open(dbName, 1)
        request.onerror = () => {
            console.log('打开数据库失败')
            reject(request.error)
        }
        request.onsuccess = () => {
            db = request.result
            resolve(db)
        }
        request.onupgradeneeded = () => {
            // 创建数据库时调用
            db = request.result
            // 创建用户会话表
            const sessionStore = db.createObjectStore('usersession', {keyPath: 'ID'})
            sessionStore.createIndex('ID', 'ID', {unique: true})
            // 创建群聊会话表
            const sessionStore2 = db.createObjectStore('groupsession', {keyPath: 'ID'})
            sessionStore2.createIndex('ID', 'ID', {unique: true})
            // 创建未读消息表
            const unReadStore = db.createObjectStore('unRead', {autoIncrement: true } )
            unReadStore.createIndex('from_id', 'from_id', {unique: false})
            unReadStore.createIndex('type', 'type', {unique: false})
            // 创建消息表
            const messageStore = db.createObjectStore('message', {autoIncrement: true })
            messageStore.createIndex('from_id', 'from_id', {unique: false})
            messageStore.createIndex('type', 'type', {unique: false})
            messageStore.createIndex('group_id', 'group_id', {unique: false})
        }
    })
}
// 插入数据
export function insertData(db: IDBDatabase, tableName: string, data: messageInt){
    const request = db.transaction([tableName], 'readwrite').objectStore(tableName).add(data)
    request.onsuccess = () => {
        console.log('插入数据成功')
    }
    request.onerror = () => {
        console.log('插入数据失败');
        
    }
}

// 获取未读消息数据
export function getunRead(db: IDBDatabase): Promise<unreadData>{
    return new Promise((resolve)=>{
        let data:unreadData = {notify:0,msg:0,notifylist:[],msglist:[]}
        const store = db.transaction(['unRead'], 'readonly').objectStore('unRead')
        const request = store.openCursor()
        request.onsuccess = () => {
            let cursor = request.result
            if (cursor){
                let d = cursor.value
                if(d.type < 3){
                    data.notify ++
                    data.notifylist.push(d)
                }else{
                    data.msg++
                    data.msglist.push(d)
                }
                cursor.continue()
            }
            resolve(data)
        }

    })
}
// 获取会话信息
export function getSession(db: IDBDatabase,tableName:string):Promise<{count:number,map:Map<number,sessionInt>}>{
    return new Promise((resolve)=>{
        let data = {count:<number>0,map:new Map<number,sessionInt>()}
        const store = db.transaction([tableName], 'readonly').objectStore(tableName)
        const request = store.openCursor()
        request.onsuccess = () => {
            let cursor = request.result
            if (cursor){
                let d = cursor.value
                data.count ++
                data.map.set(d.id,d)
                cursor.continue()
            }
            resolve(data)
        }

    })
}

// 判断是否有未读通知
export function isExistNotify(db:IDBDatabase):Promise<boolean>{
    return new Promise((resolve)=>{
        const store = db.transaction(['message'], 'readonly').objectStore('message')
        const request = store.openCursor()
        request.onsuccess = ()=>{
            let cursor = request.result
            if (cursor){
                let d = cursor.value
                if(d.type > 2){
                    resolve(true)
                    return
                }
                cursor.continue
            }
            resolve(false)
        }
    })
}

// 更新/插入会话信息
export function putSession(db:IDBDatabase,tableName:string,data:sessionInt){
    console.log(data)
    db.transaction([tableName],'readwrite').objectStore(tableName).put(data)
}

// 存储聊天信息
export function saveMessage(db:IDBDatabase,data:messageInt){
    db.transaction(['message'],'readwrite').objectStore('message').add(data)
}
