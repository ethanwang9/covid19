import {defineStore} from "pinia";

// 用户信息
interface userInfo {
    uid: string
    nickname: string
    avatar: string
    location: string
}

const useAppStore = defineStore({
    id: 'app',
    state: () => {
        // 登录UUID
        const uuid: string = ""
        // 用户凭证
        const token: string = ""
        // 用户信息
        const userInfo: userInfo = {
            uid: "",
            nickname: "",
            avatar: "",
            location: "",
        }
        return {
            uuid,
            token,
            userInfo,
        }
    },
    getters: {},
    actions: {
        // 是否登录
        isLogin(): boolean {
            if (this.token.length == 0) {
                return false
            }
            return true
        },
        // 设置用户信息
        setUserInfo(uid: string, nickname: string, avatar: string, location: string) {
            this.userInfo = {
                uid,
                nickname,
                avatar,
                location,
            }
        },
        // 清空数据
        clear() {
            this.userInfo = {
                uid: "",
                nickname: "",
                avatar: "",
                location: "",
            }
            this.token = ""
        },
        // 设置用户归属地
        setUserLocation(value: string) {
            this.userInfo.location = value
        },
    },
    persist: {
        storage: localStorage,
        paths: ['token', 'userInfo']
    },
})

export default useAppStore