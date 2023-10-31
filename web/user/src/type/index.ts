// 用户数据类型
export interface userInt {
    Name:string,
    ID:number,
    Avatar:string,
    Phone:string,
    Email:string,
    Gender:string,
    Line:string
}
// 会话类型
export interface sessionInt {
    ID:number,
    Name:string,
    Avatar:string,
    // 当前会话的最后一条消息
    lastMsg?:string,
}