// name: 公共信息查询
// author: Ethan.Wang
// desc:

// 获取微信公众号信息 - 返回
import {PromiseRes} from "./api";
import request from "./index";

export interface PublicInfoRes {
    create_at: string
    updated_at: string
    copyright: string
    gov_no1: string
    gov_no2: string
    mp_url: string
    mp_img: string
    mail: string
    blog: string
}

// 获取微信公众号信息
export const PublicInfo = (): PromiseRes<PublicInfoRes> => request({
    method: "get",
    params: {},
    url: "/public/info",
})