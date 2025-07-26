import {defineStore} from "pinia";

const useAdminStore = defineStore({
    id: 'admin',
    state: () => {
        // 后台配置项
        const cfg = {
            // 是否显示 header
            isShowHeader: true,
            // 可视化页面样式
            adminMainCustom: "height: calc(100vh - 60px);",
        }
        return {
            cfg,
        }
    },
    getters: {
        // 是否显示 header
        isShowHeader: (state) => {
            return state.cfg.isShowHeader
        },
        // 可视化页面样式
        adminMainCustom: (state) => {
            return state.cfg.adminMainCustom
        },
    },
    actions: {
        // 改变 header 是否显示
        // isShow: true->显示 false->隐藏
        // showType: true->后台数据看板显示 false->其他页面显示
        changeShowHeader(isShow: boolean, showType: boolean = false) {
            if (isShow) {
                // 显示
                if (showType) {
                    // 后台数据看板显示
                    this.cfg.isShowHeader = isShow
                    this.cfg.adminMainCustom = "padding: 0; overflow: hidden; height: calc(100vh - 60px);"
                } else {
                    // 其他页面显示
                    this.cfg.isShowHeader = isShow
                    this.cfg.adminMainCustom = "height: calc(100vh - 60px);"
                }
            } else {
                // 隐藏
                this.cfg.isShowHeader = isShow
                this.cfg.adminMainCustom = "padding: 0; overflow: hidden; height: 100%"
            }
        },
    },

})

export default useAdminStore