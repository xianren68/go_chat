import req from "./req.ts";
// 用户登录
export const login = (data:{name:string,password:string})=>req.post('/login',data)
// 获取联系人列表
export const getPersonList = ()=>req.get('/auth/relation/friendlist')
// 获取未读消息
export const getUnreadMsg = ()=>req.get('/auth/user/unread')
// 获取加入的群列表
export const getCommunityList = ()=>req.get('/auth/relation/grouplist')
// 注册用户
export const register = (data:{name:string,password:string,identity:string})=>req.post('/new',data)
// 获取emoji表情
export const getEmojis = ()=>req.get('/auth/user/emojis')