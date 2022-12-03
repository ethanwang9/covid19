import axios from "axios";
import {createSign, timestamp} from "../utils/tool";
import useStore from "../store";
import {useRouter} from "vue-router";

const request = axios.create({
    baseURL: import.meta.env.VITE_API_PATH,
})


// 请求拦截 - 发送前
request.interceptors.request.use(function (config) {
    // 添加签名
    if (config.method === "get") {
        let {params} = config
        params["timestamp"] = timestamp()
        params["sign"] = createSign(params)
        config.params = params
    } else {
        let {data} = config
        data["timestamp"] = timestamp()
        data["sign"] = createSign(data)
        config.data = data
    }

    // 添加认证
    config.headers!["Authorization"] = "Bearer " + useStore().app.token
    config.headers!["Content-Type"] = "application/x-www-form-urlencoded"

    return config;
}, function (error) {
    return Promise.reject(error);
})

// 请求拦截- 接收后
request.interceptors.response.use(function (response) {
    const store = useStore()
    const router = useRouter()

    // 判断返回接口状态码
    const code = response.data.code || undefined

    // 1. auth 返回错误
    if (code === 305) {
        // 清空本地缓存
        store.app.clear()
        setTimeout(() => {
            window.location.href = "/login"
        }, 2500)
    }

    // 2. 错误类型
    if (code > 299 || code < 200) {
        return Promise.reject(response.data)
    }

    return response.data;
}, function (error) {
    return Promise.reject(error);
})

export default request

