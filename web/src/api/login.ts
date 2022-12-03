// name: 登录接口
// author: Ethan.Wang
// desc:

import request from "./index";
import {PromiseRes} from './api'

// 获取登录链接 - 请求参数
export interface LoginUrlReq {
    back: string
}

// 获取登录链接 - 返回结构体
export interface LoginUrlRes {
    code: string
    url: string
}

// 获取登录链接请求
export const LoginUrl = (data: LoginUrlReq): PromiseRes<LoginUrlRes> => request({
    method: "post",
    url: "/auth/wx/login",
    data: data,
})

// ===================================

// 获取扫码登录状态 - 请求参数
interface LoginStatusReq {
    query: string
}

// 获取扫码登录状态 - 返回结构体
interface LoginStatusRes {
    status: boolean
    token: string
    uid: string
    nickname: string
    avatar: string
    location: string

}

// 获取扫码登录状态
export const LoginStatus = (data: LoginStatusReq): PromiseRes<LoginStatusRes> => request({
    method: "post",
    url: "/auth/wx/login/status",
    data: data,
})
