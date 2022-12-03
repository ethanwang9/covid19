// name: 系统信息接口
// author: Ethan.Wang
// desc:

import {PromiseRes} from "./api";
import request from "./index";

// 获取微信公众号信息 - 返回
export interface SysMpInfoRes {
    appid: string
    secret: string
}

// 获取微信公众号信息
export const SysMpInfo = (): PromiseRes<SysMpInfoRes> => request({
    method: "get",
    params: {},
    url: "/sys/mp",
})

// ===================================

// 设置微信公众号信息 - 请求
export interface SysMpInfoSetReq {
    appid: string
    secret: string
}

// 设置微信公众号信息
export const SysMpInfoSet = (data: SysMpInfoSetReq): PromiseRes<null> => request({
    method: "post",
    data,
    url: "/sys/mp",
})

// ===================================

// 获取系统信息 - 返回
export interface SysInfoRes {
    copyright: string
    gov_no1: string
    gov_no2: string
    mp_url: string
    mp_img: string
    mail: string
    blog: string
}

// 获取系统信息
export const SysInfo = (): PromiseRes<SysInfoRes> => request({
    method: "get",
    params: {},
    url: "/sys/info",
})

// ===================================

// 设置系统信息 - 请求
export interface SysInfoSetReq {
    copyright: string
    gov_no1: string
    gov_no2: string
    mp_url: string
    mp_img: string
    mail: string
    blog: string
}

// 设置系统信息
export const SysInfoSet = (data: SysInfoSetReq): PromiseRes<null> => request({
    method: "post",
    data,
    url: "/sys/info",
})