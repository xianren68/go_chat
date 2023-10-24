import req from "./req.ts";
// 用户登录
export const login = (data:{name:string,password:string})=>req.post('/login',data)
// 获取联系人列表
export const getPersonList = ()=>req.get('/auth/relation/friendlist')