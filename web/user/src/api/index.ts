import req from "./req.ts";
// 用户登录
export const login = (data:{name:string,password:string})=>req.post('/login',data)
