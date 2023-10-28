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
// 消息类型
export interface messageInt {
    ID:number,
    Name:string,
    Avatar:string,
    // 与这个联系人的最后一条消息
    lastMsg?:string,
}