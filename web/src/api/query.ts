// name: 查询接口
// author: Ethan.Wang
// desc:

import request from "./index";
import {PromiseRes} from './api'

// 出行政策城市列表 - 返回
export interface CityListRes {
    children: CityListResChildren[]
    cityCode: string
    label: string
}

export interface CityListResChildren {
    children: any[]
    cityCode: string
    label: string
}

// 出行政策城市列表
export const CityList = (): PromiseRes<CityListRes[]> => request({
    method: "get",
    params: {},
    url: "/query/travel/city_list",
})

// ===================================

// 获取出行政策 - 请求
export interface TravelReq {
    from: string
    to: string
}

// 获取出行政策 - 返回
export interface TravelRes {
    code: number
    from: From
    importantDict: ImportantDict
    importantNotice: string
    inKeywordDict: InKeywordDict
    message: string
    outKeywordDict: OutKeywordDict
    to: To
}

interface From {
    cityCode: string
    cityName: string
    code: number
    diseaseControlTel: any
    highInDesc: string
    isInUpdate: boolean
    isOutUpdate: boolean
    labels: Label[]
    lowInDesc: string
    medianInDesc: string
    message: string
    news: News
    outDesc: string
    provinceName: string
    riskLevel: number
    sourceContext: string
    sourceLink: SourceLink
}

interface Label {
    colour: string
    isJump: boolean
    isShow: boolean
    jumpLink: JumpLink
    label: string
}

interface JumpLink {
    appId: string
    appVer: string
    type: number
    url: string
}

interface News {
    context: string
    isShow: boolean
    jumpLink: JumpLink2
}

interface JumpLink2 {
    appId: string
    appVer: string
    type: number
    url: string
}

interface SourceLink {
    appId: string
    appVer: string
    type: number
    url: string
}

interface ImportantDict {
    isShow: boolean
    type: string
    value: string
}

interface InKeywordDict {
    isShow: boolean
    type: string
    value: string
}

interface OutKeywordDict {
    isShow: boolean
    type: string
    value: string
}

interface To {
    cityCode: string
    cityName: string
    code: number
    diseaseControlTel: any
    highInDesc: string
    isInUpdate: boolean
    isOutUpdate: boolean
    labels: Label2[]
    lowInDesc: string
    medianInDesc: string
    message: string
    news: News2
    outDesc: string
    provinceName: string
    riskLevel: number
    sourceContext: string
    sourceLink: SourceLink2
}

interface Label2 {
    colour: string
    isJump: boolean
    isShow: boolean
    jumpLink: JumpLink3
    label: string
}

interface JumpLink3 {
    appId: string
    appVer: string
    type: number
    url: string
}

interface News2 {
    context: string
    isShow: boolean
    jumpLink: JumpLink4
}

interface JumpLink4 {
    appId: string
    appVer: string
    type: number
    url: string
}

interface SourceLink2 {
    appId: string
    appVer: string
    type: number
    url: string
}


// 获取出行政策
export const Travel = (data: TravelReq): PromiseRes<TravelRes> => request({
    method: "get",
    params: data,
    url: "/query/travel",
})

// ===================================

// 获取归属地风险地区 - 返回
export interface DangerAreaLocationRes {
    update_time: string
    h_count: number
    h_list: DAList[]
    m_count: number
    m_list: DAList[]
    l_count: number
    l_list: DAList[]
}

export interface DAList {
    type: string
    province: string
    city: string
    county: string
    area_name: string
    communitys: string[]
}


// 获取归属地风险地区
export const DangerAreaLocation = (): PromiseRes<DangerAreaLocationRes> => request({
    method: "get",
    params: {},
    url: "/query/denger_area/location",
})

// ===================================

// 获取城市列表 - 返回
export interface ChinaCityListRes {
    value: string
    label: string
    children: ChinaCityListRes[]
}

// 获取城市列表 - 精确到街道
export const ChinaCityList = (): PromiseRes<ChinaCityListRes[]> => request({
    method: "get",
    params: {},
    url: "/query/danger_area/city",
})

// ===================================
// 查询风险地区 - 请求
export interface QueryDangerReq {
    province: string
    county: string
    city: string
}

// 查询风险地区 - 返回
export interface QueryDangerRes {
    info: QueryDangerResInfo,
    level: string
    update_time: string
}

export interface QueryDangerResInfo {
    type: string
    province: string
    city: string
    county: string
    area_name: string
    communitys: QueryDangerResInfo[]
}

export interface QueryDangerResInfo {
    level: string
    info: string
}

// 查询风险地区
export const QueryDanger = (data: QueryDangerReq): PromiseRes<QueryDangerRes> => request({
    method: "get",
    params: data,
    url: "/query/danger_area/query",
})

// ===================================

// 获取全国风险地区 - 返回
export interface QueryDangerAllRes {
    end_update_time: string
    hcount: number
    highlist: DAList[]
    mcount: number
    middlelist: DAList[]
    lcount: number
    lowlist: DAList[]
}

// 获取全国风险地区
export const QueryDangerAll = (): PromiseRes<QueryDangerAllRes> => request({
    method: "get",
    params: {},
    url: "/query/danger_area/all",
})