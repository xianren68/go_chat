// 用户数据类型
export interface userInt {
    Name: string,
    ID: number,
    Avatar: string,
    Phone: string,
    Email: string,
    Gender: string,
    Line: string
}
// 群组类型
export interface groupInt {
    Name: string,
    ID: number,
    OwnerId: number,
    Avatar: string,
    Line: string,
    Type: number
}
// 会话类型
export interface sessionInt {
    ID: number,
    Name: string,
    Avatar: string,
    // 当前会话的最后一条消息
    lastMsg?: string,
    // 未读消息条数
    unReadCount: number,
    // 会话类型
    type: number,
    // 最后一条记录的时间
    lastMsgTime?: number,
}

// 消息类型
export interface messageInt {
    from_id: number,
    target_id: number,
    content: string,
    group_id?: number,
    type: number,
    send_time: number,
    // 发送人名称
    send_name: string,
    // 发送人头像
    avatar: string,
    // 群聊名称
    group_name?: string,
    // 群聊头像
    group_avatar?: string,
    // 是否是markdown
    md?: boolean
}

