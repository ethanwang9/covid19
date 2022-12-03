// name: 可视化面板主要内容
// author: Ethan.Wang
// desc:

import request from "./index";
import {PromiseRes} from './api'
import {DangerAreaLocationRes} from "./query";

// 获取确诊统计详细数据 - 请求
export interface AreaStatDetailsReq {
    province: string
}

// 获取确诊统计详细数据
export const AreaStatDetails = (data: AreaStatDetailsReq): PromiseRes<string[][]> => request({
    method: "get",
    params: data,
    url: "/panel/area_stat/details",
})

// ===================================

// 获取风险地区

export interface DangerAreaReq {
    province: string
}

// 获取风险地区
export const DangerArea = (data: DangerAreaReq): PromiseRes<DangerAreaLocationRes> => request({
    method: "get",
    params: data,
    url: "/panel/area_danger",
})

// ===================================

// 获取拥挤数据

export interface AreaStatReq {
    province: string
}

export interface AreaStatRes {
    currentConfirmedCount: number
    confirmedCount: number
    deadCount: number
    curedCount: number
}

// 获取风险地区
export const AreaStat = (data: AreaStatReq): PromiseRes<AreaStatRes> => request({
    method: "get",
    params: data,
    url: "/panel/area_stat",
})

// ===================================

// 热点消息 - res
export interface HotMessageRes {
    area: string
    areaCode: string
    cityCode: string
    contentTag: string
    desc: string
    from: string
    fromDesc: string
    garea: string
    id: number
    imgUrl: string
    jumpLink: JumpLink
    publicTime: string
    tags: any[]
    timeShaftDesc: string
    title: string
}

export interface JumpLink {
    appId: string
    appVer: string
    type: number
    url: string
}

// 热点消息 - req
export interface HotMessageReq {
    province: string
}

// 热点消息
export const HotMessage = (data: HotMessageReq): PromiseRes<HotMessageRes[]> => request({
    method: "get",
    params: data,
    url: "/panel/hot_message",
})

// ===================================

// 获取历史现存确诊数据 - 请求
export interface HistoryCurrConfReq {
    province: string
}

// 获取历史现存确诊数据 - 返回
export interface HistoryCurrConfRes {
    confirmedCount: number
    confirmedIncr: number
    curedCount: number
    curedIncr: number
    currentConfirmedCount: number
    currentConfirmedIncr: number
    dateId: number
    deadCount: number
    deadIncr: number
    highDangerCount: number
    midDangerCount: number
    suspectedCount: number
    suspectedCountIncr: number
}

// 获取历史现存确诊数据
export const HistoryCurrConf = (data: HistoryCurrConfReq): PromiseRes<HistoryCurrConfRes[]> => request({
    method: "get",
    params: data,
    url: "/panel/history/currentConfirmed",
})

// ===================================
export interface GeoMapReq {
    province: string
}

// 获取地图地理数据
export const GeoMap = (data: GeoMapReq): PromiseRes<object> => request({
    method: "get",
    params: data,
    url: "/panel/geo",
})

// ===================================

// 地图数据 - 请求
export interface GeoMapDataReq {
    province: string
}

// 地图数据 - 返回
export interface GeoMapDataRes {
    value: any
    center: number[]
}

// 获取地图地理数据
export const GeoMapData = (data: GeoMapDataReq): PromiseRes<GeoMapDataRes> => request({
    method: "get",
    params: data,
    url: "/panel/geo/data",
})

// ===================================

// 城市列表 - 请求
export interface CityListReq {
    province: string
}

// 城市列表 - 返回
export interface CityListRes {
    value: string
    label: string
}

// 城市列表
export const CityList = (data: CityListReq): PromiseRes<CityListRes[]> => request({
    method: "get",
    params: data,
    url: "/panel/city_list",
})