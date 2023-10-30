const socket = {
    s:<null|WebSocket>null,
    init(){
        if(this.s == null){
            // 获取token
            let token = localStorage.getItem('token')
            if(!token){
                return
            }
            this.s = new WebSocket('ws://localhost:8080/v1/chat',[token])
            this.s.onerror = (e)=>{
                console.log(e)
            }
        }
    }
}
export default socket
