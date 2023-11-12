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
// 通过用户名查找用户
export const getUserByName = (name:string)=>req.get(`/auth/user/findbyname?name=${name}`)
// 通过群组名查找群组
export const getCommunityByName = (name:string)=>req.get(`/auth/relation/findgroupname?name=${name}`)
// 添加好友
export const newFriend = (data :{name:string})=>req.post('/auth/relation/addfriend',data)
// 加入群聊
export const joinCommunity = (data :{name:string,id:number})=>req.post('/auth/relation/joingroup',data)
// 上传图片
export const uploadImage = (data:FormData)=>req({
    method:'post',
    url:'/auth/user/uploadimg',
    data:data,
    headers:{
        'Content-Type':'multipart/form-data'
    }
})
// 更新用户信息
export const updateUserInfo = (data:any)=>req.put('/auth/user/update',data)
// 发送验证码
export const sendVerify = (data:{email:string})=>req.post('/sendverify',data)
// 修改密码
export const verfiyCode = (data:{email:string,code:string,password:string,id:string})=>req.post('/verifycode',data)
// 创建群聊
export const createCommunity = (data:{name:string})=>req.post('/auth/relation/newgroup',data)