// 是否是手机号
export function isMobile (mobile: string) {
    return /^1[3-9]\d{9}$/.test(mobile)
  }
// 是否是性别
export function isSex (sex: string) {
  return (sex === '男' || sex === '女') 
}

// 是否是邮箱
export function isEmail (email: string) {
    return /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/.test(email)
}